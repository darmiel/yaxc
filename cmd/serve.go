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
	"github.com/darmiel/yaxc/internal/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"time"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:  "serve",
	Long: `Run the YAxC server`,
	Run: func(cmd *cobra.Command, args []string) {
		// load values
		bind := viper.GetString("bind")
		defTTL := viper.GetDuration("default-ttl")
		minTTL := viper.GetDuration("min-ttl")
		maxTTL := viper.GetDuration("max-ttl")
		maxBodyLen := viper.GetInt("max-body-length")

		// validate values
		if bind == "" {
			log.Fatalln("ERROR: Empty bind address")
			return
		}

		if minTTL > maxTTL {
			log.Fatalln("MinTTL cannot be greater than MaxTTL")
			return
		}
		if minTTL > defTTL || maxTTL < defTTL {
			log.Fatalln("DefaultTTL out of range:", minTTL, "<=", defTTL, "<=", maxTTL)
			return
		}

		if maxBodyLen == 0 {
			log.Println("WARN: Infinite body length")
		}

		// redis
		redisAddr := viper.GetString("redis-addr")
		redisPass := viper.GetString("redis-pass")
		redisDB := viper.GetInt("redis-db")
		redisPrefix := viper.GetString("redis-prefix")
		if redisAddr == "" {
			log.Println("WARN: Not using redis")
		}

		// other
		enableEnc := viper.GetBool("enable-encryption")
		proxyHeader := viper.GetString("proxy-header")

		// create server & start
		s := server.NewServer(&server.YAxCConfig{
			BindAddress: bind,
			// Redis
			RedisAddress:  redisAddr,
			RedisPassword: redisPass,
			RedisDatabase: redisDB,
			RedisPrefix:   redisPrefix,
			// TTL
			DefaultTTL:    defTTL,
			MinTTL:        minTTL,
			MaxTTL:        maxTTL,
			MaxBodyLength: maxBodyLen,
			// Other
			EnableEncryption: enableEnc,
			ProxyHeader:      proxyHeader,
		})
		s.Start()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	regStrP(serveCmd, "bind", "b", ":1332", "Bind-Address")

	// redis
	regStrP(serveCmd, "redis-addr", "r", "", "Redis Address")
	regStr(serveCmd, "redis-pass", "", "Redis Password")
	regInt(serveCmd, "redis-db", 0, "Redis Database")
	regStr(serveCmd, "redis-prefix", "yaxc::", "Redis Prefix")

	// ttl
	regDurP(serveCmd, "default-ttl", "t", 60*time.Second, "Default TTL")
	regDurP(serveCmd, "min-ttl", "l", 5*time.Second, "Min TTL")
	regDurP(serveCmd, "max-ttl", "s", 5*time.Minute, "Max TTL")

	// other
	regIntP(serveCmd, "max-body-length", "x", 1024, "Max Body Length")
	regBoolP(serveCmd, "enable-encryption", "e", true, "Enable Encryption")
	regStr(serveCmd, "proxy-header", "", "Proxy Header")
}

func regStrP(cmd *cobra.Command, name, shorthand, def, usage string) {
	cmd.PersistentFlags().StringP(name, shorthand, def, usage)
	cobra.CheckErr(viper.BindPFlag(name, cmd.PersistentFlags().Lookup(name)))
}
func regStr(cmd *cobra.Command, name, def, usage string) {
	cmd.PersistentFlags().String(name, def, usage)
	cobra.CheckErr(viper.BindPFlag(name, cmd.PersistentFlags().Lookup(name)))
}
func regDurP(cmd *cobra.Command, name, shorthand string, def time.Duration, usage string) {
	cmd.PersistentFlags().DurationP(name, shorthand, def, usage)
	cobra.CheckErr(viper.BindPFlag(name, cmd.PersistentFlags().Lookup(name)))
}
func regIntP(cmd *cobra.Command, name, shorthand string, def int, usage string) {
	cmd.PersistentFlags().IntP(name, shorthand, def, usage)
	cobra.CheckErr(viper.BindPFlag(name, cmd.PersistentFlags().Lookup(name)))
}
func regInt(cmd *cobra.Command, name string, def int, usage string) {
	cmd.PersistentFlags().Int(name, def, usage)
	cobra.CheckErr(viper.BindPFlag(name, cmd.PersistentFlags().Lookup(name)))
}
func regBoolP(cmd *cobra.Command, name, shorthand string, def bool, usage string) {
	cmd.PersistentFlags().BoolP(name, shorthand, def, usage)
	cobra.CheckErr(viper.BindPFlag(name, cmd.PersistentFlags().Lookup(name)))
}
