package lexer

import (
	"testing"

	"github.com/tommy-sho/monkey/token"
)

func TestNextToken(t *testing.T) {
	in := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRANCE, "{"},
		{token.RBRANCE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(in)

	for _, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("type error: got = %v, want = %v", tok.Type, tt.expectedType)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("literal error: got = [%v], want = [%v]", tok.Literal, tt.expectedLiteral)
		}
	}
}

func Test_NextToken(t *testing.T) {
	input := `let five = 5;
let ten = 10;

let add = fn(x, y) {
x + y;
};

let result = add(five, ten);
!-/*5;
5 < 10 > 5;
if (5 < 10) {
    return true;
} else {
    return false;
};

10 == 10;
10 != 9;
"foobar"
"foo bar"
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"}, {token.IDENT, "five"}, {token.ASSIGN, "="}, {token.INT, "5"},
		{token.SEMICOLON, ";"}, {token.LET, "let"}, {token.IDENT, "ten"}, {token.ASSIGN, "="},
		{token.INT, "10"}, {token.SEMICOLON, ";"}, {token.LET, "let"}, {token.IDENT, "add"},
		{token.ASSIGN, "="}, {token.FUNCTION, "fn"}, {token.LPAREN, "("}, {token.IDENT, "x"},
		{token.COMMA, ","}, {token.IDENT, "y"}, {token.RPAREN, ")"}, {token.LBRANCE, "{"},
		{token.IDENT, "x"}, {token.PLUS, "+"}, {token.IDENT, "y"}, {token.SEMICOLON, ";"},
		{token.RBRANCE, "}"}, {token.SEMICOLON, ";"}, {token.LET, "let"}, {token.IDENT, "result"},
		{token.ASSIGN, "="}, {token.IDENT, "add"}, {token.LPAREN, "("}, {token.IDENT, "five"},
		{token.COMMA, ","}, {token.IDENT, "ten"}, {token.RPAREN, ")"}, {token.SEMICOLON, ";"},
		{token.BANG, "!"}, {token.MINUS, "-"}, {token.SLASH, "/"}, {token.ASTERISK, "*"},
		{token.INT, "5"}, {token.SEMICOLON, ";"}, {token.INT, "5"}, {token.LT, "<"},
		{token.INT, "10"}, {token.GT, ">"}, {token.INT, "5"}, {token.SEMICOLON, ";"},

		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRANCE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRANCE, "}"},
		{token.ELSE, "else"},
		{token.LBRANCE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRANCE, "}"},
		{token.SEMICOLON, ";"},

		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.NEQ, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
		{token.STRING, "foobar"},
		{token.STRING, "foo bar"},

		{token.EOF, ""},
	}

	l := New(input)

	for _, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("type error: got = %v(%v), want = %v(%v)", tok.Type, tok.Literal, tt.expectedType, tt.expectedLiteral)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("literal error: got = [%v], want = [%v]", tok.Literal, tt.expectedLiteral)
		}
	}
}
