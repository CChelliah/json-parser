package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParser_Parse(t *testing.T) {
	type fields struct {
		stack      []string
		tokenTypes map[string]TokenType
	}
	type args struct {
		tokens []string
		pos    *int
	}

	dummyPos := 0

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name: "Empty JSON Valid Tokens",
			fields: fields{
				stack: []string{},
				tokenTypes: map[string]TokenType{
					"{":  Delimiter,
					"}":  Delimiter,
					":":  Delimiter,
					"[":  Delimiter,
					"]":  Delimiter,
					"\"": Terminal,
				},
			},
			args: args{
				tokens: []string{"{", ":", "}"},
				pos:    &dummyPos,
			},
			wantErr: nil,
		},
		{
			name: "Invalid JSON Unclosed Curly Brace",
			fields: fields{
				stack: []string{},
				tokenTypes: map[string]TokenType{
					"{":  Delimiter,
					"}":  Delimiter,
					":":  Delimiter,
					"[":  Delimiter,
					"]":  Delimiter,
					"\"": Terminal,
				},
			},
			args: args{
				tokens: []string{"{"},
				pos:    &dummyPos,
			},
			wantErr: ErrInvalidJSON,
		},
		{
			name: "Invalid JSON Unopened Curly Brace",
			fields: fields{
				stack: []string{},
				tokenTypes: map[string]TokenType{
					"{":  Delimiter,
					"}":  Delimiter,
					":":  Delimiter,
					"[":  Delimiter,
					"]":  Delimiter,
					"\"": Terminal,
				},
			},
			args: args{
				tokens: []string{"}"},
				pos:    &dummyPos,
			},
			wantErr: ErrInvalidJSON,
		},
		{
			name: "Invalid JSON Unclosed Square Brace",
			fields: fields{
				stack: []string{},
				tokenTypes: map[string]TokenType{
					"{":  Delimiter,
					"}":  Delimiter,
					":":  Delimiter,
					"[":  Delimiter,
					"]":  Delimiter,
					"\"": Terminal,
				},
			},
			args: args{
				tokens: []string{"["},
				pos:    &dummyPos,
			},
			wantErr: ErrInvalidJSON,
		},
		{
			name: "Invalid JSON Unopened Square Brace",
			fields: fields{
				stack: []string{},
				tokenTypes: map[string]TokenType{
					"{":  Delimiter,
					"}":  Delimiter,
					":":  Delimiter,
					"[":  Delimiter,
					"]":  Delimiter,
					"\"": Terminal,
				},
			},
			args: args{
				tokens: []string{"]"},
				pos:    &dummyPos,
			},
			wantErr: ErrInvalidJSON,
		},
		{
			name: "Valid JSON Array",
			fields: fields{
				stack: []string{},
				tokenTypes: map[string]TokenType{
					"{":  Delimiter,
					"}":  Delimiter,
					":":  Delimiter,
					"[":  Delimiter,
					"]":  Delimiter,
					"\"": Terminal,
				},
			},
			args: args{
				tokens: []string{"{", "test", ":", "[", "A", "2", "]", "}"},
				pos:    &dummyPos,
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Parser{
				stack:      tt.fields.stack,
				tokenTypes: tt.fields.tokenTypes,
			}
			gotErr := p.Parse(tt.args.tokens, tt.args.pos)

			if tt.wantErr != nil {
				assert.ErrorIs(t, gotErr, tt.wantErr)
			} else {
				assert.Nil(t, gotErr)
			}
			dummyPos = 0
		})
	}
}
