package main

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestLexer_Tokenise(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantTokens []string
		wantErr    error
	}{
		{
			name: "Tokenise Empty JSON",
			args: args{
				r: bytes.NewReader(json.RawMessage(`{}`)),
			},
			wantTokens: []string{"{", "}"},
			wantErr:    nil,
		},
		{
			name: "Tokenise Basic JSON",
			args: args{
				r: bytes.NewReader(json.RawMessage(`{ "key" : "value" }`)),
			},
			wantTokens: []string{"{", "\"key\"", ":", "\"value\"", "}"},
			wantErr:    nil,
		},
		{
			name: "Tokenise Array JSON",
			args: args{
				r: bytes.NewReader(json.RawMessage(`{ "key" : [ test 1 3 ] }`)),
			},
			wantTokens: []string{"{", "\"key\"", ":", "[", "test", "1", "3", "]", "}"},
			wantErr:    nil,
		},
		{
			name: "Text After JSON",
			args: args{
				r: bytes.NewReader(json.RawMessage(`{"Extra value after close": true} "misplaced quoted value"`)),
			},
			wantTokens: []string{"{", "\"Extra value after close\"", ":", "true", "}", "\"misplaced quoted value\""},
			wantErr:    nil,
		},
	}

	//{"Extra value after close": true} "misplaced quoted value"

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Lexer{}
			gotTokens, err := l.Tokenise(tt.args.r)

			assert.Equal(t, tt.wantTokens, gotTokens, "Tokenise() gotTokens = %v, want %v", gotTokens, tt.wantTokens)
			assert.ErrorIs(t, err, tt.wantErr, "Tokenise() error = %v, wantErr %v", err, tt.wantErr)
		})
	}
}
