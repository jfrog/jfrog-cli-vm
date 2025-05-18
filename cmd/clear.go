package cmd

import (
	"fmt"
	"github.com/bhanurp/jfvm/cmd/utils"
	"github.com/urfave/cli/v2"
	"os"
)

var Clear = &cli.Command{
	Name:  "clear",
	Usage: "Remove all installed JFrog CLI versions",
	Action: func(c *cli.Context) error {
		err := os.RemoveAll(utils.JfvmVersions)
		if err != nil {
			return fmt.Errorf("failed to clear versions: %w", err)
		}
		fmt.Println("All versions removed.")
		return nil
	},
}
