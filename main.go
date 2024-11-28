package main

import (
	"github.com/dwiangraeni/dealls/api"
	"github.com/dwiangraeni/dealls/infra"
	"github.com/urfave/cli"
	"os"
)

const (
	// AppName Application name
	AppName = "dating-dealls"
	// AppTagLine Application tagline
	AppTagLine = "Dating Dealls Service"
)

var API = cli.Command{
	Name:     "api",
	Usage:    "Run API Server",
	HideHelp: true,
	Action: func(ctx *cli.Context) {
		api.NewServer(infra.New(ctx.GlobalString("config"))).Run()
	},
}

func main() {
	app := cli.NewApp()
	app.Name = AppName
	app.Usage = AppTagLine
	app.Version = "v1"
	app.HideHelp = true
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "config,c",
			Value:  "config/app.toml",
			Usage:  "Main config",
			EnvVar: "APP_CONFIG_FILE",
		},
	}

	app.Commands = []cli.Command{
		API,
	}

	app.Flags = append(app.Flags, []cli.Flag{}...)
	_ = app.Run(os.Args)

}
