package cmd

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zerosuxx/go-pdocker/pkg"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var runCmd = &cobra.Command{
	Use:   "run IMAGE [COMMAND] [ARG...]",
	Short: "Run a command in a new container",
	Args:  cobra.MinimumNArgs(1),
	RunE:  runHandler,
}

func runHandler(c *cobra.Command, args []string) error {
	imageName := GetImageName(args[0])
	entryPoint := ""
	if len(args) > 1 {
		entryPoint = args[1]
	}
	imageFolder := GetImageFolder(imageName)
	imagePath := filepath.Join(frozenImagesPath, imageFolder)
	containerId := getContainerId(imageName)
	containerFsPath := filepath.Join(containersPath, containerId, "fs")
	manifestFilePath := filepath.Join(imagePath, "manifest.json")

	if !pkg.IsFileExists(manifestFilePath) {
		err := pullHandler(c, args)
		if err != nil {
			return err
		}
	}

	fmt.Println("creating container...")
	layers, err := getLayersFromManifest(manifestFilePath)
	if err != nil {
		return err
	}

	err = extractLayers(layers, imagePath, containerFsPath)
	if err != nil {
		return err
	}

	fmt.Println("attaching to " + containerId + " container...")
	proot := pkg.Proot{}
	err = proot.Run(containerFsPath, "/root", entryPoint, 0, 0)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func getContainerId(imageName string) string {
	imageHash := imageName + "-" + strconv.FormatInt(time.Now().Unix(), 10)
	hash := sha1.New()
	hash.Write([]byte(imageHash))

	return fmt.Sprintf("%x", hash.Sum(nil))
}

func extractLayers(layers []string, imagePath string, containerFsPath string) error {
	for _, layerFilePath := range layers {
		file, err := os.Open(filepath.Join(imagePath, layerFilePath))
		if err != nil {
			return err
		}

		err = pkg.ExtractTarArchive(file, containerFsPath)
		if err != nil {
			return err
		}
	}

	return nil
}

func getLayersFromManifest(manifestFile string) ([]string, error) {
	manifestJson, err := os.ReadFile(manifestFile)
	if err != nil {
		return nil, err
	}
	var manifestList []ManifestResponse
	if err = json.Unmarshal(manifestJson, &manifestList); err != nil {
		return nil, err
	}

	return manifestList[0].Layers, nil
}
