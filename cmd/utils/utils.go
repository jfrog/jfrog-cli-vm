package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	ToolName    = "jfvm"
	ConfigFile  = "config"
	VersionsDir = "versions"
	BinaryName  = "jf"
	ProjectFile = ".jfrog-version"
	AliasesDir  = "aliases"
)

var (
	HomeDir      = os.Getenv("HOME")
	JfvmRoot     = filepath.Join(HomeDir, "."+ToolName)
	JfvmConfig   = filepath.Join(JfvmRoot, ConfigFile)
	JfvmVersions = filepath.Join(JfvmRoot, VersionsDir)
	JfvmAliases  = filepath.Join(JfvmRoot, AliasesDir)
)

func GetVersionFromProjectFile() (string, error) {
	fmt.Println("Attempting to read .jfrog-version file...")
	data, err := os.ReadFile(ProjectFile)
	if err != nil {
		fmt.Printf("Failed to read .jfrog-version file: %v\n", err)
		return "", err
	}
	fmt.Printf(".jfrog-version content: %s\n", string(data))
	return string(data), nil
}

func ResolveAlias(name string) (string, error) {
	path := filepath.Join(JfvmAliases, name)
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
