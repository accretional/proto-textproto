# proto-textproto

A gluon v2 EBNF grammar for the Protocol Buffers Text Format (textproto), with proto generation and a parser/validator.

## Goal

Encode the textproto language structure into protobuf at a deep, language-spec level. Textproto is used as a config language to define RPCs — the pipeline is: textproto → AST → transform → target format (OpenAPI, proto, Go stubs, etc.).

The base production rule is `Message`, consisting of zero or more `Field` entries. Each `Field` is either a `ScalarField` (name `:` value) or a `MessageField` (name with nested `BraceMessage` or `AngleMessage`). The grammar covers the full textproto spec: string concatenation, extension fields `[com.example.ext]`, Any fields `[type.googleapis.com/Type]`, lists, all numeric formats, and escape sequences.

## Approach

The grammar and proto messages are *mechanically derived* from the textproto spec — not hand-maintained. `lang/metaparser` drives a gluon v2 pipeline (`ParseEBNF` → `GrammarToAST` → `Compile`) that turns `lang/textproto.ebnf` into a `FileDescriptorProto`; protoc emits the Go types.

`lang/cmd/gengrammar` pre-compiles the EBNF into a binary `GrammarDescriptor` proto (`lang/textproto-grammar.pb`) so the parser doesn't re-parse the EBNF at runtime. `cmd/parse` loads this pre-compiled grammar and uses `metaparser.ParseCST` to validate textproto files.

Comment handling: textproto uses `#` for line comments. Gluon handles `//`, `/* */`, `(* *)` but NOT `#`. The parser strips `#` comments via a string-aware preprocessor before calling `ParseCST`.

## Directory Layout

- `lang/` — `textproto.ebnf`, embedded EBNF + pre-compiled grammar binary
- `lang/metaparser/` — compile pipeline (`Compile`, `SerializeProto`, `Grammar`)
- `lang/cmd/gengrammar/` — EBNF → binary GrammarDescriptor generator
- `cmd/parse/` — textproto file validator using pre-compiled grammar + `ParseCST`
- `proto/textproto.proto` — generated bundled proto (112 messages)
- `proto/unicode/` — unicode proto files (copied from gluon)
- `proto/gen/` — generated Go types
- `examples/valid/` — valid textproto test files (27 files)
- `examples/invalid/` — invalid textproto test files (8 files)
- `examples/external/` — textproto files copied from sibling projects (11 files)
- `docs/` — build progress and spec notes

## Build Discipline

- **NEVER build/test/run code outside of `setup.sh`, `build.sh`, `test.sh`, `LET_IT_RIP.sh`.** Add any new command to the relevant script.
- **NEVER commit or push without running these.**
- Scripts are idempotent and chained: `build.sh` runs `setup.sh`, `test.sh` runs `build.sh`, `LET_IT_RIP.sh` runs `test.sh`.
- Proto files and Go stubs are NOT regenerated during build or test.
- `build.sh` regenerates the pre-compiled grammar binary from the EBNF.

## Tooling

- Go 1.25.
- Gluon descriptors: `github.com/accretional/gluon` — reuse its proto types rather than reinventing. Sibling at `../gluon`.
- Grammar convention: uppercase rules (e.g., `Message`, `Field`) are syntactic (whitespace auto-skipped); lowercase rules (e.g., `ident`, `string_literal`) are lexical (no whitespace skipping).

## References

- textproto spec: `temp/textproto_spec.html`
- gluon: https://github.com/accretional/gluon
