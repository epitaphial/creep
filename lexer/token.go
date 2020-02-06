package lexer

const (
	//bases
	EOF     = "EOF"
	ILLEGAL = "ILLEGAL"
	IDENT   = "IDENT"

	//datas
	INT   = "INT"
	FLOAT = "FLOAT"

	//Single Operators
	ASSIGN      = "="
	PLUS        = "+"
	MINUS       = "-"
	ASTERISK    = "*"
	SLASH       = "/"
	EXCLAMATION = "!"
	GREATERTHAN = ">"
	LESSTHAN    = "<"
	AMPERSAND = "&"
	VERTICALBAR = "|"

	//Double Opertors
	EQ = "=="
	NEQ = "!="
	GTE = ">="
	LTE = "<="
	DAMPERSAND = "&&"
	DVERTICALBAR = "||"

	//Delimiters
	COMMA        = ","
	SEMICOLON    = ";"
	LBRACE       = "{"
	RBRACE       = "}"
	LBRACKET     = "["
	RBRACKET     = "]"
	LPARENTHESES = "("
	RPARENTHESES = ")"

	//keywords
	ORDER    = "ORDER"
	FUNCTION = "FUNCTION"
	RETURN = "RETURN"
	IF = "IF"
	FOR = "FOR"
	BREAK = "BREAK"
	ELSE = "ELSE"
	TRUE = "TRUE"
	FALSE = "FALSE"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fun": FUNCTION,
	"order": ORDER,
	"return":RETURN,
	"int":INT,
	"float":FLOAT,
	"if":IF,
	"else":ELSE,
	"for":FOR,
	"break":BREAK,
	"true":TRUE,
	"false":FALSE,
	}