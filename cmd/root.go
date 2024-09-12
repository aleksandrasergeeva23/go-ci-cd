package cmd

import (
	"fmt"
	"os"

	"github.com/aleksandrasergeeva23/go-ci-cd/pkg/pipeline"
	"github.com/spf13/cobra"
)

var (
	configFile string // Stores the path to the config file
)

// Define the root command
var rootCmd = &cobra.Command{
	Use:   "ci-cd",
	Short: "A simple Go-based CI/CD pipeline",
	Run: func(cmd *cobra.Command, args []string) {
		// Load the pipeline configuration
		pipe, err := pipeline.LoadPipelineConfig(configFile)
		if err != nil {
			fmt.Printf("Error loading pipeline config: %v\n", err)
			return
		}

		// Run the pipeline
		fmt.Println("Starting the pipeline...")
		pipe.Run()
		fmt.Println("Pipeline execution completed.")
	},
}

// Initialize the root command and setup flags
func init() {
	// Add a flag to allow users to specify the configuration file
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "config/pipeline.yaml", "Pipeline config file")
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
