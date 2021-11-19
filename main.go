package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"strings"
)

func main() {
	app := cli.NewApp()
	app.Name = "highwayhash-cli"
	app.Usage = "fast calc highwayhash"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "key, k",
			Value: strings.Repeat("a", 32),
			Usage: "Key of this hash, should be 32 bit long",
		},
		cli.StringFlag{
			Name:  "path, p",
			Value: "test.bin",
			Usage: "the file you want to hash",
		},
	}

	app.Action = func(c *cli.Context) error {
		key := []byte(c.String("key"))
		path := c.String("path")

		checksum, err := Hash256(key, path)
		if err != nil {
			os.Exit(2)
		}

		fmt.Print(checksum)
		return nil
	}

	app.Run(os.Args)
}
