package internal

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/bhanurp/jfvm/cmd/utils"
)

func mapPlatform(goos, arch string) (string, error) {
	switch goos {
	case "darwin":
		if arch == "arm64" {
			return "mac-arm64", nil
		}
		if arch == "amd64" {
			return "mac-386", nil
		}
	case "linux":
		if arch == "amd64" {
			return "linux-amd64", nil
		}
	case "windows":
		if arch == "amd64" {
			return "windows-amd64", nil
		}
	}
	return "", fmt.Errorf("unsupported platform: %s-%s", goos, arch)
}

func DownloadAndInstall(version string) error {
	platform, err := mapPlatform(runtime.GOOS, runtime.GOARCH)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://releases.jfrog.io/artifactory/jfrog-cli/v2-jf/%s/jfrog-cli-%s/jf", version, platform)
	fmt.Printf("📥 Downloading from: %s\n", url)

	dir := filepath.Join(utils.JfvmVersions, version)
	os.MkdirAll(dir, 0755)
	binPath := filepath.Join(dir, utils.BinaryName)

	out, err := os.Create(binPath)
	if err != nil {
		return fmt.Errorf("failed to create binary file: %w", err)
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("http request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download: %s", resp.Status)
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to write binary: %w", err)
	}

	if err := os.Chmod(binPath, 0755); err != nil {
		return fmt.Errorf("chmod failed: %w", err)
	}

	if runtime.GOOS == "darwin" {
		_ = exec.Command("xattr", "-c", binPath).Run()
	}

	return nil
}
