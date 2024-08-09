#!/bin/bash

basedir=$(
	cd $(dirname "$0") || exit
	pwd
)

echo ${basedir}

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

    cd ${project_path}
    make build

    cd ${project_path}/deploy
    docker-compose -f docker-compose.yaml down
    rm -rf ./data
    docker-compose -f docker-compose.yaml up -d

    echo "wait 45s for graphql engine start..."
    for ((i=45; i>0; i-=3))
    do
        echo "please wait ${i}s..."
        sleep 3
    done

    cd $project_path/hasura
    hasura metadata apply --endpoint http://localhost:8080 --admin-secret 123456

    cd $project_path/deploy
    cp config.yaml ./data
    ${bin} parse genesis-file --genesis-file-path ./genesis.json --home ./data
    #nohup ${bin} start --home ./data >./bdjuno.log 2>&1 &
    ${bin} start --home ./data

    echo "===== end ===="
    ;;
*)
    echo "Usage: dev.sh build | init | reset | start | stop"
    ;;
esac
