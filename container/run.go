package container

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func Run(command []string) {

	fmt.Println("Starting container...")

	cmd := exec.Command("/proc/self/exe", append([]string{"init"}, command...)...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS |
			syscall.CLONE_NEWPID |
			syscall.CLONE_NEWNS,
	}

	err := cmd.Run()
	if err != nil {
		fmt.Println("error:", err)
	}
}
