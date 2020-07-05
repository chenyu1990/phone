package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"phone-area/app/area"
)

func main()  {
	app := cli.NewApp()
	app.Name = "Phone"
	app.Version = "1.0.0"
	app.Usage = "手机号码"
	app.Commands = []*cli.Command{
		newCmd(),
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func newCmd() *cli.Command {
	return &cli.Command{
		Name:  "area",
		Usage: "手机号归属地查询",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "file",
				Aliases:  []string{"f"},
				Usage:    "手机号码文件",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			return area.NewArea(c.String("file")).Run()
		},
	}
}