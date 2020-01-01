package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/tommy-sho/monkey/lexer"
	"github.com/tommy-sho/monkey/parser"
)

const PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParseError(out, p.Errors())
			continue
		}

		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}

func printParseError(out io.Writer, errors []string) {
	fmt.Println("----error-------")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
