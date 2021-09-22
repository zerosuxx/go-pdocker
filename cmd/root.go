package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use: 	"pdocker",
	Short:	"PDocker - PRoot based docker client",
}

func Execute() {
	rootCmd.Version = "1.0.4"
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
