package cmd

import (
	"cli-urfave-exemplo/functions"

	"github.com/urfave/cli"
)

func NewCMD() *cli.App {

	cmd := cli.App{}

	cmd.Name = "minha cli"
	cmd.Usage = "cli para resolver dns"

	cmd.Commands = []cli.Command{
		{
			Name:  "resolver",
			Usage: "Resolver dns",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:     "host",
					Usage:    "Host a ser resolvido",
					Required: true,
				},
			},
			Action: functions.ResolverDNS,
		},
	}

	return &cmd

}
