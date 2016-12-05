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

func printAlphabets(c *cli.Context) {
	for name, a := range alphabets.All {
		fmt.Fprintf(c.App.Writer, "%s (%s)\n", name, a["desc"])
	}
}

func main() {
	var output []string
	var cmap alphabets.Alphabet

	cli.VersionPrinter = printFullVersion

	app := cli.NewApp()
	app.Name = "acrophone"
	app.Usage = "convert text to phonetic spelling"
	app.UsageText = "acrophone [options] [input text]"
	app.HideHelp = true
	app.Version = version
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "alphabet, a",
			Usage: "Phonetic alphabet to use",
			Value: "nato",
		},
		cli.BoolFlag{
			Name:  "list, l",
			Usage: "List available phonetic alphabets",
		},
	}
	app.Action = func(c *cli.Context) {
		if c.Bool("list") {
			printAlphabets(c)
			os.Exit(0)
		}

		if len(c.Args()) < 1 {
			fmt.Println("no argument provided")
			os.Exit(1)
		}

		if _, ok := alphabets.All[c.String("alphabet")]; ok == false {
			fmt.Printf("invalid alphabet: %s\n", c.String("alphabet"))
			os.Exit(1)
		}

		cmap = alphabets.All[c.String("alphabet")]

		var ignored []string
		input := strings.Join(c.Args(), "")
		for _, char := range input {
			result, err := cmap.Lookup(string(char))
			if err != nil {
				ignored = append(ignored, string(char))
			} else {
				output = append(output, result)
			}
		}
		if len(ignored) > 0 {
			fmt.Printf("warning: ignored non-standard characters: %s\n", strings.Join(ignored, ","))
		}

		fmt.Println(strings.Join(output, " "))
	}

	app.Run(os.Args)
}
