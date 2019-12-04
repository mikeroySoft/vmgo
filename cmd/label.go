package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// labelCmd represents the label command
var labelCmd = &cobra.Command{
	Use:   "label [--overwrite] (-f FILENAME | TYPE NAME) KEY_1=VAL_1 ... KEY_N=VAL_N [--resource-version=version]",
	Short: "Update the label of a given resource",
	Long: `Update the labels on a resource.

 * A label key and value must begin with a letter or number, and may contain letters, numbers, hyphens, dots, and underscores, up to %[1]d characters each.
 * Optionally, the key can begin with a DNS subdomain prefix and a single '/', like example.com/my-app
 * If --overwrite is true, then existing labels can be overwritten, otherwise attempting to overwrite a label will result in an error.
 * If --resource-version is specified, then updates will use this resource version, otherwise the existing resource-version will be used.
 `,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("label called")
	},
}

func init() {
	rootCmd.AddCommand(labelCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// labelCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// labelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
