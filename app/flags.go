package app

import "github.com/urfave/cli/v2"

var flags = []cli.Flag{
	&cli.StringFlag{
		Name:    "header",
		Aliases: []string{"hd"},
		Value:   "",
		Usage:   "Inform .csv header values (comma separated)",
	},
	&cli.StringFlag{
		Name: "separator",
		Aliases: []string{"s"},
        Value:   ",",
        Usage:   "Inform .csv separator",
	},
	&cli.StringFlag{
		Name: "output",
		Aliases: []string{"o"},
		Value:   "json",
        Usage:   "Inform output file extension",
	},
}

func Flags() []cli.Flag {
	return flags
}
