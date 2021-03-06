// +build server

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
	"github.com/darmiel/yaxc/internal/common"
	"github.com/darmiel/yaxc/internal/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
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
		maxBodyLen := viper.GetInt64("max-body-length")

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
		redisPrefixVal := viper.GetString("redis-prefix-value")
		redisPrefixHsh := viper.GetString("redis-prefix-hash")
		if redisAddr == "" {
			log.Println("WARN: Not using redis")
		}

		// other
		enableEnc := viper.GetBool("enable-encryption")
		proxyHeader := viper.GetString("proxy-header")
		jwt := viper.GetString("jwt")

		// create server & start
		s := server.NewServer(&server.YAxCConfig{
			BindAddress: bind,
			// Redis
			RedisAddress:   redisAddr,
			RedisPassword:  redisPass,
			RedisDatabase:  redisDB,
			RedisPrefixVal: redisPrefixVal,
			RedisPrefixHsh: redisPrefixHsh,
			// TTL
			DefaultTTL:    defTTL,
			MinTTL:        minTTL,
			MaxTTL:        maxTTL,
			MaxBodyLength: maxBodyLen,
			// Other
			EnableEncryption: enableEnc,
			ProxyHeader:      proxyHeader,
			JWTSign:          []byte(jwt),
		})
		go s.Start()

		fmt.Println(common.StyleInfo(), "Started clipboard-server. Press CTRL-C to stop.")
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
		<-sc

		// shutdown gracefully
		fmt.Println(common.StyleInfo(), "Stopping clipboard-server.")
		if s.App != nil {
			if err := s.App.Shutdown(); err != nil {
				fmt.Println(common.StyleWarn(), "Error shutting down:", err)
			} else {
				fmt.Println(common.StyleInfo(), "OK: Shut down")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	regStrP(serveCmd, "bind", "b", ":1332", "Bind-Address")

	// redis
	regStrP(serveCmd, "redis-addr", "r", "", "Redis Address")
	regStr(serveCmd, "redis-pass", "", "Redis Password")
	regInt(serveCmd, "redis-db", 0, "Redis Database")
	regStr(serveCmd, "redis-prefix-value", "yaxc::val::", "Redis Prefix (Value)")
	regStr(serveCmd, "redis-prefix-hash", "yaxc::hash::", "Redis Prefix (Hash)")

	// ttl
	regDurP(serveCmd, "default-ttl", "t", 60*time.Second, "Default TTL")
	regDurP(serveCmd, "min-ttl", "l", 5*time.Second, "Min TTL")
	regDurP(serveCmd, "max-ttl", "s", 60*time.Minute, "Max TTL")

	// other
	regInt64P(serveCmd, "max-body-length", "x", 8192, "Max Body Length")
	regBoolP(serveCmd, "enable-encryption", "e", true, "Enable Encryption")
	regStr(serveCmd, "proxy-header", "", "Proxy Header")
	regStr(serveCmd, "jwt", "", "JWT-Token")
}
