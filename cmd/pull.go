package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zerosuxx/go-pdocker/pkg"
	"os"
	"os/exec"
	"strings"
)

var pullCmd = &cobra.Command{
	Use:   "pull NAME[:TAG|@DIGEST]",
	Short: "Pull an image from a registry",
	Args:  cobra.MinimumNArgs(1),
	RunE:  pullHandler,
}

func pullHandler(_ *cobra.Command, args []string) error {
	if !pkg.IsFileExists(imagePullScriptPath) {
		return errors.New("image pull script not installed, run `install` command first")
	}

	imageName := GetImageName(args[0])
	imageFolder := GetImageFolder(imageName)

	frozenImagePath := frozenImagesPath + "/" + imageFolder
	cmd := exec.Command("bash", imagePullScriptPath, frozenImagePath, imageName)
	err := pkg.RunCommand(cmd, os.Stdin, OutputFilter(os.Stdout), OutputFilter(os.Stderr))
	if err != nil {
		return err
	}

	fmt.Println(imageName + " downloaded.")

	return nil
}

func init() {
	rootCmd.AddCommand(pullCmd)
}

type ManifestResponse struct {
	Layers []string
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

func GetImageName(imageName string) string {
	if !strings.Contains(imageName, ":") {
		imageName += ":latest"
	}
	return imageName
}

func GetImageFolder(imageName string) string {
	return strings.Replace(imageName, ":", "/", 1)
}
