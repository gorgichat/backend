#!/bin/sh

set -e

export operation="${1:-build}"

case "${1}" in
	"clean"|"c")
		export operation="clean"
		[ -d "build" ] && rm -rf "build"
	;;
	*)
		[ ! -d "build" ] && mkdir "build"
		go build -o "build/darkcukka" "src/main.go"
	;;
esac
echo "${0##*/}: ${operation}: operation completed."
