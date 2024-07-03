package action

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func Action(c *cli.Context) error {
	if c.NArg() < 1 {
        return fmt.Errorf("requires input file path")
    }
	_inputPath := c.Args().Get(0)

	_header := c.String("header")
	_separator := c.String("separator")
	_outputExtension := c.String("output")
	Convert(_header, _inputPath, _separator, _outputExtension)
    return nil
}

