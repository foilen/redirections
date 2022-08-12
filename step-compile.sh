#!/bin/bash

set -e

RUN_PATH="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $RUN_PATH

echo ----[ Compile ]----
CGO_ENABLED=0 go build -ldflags "-s -w -extldflags '-static'" -o build/bin/redirections ./main
