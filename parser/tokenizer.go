package parser

import (
	"errors"
)

func tokenize(text string) ([]string, error) {
	stringless, err := removeStringConstant(text)
	if err != nil {
		return nil, err
	}

	return tokenize_stringless(stringless), nil
}

func removeStringConstant(text string) (string, error) {
	i := 0
	state := 0

	out := ""

	for {
		if i == len(text) {
			break
		}

		cur := text[i]

		if state == 0 {
			if isDoubleQuote(cur) {
				state = 1
			} else if isSingleQuote(cur) {
				state = 2
			} else {
				out += string(cur)
			}

		} else if state == 1 {
			if isDoubleQuote(cur) {
				state = 0
			} else if isBackslash(cur) {
				state = 3

			}
		} else if state == 2 {
			if isSingleQuote(cur) {
				state = 0
			} else if isBackslash(cur) {
				state = 4
			}
		} else if state == 3 {
			state = 1

		} else if state == 4 {
			state = 2
		}

		i++
	}

	if state != 0 {
		// leave this funky return string in case error is not checked by caller
		// if we return "", this could make its way into a sql query which selects ALL!
		return "** there was a typo in the encoded query **", errors.New("string ended abruptly in encoded query")
	}

	return out, nil
}

func tokenize_stringless(text string) []string {
	i := 0
	state := 0

	tokens := make([]string, 0)
	token := ""

	for {
		if i == len(text) {
			break
		}

		cur := text[i]

		if state == 0 {
			if isSpace(cur) {
				// ignore
			} else if isStartingTokenChar(cur) {
				// non-numeric token now
				token = string(cur)
				state = 1
			} else if isDigit(cur) {
				// a numeric token
				token = string(cur)
				state = 2
			} else if isOperator(cur) {
				//
				token = string(cur)
				state = 3
			}
		} else if state == 1 {
			// non-numeric token
			if isTokenChar(cur) {
				token += string(cur)
			} else {
				// we finished reading token
				tokens = append(tokens, token)
				token = ""
				state = 0
				i--
			}
		} else if state == 2 {
			// numeric token
			if isNumericTokenChar(cur) {
				token += string(cur)
			} else {
				// we finished reading token
				tokens = append(tokens, token)
				token = ""
				state = 0
				i--
			}
		} else if state == 3 {
			// operator
			if isOperator(cur) {
				token += string(cur)
			} else {
				// we finished reading operator
				tokens = append(tokens, token)
				token = ""
				state = 0
				i--
			}
		}

		i++
	}

	if token != "" {
		tokens = append(tokens, token)
	}

	return tokens
}
