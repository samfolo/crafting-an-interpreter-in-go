package lexer

import (
	token "monkey/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	// set lexer on first character upon initialisation:
	l.readChar()

	return l
}

func (lex *Lexer) readChar() {
	if lex.readPosition >= len(lex.input) {
		lex.ch = 0
	} else {
		lex.ch = lex.input[lex.readPosition]
	}

	lex.position = lex.readPosition
	lex.readPosition += 1
}

func (lex *Lexer) NextToken() token.Token {
	var tok token.Token

	lex.skipWhitespace()

	switch lex.ch {
	case '=':
		// handle "=="
		if lex.peekChar() == '=' {
			ch := lex.ch
			lex.readChar()
			literal := string(ch) + string(lex.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}

		} else {
			tok = newToken(token.ASSIGN, lex.ch)
		}
	case ';':
		tok = newToken(token.SEMICOLON, lex.ch)
	case ':':
		tok = newToken(token.COLON, lex.ch)
	case ',':
		tok = newToken(token.COMMA, lex.ch)
	case '+':
		tok = newToken(token.PLUS, lex.ch)
	case '-':
		tok = newToken(token.MINUS, lex.ch)
	case '(':
		tok = newToken(token.LPAREN, lex.ch)
	case ')':
		tok = newToken(token.RPAREN, lex.ch)
	case '{':
		tok = newToken(token.LBRACE, lex.ch)
	case '}':
		tok = newToken(token.RBRACE, lex.ch)
	case '[':
		tok = newToken(token.LBRACKET, lex.ch)
	case ']':
		tok = newToken(token.RBRACKET, lex.ch)
	case '!':
		// handle "!="
		if lex.peekChar() == '=' {
			ch := lex.ch
			lex.readChar()
			literal := string(ch) + string(lex.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}

		} else {
			tok = newToken(token.BANG, lex.ch)
		}
	case '/':
		tok = newToken(token.SLASH, lex.ch)
	case '*':
		tok = newToken(token.ASTERISK, lex.ch)
	case '<':
		tok = newToken(token.LT, lex.ch)
	case '>':
		tok = newToken(token.GT, lex.ch)
	case '"':
		tok.Literal = lex.readString()
		tok.Type = token.STRING
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(lex.ch) {
			tok.Literal = lex.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(lex.ch) {
			tok.Literal = lex.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, lex.ch)
		}
	}

	lex.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// decide here what can be used in variable and function names
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (lex *Lexer) readIdentifier() string {
	position := lex.position

	for isLetter(lex.ch) {
		lex.readChar()
	}

	return lex.input[position:lex.position]
}

func (lex *Lexer) readNumber() string {
	position := lex.position

	for isDigit(lex.ch) {
		lex.readChar()
	}

	return lex.input[position:lex.position]
}

func (lex *Lexer) skipWhitespace() {
	for lex.ch == ' ' || lex.ch == '\t' || lex.ch == '\n' || lex.ch == '\r' {
		lex.readChar()
	}
}

func (lex *Lexer) peekChar() byte {
	if lex.readPosition >= len(lex.input) {
		return 0
	} else {
		return lex.input[lex.readPosition]
	}
}

func (lex *Lexer) readString() string {
	position := lex.readPosition // position + 1

	for {
		lex.readChar()
		if lex.ch == '"' || lex.ch == 0 {
			break
		}
	}

	return lex.input[position:lex.position]
}
