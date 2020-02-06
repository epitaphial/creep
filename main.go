package main

import(
	"os"
	_"fmt"
	"creep/repl"
)

func main(){
	if len(os.Args) == 1{
		//enter the repl mode
		repl.Entrance()
	}else{
		//interpret by file
	}
}