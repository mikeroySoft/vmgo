package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// attachCmd represents the attach command
var attachCmd = &cobra.Command{
	Use:   "attach (POD | TYPE/NAME) -c CONTAINER",
	Short: "Attaches to the console of a running virtual machine or pod",
	Long: `Use 'attach' to gain direct access to the console of the virtual machine.
The function is similar to VMRC windows, but is done from this terminal, effectively behaving like an SSH session.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("attach called")
	},
}

func init() {
	rootCmd.AddCommand(attachCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// attachCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// attachCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
