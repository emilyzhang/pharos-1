package cmd

import (
	"fmt"
	"os"

	"github.com/lob/pharos/pkg/pharos/cli"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// CurrentCmd is the pharos clusters current command.
var CurrentCmd = &cobra.Command{
	Use:   "current",
	Short: "Print current cluster",
	Long:  "Prints current cluster in the designated kubeconfig file.",
	RunE:  func(cmd *cobra.Command, args []string) error { return runCurrent(file) },
}

func runCurrent(kubeConfigFile string) error {
	clusterName, err := cli.CurrentCluster(kubeConfigFile)
	if err != nil {
		return errors.Wrap(err, "unable to retrieve cluster")
	}
	fmt.Println(clusterName)
	return nil
}

func init() {
	CurrentCmd.Flags().StringVarP(&file, "file", "f", os.Getenv("HOME")+"/.kube/config", "specify kubeconfig file (defaults to $HOME/.kube/config)")
}
