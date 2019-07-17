package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.PersistentFlags().BoolP("help", "", false, "Help default flag")
}

var rootCmd = &cobra.Command{
	Use: "pushshift",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello world :)")
	},
}

// Execute executes the command line process
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
