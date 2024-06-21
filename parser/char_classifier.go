package parser

import "unicode"

func isAlphabet(cur uint8) bool {
	return unicode.IsLetter(rune(cur))
}

func isDigit(cur uint8) bool {
	return unicode.IsDigit(rune(cur))
}

func isSpace(cur uint8) bool {
	return unicode.IsSpace(rune(cur))
}

func isOperator(cur uint8) bool {
	return cur == '!' || cur == '=' ||
		cur == '<' || cur == '>' || cur == ')' || cur == '('
}

func isDot(cur uint8) bool {
	return cur == '.'
}

func isNumericTokenChar(cur uint8) bool {
	return isDigit(cur) || isDot(cur)
}

func isUnderScore(cur uint8) bool {
	return cur == '_'
}

func isDollar(cur uint8) bool {
	return cur == '$'
}

func isDoubleQuote(cur uint8) bool {
	return cur == '"'
}

func isSingleQuote(cur uint8) bool {
	return cur == '\''
}

func isBackslash(cur uint8) bool {
	return cur == '\\'
}

func isAlphaNumeric(cur uint8) bool {
	return isAlphabet(cur) || isDigit(cur)
}

func isStartingTokenChar(cur uint8) bool {
	return isAlphabet(cur) || isUnderScore(cur) || isDollar(cur)
}

func isTokenChar(cur uint8) bool {
	return isAlphaNumeric(cur) || isDot(cur) || isUnderScore(cur)
}
