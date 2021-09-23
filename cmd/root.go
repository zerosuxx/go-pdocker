package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "pdocker",
	Short: "PDocker - PRoot based docker client",
}

var Version = "development"

func Execute() {
	rootCmd.Version = Version
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
