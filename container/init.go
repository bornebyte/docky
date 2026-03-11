package container

import (
	"syscall"
)

func SetHostname(name string) {
	syscall.Sethostname([]byte(name))
}
