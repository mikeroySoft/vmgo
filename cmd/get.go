package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get [(-o|--output=)json|yaml|wide|custom-columns=...|jsonpath=...|jsonpath-file=...] [flags]",
	Short: "Prints some detail about one or more specified resources.",
	Long: `Display one or many resources
Prints a table of the most important information about the specified resources.

If the desired resource type is namespaced you will only see results in your current
namespace unless you pass --all-namespaces.

Uninitialized objects are not shown unless --include-uninitialized is passed.
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get called")
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
