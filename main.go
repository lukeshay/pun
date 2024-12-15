package main

import (
	"os"
	"os/exec"
	"log"
)

func main() {
	if len(os.Args) < 1 {
		log.Fatal("You must provide a command to proxy.")
	}

	cmd := exec.Command("bun", os.Args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		log.Fatalf("Command failed: %v", err)
	}
} 