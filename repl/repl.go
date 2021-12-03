package repl

import (
	"bufio"
	"fmt"
	"io"

	"monkey/lexer"
	"monkey/parser"
)

const PROMPT = ">> "

const HEADER = `OH NO.....ERRORS.`

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		if line == "exit" {
			break
		}

		lex := lexer.New(line)
		p := parser.New(lex)
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, HEADER+"\n")
	io.WriteString(out, "parser errors:\n")
	for _, message := range errors {
		io.WriteString(out, "\t"+message+"\n")
	}
}
