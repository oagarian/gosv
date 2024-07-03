package app

import (
	"os"

	"github.com/oagarian/gosv/app/action"
	"github.com/urfave/cli/v2"
)

func Run() error {
	app := &cli.App{
		Name: "gosv",
		Usage: "Convert .csv files to another extension",
		Flags: Flags(),
		Action: action.Action,
	}

    return app.Run(os.Args)
}