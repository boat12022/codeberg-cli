package cmd

import (
	"fmt"

	"codeberg.org/13thab/codeberg-cli/internal/codeberg"
	"github.com/spf13/cobra"
)

var repoCmd = &cobra.Command{
	Use:   "repo <username> <repo>",
	Short: "Get information about a Codeberg repository",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		client := codeberg.NewClient()

		username := args[0]
		repoName := args[1]

		repo, err := client.GetRepo(username, repoName)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Repository: %s\n", repo.Name)
		fmt.Printf("Description: %s\n", repo.Description)
		fmt.Printf("Owner: %s\n", repo.Owner.Username)
		fmt.Printf("Language: %s\n", repo.Language)
		fmt.Printf("Stars: %d\n", repo.StarsCount)
		fmt.Printf("Forks: %d\n", repo.ForksCount)
		fmt.Printf("Watchers: %d\n", repo.WatchersCount)
		fmt.Printf("Created: %s\n", repo.CreatedAt)
		fmt.Printf("Updated: %s\n", repo.UpdatedAt)
		fmt.Printf("URL: %s\n", repo.HTMLURL)

	},
}

func init() {
	rootCmd.AddCommand(repoCmd)
}
