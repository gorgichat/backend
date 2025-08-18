#!/bin/sh

set -e

export operation="${1:-build}"

echo "${0##*/}: operation started..."

case "${1}" in
	"clean"|"c")
		export operation="clean"
		[ -d "build" ] && rm -rf "build"
	;;
	*)
		[ ! -d "build" ] && mkdir "build"
		go build -o "build/gorgi" "src/main.go"
	;;
esac

echo "${0##*/}: ${operation}: operation completed."
