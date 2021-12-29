#!/bin/bash

# download and install CompileDaemon without modifying go.mod
go install github.com/githubnemo/CompileDaemon@latest

# use ./main for command instead of "make run", make run exits after starting the application, and CompileDaemon will not have a reference to the application restart it on change
CompileDaemon --build="make build" --command="./main" --include=".env" --include="Makefile" --include="config.yaml"