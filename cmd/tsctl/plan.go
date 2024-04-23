package main

import (
	"github.com/spf13/cobra"
	"github.com/trustacks/trustacks/internal"
)

var (
	planCmdName   string
	planCmdSource string
	planCmdForce  bool
)

var planCmd = &cobra.Command{
	Use:   "plan",
	Short: "Generate an action plan",
	Run: func(_ *cobra.Command, _ []string) {
		if err := internal.PlanCmd(planCmdSource, planCmdName, planCmdForce); err != nil {
			commandFailure(err)
		}
	},
}

func init() {
	planCmd.Flags().StringVar(&planCmdName, "name", "trustacks.plan", "plan name")
	planCmd.Flags().StringVar(&planCmdSource, "source", "./", "path to the application source")
	planCmd.Flags().BoolVar(&planCmdForce, "force", false, "overwrite the existing plan file")
	rootCmd.AddCommand(planCmd)
}
