package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(tweetCmd)
}

var tweetCmd = &cobra.Command{
	Use:   "tweet",
	Short: "Tweet current tasks",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
