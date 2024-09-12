package pipeline

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"time"

	"gopkg.in/yaml.v2"
)

// Stage represents each stage in the CI/CD pipeline
type Stage struct {
	Name   string `yaml:"name"`   // Name of the pipeline stage (e.g., Build, Test, Deploy)
	Script string `yaml:"script"` // Script to be executed for this stage
}

// Pipeline represents the entire pipeline with multiple stages
type Pipeline struct {
	Stages []Stage `yaml:"stages"` // Collection of stages that make up the pipeline
}

// LoadPipelineConfig loads the pipeline configuration from a YAML file
func LoadPipelineConfig(file string) (*Pipeline, error) {
	// Read the YAML file
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	// Unmarshal the YAML into the Pipeline struct
	var pipeline Pipeline
	err = yaml.Unmarshal(data, &pipeline)
	if err != nil {
		return nil, err
	}

	return &pipeline, nil
}

// Run executes each stage in the pipeline sequentially
func (p *Pipeline) Run() {
	for _, stage := range p.Stages {
		fmt.Printf("Running stage: %s\n", stage.Name)
		start := time.Now()

		// Execute the script for this stage
		cmd := exec.Command("/bin/sh", stage.Script)
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Error running stage %s: %v\n", stage.Name, err)
			continue
		}

		// Print output from the script execution
		fmt.Printf("Output:\n%s\n", string(out))
		fmt.Printf("Stage %s completed in %v\n", stage.Name, time.Since(start))
	}
}
