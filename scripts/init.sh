#!/bin/sh

echo 'create docker network'

docker network create address-networks

echo 'create docker volume'

docker volume create --name=address-go