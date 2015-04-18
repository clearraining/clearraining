package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func getDiscovery(c *cli.Context) string {
	if len(c.Args()) == 1 { //命令行参数
		return c.Args()[0]
	}
	return os.Getenv("SWARM_DISCOVERY")
}
