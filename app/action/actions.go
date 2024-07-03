package action

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func Action(c *cli.Context) error {
	if c.NArg() < 1 {
        return fmt.Errorf("requires input file path")
    }
	inputPath := c.Args().Get(0)
	Convert(c.String("header"), inputPath, c.String("separator"))
    return nil
}

