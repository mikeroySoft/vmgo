package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion SHELL",
	Short: "Installs vmgo shell completion commands",
	Long: `Completion will install the shell completion vocabulary in your currently running shell.

This will allow you to explore all of the vmgo commands without having to read the entire document first.
Simply begin a command, hit tab, and you will be presented the available options for that particular command or operation.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("completion called")
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
