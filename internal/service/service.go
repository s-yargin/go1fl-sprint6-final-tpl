package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

var ErrEmptyInput = errors.New("input is empty")

// Convert determines whether input is Morse code and converts it to the other
// representation. A Morse message may contain dots, dashes and whitespace only.
func Convert(input string) (string, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return "", ErrEmptyInput
	}

	if isMorse(input) {
		return morse.ToText(input), nil
	}

	return morse.ToMorse(input), nil
}

func isMorse(input string) bool {
	return strings.Trim(input, ".- \t\r\n") == "" && strings.ContainsAny(input, ".-")
}
