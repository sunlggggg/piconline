package greet

import (
	"os"
	"gopkg.in/urfave/cli.v2"
	"fmt"
)

func main() {
	app := &cli.App{
		Name: "greet",
		Usage: "figgoht the loneliness!",
		Action: func(c *cli.Context) error {
			fmt.Println("Hello friend!")
			return nil
		},
	}

	app.Run(os.Args)
}