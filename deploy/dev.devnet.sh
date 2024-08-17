#!/bin/bash

basedir=$(
    cd $(dirname "$0") || exit
    pwd
)

echo "${basedir}"

project_path=${basedir}/../
bin=${project_path}/build/bdjuno

#GENESIS_PATH=/Users/lcq/Code/evmos/evmos/build/nodes/node0/evmosd/config/genesis.json

cmd=$1

case ${cmd} in
build)
    echo "===== build ===="

    echo "===== end ===="
    ;;
init)
    echo "===== init ===="

    echo "===== end ===="
    ;;
start)
    echo "===== start ===="

    echo "===== end ===="
    ;;
stop)
    echo "===== stop ===="

    echo "===== end ===="
    ;;
reset)
    echo "===== reset ===="

    killall bdjuno

    cd "${project_path}" || exit
    make build

    cd "${project_path}"/deploy || exit
    docker compose -f docker-compose.devnet.yaml down
    rm -rf ./data
    docker compose -f docker-compose.devnet.yaml up -d

    echo "wait 45s for graphql engine start..."
    for ((i = 20; i > 0; i -= 3)); do
        echo "please wait ${i}s..."
        sleep 3
    done

    echo "Importing the Hasura metadata..."
    cd "$project_path"/hasura || exit
    hasura metadata apply --endpoint http://localhost:9090 --admin-secret mechain

    echo "Initializing the configuration..."
    cd "$project_path"/deploy || exit
    cp config.devnet.yaml ./data/config.yaml
    ${bin} parse genesis-file --genesis-file-path ./genesis.json --home ./data

    echo "run BDjuno...."
    nohup "${bin}" start --home ./data >./data/bdjuno.log 2>&1 &

    echo "===== end ===="
    ;;
*)
    echo "Usage: dev.sh build | init | reset | start | stop"
    ;;
esac
