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
	"log"

	"github.com/spf13/cobra"
)

var (
	getAnywherePath string
	getPassphrase   string
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:  "get",
	Long: `Get (encrypted) contents from /:path`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Loading contents for path", getAnywherePath, "with passphrase", getPassphrase)
		content, err := api.API().GetContent(getAnywherePath, getPassphrase)
		if err != nil {
			log.Fatalln("Error receiving contents:", err)
			return
		}

		log.Println("Received contents (", len(content), "bytes ):")
		fmt.Println(content)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.Flags().StringVarP(&getAnywherePath, "anywhere", "a", "", "Path (Anywhere)")
	getCmd.Flags().StringVarP(&getPassphrase, "passphrase", "s", "", "Encryption Key")
}
