package repl

import(
	"fmt"
    "os"
	"bufio"
	"creep/lexer"
)

func Entrance(){
	fmt.Println("Welcome to the creep language!")
	for{
		fmt.Print("->")
		reader := bufio.NewReader(os.Stdin)
		stringToDeal, _ := reader.ReadString('\n')
		currLexer := lexer.NewLexer(stringToDeal)
		for{
			currtok := currLexer.NextToken()
			fmt.Println("Token Type:",currtok.Type," Token Literal:",currtok.Literal)
			if currtok.Type == lexer.EOF{
				break
			}
		}
	}
}