package calc

import "errors"

var (
	ErrInvalidExpression  = errors.New("invalid expression")
	ErrDivisionByZero     = errors.New("division by zero")
	ErrOutOfTokens        = errors.New("out of tokens")
	ErrInvalidParentheses = errors.New("invalid parentheses")
)
