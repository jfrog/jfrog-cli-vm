package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/jfrog/jfrog-cli-vm/cmd/utils"
	"github.com/urfave/cli/v2"
)

var Alias = &cli.Command{
	Name:  "alias",
	Usage: "Manage aliases for JFrog CLI versions",
	Subcommands: []*cli.Command{
		{
			Name:      "set",
			Usage:     "Set an alias (e.g., prod => 2.57.0)",
			ArgsUsage: "<alias> <version>",
			Action: func(c *cli.Context) error {
				if c.Args().Len() != 2 {
					return cli.Exit("Usage: jfvm alias set <alias> <version>", 1)
				}
				alias, version := c.Args().Get(0), c.Args().Get(1)

				// Prevent using "latest" as an alias since it's a reserved keyword
				if strings.ToLower(alias) == "latest" {
					return cli.Exit("'latest' is a reserved keyword and cannot be used as an alias", 1)
				}

				os.MkdirAll(utils.JfvmAliases, 0755)
				return os.WriteFile(filepath.Join(utils.JfvmAliases, alias), []byte(version), 0644)
			},
		},
		{
			Name:      "get",
			Usage:     "Get the version mapped to an alias",
			ArgsUsage: "<alias>",
			Action: func(c *cli.Context) error {
				if c.Args().Len() != 1 {
					return cli.Exit("Usage: jfvm alias get <alias>", 1)
				}
				version, err := utils.ResolveAlias(c.Args().Get(0))
				if err != nil {
					return err
				}
				fmt.Println(version)
				return nil
			},
		},
		{
			Name:      "remove",
			Usage:     "Remove an alias",
			ArgsUsage: "<alias>",
			Action: func(c *cli.Context) error {
				if c.Args().Len() != 1 {
					return cli.Exit("Usage: jfvm alias remove <alias>", 1)
				}
				return os.Remove(filepath.Join(utils.JfvmAliases, c.Args().Get(0)))
			},
		},
	},
}
