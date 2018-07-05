#!/usr/bin/env bash

set -x

function init_dbs() {
    while IFS='' read -r line || [[ -n "$line" ]]; do
        curl -X PUT http://${COUCH_URL}:${COUCH_PORT}/${line}
    done < "$1"
}

source ./environment.env
init_dbs ./resources/dbs/dbs.txt