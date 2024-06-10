package main

import "errors"

var (
	ErrInvalidJSON = errors.New("invalid json")
	Terminal       = TokenType("terminal")
	Delimiter      = TokenType("delimiter")
)

type (
	TokenType string

	Parser struct {
		stack      []string
		tokenTypes map[string]TokenType
	}
)

func NewParser() Parser {

	tt := map[string]TokenType{
		"{":  Delimiter,
		"}":  Delimiter,
		":":  Delimiter,
		"[":  Delimiter,
		"]":  Delimiter,
		"\"": Terminal,
	}

	return Parser{
		stack:      []string{},
		tokenTypes: tt,
	}
}

func (p Parser) Parse(tokens []string, pos *int) error {
	for *pos < len(tokens) {
		var tokenType TokenType
		token := tokens[*pos]
		switch {
		case p.tokenTypes[token] == Delimiter || token == "true" || token == "false":
			tokenType = Delimiter
		case p.tokenTypes[string(tokens[*pos][0])] == Terminal:
			tokenType = Terminal
		default:
			tokenType = Terminal
		}
		length := len(p.stack)
		if tokenType == Delimiter && (token == "{" || token == "[") {
			*pos++
			p.stack = append(p.stack, token)
		} else if tokenType == Delimiter && token == ":" {
			*pos++
			p.stack = append(p.stack, token)
		} else if tokenType == Delimiter && len(p.stack) >= 2 &&
			(token == "}" && p.stack[length-2] == "{" || token == "]" && p.stack[length-2] == "[") && p.stack[length-1] == ":" {
			*pos++
			p.stack = p.stack[:(length - 2)]
		} else if tokenType == Delimiter && len(p.stack) == 1 && (token == "}" && p.stack[length-1] == "{" ||
			token == "]" && p.stack[length-1] == "[") {
			*pos++
			p.stack = p.stack[:(length - 1)]
		} else if tokenType == Delimiter &&
			len(p.stack) >= 1 &&
			token == "]" && p.stack[length-1] == "[" {
			*pos++
			p.stack = p.stack[:(length - 1)]
		} else if tokenType == Terminal {
			*pos++
		} else {
			return ErrInvalidJSON
		}
	}

	if len(p.stack) != 0 {
		return ErrInvalidJSON
	}
	return nil
}
