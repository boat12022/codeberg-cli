package cmd

import (
	"fmt"

	"codeberg.org/13thab/codeberg-cli/internal/codeberg"
	"github.com/spf13/cobra"
)

var userCmd = &cobra.Command{
	Use:   "user <username>",
	Short: "Get information about a Codeberg user",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client := codeberg.NewClient()

		username := args[0]
		user, err := client.GetUser(username)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Username: %s\n", user.Username)
		fmt.Printf("Starred Repos: %d\n", user.StarredReposCount)
		fmt.Printf("Followers: %d\n", user.FollowersCount)
		fmt.Printf("Following: %d\n", user.FollowingCount)
		fmt.Printf("Created: %s\n", user.Created)
	},
}

func init() {
	rootCmd.AddCommand(userCmd)
}
