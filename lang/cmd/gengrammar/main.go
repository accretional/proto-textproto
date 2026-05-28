// Command gengrammar reads lang/textproto.ebnf and emits a
// pre-compiled GrammarDescriptor binary via gluon's v2 metaparser.ParseEBNF.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"google.golang.org/protobuf/proto"

	metaparserv2 "github.com/accretional/gluon/v2/metaparser"
)

func main() {
	ebnfPath := flag.String("ebnf", "lang/textproto.ebnf", "EBNF source")
	outPath := flag.String("out", "lang/textproto-grammar.pb", "grammar descriptor output (binary proto)")
	flag.Parse()

	src, err := os.ReadFile(*ebnfPath)
	if err != nil {
		log.Fatalf("read ebnf %s: %v", *ebnfPath, err)
	}
	doc := metaparserv2.WrapString(string(src))
	doc.Name = *ebnfPath

	gd, err := metaparserv2.ParseEBNF(doc)
	if err != nil {
		log.Fatalf("ParseEBNF: %v", err)
	}

	blob, err := proto.Marshal(gd)
	if err != nil {
		log.Fatalf("proto.Marshal: %v", err)
	}
	if err := os.WriteFile(*outPath, blob, 0o644); err != nil {
		log.Fatalf("write %s: %v", *outPath, err)
	}
	fmt.Printf("wrote %s (%d rules, %d bytes)\n", *outPath, len(gd.GetRules()), len(blob))
}
