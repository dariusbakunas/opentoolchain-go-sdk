#!/bin/bash

openapi-sdkgen.sh generate -g ibm-go -i docs/openapi.yaml -o .
go fmt ./...
patch opentoolchainv1/open_toolchain_v1.go < opentoolchainv1/patch
