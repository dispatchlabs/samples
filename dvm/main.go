package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/dispatchlabs/dvm/test"
	"C"
)

var (
	batch       = flag.Bool("b", false, "batch (non-interactive) mode")
	optimized   = flag.Bool("opt", false, "add some optimization passes")
	printTokens = flag.Bool("tok", true, "print tokens")
	printAst    = flag.Bool("ast", false, "print abstract syntax tree")
	printLLVMIR = flag.Bool("llvm", false, "print LLVM generated code")
)

func main() {
	flag.Parse()
	if *optimized {
		test.Optimize()
	}

	lex := test.Lex()
	tokens := lex.Tokens()
	if *printTokens {
		tokens = test.DumpTokens(lex.Tokens())
	}

	// add files for the lexer to lex
	go func() {
		// command line filenames
		for _, fn := range flag.Args() {
			f, err := os.Open(fn)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(-1)
			}
			lex.Add(f)
		}

		// stdin
		if !*batch {
			lex.Add(os.Stdin)
		}
		lex.Done()
	}()

	nodes := test.Parse(tokens)
	nodesForExec := nodes
	if *printAst {
		nodesForExec = test.DumpTree(nodes)
	}

	test.Exec(nodesForExec, *printLLVMIR)
}
