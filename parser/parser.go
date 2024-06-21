package parser

import (
	"strings"
)

func ParseDotted(text string) ([]string, error) {
	tokens, err := tokenize(text)
	if err != nil {
		return nil, err
	}

	dotted := make([]string, 0)

	for _, token := range tokens {
		if strings.Contains(token, ".") && isValidDottedToken(token) {
			dotted = append(dotted, token)
		}
	}

	return dotted, nil
}

func isValidDottedToken(token string) bool {
	state := 0
	i := 0

	for {
		if i == len(token) {
			break
		}

		cur := token[i]

		if state == 0 {
			// first char!
			if !isStartingTokenChar(cur) {
				return false
			}

			state = 1

		} else if state == 1 {
			if isDot(cur) {
				state = 2
			}

			if !isTokenChar(cur) {
				return false
			}
		} else if state == 2 {
			if isDot(cur) {
				return false
			}

			if !isTokenChar(cur) {
				return false
			}

			state = 1
		}

		i++
	}

	if state != 1 {
		return false
	}

	return true
}
