/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	// "errors" 
	"fmt"
	"os"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	cfgFile     string
	userLicense string
	RootCmd = rootCmd
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cobra",
	Short: "A brief description of your application",
	Long:  `A longer description of your application`,
	// Validation of positional argument
	// Args: cobra.MinimumNArgs(1),
	Args: func(cmd *cobra.Command, args []string) error { // custom validation
		// if len(args) < 1 {
		// 	return errors.New("requires a color argument")
		// }
		return nil
	},

	// PersistentPreRun 所有子命令在run前都要执行此func
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("RootCmd PersistentPreRun with args: %v\n", args)
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("RootCmd PreRun with args: %v\n", args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("RootCmd Run with args: %v\n", args)
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("RootCmd PostRun with args: %v\n", args)
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("RootCmd PersistentPostRun with args: %v\n", args)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
	// cobra.CheckErr() 等同于如下
	// if err := rootCmd.Execute(); err != nil {
	//   fmt.Fprintln(os.Stderr, err)
	//   os.Exit(1)
	// }
}

func init() {
	cobra.OnInitialize(initConfig)

	// PersistentFlags 在当前cmd及其下所有子cmd均可用该flag， root下定义的flag则为全局flag
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	// mark config flag为必须的
	// rootCmd.MarkPersistentFlagRequired("config")

	rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "author name for copyright attribution")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")

	// Flags 局部flag，仅对设置该flag的cmd有效
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// toggle标记为必须的
	// rootCmd.MarkFlagRequired("toggle")

	// 通过viper设置flag
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	viper.SetDefault("license", "apache")
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

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".cobra")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
