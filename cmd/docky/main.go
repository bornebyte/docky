package main

import (
	"docky/container"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: docky <command>")
		return
	}

	switch os.Args[1] {
	case "run":
		if len(os.Args) < 3 {
			fmt.Println("usage: docky run <command>")
			return
		}
		container.Run(os.Args[2:])
	case "init":
		container.Init()
	default:
		fmt.Println("unknown command")
	}
}
