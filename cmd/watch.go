/*
Copyright © 2021 darmiel <hi@d2a.io>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/atotto/clipboard"
	"github.com/darmiel/yaxc/internal/api"
	"github.com/darmiel/yaxc/internal/common"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

var (
	watchAnywherePath string
	watchPassphrase   string
)

var (
	errors int

	lastClipboardData   string
	lastClipboardHash   string
	ignoreClipboardHash string

	mu    = sync.Mutex{}
	errMu = sync.Mutex{}
)

// watchCmd represents the watch command
var watchCmd = &cobra.Command{
	Use:  "watch",
	Long: `Watch Clipboard`,
	Run: func(cmd *cobra.Command, args []string) {

		d := make(chan int, 1)

		// start watcher
		// clipboard
		go watchClipboard(watchAnywherePath, watchPassphrase, d)
		go watchServer(watchAnywherePath, watchPassphrase, d)

		log.Println("Started clipboard-watcher. Press CTRL-C to stop.")
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
		<-sc

		// stop watchers
		d <- 1
		d <- 1

		log.Println("Stopped.")
	},
}

func handleErr(err error, m string) {
	errMu.Lock()
	errors++
	errMu.Unlock()
	log.Println("[err]", "(", m, ") ::", errors, "errors:", err)
}

func watchClipboard(path, pass string, done chan int) {
	a := api.API()
	t := time.NewTicker(100 * time.Millisecond)

	for {
		select {
		case <-t.C:
			mu.Lock()

			data, err := clipboard.ReadAll()
			if err != nil {
				mu.Unlock()
				continue
			}
			errors = 0

			// strip data
			data = strings.TrimSpace(data)

			if lastClipboardData == data {
				mu.Unlock()
				continue
			}
			log.Println("[o] Changed!")

			// calculate new hash
			lastClipboardData = data
			lastClipboardHash = common.MD5Hash(data)

			// check if server has current clipboard
			serverHash, err := a.GetHash(path)
			if err != nil && err != api.ErrErrResponse {
				handleErr(err, "read-server-hash")
				mu.Unlock()
				continue
			}

			if err != api.ErrErrResponse {
				if serverHash == lastClipboardHash {
					log.Println("[ ~ ] (rea) Server Hash == Local Hash")
					mu.Unlock()
					continue
				}
			}
			ignoreClipboardHash = serverHash

			// update server hash
			if err := a.SetContent(path, pass, data); err != nil {
				mu.Unlock()
				handleErr(err, "set-server-content")
				continue
			}

			log.Println("[ ok] Updated contents.")
			mu.Unlock()
			break

		case <-done:
			log.Println("[ ok] Stopping watch clipboard")
			return
		}
	}
}

func watchServer(path, pass string, done chan int) {
	a := api.API()
	t := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-t.C:

			time.Sleep(1 * time.Second)
			mu.Lock()

			hash, err := a.GetHash(path)
			if err != nil {
				if err != api.ErrErrResponse {
					handleErr(err, "read-server-hash")
				}
				mu.Unlock()
				continue
			}

			if hash == lastClipboardHash || hash == ignoreClipboardHash {
				mu.Unlock()
				continue
			}

			// get data
			data, err := a.GetContent(path, pass)
			if err != nil {
				mu.Unlock()
				handleErr(err, "read-server-contents")
				continue
			}

			// empty
			if strings.TrimSpace(data) == "" {
				log.Println("[wrn] Received empty data")
				mu.Unlock()
				continue
			}

			log.Println("[ ~ ] Received new data:", data)

			lastClipboardData = data
			lastClipboardHash = common.MD5Hash(data)

			if err := clipboard.WriteAll(data); err != nil {
				ignoreClipboardHash = hash
				handleErr(err, "write-clipboard")
				mu.Unlock()
				continue
			}

			log.Println("[ ok] Wrote client content")
			mu.Unlock()
			break

		case <-done:
			log.Println("[ ok] Stopping watch server")
			return
		}
	}
}

func init() {
	rootCmd.AddCommand(watchCmd)

	watchCmd.Flags().StringVarP(&watchAnywherePath, "anywhere", "a", "", "Path (Anywhere)")
	cobra.CheckErr(cobra.MarkFlagRequired(watchCmd.Flags(), "anywhere"))

	watchCmd.Flags().StringVarP(&watchPassphrase, "passphrase", "s", "", "Encryption Key")
}
