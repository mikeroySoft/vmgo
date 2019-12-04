package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var filename, directory string

// applyCmd represents the apply command
var applyCmd = &cobra.Command{
	Use:   "apply (-f FILENAME | -k DIRECTORY)",
	Short: "Apply a configuration to a resource by filename or stdin.",
	Long: `Apply a configuration to a resource by filename or stdin.

The resource name must be specified. 

This resource will be created if it doesn't exist yet. 

To use 'apply', always create the resource initially with either 'apply' or 'create --save-config'
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("applying configuration:")
		fmt.Println(filename, directory)

	},
}

func init() {
	rootCmd.AddCommand(applyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// applyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	applyCmd.Flags().StringVarP(&filename, "filename", "f", "", "specify a file to apply (YAML or JSON)")
	applyCmd.Flags().StringVarP(&directory, "directory", "k", "", "path to a directory containing .yaml or .json files to apply")
	applyCmd.MarkFlagRequired("filename")
}
