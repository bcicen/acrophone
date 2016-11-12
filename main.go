package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/bcicen/acrophone/alphabets"
	"github.com/codegangsta/cli"
)

var (
	build   = ""
	version = "dev-build"
)

func printFullVersion(c *cli.Context) {
	fmt.Fprintf(c.App.Writer, "%v version %v, build %v\n", c.App.Name, c.App.Version, build)
}

func main() {
	var output []string
	var cmap alphabets.Alphabet

	cli.VersionPrinter = printFullVersion

	app := cli.NewApp()
	app.Name = "acrophone"
	app.Usage = "convert text to phonetic spelling"
	app.UsageText = "acrophone [options] [input text]"
	app.Version = version
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "alphabet, a",
			Usage: "Phonetic alphabet to use",
			Value: "nato",
		},
	}
	app.Action = func(c *cli.Context) {
		if len(c.Args()) < 1 {
			fmt.Println("no argument provided")
			os.Exit(1)
		}

		if _, ok := alphabets.All[c.String("alphabet")]; ok == false {
			fmt.Printf("invalid alphabet: %s\n", c.String("alphabet"))
			os.Exit(1)
		}

		cmap = alphabets.All[c.String("alphabet")]

		input := strings.Join(c.Args(), "")
		for _, char := range input {
			result, err := cmap.Lookup(string(char))
			if err != nil {
				panic(err)
			}
			output = append(output, result)
		}

		fmt.Println(strings.Join(output, " "))
	}

	// hacky way of getting around reserved command/arg words
	var newArgs []string
	for _, arg := range os.Args {
		if arg == "help" || arg == "version" {
			arg += "."
		}
		newArgs = append(newArgs, arg)
	}

	app.Run(newArgs)
}
