package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	app := cli.App{
		Name: "hello",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "lang",
				Value: "english", //
				Usage: "language for the greeting",
			},
		},
		Action: func(c *cli.Context) error {
			name := "world"
			if c.NArg() > 0 {
				name = c.Args().Get(0)
			}

			if c.String("lang") == "english" {
				fmt.Println("hello", name)
			} else {
				fmt.Println("你好", name)
			}
			return nil
		},
		Usage: "hello world example,I am here!",
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("app run failed,error:", err)
		return
	}

	fmt.Println("app exit")
}
