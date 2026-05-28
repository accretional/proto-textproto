package main

import (
	"fmt"
	"os"

	"github.com/accretional/proto-textproto/lang/metaparser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: ebnf2proto <output.proto>\n")
		os.Exit(1)
	}
	outPath := os.Args[1]

	result, err := metaparser.Compile(metaparser.EBNF())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Compile: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Parsed %d rules from EBNF\n", result.RuleCount)
	fmt.Printf("AST root has %d children\n", result.ASTChildCount)
	fmt.Printf("Generated proto with %d messages, %d dependencies\n",
		len(result.FileDescriptor.GetMessageType()), len(result.FileDescriptor.GetDependency()))

	protoText := metaparser.SerializeProto(result.FileDescriptor)

	if err := os.WriteFile(outPath, []byte(protoText), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "write %s: %v\n", outPath, err)
		os.Exit(1)
	}
	fmt.Printf("Wrote %s\n", outPath)
}
