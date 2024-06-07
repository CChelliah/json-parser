package main

import (
	"errors"
	"os"
)

type Validator struct {
	Lexer  Lexer
	Parser Parser
}

func NewJSONValidator(lexer Lexer, parser Parser) Validator {
	return Validator{
		Lexer:  lexer,
		Parser: parser,
	}
}

func (v Validator) Validate(file *os.File) (result int, err error) {

	tokens, err := v.Lexer.Tokenise(file)

	if err != nil {
		return 0, err
	}

	pos := 0
	err = v.Parser.Parse(tokens, &pos)

	switch {
	case err != nil && errors.Is(err, ErrInvalidJSON):
		return 0, nil
	case err != nil:
		return 0, err
	}

	return 1, nil
}
