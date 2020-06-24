#!/bin/sh

# To filter out requests that returned an OK code
go run cmd/kubernetes-server/main.go --scheme=http --port 8001 | awk '{ if( $4 != 200 && $4 != 0) print $0 }'
