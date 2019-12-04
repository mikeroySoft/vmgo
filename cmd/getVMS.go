package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// getVMCmd represents the getVM command
var getVMSCmd = &cobra.Command{
	Use:     "vms (VMX PATH) [flags]",
	Aliases: []string{"vm"},
	Short:   "Prints a list of all VMs in the current host inventory.",
	Long: `'vms' is a resource type that can be retrieved with 'get'.
Specifying a name or path to .vmx file will print details about that vm object.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("getVMS called")
	},
}

func init() {
	getCmd.AddCommand(getVMSCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getVMCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getVMCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
