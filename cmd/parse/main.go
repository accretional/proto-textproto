package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/accretional/proto-textproto/lang"
	pb "github.com/accretional/gluon/v2/pb"

	"github.com/accretional/gluon/v2/metaparser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: parse <input.textproto> [-v]\n")
		os.Exit(1)
	}
	inputPath := os.Args[1]

	// Use embedded grammar
	grammarDoc := metaparser.WrapString(string(lang.EBNF))
	gd, err := metaparser.ParseEBNF(grammarDoc)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ParseEBNF: %v\n", err)
		os.Exit(1)
	}

	// Load input
	inputSrc, err := os.ReadFile(inputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "read input %s: %v\n", inputPath, err)
		os.Exit(1)
	}

	// Strip # comments (textproto uses # for line comments, but
	// gluon's parser only handles //, /* */, and (* *) comments)
	cleaned := stripComments(string(inputSrc))

	// Parse input against grammar
	srcDoc := metaparser.WrapString(cleaned)
	ast, err := metaparser.ParseCST(&pb.CstRequest{
		Grammar:  gd,
		Document: srcDoc,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "FAIL: %s\n  %v\n", inputPath, err)
		os.Exit(1)
	}

	fmt.Printf("PASS: %s\n", inputPath)
	if len(os.Args) > 2 && os.Args[2] == "-v" {
		printAST(ast.GetRoot(), 0)
	}
}

// stripComments removes # line comments from textproto source,
// being careful not to strip inside single- or double-quoted strings.
func stripComments(src string) string {
	var b strings.Builder
	b.Grow(len(src))

	inDouble := false
	inSingle := false
	inComment := false
	escaped := false

	for i := 0; i < len(src); i++ {
		ch := src[i]

		if escaped {
			if !inComment {
				b.WriteByte(ch)
			}
			escaped = false
			continue
		}

		if ch == '\\' && (inDouble || inSingle) {
			escaped = true
			b.WriteByte(ch)
			continue
		}

		if inComment {
			if ch == '\n' {
				inComment = false
				b.WriteByte(ch) // preserve the newline
			}
			continue
		}

		if ch == '"' && !inSingle {
			inDouble = !inDouble
			b.WriteByte(ch)
			continue
		}

		if ch == '\'' && !inDouble {
			inSingle = !inSingle
			b.WriteByte(ch)
			continue
		}

		if ch == '#' && !inDouble && !inSingle {
			inComment = true
			continue
		}

		b.WriteByte(ch)
	}

	return b.String()
}

func printAST(node *pb.ASTNode, depth int) {
	if node == nil {
		return
	}
	indent := strings.Repeat("  ", depth)
	if node.GetValue() != "" {
		fmt.Printf("%s%s: %q\n", indent, node.GetKind(), node.GetValue())
	} else {
		fmt.Printf("%s%s\n", indent, node.GetKind())
	}
	for _, child := range node.GetChildren() {
		printAST(child, depth+1)
	}
}
