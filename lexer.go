package main

import (
	"bufio"
	"io"
)

type Lexer struct {
}

func NewLexer() Lexer {
	return Lexer{}
}

func (l *Lexer) Tokenise(r io.Reader) (tokens []string, err error) {

	s := bufio.NewScanner(r)
	s.Split(splitJSON)

	tokens = []string{}

	for s.Scan() {
		tokens = append(tokens, s.Text())
	}

	return tokens, nil
}

func splitJSON(data []byte, eof bool) (advance int, tokenBytes []byte, err error) {

	token := ""

	if eof {
		return 0, nil, io.EOF
	}

	for i := 0; i < len(data); i++ {
		advance++
		if isWhiteSpace(data[i]) {
			continue
		}
		token += string(data[i])

		switch {
		case isDelimiter(token):
			return advance, []byte(token), nil
		case isLiteral(token):
			advance++
			i++
			for advance < len(data) && string(data[i]) != "\"" {
				token += string(data[i])
				i++
				advance++
			}
			token += string(data[i])
			return advance, []byte(token), nil
		default:
			advance++
			i++
			for advance < len(data) &&
				string(data[i]) != " " &&
				string(data[i]) != "}" &&
				string(data[i]) != "{" &&
				string(data[i]) != "]" &&
				string(data[i]) != "[" {
				token += string(data[i])
				i++
				advance++
			}
			return advance, []byte(token), nil
		}
	}

	return 0, nil, nil
}

func isWhiteSpace(ch byte) bool {
	return string(ch) == " " || string(ch) == "\n" || string(ch) == "\r" || string(ch) == "\t" || string(ch) == "\v"
}

func isLiteral(token string) bool {
	return token == "\""
}

func isDelimiter(token string) bool {

	delimiters := map[string]interface{}{
		"{": nil,
		"}": nil,
		"[": nil,
		"]": nil,
		":": nil,
	}

	if _, ok := delimiters[token]; ok {
		return true
	}
	return false
}
