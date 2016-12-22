package main;

import (
	_ "fmt"
	"github.com/urfave/cli"
	"os"
	"fmt"
	"github.com/hashicorp/consul/api"
)

func main() {
	app := cli.NewApp();
	app.Usage = "Hello world"

	app.Commands = []cli.Command {
		{
			Name : "test",
			ShortName: "t",
			Usage: "It is a test!",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name: "test, t",
					Usage: "Some number",

				},
			},
			Action: doSomething,
		},
	}

	app.CommandNotFound = func(c *cli.Context, command string) {
		cli.ShowAppHelp(c);
	}


	app.Run(os.Args)
}


func doSomething(c *cli.Context) error {
	val := c.Int("test");
	fmt.Printf("You ran the test with %d \n", val)

	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}

	kv := client.KV()
	pair, _, _ := kv.Get("foo/bar/baz", nil)
	fmt.Printf("KV: %s", pair.Value)

	return nil
}