#!/bin/bash

set -x 

go mod tidy

go build -o ./main  

docker build -t websocket-golang-test:v0.1  .
