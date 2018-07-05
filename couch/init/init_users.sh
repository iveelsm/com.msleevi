#!/usr/bin/env bash

set -x

function init_users() {
    res=$(cat $1 | jq '.')
    echo $res
    for row in $(jq '.[]' $1); do
        echo ${row}
        echo ${row} | base64 --decode | jq -r ${1}
        # curl -H 'Content-Type: application/json' \
        #     -X PUT http://${COUCH_URL}:${COUCH_PORT}/${DATABASE} \
        #     -D ${row}
    done
}

source ./environment.env
source ./resources/users/environment.env
init_users ./resources/users/users.json