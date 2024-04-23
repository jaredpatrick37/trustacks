package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/trustacks/trustacks/internal"
)

var (
	configInitCmdFromPlan string
)

var configInitCmd = &cobra.Command{
	Use:   "config-inits",
	Short: "create a Configu schema from an action plan",
	Run: func(_ *cobra.Command, _ []string) {
		if err := internal.ConfigInitCmd(configInitCmdFromPlan); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	configInitCmd.Flags().StringVar(&configInitCmdFromPlan, "from-plan", "trustacks.plan", "path to the plan file")
	rootCmd.AddCommand(configInitCmd)
}
