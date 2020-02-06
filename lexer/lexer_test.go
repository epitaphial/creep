package lexer

import (
	"testing"
)

func Test_lexer(t *testing.T) {
	//test1
	testString1 := `=+-*/!><&|,;{}[]()`
	testResult1 := []Token{
		Token{ASSIGN, "="},
		Token{PLUS, "+"},
		Token{MINUS, "-"},
		Token{ASTERISK, "*"},
		Token{SLASH, "/"},
		Token{EXCLAMATION,"!"},
		Token{GREATERTHAN,">"},
		Token{LESSTHAN,"<"},
		Token{AMPERSAND,"&"},
		Token{VERTICALBAR,"|"},
		Token{COMMA, ","},
		Token{SEMICOLON, ";"},
		Token{LBRACE, "{"},
		Token{RBRACE, "}"},
		Token{LBRACKET, "["},
		Token{RBRACKET, "]"},
		Token{LPARENTHESES, "("},
		Token{RPARENTHESES, ")"},
		Token{EOF,""},
	}
	nlexer1 := NewLexer(testString1)
	for key, token := range testResult1 {
		currtok := nlexer1.NextToken()
		if currtok.Type != token.Type {
			t.Fatalf("1-tests[%d] - tokentype wrong. expected=%q, got=%q", key, token.Type, currtok.Type)
		}
		if currtok.Literal != token.Literal {
			t.Fatalf("1-tests[%d] - literal wrong. expected=%q, got=%q", key, token.Literal, currtok.Literal)
		}
	}
	//test2
	testString2 := `
	order a = 1;
	addone = fun(a){
		a = a + 1;
		return a;
	};
	order b = addone(a);
	order Typeone = 556.886;
	`
	testResult2 := []Token{
		Token{ORDER, "order"},
		Token{IDENT, "a"},
		Token{ASSIGN, "="},
		Token{INT,"1"},
		Token{SEMICOLON, ";"},
		Token{IDENT, "addone"},
		Token{ASSIGN, "="},
		Token{FUNCTION, "fun"},
		Token{LPARENTHESES,"("},
		Token{IDENT,"a"},
		Token{RPARENTHESES,")"},
		Token{LBRACE,"{"},
		Token{IDENT,"a"},
		Token{ASSIGN,"="},
		Token{IDENT,"a"},
		Token{PLUS, "+"},
		Token{INT, "1"},
		Token{SEMICOLON, ";"},
		Token{RETURN, "return"},
		Token{IDENT, "a"},
		Token{SEMICOLON, ";"},
		Token{RBRACE, "}"},
		Token{SEMICOLON, ";"},
		Token{ORDER, "order"},
		Token{IDENT, "b"},
		Token{ASSIGN, "="},
		Token{IDENT, "addone"},
		Token{LPARENTHESES, "("},
		Token{IDENT, "a"},
		Token{RPARENTHESES, ")"},
		Token{SEMICOLON, ";"},
		Token{ORDER, "order"},
		Token{IDENT, "Typeone"},
		Token{ASSIGN, "="},
		Token{FLOAT,"556.886"},
		Token{SEMICOLON, ";"},
		Token{EOF,""},
	}
	nlexer2 := NewLexer(testString2)
	for key, token := range testResult2 {
		currtok := nlexer2.NextToken()
		if currtok.Type != token.Type {
			t.Fatalf("2-tests[%d] - tokentype wrong. expected=%q, got=%q", key, token.Type, currtok.Type)
		}
		if currtok.Literal != token.Literal {
			t.Fatalf("2-tests[%d] - literal wrong. expected=%q, got=%q", key, token.Literal, currtok.Literal)
		}
	}
	
	//test3
	testString3 := `
		order ifbiggerthanzero = fun(num){
			if (num > 0){
				return true;
			}else{
				return false;
			}
		};
		3>255.54<6;
	`
	testResult3 := []Token{
		Token{ORDER, "order"},
		Token{IDENT, "ifbiggerthanzero"},
		Token{ASSIGN, "="},
		Token{FUNCTION,"fun"},
		Token{LPARENTHESES, "("},
		Token{IDENT, "num"},
		Token{RPARENTHESES,")"},
		Token{LBRACE,"{"},
		Token{IF,"if"},
		Token{LPARENTHESES, "("},
		Token{IDENT, "num"},
		Token{GREATERTHAN, ">"},
		Token{INT, "0"},
		Token{RPARENTHESES, ")"},
		Token{LBRACE,"{"},
		Token{RETURN, "return"},
		Token{TRUE, "true"},
		Token{SEMICOLON, ";"},
		Token{RBRACE,"}"},
		Token{ELSE,"else"},
		Token{LBRACE,"{"},
		Token{RETURN, "return"},
		Token{FALSE, "false"},
		Token{SEMICOLON, ";"},	
		Token{RBRACE,"}"},
		Token{RBRACE,"}"},
		Token{SEMICOLON, ";"},
		Token{INT, "3"},
		Token{GREATERTHAN, ">"},
		Token{FLOAT, "255.54"},
		Token{LESSTHAN, "<"},
		Token{INT, "6"},
		Token{SEMICOLON, ";"},
		Token{EOF,""},
	}
	nlexer3 := NewLexer(testString3)
	for key, token := range testResult3 {
		currtok := nlexer3.NextToken()
		if currtok.Type != token.Type {
			t.Fatalf("3-tests[%d] - tokentype wrong. expected=%q, got=%q", key, token.Type, currtok.Type)
		}
		if currtok.Literal != token.Literal {
			t.Fatalf("3-tests[%d] - literal wrong. expected=%q, got=%q", key, token.Literal, currtok.Literal)
		}
	}

	//test4
	testString4 := `
	order a = 1;
	for(a;a<=4;a=a+1){
		if(a == 2 && a !=4){
			break;
		}
	}
	a = a & 2 | 4;
	`
	testResult4 := []Token{
		Token{ORDER, "order"},
		Token{IDENT, "a"},
		Token{ASSIGN, "="},
		Token{INT,"1"},
		Token{SEMICOLON, ";"},
		Token{FOR, "for"},
		Token{LPARENTHESES, "("},
		Token{IDENT, "a"},
		Token{SEMICOLON, ";"},
		Token{IDENT,"a"},
		Token{LTE,"<="},
		Token{INT,"4"},
		Token{SEMICOLON, ";"},
		Token{IDENT,"a"},
		Token{ASSIGN, "="},
		Token{IDENT,"a"},
		Token{PLUS,"+"},
		Token{INT,"1"},
		Token{RPARENTHESES,")"},
		Token{LBRACE,"{"},
		Token{IF,"if"},
		Token{LPARENTHESES,"("},
		Token{IDENT,"a"},
		Token{EQ,"=="},
		Token{INT,"2"},
		Token{DAMPERSAND, "&&"},
		Token{IDENT, "a"},
		Token{NEQ, "!="},
		Token{INT,"4"},
		Token{RPARENTHESES,")"},
		Token{LBRACE,"{"},
		Token{BREAK, "break"},
		Token{SEMICOLON, ";"},	
		Token{RBRACE,"}"},
		Token{RBRACE,"}"},
		//a = a & 2 | 4;
		Token{IDENT,"a"},
		Token{ASSIGN,"="},
		Token{IDENT,"a"},
		Token{AMPERSAND,"&"},
		Token{INT,"2"},
		Token{VERTICALBAR,"|"},
		Token{INT,"4"},
		Token{SEMICOLON,";"},
		Token{EOF,""},
	}
	nlexer4 := NewLexer(testString4)
	for key, token := range testResult4 {
		currtok := nlexer4.NextToken()
		if currtok.Type != token.Type {
			t.Fatalf("4-tests[%d] - tokentype wrong. expected=%q, got=%q", key, token.Type, currtok.Type)
		}
		if currtok.Literal != token.Literal {
			t.Fatalf("4-tests[%d] - literal wrong. expected=%q, got=%q", key, token.Literal, currtok.Literal)
		}
	}
}
