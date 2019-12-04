package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start [OPTIONS] CONTAINER | VMX [CONTAINER | VMX...]",
	Short: "Start a virtual machine or container image",
	Long: `vmgo start will launch a new virtual machine or container image in the background.
Use -a or --attach to attach STDOUT/STDERR and forward signals (Containers or Linux VMs only).
Can also be used to start the container runtime without running a container.
It used without options, it will prompt and ask if you would like to start the container runtime.

Options:
  -c,  							Start the VMware Container runtime
  -a, --attach                  Attach STDOUT/STDERR and forward signals
  -i, --interactive             Attach container's STDIN

`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`
😄  vmgo container runtime v0.0.1 on Darwin 10.14.6
🔥  Creating podVM (CPUs=4, Memory=8192MB, Disk=20000MB) ...
🐳  Preparing runtime ...
🚜  Pulling images ...
🚀  Launching container services ...
⌛  Waiting for: apiserver
🏄  Done! `)
		fmt.Println("")
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
