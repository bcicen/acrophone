package alphabets

import (
	"fmt"
	"strings"
)

type Alphabet map[string]string

func (a Alphabet) Lookup(char string) (string, error) {
	char = strings.ToLower(char)

	if len(char) > 1 {
		return "", fmt.Errorf("Lookup() accepts single character argument only")
	}

	if res, ok := a[char]; ok == false {
		return "", fmt.Errorf("unable to find match for character: %s", char)
	} else {
		return res, nil
	}
}
