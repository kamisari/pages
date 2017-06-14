#!/bin/sh
set -eu
# cds(change to script directory)
cd "$(dirname "$(readlink -f "$0")")"
go run ./genlinks.go
# EOF
