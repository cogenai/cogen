package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/cogenai/cogen/cogen"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "cogen",
		Short: "cogen is an autocoding command line system",
		Long: `cogen is a CLI application that helps you manage your project definitions,
plans, and refinements.`,
	}

	initCmd := &cobra.Command{
		Use:   "init [project_name]",
		Short: "Initialize a new cogen project",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			projectName := args[0]
			err := cogen.Init(projectName)
			if err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}
		},
	}

	defineCmd := &cobra.Command{
		Use:   "define",
		Short: "Define the project configuration",
		Run: func(cmd *cobra.Command, args []string) {
			err := cogen.Define()
			if err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}
		},
	}

	planCmd := &cobra.Command{
		Use:   "plan",
		Short: "Create a project plan",
		Run: func(cmd *cobra.Command, args []string) {
			err := cogen.Plan()
			if err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}
		},
	}

	refineCmd := &cobra.Command{
		Use:   "refine",
		Short: "Refine the project",
		Run: func(cmd *cobra.Command, args []string) {
			err := cogen.Refine()
			if err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}
		},
	}

	rootCmd.AddCommand(initCmd, defineCmd, planCmd, refineCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
