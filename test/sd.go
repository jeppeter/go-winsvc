package main

import (
	"fmt"
	"os"
	"os/exec"
	//"unsafe"
)

func ShutdownSystem(rebflag int) {
	c := exec.Command("shutdown.exe", "-s", "-t", "0")
	if rebflag > 0 {
		c = exec.Command("shutdown.exe", "-r", "-t", "0")
	}

	err := c.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "can not run shutdown")
	}
}

func main() {
	shutdown := 0
	if len(os.Args) > 1 {
		shutdown = 1
	}
	ShutdownSystem(shutdown)
}
