package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// pushCmd represents the push command
var pushCmd = &cobra.Command{
	Use:   "push [OPTIONS] NAME[:TAG]",
	Short: "Uploads a resource to the specified remote repository.",
	Long: `'vmgo push' can upload pre-built container images and virtual machine template to remote repositories.
Defaults imply a container image and Docker Hub as the repository, which can be changed using the '-r' flag.
Virtual machines can be specified using the --vm flag, however a repository is required.

`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("push called")
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pushCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pushCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
