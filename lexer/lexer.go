package lexer

import(
	_"fmt"
)

type Status int
const(
	LEXICALERROR = iota
	IN_DEC_LEX
	IN_FRAC_LEX
	IN_HEX_LEX
	IN_OCT_LEX
	IN_BIN_LEX
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
			if isNumber(lexer.CurrentChar){//deal with number
				numberLiteral,status := lexer.readNumber()
				if status == IN_DEC_LEX{
					token = newToken(INT,numberLiteral)
				}else if status == IN_FRAC_LEX{
					token = newToken(FLOAT,numberLiteral)
				}else if status == IN_HEX_LEX{
					token = newToken(HEX,numberLiteral)
				}else if status == IN_OCT_LEX{
					token = newToken(OCT,numberLiteral)
				}else if status == IN_BIN_LEX{
					token = newToken(BIN,numberLiteral)
				}else if status == LEXICALERROR{
					token = newToken(ILLEGAL,numberLiteral)
				}
			}else if isAlpha(lexer.CurrentChar){//deal with identify and key word
				word := lexer.readWord()
				tokenType,ok := findKeywords(word)
				if ok{//is keyword
					token = newToken(tokenType,word)
				}else{//is identify
					token = newToken(IDENT,word)
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
	for isAlpha(lexer.CurrentChar) || isNumber(lexer.CurrentChar){
		word = word + string(lexer.CurrentChar)
		if isAlpha(lexer.peepNextChar()) || isNumber(lexer.peepNextChar()){
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
	if lexer.CurrentChar == '0'{//0,0x,0o,0b(integers only)
		switch lexer.peepNextChar(){
		case 'b':
			status = IN_BIN_LEX
			numberLiteral = "0b"
			lexer.nextChar()
			lexer.nextChar()
			for lexer.CurrentChar == '0' || lexer.CurrentChar == '1'{
				numberLiteral = numberLiteral + string(lexer.CurrentChar)
				if isNumber(lexer.peepNextChar()){
					if lexer.peepNextChar() == '0' || lexer.peepNextChar()=='1'{
						lexer.nextChar()
					}else{
						status = LEXICALERROR
						break
					}
				}else if isAlpha(lexer.peepNextChar()){
					status = LEXICALERROR
					break
				}else{
					break
				}
			}
		case 'o':
			status = IN_OCT_LEX
			numberLiteral = "0o"
			lexer.nextChar()
			lexer.nextChar()
			for lexer.CurrentChar >= '0' && lexer.CurrentChar <= '7'{
				numberLiteral = numberLiteral + string(lexer.CurrentChar)
				if isNumber(lexer.peepNextChar()){
					if lexer.peepNextChar() >= '0' && lexer.peepNextChar()<='7'{
						lexer.nextChar()
					}else{
						status = LEXICALERROR
						break
					}
				}else if isAlpha(lexer.peepNextChar()){
					status = LEXICALERROR
					break
				}else{
					break
				}
			}
		case 'x':
			status = IN_HEX_LEX
			numberLiteral = "0x"
			lexer.nextChar()
			lexer.nextChar()
			for isNumber(lexer.CurrentChar) || (lexer.CurrentChar >= 'a' || lexer.CurrentChar <= 'f'){
				numberLiteral = numberLiteral + string(lexer.CurrentChar)
				if isNumber(lexer.peepNextChar()){
					lexer.nextChar()
				}else if isAlpha(lexer.peepNextChar()){
					if lexer.peepNextChar() >= 'a' && lexer.peepNextChar() <= 'f'{
						lexer.nextChar()
					}else{
						status = LEXICALERROR
						break
					}
				}else{
					break
				}
			}
		case '.':
			status = IN_FRAC_LEX
			numberLiteral = "0."
			lexer.nextChar()
			lexer.nextChar()
			for isNumber(lexer.CurrentChar){
				numberLiteral = numberLiteral + string(lexer.CurrentChar)
				if isNumber(lexer.peepNextChar()){
					lexer.nextChar()
				}else if isAlpha(lexer.peepNextChar()){
					status = LEXICALERROR
					break
				}else{
					break
				}
			}
		default:
			if isAlpha(lexer.peepNextChar()) || isNumber(lexer.peepNextChar()){
				status = LEXICALERROR
			}else{
				numberLiteral = numberLiteral + string(lexer.CurrentChar)
				status = IN_DEC_LEX
			}
		}
	}else{//decimal integers and float
		status = IN_DEC_LEX
		for isNumber(lexer.CurrentChar) || lexer.CurrentChar == '.'{
			numberLiteral = numberLiteral + string(lexer.CurrentChar)
			if isNumber(lexer.peepNextChar()){
				lexer.nextChar()
			}else if lexer.peepNextChar() == '.'{
				if status == IN_DEC_LEX{
					status = IN_FRAC_LEX
					lexer.nextChar()
				}else{
					status = LEXICALERROR
					break
				}
			}else if isAlpha(lexer.peepNextChar()){
				status = LEXICALERROR
				break
			}else{
				break
			}
		}
	}
	return
}

func newToken(tokenType TokenType,literal string)(token Token){
	token.Type = tokenType
	token.Literal = literal
	return
}

//a-z,A-Z,_
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