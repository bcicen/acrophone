package main

import (
	"fmt"
	"os"
	"strings"
)

var NatoMap = map[string]string{
	"-":        "Dash",
	"0":        "Zero",
	"1":        "One",
	"2":        "Two",
	"3":        "Three",
	"4":        "Four",
	"5":        "Five",
	"6":        "Six",
	"7":        "Seven",
	"8":        "Eight",
	"9":        "Niner",
	"a":        "Alfa",
	"b":        "Bravo",
	"c":        "Charlie",
	"d":        "Delta",
	"e":        "Echo",
	"f":        "Foxtrot",
	"g":        "Golf",
	"h":        "Hotel",
	"i":        "India",
	"j":        "Juliett",
	"k":        "Kilo",
	"l":        "Lima",
	"m":        "Mike",
	"n":        "November",
	"o":        "Oscar",
	"p":        "Papa",
	"q":        "Quebec",
	"r":        "Romeo",
	"s":        "Sierra",
	"t":        "Tango",
	"u":        "Uniform",
	"v":        "Victor",
	"w":        "Whiskey",
	"x":        "X-ray or Xray",
	"y":        "Yankee",
	"z":        "Zulu",
	".":        "Point",
	"EOL":      "Stop",
	"hundred":  "Hundred",
	"thousand": "Thousand",
}

type Acrophone struct {
	charMap map[string]string
}

func (a *Acrophone) lookup(char string) (string, error) {
	if len(char) > 1 {
		return "", fmt.Errorf("lookup() accepts single character argument only")
	}
	if res, ok := a.charMap[char]; ok == false {
		return "", fmt.Errorf("unable to find match for character: %s", char)
	} else {
		return res, nil
	}
}

func main() {
	var output []string

	if len(os.Args) < 2 {
		fmt.Println("usage: acrophone <input>")
		os.Exit(1)
	}

	a := &Acrophone{NatoMap}

	input := strings.Join(os.Args[1:], "")
	for _, char := range input {
		result, err := a.lookup(string(char))
		if err != nil {
			panic(err)
		}
		output = append(output, result)
	}

	fmt.Println(strings.Join(output, " "))
}
