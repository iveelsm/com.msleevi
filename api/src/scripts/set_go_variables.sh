#!/usr/bin/env bash

set_go_path() {
  local dir=$(cd ../../ && pwd)
  export GOPATH=$dir
}

set_go_path
echo $GOPATH
export GOBIN=${GOPATH}/bin
