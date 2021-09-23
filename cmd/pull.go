package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"strings"
)

var pullCmd = &cobra.Command{
	Use:   "pull NAME[:TAG|@DIGEST]",
	Short: "Pull an image from a registry",
	Args:  cobra.MinimumNArgs(1),
	RunE:  runPull,
}

func runPull(_ *cobra.Command, args []string) error {
	if !CheckImagePullScriptExists() {
		return errors.New("image pull script not installed, run `install` command first")
	}

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
	cmd.Stdout = OutputFilter(os.Stdout)
	cmd.Stderr = OutputFilter(os.Stderr)
	err := cmd.Run()
	if err != nil {
		return err
	}
	_ = cmd.Wait()

	return nil
}

type ProxyWriter struct {
	file *os.File
}

func OutputFilter(file *os.File) *ProxyWriter {
	return &ProxyWriter{
		file: file,
	}
}

func (w *ProxyWriter) Write(p []byte) (int, error) {
	if strings.Contains(string(p), "type: go: not found") ||
		strings.Contains(string(p), "type: dpkg: not found") ||
		strings.Contains(string(p), "Download of images") ||
		strings.Contains(string(p), "Docker daemon:") ||
		strings.Contains(string(p), "docker load") {

		return len(p), nil
	}

	return w.file.Write(p)
}
