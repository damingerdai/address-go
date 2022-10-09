#!/bin/bash

sha=$(git rev-parse HEAD)

docker build -t address-go:${sha} .