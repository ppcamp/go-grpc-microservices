package configs

import (
	"github.com/urfave/cli/v2"
)

var Flags = []cli.Flag{
	&cli.StringFlag{
		Name:        "app-port",
		Destination: &ServerPort,
		EnvVars:     []string{"APP_PORT"},
		Value:       ":8080",
	},
	&cli.StringFlag{
		Name:        "path",
		Destination: &Path,
		EnvVars:     []string{"APP_PATH"},
		Value:       "../public",
	},
}
