package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "poolcog",
	Short: "Poolcog it's a cognito user pool export CLI",
	Long:  `A very useful way to export your users pool from AWS cognito user pool service.`,
	Run: func(cmd *cobra.Command, arg []string) {
		fmt.Println("TEST command")
	},
}

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
