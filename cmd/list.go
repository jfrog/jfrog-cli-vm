package cmd

import (
	"fmt"
	"github.com/bhanurp/jfvm/cmd/utils"
	"github.com/urfave/cli/v2"
	"os"
)

var List = &cli.Command{
	Name:  "list",
	Usage: "List all installed JFrog CLI versions",
	Action: func(c *cli.Context) error {
		currentData, _ := os.ReadFile(utils.JfvmConfig)
		current := string(currentData)

		entries, err := os.ReadDir(utils.JfvmVersions)
		if err != nil {
			return err
		}

		fmt.Println("Installed versions:")
		for _, entry := range entries {
			if entry.IsDir() {
				version := entry.Name()
				mark := ""
				if version == current {
					mark = " (current)"
				}
				fmt.Printf(" - %s%s\n", version, mark)
			}
		}
		return nil
	},
}
