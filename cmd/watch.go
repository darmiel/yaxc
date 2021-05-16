// +build client

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
	"fmt"
	"github.com/darmiel/yaxc/internal/client"
	"github.com/darmiel/yaxc/internal/common"
	"github.com/muesli/termenv"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	watchAnywherePath string
	watchPassphrase   string
	watchIgnoreServer bool
	watchIgnoreClient bool
	watchUseBase64    bool
)

// watchCmd represents the watch command
var watchCmd = &cobra.Command{
	Use:  "watch",
	Long: `Watch Clipboard`,
	Run: func(cmd *cobra.Command, args []string) {
		check := client.NewCheck(watchAnywherePath, watchPassphrase, watchUseBase64)
		done := make(chan bool)

		fmt.Println(common.StyleInfo(), "Starting Watchers:")

		if !watchIgnoreServer {
			fmt.Println(common.StyleInfo(), "*",
				termenv.String("Server").Foreground(common.Profile().Color("#D290E4")),
				"->",
				termenv.String("Client").Foreground(common.Profile().Color("#DBAB79")))

			go client.WatchServer(check, 1*time.Second, done)
		}

		if !watchIgnoreClient {
			fmt.Println(common.StyleInfo(), "*",
				common.WordServer(),
				"<-",
				common.WordClient())

			go client.WatchClient(check, 50*time.Millisecond, done)
		}

		if watchIgnoreServer && watchIgnoreClient {
			fmt.Println(common.StyleWarn(),
				"Ignored",
				termenv.String("Client").Foreground(common.Profile().Color("#DBAB79")),
				"AND",
				termenv.String("Server").Foreground(common.Profile().Color("#D290E4")),
				". That's probably a bad idea.")
		}

		if watchUseBase64 {
			fmt.Println(common.StyleInfo(), "Using",
				termenv.String("Base64").Foreground(common.Profile().Color("#A8CC8C")))
		}

		fmt.Println(common.StyleInfo(), "Started clipboard-watcher. Press CTRL-C to stop.")
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
		<-sc

		// Stopping server watcher
		done <- true
	},
}

func init() {
	rootCmd.AddCommand(watchCmd)

	watchCmd.Flags().StringVarP(&watchAnywherePath, "anywhere", "a", "", "Path (Anywhere)")
	cobra.CheckErr(cobra.MarkFlagRequired(watchCmd.Flags(), "anywhere"))

	watchCmd.Flags().StringVarP(&watchPassphrase, "passphrase", "s", "", "Encryption Key")
	watchCmd.Flags().BoolVar(&watchIgnoreServer, "ignore-server", false, "Ignore Server Updates")
	watchCmd.Flags().BoolVar(&watchIgnoreClient, "ignore-client", false, "Ignore Client Updates")
	watchCmd.Flags().BoolVarP(&watchUseBase64, "base64", "b", false, "Use Base64?")
}
