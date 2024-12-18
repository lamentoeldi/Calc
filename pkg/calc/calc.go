package calc

import (
	"strconv"
	"unicode"
)

func Calc(expression string) (float64, error) {
	// Recursive descent, why not?

	tokens, err := func() ([]string, error) {
		var tokens []string
		var currentToken string

		for _, r := range expression {
			if unicode.IsSpace(r) {
				continue
			}
			if unicode.IsDigit(r) || r == '.' {
				currentToken += string(r)
			} else {
				if currentToken != "" {
					tokens = append(tokens, currentToken)
					currentToken = ""
				}
				if r == '+' || r == '-' || r == '*' || r == '/' || r == '(' || r == ')' {
					tokens = append(tokens, string(r))
				} else {
					return nil, ErrInvalidExpression
				}
			}
		}
		if currentToken != "" {
			tokens = append(tokens, currentToken)
		}
		return tokens, nil
	}()
	if err != nil {
		return 0, err
	}

	return func() (float64, error) {
		var currentTokenIndex int

		var parseExpression func() (float64, error)
		var parseTerm func() (float64, error)
		var parseFactor func() (float64, error)

		parseExpression = func() (float64, error) {
			var leftTerm float64

			leftTerm, err = parseTerm()
			if err != nil {
				return 0, err
			}

			for currentTokenIndex < len(tokens) {
				if tokens[currentTokenIndex] == "+" || tokens[currentTokenIndex] == "-" {
					operator := tokens[currentTokenIndex]
					currentTokenIndex++
					rightTerm, err := parseTerm()
					if err != nil {
						return 0, err
					}
					switch operator {
					case "+":
						leftTerm += rightTerm
					case "-":
						leftTerm -= rightTerm
					}
				} else {
					break
				}
			}
			return leftTerm, nil
		}

		parseTerm = func() (float64, error) {
			var leftTerm float64

			leftTerm, err = parseFactor()
			if err != nil {
				return 0, err
			}

			for currentTokenIndex < len(tokens) {
				if tokens[currentTokenIndex] == "*" || tokens[currentTokenIndex] == "/" {
					operator := tokens[currentTokenIndex]
					currentTokenIndex++
					rightTerm, err := parseFactor()
					if err != nil {
						return 0, err
					}
					if operator == "*" {
						leftTerm *= rightTerm
					} else {
						if rightTerm == 0 {
							return 0, ErrDivisionByZero
						}
						leftTerm /= rightTerm
					}
				} else {
					break
				}
			}
			return leftTerm, nil
		}

		parseFactor = func() (float64, error) {
			if currentTokenIndex >= len(tokens) {
				return 0, ErrOutOfTokens
			}
			if tokens[currentTokenIndex] == "-" {
				currentTokenIndex++

				value, err := parseFactor()
				if err != nil {
					return 0, err
				}

				return -value, nil
			}
			if tokens[currentTokenIndex] == "(" {
				currentTokenIndex++

				value, err := parseExpression()
				if err != nil {
					return 0, err
				}

				if currentTokenIndex >= len(tokens) || tokens[currentTokenIndex] != ")" {
					return 0, ErrInvalidParentheses
				}
				currentTokenIndex++
				return value, nil
			}

			num, err := strconv.ParseFloat(tokens[currentTokenIndex], 64)
			if err != nil {
				return 0, err
			}

			currentTokenIndex++
			return num, nil
		}
		return parseExpression()
	}()
}
