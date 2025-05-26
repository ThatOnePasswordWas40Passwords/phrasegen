#!/usr/bin/env bash

set -euo pipefail

os=$(go env GOOS)
if [[ "${os:-}" == "" ]]; then
    echo "ERR - missing 'go env GOOS'?"
    exit 1
fi

arch=$(go env GOARCH)
if [[ "${arch:-}" == "" ]]; then
    echo "ERR - missing 'go env GOARCH'?"
    exit 1
fi

platform="${os}/${arch}"
if [[ "$(go tool dist list)" != *"${platform}"* ]]; then
    echo "ERR - unsupported platform '${platform}'; refer to 'go tool dist list'"
    exit 1
fi

./binaries/phrasegen.${os}.${arch} -i "Some sample input!" -no-strip -join-str "-" -size 2