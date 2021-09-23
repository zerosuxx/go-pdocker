package pkg

import (
	"os"
	"os/exec"
	"strconv"
)

type Proot struct {
	command string
}

func (p *Proot) Run(rootFs string, workDir string, entryPoint string, userId int, groupId int) error {
	var prootArgs []string
	prootArgs = append(prootArgs, "--link2symlink")
	prootArgs = append(prootArgs, "--change-id="+strconv.Itoa(userId)+":"+strconv.Itoa(groupId))
	prootArgs = append(prootArgs, "--rootfs="+rootFs)
	prootArgs = append(prootArgs, "--pwd="+workDir)
	prootArgs = append(prootArgs, "--kill-on-exit")
	prootArgs = append(prootArgs, "--mount=/dev")
	prootArgs = append(prootArgs, "--mount=/proc")
	prootArgs = append(prootArgs, "--mount="+rootFs+"/root:/dev/shm")
	if IsFileExists("/sdcard") {
		prootArgs = append(prootArgs, "--mount=/sdcard")
	}
	prootArgs = append(prootArgs, "/usr/bin/env")
	prootArgs = append(prootArgs, "-i")
	prootArgs = append(prootArgs, "HOME="+workDir)
	prootArgs = append(prootArgs, "TERM="+os.Getenv("TERM"))
	prootArgs = append(prootArgs, "LANG=en_US.UTF-8")
	prootArgs = append(prootArgs, "LC_ALL=C")
	prootArgs = append(prootArgs, "LANGUAGE=en_US")
	prootArgs = append(prootArgs, "PATH=/bin:/usr/bin:/sbin:/usr/sbin")

	prootArgs = append(prootArgs, "/bin/sh")
	prootArgs = append(prootArgs, "--login")
	if entryPoint != "" {
		prootArgs = append(prootArgs, "-c")
		prootArgs = append(prootArgs, entryPoint)
	}

	_ = os.Setenv("LD_PRELOAD", "")
	prootCommand := exec.Command(
		"proot",
		prootArgs...,
	)
	err := RunCommand(prootCommand, os.Stdin, os.Stdout, os.Stderr)
	if err != nil {
		return err
	}

	return nil
}
