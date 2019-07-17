package main

import (
	"demo5_cli_server/cmd/docker"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "testcli"
	app.Version = "v1.0"
	app.Commands = []cli.Command{
		docker.Command,
	}

	app.Run(os.Args)
}
