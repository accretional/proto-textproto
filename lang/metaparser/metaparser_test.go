package metaparser

import (
	"strings"
	"testing"
)

func TestEBNF(t *testing.T) {
	src := EBNF()
	if len(src) == 0 {
		t.Fatal("EBNF source is empty")
	}
	if !strings.Contains(string(src), "Message") {
		t.Error("EBNF source missing Message rule")
	}
}

func TestCompile(t *testing.T) {
	result, err := Compile(EBNF())
	if err != nil {
		t.Fatalf("Compile: %v", err)
	}
	if result.RuleCount != 84 {
		t.Errorf("expected 84 rules, got %d", result.RuleCount)
	}
	if result.ASTChildCount != 84 {
		t.Errorf("expected 84 AST children, got %d", result.ASTChildCount)
	}
	fdp := result.FileDescriptor
	if fdp.GetPackage() != "textproto" {
		t.Errorf("expected package textproto, got %s", fdp.GetPackage())
	}
	if fdp.GetName() != "textproto.proto" {
		t.Errorf("expected file name textproto.proto, got %s", fdp.GetName())
	}
	if got := len(fdp.GetMessageType()); got != 112 {
		t.Errorf("expected 112 messages, got %d", got)
	}
	deps := fdp.GetDependency()
	if len(deps) != 1 || deps[0] != "unicode/utf_8.proto" {
		t.Errorf("expected [unicode/utf_8.proto] dependency, got %v", deps)
	}
}

func TestCompileExpectedMessages(t *testing.T) {
	result, err := Compile(EBNF())
	if err != nil {
		t.Fatalf("Compile: %v", err)
	}
	fdp := result.FileDescriptor

	expected := []string{
		"Message", "Field", "ScalarField", "MessageField",
		"ScalarValue", "MessageValue", "BraceMessage", "AngleMessage",
		"StringVal", "FloatVal", "Ident",
	}

	msgNames := make(map[string]bool)
	for _, msg := range fdp.GetMessageType() {
		msgNames[msg.GetName()] = true
	}

	for _, name := range expected {
		if !msgNames[name] {
			t.Errorf("missing expected message %s", name)
		}
	}
}

func TestGrammar(t *testing.T) {
	gd, err := Grammar()
	if err != nil {
		t.Fatalf("Grammar: %v", err)
	}
	if got := len(gd.GetRules()); got != 84 {
		t.Errorf("expected 84 rules, got %d", got)
	}
	// Verify key rules exist
	ruleNames := make(map[string]bool)
	for _, r := range gd.GetRules() {
		ruleNames[r.GetName()] = true
	}
	for _, name := range []string{"Message", "Field", "ScalarValue", "ident", "string_literal"} {
		if !ruleNames[name] {
			t.Errorf("missing expected rule %s", name)
		}
	}
}

func TestSerializeProto(t *testing.T) {
	result, err := Compile(EBNF())
	if err != nil {
		t.Fatalf("Compile: %v", err)
	}

	output := SerializeProto(result.FileDescriptor)

	checks := []string{
		`syntax = "proto3"`,
		"package textproto;",
		`import "unicode/utf_8.proto"`,
		"message Message {",
	}
	for _, check := range checks {
		if !strings.Contains(output, check) {
			t.Errorf("serialized output missing %q", check)
		}
	}
}
