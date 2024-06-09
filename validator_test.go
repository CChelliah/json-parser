package main

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestValidator_Validate(t *testing.T) {
	type fields struct {
		Lexer  Lexer
		Parser Parser
	}
	type args struct {
		content []byte
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult int
		wantErr    error
	}{
		{
			name: "Empty JSON",
			args: args{
				content: []byte("{}"),
			},
			wantResult: 1,
			wantErr:    nil,
		},
		{
			name: "String Outside JSON",
			args: args{
				content: []byte("{\"Extra value after close\": true} \"misplaced quoted value\""),
			},
			wantResult: 0,
			wantErr:    ErrInvalidJSON,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Validator{
				Lexer:  tt.fields.Lexer,
				Parser: tt.fields.Parser,
			}

			file, err := createTestFile(tt.args.content)

			if err != nil {
				panic(err)
			}

			gotResult, err := v.Validate(file)

			assert.Equal(t, tt.wantResult, gotResult)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}

func createTestFile(content []byte) (file *os.File, err error) {

	file, err = os.CreateTemp("", "")

	if err != nil {
		return file, errors.New("failed to create temp file")
	}

	_, err = file.Write(content)
	if err != nil {
		return file, errors.New("failed to create temp file")
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		return file, errors.New("failed to seek temp file")
	}
	return file, nil
}
