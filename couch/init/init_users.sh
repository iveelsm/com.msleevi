#!/usr/bin/env bash

set -ex

function init_users() {
    while IFS="=" read -r key value; do
        echo ${key}
        echo ${value}
        curl -H 'Content-Type: application/json' \
            -X POST http://${COUCH_URL}:${COUCH_PORT}/${DATABASE}/ \
            -D $value
    done < <(jq -r ".users | to_entries | map(\"\(.value.id)=\(.value)\")|.[]" $1)
}

source ./environment.env
source ./resources/users/environment.env
init_users ./resources/users/users.json