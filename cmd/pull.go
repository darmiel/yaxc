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
	"github.com/darmiel/yaxc/internal/api"
	"github.com/darmiel/yaxc/internal/common"
	"github.com/spf13/cobra"
)

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Force Pull Clipboard From Server",
	Long:  "Force Pull Clipboard From Server",
	Run: func(cmd *cobra.Command, args []string) {
		body, err := api.API().GetContent(fpAnywherePath, fpSecret)
		if err != nil {
			fmt.Println(common.StyleWarn(), "Error retrieving contents from server:", err)
			return
		}
		if err := common.WriteClipboard(body, fpBase64); err != nil {
			fmt.Println(common.StyleWarn(), "Error writing clipboard:", err)
			return
		}
		fmt.Println(common.StyleInfo(), "Read <-", common.Color(common.PrettyLimit(body, 32), "66C2CD"))
	},
}

func init() {
	forceCmd.AddCommand(pullCmd)
}
