package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run [OPTIONS] IMAGE [COMMAND] [ARG...]",
	Short: "Run commands within a virtual machine or container.",
	Long: `vmgo run is used to execute commands within a container or VM in a non-interactive way.
If the container is not running it will start it, with options to stop it after the command is complete.

Examples:
run 'ls' in a container using the 'my-image' image, and stop the running image when finished
 vmgo run --rm my-image ls

run 'ls' in the 'my-vmx' virtual machine type (-t vm) that is currently running
 vmgo run -t vm my-vmx ls

 If no commands are given it will start the vm or container in the background.
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run called")
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
