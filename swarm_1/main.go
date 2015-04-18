package main

import (
	"fmt"
	"os"
	"path"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/docker/swarm_1/discovery/token"
	"github.com/docker/swarm_1/version"
	"github.com/docker/swarm_1/discovery"
)

func main() {
	app := cli.NewApp()
	app.Name = path.Base(os.Args[0]) //返回路径的最后一个元素（string）
	app.Usage = "a Docker-native clustering system"
	app.Version = version.VERSION + "(" + version.GITCOMMIT + ")"

	app.Author = ""
	app.Email = ""

	app.Flags = []cli.Flag{ //定义app的标识
		cli.BoolFlag{
			Name:   "debug",
			Usage:  "debug mode",
			EnvVar: "DEBUG",
		},
		cli.StringFlag{
			Name:  "log-level, l",
			Value: "info",
			Usage: fmt.Sprintf("Log level (options: debug, warn, error, fatal, panic)"), //定义日志的级别类型
		},
	}

	//logs
	app.Before = func(c *cli.Context) error { //在子命令run之前,   context中包括 *App 和 Command
		log.SetOutput(os.Stderr)                            //设置标准的日志输出
		level, err := log.ParseLevel(c.String("log-level")) //解析日志的级别（fatal, error, warn, info, debug）
		if err != nil {
			log.Fatalf(err.Error())
		}
		log.SetLevel(level) //设置日志级别

		if !c.IsSet("log-level") && !c.IsSet("l") && c.Bool("debug") { //IsSet()查明是否标志是确切的set
			log.SetLevel(log.DebugLevel)
		}
		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:      "create",
			ShortName: "c",
			Usage:     "create a cluster",
			Action: func(c *cli.Context) {
				discovery := &token.Discovery{}
				discovery.Initialize("", 0)
				token, err := discovery.CreateCluster()
				if len(c.Args()) != 0 {   //create 后面不能跟参数
					log.Fatalf("the 'create' command takes no arguments, See '%s create --help'.", c.App.Name)
				}
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println(token)
			},
		},

		{
			Name: "list",
			ShortName: "l",
			Usage: "list node in a cluster",
			Action: func(c *cli.Context) {
				dflag := getDiscovery(c)  //返回命令行的第一个参数,即token字串
				if dflag == "" {   //第一个参数为空
					log.Fatalf("discovery required to list a cluster. See '%s list --help'.", c.App.Name)
				}

				d, err := discovery.New(dflag, 0)
				
			}
		},



	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
