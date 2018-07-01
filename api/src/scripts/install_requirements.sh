#!/usr/bin/env bash


install_packages() {
  while IFS='' read -r line || [[ -n "${line}" ]]; do
    echo -n "Installing $line... "
    go get $line >> /dev/null
    echo "$line installed"
  done < $1
}

echo "GOPATH: $(go env GOPATH)"
install_packages ./requirements.txt
