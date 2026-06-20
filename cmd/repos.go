package cmd

import (
	"fmt"

	"codeberg.org/13thab/codeberg-cli/internal/codeberg"
	"github.com/spf13/cobra"
)

var reposCmd = &cobra.Command{
	Use:   "repos <username>",
	Short: "List repositories of a user",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client := codeberg.NewClient()

		username := args[0]
		repos, err := client.GetRepos(username)
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, repo := range repos {
			fmt.Printf(
				"%s (%s) ⭐ %d\n",
				repo.Name,
				repo.Language,
				repo.StarsCount,
			)
		}
	},
}

func init() {
	rootCmd.AddCommand(reposCmd)
}
