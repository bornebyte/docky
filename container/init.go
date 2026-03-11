package container

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func Init() {
	fmt.Println("Container init process running")

	syscall.Sethostname([]byte("docky-container"))
	syscall.Mount("proc", "/proc", "proc", 0, "")

	command := os.Args[2:]

	cmd := exec.Command(command[0], command[1:]...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		fmt.Println("error : ", err)
	}
}
