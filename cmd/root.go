package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "vmgo",
	Short: "Using Cobra and Resty",
	Long: `
vmgo is a CLI for interacting with VMware Fusion and Workstation
  `,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },

}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	//fmt.Println("if cfgFile, viper.SetConfigFile(cfgFile)")
	if cfgFile != "" {
		// Use config file from the flag.
		fmt.Println(cfgFile)
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory
		//fmt.Println("Find home directory")
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".vmgo".
		//fmt.Println("Search config in home directory with name .vmgo")
		viper.AddConfigPath(home)
		viper.SetConfigName(".vmgo")
		// make sure we're reading yaml
		viper.SetConfigType("yaml")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in, write to it
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
		// viper.WriteConfig()

	}
}
