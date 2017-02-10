package alphabets

import (
	"fmt"
	"strings"
)

var All = map[string]Alphabet{
	"nato":    Nato,
	"nypd":    NYPD,
	"spanish": Spanish,
}

type LookupError struct {
	char string
}

func (n LookupError) Error() string {
	return fmt.Sprintf("unable to find match for character: %s", n.char)
}

type Alphabet map[string]string

func (a Alphabet) Lookup(char string) (string, error) {
	char = strings.ToLower(char)

	if len(char) > 1 {
		return "", fmt.Errorf("Lookup() accepts single character argument only")
	}

	if res, ok := a[char]; ok == false {
		err := LookupError{char}
		return "", err
	} else {
		return res, nil
	}
}
