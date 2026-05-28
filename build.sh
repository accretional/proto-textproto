#!/usr/bin/env bash
set -euo pipefail

echo "=== proto-textproto build ==="

# Run setup (idempotent)
bash setup.sh

# Check that proto files exist (they should be pre-generated, not regenerated)
if [ ! -f proto/textproto.proto ]; then
  echo "ERROR: proto/textproto.proto not found."
  echo "Generate proto files first using lang/metaparser."
  exit 1
fi

if [ ! -f proto/gen/textproto.pb.go ]; then
  echo "ERROR: proto/gen/textproto.pb.go not found. Compile proto files first:"
  echo "  protoc --go_out=. --go_opt=module=github.com/accretional/proto-textproto -I proto proto/textproto.proto proto/unicode/utf_8.proto proto/unicode/ascii.proto"
  exit 1
fi

# Build binaries
mkdir -p bin

echo "Building cmd/parse..."
go build -o bin/parse ./cmd/parse

echo "=== build complete ==="
