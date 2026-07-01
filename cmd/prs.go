package cmd

import (
	"fmt"

	"codeberg.org/13thab/codeberg-cli/internal/codeberg"
	"github.com/spf13/cobra"
)

var prsCmd = &cobra.Command{
	Use:   "prs <username> <repo>",
	Short: "List pull requests for a repository",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		client := codeberg.NewClient()

		username := args[0]
		repo := args[1]

		state, err := cmd.Flags().GetString("state")
		if err != nil {
			fmt.Println(err)
			return
		}

		limit, err := cmd.Flags().GetInt("limit")
		if err != nil {
			fmt.Println(err)
			return
		}

		prs, err := client.GetPullRequests(username, repo, state, limit)
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, pr := range prs {
			fmt.Printf(
				"#%d [%s] %s\n",
				pr.Number,
				pr.State,
				pr.Title,
			)
		}
	},
}

func addCommonListFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("state", "s", "open", "Filter by state")
	cmd.Flags().IntP("limit", "l", 10, "Number of results")
}

func init() {
	rootCmd.AddCommand(prsCmd)
	addCommonListFlags(prsCmd)
}
