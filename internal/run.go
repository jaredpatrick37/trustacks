package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"dagger.io/dagger"
	"github.com/charmbracelet/log"
	"github.com/trustacks/trustacks/pkg/engine"
)

type RunCmdOptions struct {
	Source              string
	Plan                string
	Stages              []string
	IgnoreMissingInputs bool
	Prerelease          bool
}

func removeReleaseStage(stages []string) []string {
	stageIndex := -1
	for i, stage := range stages {
		if stage == engine.GetStage(engine.ReleaseStage) {
			stageIndex = i
		}
	}
	if stageIndex >= 0 {
		stages = append(stages[:stageIndex], stages[stageIndex+1:]...)
	}
	return stages
}

func RunCmd(options *RunCmdOptions) error {
	var planData map[string]interface{}
	if _, err := os.Stat(options.Plan); os.IsNotExist(err) {
		actionPlan, err := engine.New().CreateActionPlan(options.Source)
		if err != nil {
			return fmt.Errorf("failed creating the action plan: %s", err)
		}
		spec, err := actionPlan.ToJSON()
		if err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(spec), &planData); err != nil {
			return fmt.Errorf("failed unmarshaling the action plan data: %s", err)
		}
	} else {
		planJSON, err := os.ReadFile(options.Plan)
		if err != nil {
			return fmt.Errorf("failed opening plan file: %s", err)
		}
		if err := json.Unmarshal(planJSON, &planData); err != nil {
			return fmt.Errorf("failed parsing plan file: %s", err)
		}
	}
	spec, err := json.Marshal(planData)
	if err != nil {
		return fmt.Errorf("failed converting plan file to spec: %s", err)
	}
	clientOpts := []dagger.ClientOpt{}
	if os.Getenv("VERBOSE") != "" {
		clientOpts = append(clientOpts, dagger.WithLogOutput(os.Stdout))
	}
	client, err := dagger.Connect(context.Background(), clientOpts...)
	if err != nil {
		return fmt.Errorf("failed connecting to the dagger agent")
	}
	if options.Prerelease {
		removeReleaseStage(options.Stages)
	}
	if err := engine.Run(engine.RunArgs{
		Source:              options.Source,
		Spec:                string(spec),
		Client:              client,
		Stages:              options.Stages,
		IgnoreMissingInputs: options.IgnoreMissingInputs,
	}); err != nil {
		log.Error("", "err", err)
		os.Exit(1)
	}
	return nil
}
