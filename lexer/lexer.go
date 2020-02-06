package lexer

import(
	_"fmt"
)

type Status int
const(
	LEXICALERROR = iota
	IN_INT_LEX
	IN_FRAC_LEX
)

type Lexer struct{
	StringToDeal string
	CurrentPosition int
	NextPosition int
	CurrentChar byte
}

func NewLexer(strTodeal string)(lexer *Lexer){
	lexer = &Lexer{StringToDeal:strTodeal,CurrentPosition:-1,NextPosition:0}
	return
}

func (lexer *Lexer)NextToken()(token Token){
	lexer.nextChar()
	lexer.eatSpace()
	switch lexer.CurrentChar{
		case '=':
			if lexer.peepNextChar() == '='{
				tmpStr := string(lexer.CurrentChar)
				lexer.nextChar()
				tmpStr = tmpStr + string(lexer.CurrentChar)
				token = newToken(EQ,tmpStr)
			}else{
				token = newToken(ASSIGN,string(lexer.CurrentChar))
			}
		case '+':
			token = newToken(PLUS,string(lexer.CurrentChar))
		case '-':
			token = newToken(MINUS,string(lexer.CurrentChar))
		case '*':
			token = newToken(ASTERISK,string(lexer.CurrentChar))
		case '/':
			token = newToken(SLASH,string(lexer.CurrentChar))
		case '!':
			if lexer.peepNextChar() == '='{
				tmpStr := string(lexer.CurrentChar)
				lexer.nextChar()
				tmpStr = tmpStr + string(lexer.CurrentChar)
				token = newToken(NEQ,tmpStr)
			}else{
				token = newToken(EXCLAMATION,string(lexer.CurrentChar))
			}
		case '>':
			if lexer.peepNextChar() == '='{
				tmpStr := string(lexer.CurrentChar)
				lexer.nextChar()
				tmpStr = tmpStr + string(lexer.CurrentChar)
				token = newToken(GTE,tmpStr)
			}else{
				token = newToken(GREATERTHAN,string(lexer.CurrentChar))
			}
		case '<':
			if lexer.peepNextChar() == '='{
				tmpStr := string(lexer.CurrentChar)
				lexer.nextChar()
				tmpStr = tmpStr + string(lexer.CurrentChar)
				token = newToken(LTE,tmpStr)
			}else{
				token = newToken(LESSTHAN,string(lexer.CurrentChar))
			}
		case '&':
			if lexer.peepNextChar() == '&'{
				tmpStr := string(lexer.CurrentChar)
				lexer.nextChar()
				tmpStr = tmpStr + string(lexer.CurrentChar)
				token = newToken(DAMPERSAND,tmpStr)
			}else{
				token = newToken(AMPERSAND,string(lexer.CurrentChar))
			}
		case '|':
			if lexer.peepNextChar() == '|'{
				tmpStr := string(lexer.CurrentChar)
				lexer.nextChar()
				tmpStr = tmpStr + string(lexer.CurrentChar)
				token = newToken(DVERTICALBAR,tmpStr)
			}else{
				token = newToken(VERTICALBAR,string(lexer.CurrentChar))
			}
		case ',':
			token = newToken(COMMA,string(lexer.CurrentChar))
		case ';':
			token = newToken(SEMICOLON,string(lexer.CurrentChar))
		case '{':
			token = newToken(LBRACE,string(lexer.CurrentChar))
		case '}':
			token = newToken(RBRACE,string(lexer.CurrentChar))
		case '[':
			token = newToken(LBRACKET,string(lexer.CurrentChar))
		case ']':
			token = newToken(RBRACKET,string(lexer.CurrentChar))
		case '(':
			token = newToken(LPARENTHESES,string(lexer.CurrentChar))
		case ')':
			token = newToken(RPARENTHESES,string(lexer.CurrentChar))
		case 0:
			token = newToken(EOF,"")
		default:
			if isVisibleChar(lexer.CurrentChar){
				//deal with number
				if isNumber(lexer.CurrentChar){
					numberLiteral,status := lexer.readNumber()
					if status == LEXICALERROR{
						token = newToken(ILLEGAL,string(lexer.CurrentChar))
					}else if status == IN_INT_LEX{
						token = newToken(INT,numberLiteral)
					}else if status == IN_FRAC_LEX{
						token = newToken(FLOAT,numberLiteral)
					}
				}else{//deal with identify and key word
					word := lexer.readWord()
					tokenType,ok := findKeywords(word)
					if ok{//is keyword
						token = newToken(tokenType,word)
					}else{//is identify
						token = newToken(IDENT,word)
					}
				}
			}else{
				token = newToken(ILLEGAL,string(lexer.CurrentChar))
			}
		}
	return
}

func (lexer * Lexer)nextChar(){
	lexer.CurrentPosition = lexer.NextPosition
	if lexer.CurrentPosition >= len(lexer.StringToDeal){
		lexer.CurrentChar = 0
	}else{
		lexer.CurrentChar = lexer.StringToDeal[lexer.CurrentPosition]
		lexer.NextPosition += 1
	}
}

func (lexer * Lexer)eatSpace(){
	for{
		if lexer.CurrentChar == ' ' || lexer.CurrentChar == '\n' || lexer.CurrentChar == '\t' || lexer.CurrentChar == '\r'{
			lexer.nextChar()
		}else{
			break
		}
	}
}

func (lexer *Lexer)readWord()(word string){
	word = ""
	for isAlpha(lexer.CurrentChar){
			word = word + string(lexer.CurrentChar)
			if isAlpha(lexer.peepNextChar()){
				lexer.nextChar()
			}else{
				break
			}
	}
	return
}

func (lexer *Lexer)peepNextChar()(nextChar byte){
	return lexer.StringToDeal[lexer.NextPosition]
}

func (lexer *Lexer)readNumber()(numberLiteral string,status Status){
	numberLiteral = ""
	status = IN_INT_LEX
	if lexer.CurrentChar == '0'{
		if isAlpha(lexer.peepNextChar()) || isNumber(lexer.peepNextChar()){
			status = LEXICALERROR
		}else if lexer.peepNextChar() == '.'{
			status = IN_FRAC_LEX
			lexer.nextChar()
			numberLiteral = numberLiteral + "0."
			numberLiteral = numberLiteral + lexer.readFrac()
			if isAlpha(lexer.peepNextChar()) || lexer.peepNextChar() == '.'{
				status = LEXICALERROR
			}
		}else{
			numberLiteral = numberLiteral + "0"
		}
	}else{
		numberLiteral = numberLiteral + lexer.readInt()
		if lexer.peepNextChar() == '.'{
			status = IN_FRAC_LEX
			lexer.nextChar()
			numberLiteral = numberLiteral + "."
			numberLiteral = numberLiteral + lexer.readFrac()
			if isAlpha(lexer.peepNextChar()) || lexer.peepNextChar() == '.'{
				status = LEXICALERROR
			}
		}else if isAlpha(lexer.peepNextChar()){
			status = LEXICALERROR
		}
	}
	return
}

func (lexer *Lexer)readInt()(intLiteral string){
	intLiteral = ""
	for isNumber(lexer.CurrentChar){
		intLiteral = intLiteral + string(lexer.CurrentChar)
		if isNumber(lexer.peepNextChar()){
			lexer.nextChar()
		}else{
			break
		}
	}
	return
}

func (lexer *Lexer)readFrac()(fracLiteral string){
	fracLiteral = ""
	lexer.nextChar()
	for isNumber(lexer.CurrentChar){
		fracLiteral = fracLiteral + string(lexer.CurrentChar)
		if isNumber(lexer.peepNextChar()){
			lexer.nextChar()
		}else{
			break
		}
	}
	return
}

func newToken(tokenType TokenType,literal string)(token Token){
	token.Type = tokenType
	token.Literal = literal
	return
}

func isVisibleChar(currentChar byte)(bool){
	return currentChar > 32 && currentChar < 127
}

func isAlpha(currentChar byte)(bool){
	return (currentChar >= 65 && currentChar <= 90)||(currentChar >= 97 && currentChar <= 122)||(currentChar == 95)
}

func isNumber(currentChar byte)(bool){
	return (currentChar >= 48 && currentChar <= 57)
}

func isSpace(currentChar byte)(bool){
	return currentChar == ' ' || currentChar == '\n' || currentChar == '\t' || currentChar == '\r'
}

func findKeywords(word string)(tokenType TokenType,isKeyword bool){
	tokenType,isKeyword = keywords[word]
	return
}