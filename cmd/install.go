package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zerosuxx/go-pdocker/pkg"
	"os"
)

var containersPath = GetAppPath() + "/containers"
var frozenImagesPath = GetAppPath() + "/frozen-images"
var imagePullScriptPath = GetAppPath() + "/download-frozen-image-v2.sh"
var frozenImageDownloaderUrl = "https://raw.githubusercontent.com/zerosuxx/moby/master/contrib/download-frozen-image-v2.sh"

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install application files",
	RunE:  installHandler,
}

func installHandler(_ *cobra.Command, _ []string) error {
	err := os.MkdirAll(GetAppPath(), os.ModePerm)
	if err != nil {
		return err
	}

	err = os.MkdirAll(frozenImagesPath, os.ModePerm)
	if err != nil {
		return err
	}

	err = pkg.DownloadFile(frozenImageDownloaderUrl, imagePullScriptPath)
	if err != nil {
		return err
	}

	fmt.Println("installation completed.")

	return nil
}

func init() {
	rootCmd.AddCommand(installCmd)
}

func GetAppPath() string {
	return os.Getenv("PREFIX") + "/var/lib/pdocker"
}
