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
	"github.com/spf13/cobra"
)

var (
	fpAnywherePath string
	fpSecret       string
	fpBase64       bool
	fpHideSecret   bool
	fpHideURL      bool
)

// forceCmd represents the forcePush command
var forceCmd = &cobra.Command{
	Use:   "force",
	Short: "One-Time push/pull to YAxC",
	Long:  `One-Time push/pull to YAxC`,
}

func init() {
	rootCmd.AddCommand(forceCmd)

	forceCmd.PersistentFlags().StringVarP(&fpAnywherePath, "anywhere", "a", "", "Anywhere Path")
	forceCmd.PersistentFlags().StringVarP(&fpSecret, "secret", "s", "", "Encryption Key")

	forceCmd.PersistentFlags().BoolVarP(&fpBase64, "base64", "b", false, "Use Base64")
	forceCmd.PersistentFlags().BoolVarP(&fpHideSecret, "hide-secret", "S", false, "Hide Secret")
	forceCmd.PersistentFlags().BoolVarP(&fpHideURL, "hide-url", "U", false, "Hide URL")

	cobra.CheckErr(cobra.MarkFlagRequired(forceCmd.PersistentFlags(), "anywhere"))
}
