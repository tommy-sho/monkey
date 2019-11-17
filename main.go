package main

import (
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/tommy-sho/monkey/repl"
)

func main() {
	userName, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Hello %s!, This is the monky languege\n", userName)
	repl.Start(os.Stdin, os.Stdout)
}
