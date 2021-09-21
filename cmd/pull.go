package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"strings"
)

var pullCmd = &cobra.Command{
	Use:   	"pull NAME[:TAG|@DIGEST]",
	Short: 	"Pull an image from a registry",
	Args: 	cobra.MinimumNArgs(1),
	RunE: 	runPull,
}

func runPull(_ *cobra.Command, args []string) error {
	imageName := args[0]
	if !strings.Contains(imageName, ":") {
		imageName += ":latest"
	}

	frozenImagePath := frozenImagesPath + "/" + strings.Replace(imageName, ":", "/", 1)
	cmd := exec.Command("bash", imagePullScriptPath, frozenImagePath, imageName)
	err := runCommand(cmd)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	rootCmd.AddCommand(pullCmd)
}

func runCommand(cmd *exec.Cmd) error {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	_ = cmd.Wait()

	return nil
}
