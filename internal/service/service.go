package service

import (
	"strings"
	"unicode"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func AutoDetectAndConvert(input string) (string, error) {
	if strings.TrimSpace(input) == "" {
		return "", nil
	}

	if isMorse(input) {
		result := morse.ToText(input)
		return result, nil
	} else {
		result := morse.ToMorse(input)
		return result, nil
	}
}

func isMorse(input string) bool {
	if strings.TrimSpace(input) == "" {
		return false
	}

	hasMorseSymbols := strings.Contains(input, ".") || strings.Contains(input, "-")
	if !hasMorseSymbols {
		return false
	}

	for _, char := range input {
		if char != '.' && char != '-' && char != '/' && !unicode.IsSpace(char) {
			return false
		}
	}

	return true
}
