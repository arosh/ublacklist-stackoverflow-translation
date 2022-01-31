#!/bin/sh
set -u

for FILE in testdata/*.yml; do
  echo "Run test by ${FILE}" >&2
  go clean -testcache
  YAML="${FILE}" go test ./...
  if [ $? = 0 ]; then
    echo "The test was unexpectedly successful" >&2
    exit 1
  fi
done
