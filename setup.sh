#!/usr/bin/env bash
set -euo pipefail

echo "=== proto-textproto setup ==="

# Check Go
if ! command -v go &>/dev/null; then
  echo "ERROR: Go is not installed. Install Go 1.21+ first."
  exit 1
fi
echo "Go: $(go version)"

# Check protoc
if ! command -v protoc &>/dev/null; then
  echo "ERROR: protoc is not installed. Install protobuf compiler first."
  echo "  brew install protobuf  # macOS"
  exit 1
fi
echo "protoc: $(protoc --version)"

# Install protoc-gen-go if missing
if ! command -v protoc-gen-go &>/dev/null; then
  echo "Installing protoc-gen-go..."
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
else
  echo "protoc-gen-go: $(which protoc-gen-go)"
fi

# Check that gluon is available locally
if [ ! -d "../gluon/v2" ]; then
  echo "ERROR: ../gluon/v2 not found. Clone gluon next to this repo."
  exit 1
fi
echo "gluon v2: found at ../gluon/v2"

# Tidy module
echo "Running go mod tidy..."
go mod tidy

echo "=== setup complete ==="
