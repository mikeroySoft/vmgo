package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// explainCmd represents the explain command
var explainCmd = &cobra.Command{
	Use:   "explain RESOURCE",
	Short: "Prints more detail about a given resource.",
	Long: `Prints detail about a given resource.
Provides more detail than 'get' and less detail than 'describe'.
Verbosity can be set with the -v flag`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("explain called")
	},
}

func init() {
	rootCmd.AddCommand(explainCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// explainCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// explainCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
