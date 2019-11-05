package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "snail-bot",
	Short: "Snail Bot is Assignment management bot for \"denden\"s",
	Long: `An assignment management Twtiter bot in Go.
                Source Code is available at https://github.com/dora1998/snail-bot`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
