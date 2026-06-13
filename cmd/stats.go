package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var statsCmd = &cobra.Command{
	Use:   "stats <username>",
	Short: "Show Codeberg user statistics",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]

		fmt.Printf("Showing stats for %s\n", username)
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)
}
