#!/usr/bin bash
set -e

GOOS=windows go build -o bin/bm.exe
GOOS=darwin go build -o bin/bm
