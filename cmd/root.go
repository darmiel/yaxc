/*
Copyright © 2021 darmiel

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
	"github.com/spf13/cobra"
	"os"
	"time"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	cfgFile string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "yaxc",
	Short:   "Yet Another Cross Clipboard",
	Version: "1.1.0",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.yaxc.yaml)")
	regStr(rootCmd, "server", "https://yaxc.d2a.io", "URL of API-Server")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".yaxc" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".yaxc")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

// misc

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
