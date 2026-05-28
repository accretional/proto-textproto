// Package lang embeds the textproto EBNF grammar.
package lang

import _ "embed"

//go:embed textproto.ebnf
var EBNF []byte
