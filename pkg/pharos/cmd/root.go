package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:     "pharos",
	Short:   "A tool for managing kubeconfig files.",
	Long:    `Pharos is a tool for cluster discovery and distribution of kubeconfig files.`,
	Version: "1.0",
}

// ClustersCmd is the pharos clusters command.
var ClustersCmd = &cobra.Command{Use: "clusters"}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().BoolP("version", "v", false, "print pharos version number")

	// Prevent usage message from being printed out upon command error.
	rootCmd.SilenceUsage = true

	// Add child commands.
	rootCmd.AddCommand(ClustersCmd)
	ClustersCmd.AddCommand(CurrentCmd)
}
