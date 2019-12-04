package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// cpCmd represents the cp command
var cpCmd = &cobra.Command{
	Use:     "cp",
	Aliases: []string{"copy"},
	Short:   "Copies the specified file from the current HOST into the specified path within the virtual machine or pod",
	Long: `cp is used to copy files from host to guest.
vmgo also is aware of containers running within podVMs, and can copy files directly into the running container's file system using the -c flag.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cp called")
	},
}

func init() {
	rootCmd.AddCommand(cpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
