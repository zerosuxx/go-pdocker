package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"strings"
)

var doctorCmd = &cobra.Command{
	Use:   	"doctor",
	Short: 	"Check system requirements",
	Run: 	runDoctor,
}

func runDoctor(_ *cobra.Command, _ []string) {
	prootCommand := exec.Command("proot", "--help")
	prootOutput, prootError := prootCommand.Output()

	if prootError != nil {
		fmt.Println("[ ] PRoot is not installed (See: https://proot-me.github.io)")
	} else {
		version := strings.TrimLeft(strings.Split(string(prootOutput), ":")[0], "proot ")
		fmt.Println(fmt.Sprintf("[X] PRoot (%s) installed", version))
	}

	bashCommand := exec.Command("bash", "--version")
	bashOutput, bashError := bashCommand.Output()

	if bashError != nil {
		fmt.Println("[ ] Bash is not installed")
	} else {
		firstLine := strings.Split(string(bashOutput), "\n")[0]
		version := strings.Split(firstLine, "version ")[1]
		fmt.Println(fmt.Sprintf("[X] Bash (%s) installed", version))
	}

	if _, err := os.Stat(imagePullScriptPath); err != nil {
		fmt.Println("[ ] Pull script is not installed")
	} else {
		fmt.Println(fmt.Sprintf("[X] Pull script (%s) installed", imagePullScriptPath))
	}
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}