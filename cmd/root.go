package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "posixer",
	Short: `posixer command provides information about the Posix-Shell`,
}

func exitError(msg interface{}) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}

// Execute run posixer command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		exitError(err)
	}
}
