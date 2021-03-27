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
	"sync"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

var (
	watchAnywherePath string
	watchPassphrase   string
)

// watchCmd represents the watch command
var watchCmd = &cobra.Command{
	Use:  "watch",
	Long: `Watch Clipboard`,
	Run: func(cmd *cobra.Command, args []string) {

		// start watcher
		// clipboard
		go watchClipboard(watchAnywherePath, watchPassphrase)
		go watchServer(watchAnywherePath, watchPassphrase)

		log.Println("Started clipboard-watcher. Press CTRL-C to stop.")
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
		<-sc
		log.Println("Stopped.")
	},
}

var (
	errors int

	lastClipboardData string
	lastClipboardHash string

	mu    = sync.Mutex{}
	errMu = sync.Mutex{}
)

func handleErr(err error, m string) {
	errMu.Lock()
	errors++
	e := errors
	errMu.Unlock()

	if e > 5 {
		log.Fatalln("[err]", "(", m, ") :: Error-Limit exceeded:", errors, "::", err)
		return
	}
	log.Println("[err] ERROR reading clipboard:", err, errors, "/ 6")
}

func watchClipboard(path, pass string) {
	a := api.API()
	for {
		time.Sleep(100 * time.Millisecond)
		mu.Lock()

		data, err := clipboard.ReadAll()
		if err != nil {
			mu.Unlock()
			handleErr(err, "clipboard-read")
			continue
		}
		errors = 0

		if lastClipboardData == data {
			mu.Unlock()
			continue
		}

		// calculate new hash
		lastClipboardData = data
		lastClipboardHash = common.MD5Hash(data)

		// check if server has current clipboard
		serverHash, err := a.GetHash(path)
		if err != nil {
			mu.Unlock()
			handleErr(err, "read-server-hash")
			continue
		}

		if serverHash == lastClipboardHash {
			mu.Unlock()
			log.Println("[ ~ ] (rea) Server Hash == Local Hash")
			continue
		}

		// update server hash
		if err := a.SetContent(path, pass, data); err != nil {
			mu.Unlock()
			handleErr(err, "set-server-content")
			continue
		}

		log.Println("[ ok] Updated contents.")
		mu.Unlock()
	}
}

func watchServer(path, pass string) {
	a := api.API()
	for {
		time.Sleep(1 * time.Second)
		mu.Lock()

		hash, err := a.GetHash(path)
		if err != nil {
			mu.Unlock()
			handleErr(err, "read-server-hash")
			continue
		}

		if hash == lastClipboardHash {
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

		log.Println("[ ~ ] Received new data:", data)

		lastClipboardData = data
		lastClipboardHash = common.MD5Hash(data)

		if err := clipboard.WriteAll(data); err != nil {
			mu.Unlock()
			handleErr(err, "write-clipboard")
			continue
		}

		log.Println("[ ok] Wrote client content")
		mu.Unlock()
	}
}

func init() {
	rootCmd.AddCommand(watchCmd)

	watchCmd.Flags().StringVarP(&watchAnywherePath, "anywhere", "a", "", "Path (Anywhere)")
	cobra.CheckErr(cobra.MarkFlagRequired(watchCmd.Flags(), "anywhere"))

	watchCmd.Flags().StringVarP(&watchPassphrase, "passphrase", "s", "", "Encryption Key")
}
