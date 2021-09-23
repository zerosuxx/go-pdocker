package pkg

import (
	"io"
	"os/exec"
)

func RunCommand(cmd *exec.Cmd, stdin io.Reader, stdout io.Writer, stderr io.Writer) error {
	cmd.Stdin = stdin
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	_ = cmd.Wait()

	return nil
}
