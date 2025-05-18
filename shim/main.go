package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	home := os.Getenv("HOME")
	configPath := filepath.Join(home, ".jfvm", "config")
	data, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "No current version set. Run `jfvm use <version>` first.\n")
		os.Exit(1)
	}

	version := strings.TrimSpace(string(data))
	bin := filepath.Join(home, ".jfvm", "versions", version, "jf")

	fmt.Printf("[shim] Executing version: %s\n", version)
	fmt.Printf("[shim] Full binary path: %s\n", bin)

	cmd := exec.Command(bin, os.Args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "[shim] Error executing binary: %v\n", err)
		os.Exit(1)
	}
}
