package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
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
	version := strings.TrimSpace(string(data))
	fmt.Printf(".jfrog-version content: %s\n", version)
	return version, nil
}

func ResolveAlias(name string) (string, error) {
	path := filepath.Join(JfvmAliases, name)
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}

// ResolveVersionOrAlias attempts to resolve an alias first, then falls back to the original name
func ResolveVersionOrAlias(name string) (string, error) {
	// Try to resolve as alias first
	resolved, err := ResolveAlias(name)
	if err == nil {
		return strings.TrimSpace(resolved), nil
	}

	// If not an alias, return the original name
	return name, nil
}

// CheckVersionExists verifies that a version directory and binary exist
func CheckVersionExists(version string) error {
	versionDir := filepath.Join(JfvmVersions, version)
	binaryPath := filepath.Join(versionDir, BinaryName)

	// Check if version directory exists
	if _, err := os.Stat(versionDir); os.IsNotExist(err) {
		return fmt.Errorf("version directory does not exist")
	}

	// Check if binary exists
	if _, err := os.Stat(binaryPath); os.IsNotExist(err) {
		return fmt.Errorf("binary not found in version directory")
	}

	return nil
}

// GetLatestVersion fetches the latest version from GitHub API
func GetLatestVersion() (string, error) {
	// Use GitHub API to get the latest release
	url := "https://api.github.com/repos/jfrog/jfrog-cli/releases/latest"
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to fetch latest version: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch latest version: HTTP %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}
	content := string(body)
	tagNameIndex := strings.Index(content, `"tag_name":"`)
	if tagNameIndex == -1 {
		return "", fmt.Errorf("could not find tag_name in response")
	}

	// Extract the version starting after "tag_name":"
	startIndex := tagNameIndex + len(`"tag_name":"`)
	endIndex := strings.Index(content[startIndex:], `"`)
	if endIndex == -1 {
		return "", fmt.Errorf("could not parse tag_name value")
	}

	version := content[startIndex : startIndex+endIndex]
	if !strings.HasPrefix(version, "v2.") {
		return "", fmt.Errorf("invalid version format: %s", version)
	}
	version = strings.TrimPrefix(version, "v")

	return version, nil
}
