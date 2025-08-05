#!/bin/sh

# Load environment variables from .env file
if [ -f .env ]; then
    export $(grep -v '^#' .env | xargs)
fi

# check port is available
if [ -n "$(lsof -t -i TCP:3000)" ]; then
    kill $(lsof -t -i TCP:3000)
fi

go run ./main.go
