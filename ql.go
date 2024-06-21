package dotsql

import "github.com/amortaza/dotsql/parser"

func ToSQL(dsql string) (string, error) {
	if dsql == "" {
		return "", nil
	}

	tokens := parser.Tokenize(dsql)
	dotted := parser.ParseDotted(tokens)

	return "", nil
}
