// +build dev

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
	"github.com/darmiel/yaxc/internal/whitelist"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

// jwtgenCmd represents the jwtgen command
var jwtgenCmd = &cobra.Command{
	Use:   "jwtgen",
	Short: "Generate JWT Token",
	Run: func(cmd *cobra.Command, args []string) {
		secret := viper.GetString("secret")
		maxBody := viper.GetInt64("max-body")
		audience := viper.GetString("audience")
		issuer := viper.GetString("issuer")
		count := viper.GetInt("count")

		fmt.Println("🔨 Generating", count, "JWT-Tokens ...")
		for i := 0; i < count; i++ {
			signed, err := whitelist.GenerateToken(secret,
				audience,
				issuer,
				maxBody)
			if err != nil {
				log.Fatalln("Error signing:", err)
				return
			}
			fmt.Println(signed)
		}
		fmt.Println("🤗 Done!")
	},
}

func init() {
	rootCmd.AddCommand(jwtgenCmd)
	regStrP(jwtgenCmd, "secret", "s", "", "Secret")
	regStrP(jwtgenCmd, "audience", "a", "", "Audience")
	regStrP(jwtgenCmd, "Issuer", "i", "", "Issuer")
	regInt64P(jwtgenCmd, "max-body", "x", 1_000_000, "Max-Body")
	regIntP(jwtgenCmd, "count", "c", 1, "Token Count")
}
