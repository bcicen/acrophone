package main

import (
	"fmt"
	"strings"
)

type charMap map[string]string

func (cmap charMap) Lookup(char string) (string, error) {
	char = strings.ToLower(char)

	if len(char) > 1 {
		return "", fmt.Errorf("Lookup() accepts single character argument only")
	}

	if res, ok := cmap[char]; ok == false {
		return "", fmt.Errorf("unable to find match for character: %s", char)
	} else {
		return res, nil
	}
}

var NatoMap = charMap{
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
	"EOL":      "<Stop>",
	"hundred":  "Hundred",
	"thousand": "Thousand",
}
