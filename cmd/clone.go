package cmd

import (
	"github.com/spf13/cobra"
)

// cloneCmd represents the clone command
var cloneCmd = &cobra.Command{
	Use:   "clone <username> <repo>",
	Short: "Clone a Codeberg repository",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(cloneCmd)
}
