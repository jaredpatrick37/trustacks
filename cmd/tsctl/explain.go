package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/trustacks/trustacks/internal"
)

var explainCmd = &cobra.Command{
	Use:   "explain",
	Short: "Explain an action plan",
	Run: func(_ *cobra.Command, args []string) {
		var path string
		if len(args) > 0 {
			path = args[0]
		}
		if err := internal.ExplainCmd(path); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(explainCmd)
}
