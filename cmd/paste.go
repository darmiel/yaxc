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
	"github.com/darmiel/yaxc/internal/api"
	"github.com/darmiel/yaxc/internal/common"
	"log"

	"github.com/spf13/cobra"
)

var (
	pasteAnywherePath string
	pastePassphrase   string
)

// pasteCmd represents the paste command
var pasteCmd = &cobra.Command{
	Use:  "paste",
	Long: `Paste piped content to /:path`,
	Run: func(cmd *cobra.Command, args []string) {
		// Read pipe
		pipe, err := common.ReadPipe()
		if err == common.NotPiped {
			log.Fatalln("The command is intended to work with pipes.")
			return
		}

		if pipe == "" {
			log.Fatalln("Empty input.")
			return
		}

		if err := api.API().SetContent(pasteAnywherePath, pastePassphrase, pipe); err != nil {
			log.Fatalln("ERROR ::", err)
			return
		}

		fmt.Println("Successfully uploaded contents to", pasteAnywherePath)
	},
}

func init() {
	rootCmd.AddCommand(pasteCmd)

	pasteCmd.Flags().StringVarP(&pasteAnywherePath, "anywhere", "a", "", "Path (Anywhere)")
	pasteCmd.Flags().StringVarP(&pastePassphrase, "passphrase", "s", "", "Encryption Key")

	regStrP(pasteCmd, "passphrase", "s", "", "Secret")
}
