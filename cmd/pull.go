package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull [OPTIONS] NAME[:TAG|@DIGEST]",
	Short: "Downloads an image (container or VM) from a remote repository.",
	Long: `Pull can be used to get a varity of resources from remote repositories such as Docker Hub or Harbor.

Options:
 -a, --all-tags                Download all tagged images in the repository
 --disable-content-trust   	   Skip image verification (default true)
 --platform string             Set platform if server is multi-platform capable
 -q, --quiet                   Suppress verbose output
 --vm 						   Specify the VM object type
 -r, --repository			   Specify the remote repository (required when using the '--vm' flag)


 Resource types:
 vm
 image
 
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pull called")
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pullCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pullCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
