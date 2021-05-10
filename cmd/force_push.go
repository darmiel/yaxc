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
	"github.com/spf13/cobra"
	"strings"
)

var (
	fpAnywherePath string
	fpSecret       string
	fpBase64       bool
	fpHideSecret   bool
	fpHideURL      bool
)

// forcePushCmd represents the forcePush command
var forcePushCmd = &cobra.Command{
	Use:   "force-push",
	Short: "One-Time push to YAxC",
	Long:  `One-Time push to YAxC`,
	Run: func(cmd *cobra.Command, args []string) {
		cb, err := common.GetClipboard(fpBase64)
		if err != nil {
			fmt.Println(common.StyleWarn(), "Error retrieving contents:", err)
			return
		}

		if err := api.API().SetContent(fpAnywherePath, fpSecret, cb); err != nil {
			fmt.Println(common.StyleWarn(), "Error uploading contents:", err)
			return
		}

		fmt.Println(common.StyleInfo(), "Sent ->", common.Color("/"+fpAnywherePath, "A8CC8C"))
		if len(fpSecret) != 0 {
			var secret string
			if fpHideSecret {
				secret = strings.Repeat("*", len(fpSecret)) + " (hidden)"
			} else {
				secret = fpSecret
			}
			fmt.Println(common.StyleInfo(), "🔐", common.Color(secret, "A8CC8C"))
		}

		if !fpHideURL {
			fmt.Println(common.StyleDebug(), "URL:", api.API().UrlGetContents(fpAnywherePath, fpSecret))
		}
	},
}

func init() {
	rootCmd.AddCommand(forcePushCmd)

	forcePushCmd.PersistentFlags().StringVarP(&fpAnywherePath, "anywhere", "a", "", "Anywhere Path")
	forcePushCmd.PersistentFlags().StringVarP(&fpSecret, "secret", "s", "", "Encryption Key")

	forcePushCmd.PersistentFlags().BoolVarP(&fpBase64, "base64", "b", false, "Use Base64")
	forcePushCmd.PersistentFlags().BoolVarP(&fpHideSecret, "hide-secret", "S", false, "Hide Secret")
	forcePushCmd.PersistentFlags().BoolVarP(&fpHideURL, "hide-url", "U", false, "Hide URL")

	cobra.CheckErr(cobra.MarkFlagRequired(forcePushCmd.PersistentFlags(), "anywhere"))
}
