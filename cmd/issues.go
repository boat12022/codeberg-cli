package cmd

import (
	"fmt"

	"codeberg.org/13thab/codeberg-cli/internal/codeberg"
	"github.com/spf13/cobra"
)

var (
	state string
	limit int
)

var issuesCmd = &cobra.Command{
	Use:   "issues <username> <repo>",
	Short: "Get issues for a Codeberg repository",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		client := codeberg.NewClient()

		username := args[0]
		repoName := args[1]

		issues, err := client.GetIssues(username, repoName, state, limit)
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, issue := range issues {
			fmt.Printf("#%d %s\n", issue.Number, issue.Title)
		}
	},
}

func init() {
	rootCmd.AddCommand(issuesCmd)

	issuesCmd.Flags().StringVarP(
		&state,
		"state",
		"s",
		"open",
		"Filter issues by state",
	)
	issuesCmd.Flags().IntVarP(
		&limit,
		"limit",
		"l",
		10,
		"Number of issues to display")
}
