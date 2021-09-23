package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"strings"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check system requirements",
	Run:   runDoctor,
}

func runDoctor(_ *cobra.Command, _ []string) {
	prootCommand := exec.Command("proot", "--help")
	prootOutput, prootError := prootCommand.Output()
	printRequirementResult("proot is not installed (See: https://proot-me.github.io)", prootOutput, prootError, func(output string) string {
		version := strings.TrimLeft(strings.Split(output, ":")[0], "proot ")
		return fmt.Sprintf("proot (%s) installed", version)
	})

	bashCommand := exec.Command("bash", "--version")
	bashOutput, bashError := bashCommand.Output()

	printRequirementResult("bash is not installed", bashOutput, bashError, func(output string) string {
		firstLine := strings.Split(output, "\n")[0]
		version := strings.Split(firstLine, " ")[3]
		return fmt.Sprintf("bash (%s) installed", version)
	})

	curlCommand := exec.Command("curl", "--version")
	curlOutput, curlError := curlCommand.Output()

	printRequirementResult("curl is not installed", curlOutput, curlError, func(output string) string {
		firstLine := strings.Split(output, "\n")[0]
		version := strings.Split(firstLine, " ")[1]
		return fmt.Sprintf("curl (%s) installed", version)
	})

	jqCommand := exec.Command("jq", "--version")
	jqOutput, jqError := jqCommand.Output()

	printRequirementResult("jq is not installed", jqOutput, jqError, func(output string) string {
		version := strings.Trim(output, "\n")
		return fmt.Sprintf("jq (%s) installed", version)
	})

	_, imagePullScriptError := os.Stat(imagePullScriptPath)
	printRequirementResult("image pull script is not installed", []byte{}, imagePullScriptError, func(_ string) string {
		return fmt.Sprintf("image pull script (%s) installed", imagePullScriptPath)
	})

}

func init() {
	rootCmd.AddCommand(doctorCmd)
}

type successTextTransformer func(string) string

func printRequirementResult(errorText string, output []byte, err error, transformer successTextTransformer) {
	if err != nil {
		fmt.Println("[ ] " + errorText)
	} else {
		successText := transformer(string(output))
		fmt.Println(fmt.Sprintf("[X] " + successText))
	}
}

func CheckImagePullScriptExists() bool {
	_, imagePullScriptError := os.Stat(imagePullScriptPath)

	return imagePullScriptError == nil
}
