package cmd

import (
	"fmt"
	"os"
	"upm/ui"

	"github.com/spf13/cobra"
)

var (
	gitHash string
	version string

	rootCmd = &cobra.Command{
		Use: "UnityPackageManager",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return nil
			} else if len(args) == 1 {
				if _, err := os.Stat(args[0]); err == nil || os.IsExist(err) {
					return nil
				}
			}
			return fmt.Errorf("unknown command %q for %q", args[0], cmd.CommandPath())
		},
		Run: func(cmd *cobra.Command, args []string) {
			var filename string
			if len(args) > 0 {
				filename = args[0]
			}
			ui.Start(filename)
		},
	}
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version info of UnityPackageManager",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("UnityPackageManager version %s, build %s\n", version, gitHash)
		},
	}
)

func Execute() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.Execute()
}
