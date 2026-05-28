// Package lang embeds the textproto EBNF grammar and pre-compiled
// GrammarDescriptor binary.
package lang

import _ "embed"

//go:embed textproto.ebnf
var EBNF []byte

//go:embed textproto-grammar.pb
var Grammar []byte
