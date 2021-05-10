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
	"strings"

	"github.com/spf13/cobra"
)

// pushCmd represents the push command
var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Force Push Clipboard From Server",
	Long:  "Force Push Clipboard To Server",
	Run: func(cmd *cobra.Command, args []string) {
		cb, err := common.GetClipboard(fpBase64)
		if err != nil {
			fmt.Println(common.StyleWarn(), "Error retrieving contents from clipboard:", err)
			return
		}
		if err := api.API().SetContent(fpAnywherePath, fpSecret, cb); err != nil {
			fmt.Println(common.StyleWarn(), "Error uploading contents:", err)
			return
		}

		fmt.Println(common.StyleInfo(), "Sent ->",
			common.Color(common.PrettyLimit(cb, 32), "66C2CD"), "->",
			common.Color("/"+fpAnywherePath, "A8CC8C"))

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
	forceCmd.AddCommand(pushCmd)
}
