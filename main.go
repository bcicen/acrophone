package main

import (
	"fmt"
	"os"
	"strings"
)

type Acrophone struct {
	charMap map[string]string
}

func (a *Acrophone) Lookup(char string) (string, error) {
	char = strings.ToLower(char)

	if len(char) > 1 {
		return "", fmt.Errorf("Lookup() accepts single character argument only")
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
		result, err := a.Lookup(string(char))
		if err != nil {
			panic(err)
		}
		output = append(output, result)
	}

	fmt.Println(strings.Join(output, " "))
}
