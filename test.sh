#!/usr/bin/env bash
set -euo pipefail

echo "=== proto-textproto test ==="

# Build first (idempotent)
bash build.sh

PASS=0
FAIL=0
ERRORS=""

echo ""
echo "--- Valid textproto files (should PASS) ---"
for textproto in examples/valid/*.textproto; do
  base=$(basename "$textproto" .textproto)
  if ./bin/parse "$textproto" 2>/dev/null; then
    echo "PASS: $base"
    PASS=$((PASS + 1))
  else
    echo "FAIL: $base (expected PASS but got rejection)"
    FAIL=$((FAIL + 1))
    ERRORS="${ERRORS}\n  ${base}: expected valid, got rejected"
  fi
done

echo ""
echo "--- Invalid textproto files (should FAIL) ---"
for textproto in examples/invalid/*.textproto; do
  base=$(basename "$textproto" .textproto)
  if ./bin/parse "$textproto" 2>/dev/null; then
    echo "FAIL: $base (expected rejection but got PASS)"
    FAIL=$((FAIL + 1))
    ERRORS="${ERRORS}\n  ${base}: expected invalid, got accepted"
  else
    echo "PASS: $base (correctly rejected)"
    PASS=$((PASS + 1))
  fi
done

echo ""
echo "Results: ${PASS} passed, ${FAIL} failed"

if [ "$FAIL" -gt 0 ]; then
  echo -e "Failures:${ERRORS}"
  exit 1
fi

if [ "$PASS" -eq 0 ]; then
  echo "WARNING: No tests ran!"
  exit 1
fi

echo "=== all tests passed ==="
