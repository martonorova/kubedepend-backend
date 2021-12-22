#!/bin/bash

# download and install CompileDaemon without modifying go.mod
go install github.com/githubnemo/CompileDaemon@latest

CompileDaemon --build="make build" --command="make run" --include=".env" --include="Makefile"