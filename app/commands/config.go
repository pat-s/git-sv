package commands

import (
	"fmt"

	"github.com/thegeeklab/git-sv/v2/app"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
)

func ConfigDefaultHandler() cli.ActionFunc {
	return func(c *cli.Context) error {
		cfg := app.GetDefault()

		content, err := yaml.Marshal(&cfg)
		if err != nil {
			return err
		}

		fmt.Println(string(content))

		return nil
	}
}

func ConfigShowHandler(cfg *app.Config) cli.ActionFunc {
	return func(c *cli.Context) error {
		content, err := yaml.Marshal(cfg)
		if err != nil {
			return err
		}

		fmt.Println(string(content))

		return nil
	}
}
