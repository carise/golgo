#!/bin/bash
GOPATH=$(pwd)

export GOPATH
export PATH=$PATH:$(go env GOPATH)/bin
