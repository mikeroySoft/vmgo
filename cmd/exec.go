package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "exec (VM | POD | TYPE/NAME) [-c CONTAINER] [flags] -- COMMAND [args...]",
	Short: "Execute a command in a VM or a Container",
	Long: `Executes a given command within a specified VM, POD or Container
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("exec called")
	},
}

func init() {
	rootCmd.AddCommand(execCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// execCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// execCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
