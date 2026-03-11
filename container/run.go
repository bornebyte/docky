package container

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func Run(command []string) {
	fmt.Println("Running container : ", command)
	cmd := exec.Command(command[0], command[1:]...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS |
			syscall.CLONE_NEWPID |
			syscall.CLONE_NEWNS,
	}

	SetHostname("docky-container")

	err := cmd.Run()

	if err != nil {
		fmt.Println("error : ", err)
	}
}
