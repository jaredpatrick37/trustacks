package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/trustacks/trustacks/internal"
	"github.com/trustacks/trustacks/pkg/engine"
)

var (
	runCmdSource              string
	runCmdStages              []string
	runCmdIgnoreMissingInputs bool
	runCmdPrerelease          bool
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run an action plan",
	Run: func(_ *cobra.Command, args []string) {
		planFile := "trustacks.plan"
		if len(args) > 0 {
			planFile = args[0]
		}
		if err := internal.RunCmd(&internal.RunCmdOptions{
			Source:              runCmdSource,
			Plan:                planFile,
			Stages:              runCmdStages,
			IgnoreMissingInputs: runCmdIgnoreMissingInputs,
			Prerelease:          runCmdPrerelease,
		}); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	runCmd.Flags().StringVar(&runCmdSource, "source", "./", "application source path")
	runCmd.Flags().StringSliceVar(
		&runCmdStages,
		"stages",
		[]string{
			engine.GetStage(engine.CommitStage),
			engine.GetStage(engine.AcceptanceStage),
			engine.GetStage(engine.NonFunctionalStage),
			engine.GetStage(engine.DeployStage),
		},
		"activity phases to run during the action plan",
	)
	runCmd.Flags().BoolVar(&runCmdIgnoreMissingInputs, "ignore-missing-inputs", false, "ignore missing action inputs")
	runCmd.Flags().BoolVar(&runCmdPrerelease, "prerelease", false, "run the pipeline without the production release actions")
	rootCmd.AddCommand(runCmd)
}
