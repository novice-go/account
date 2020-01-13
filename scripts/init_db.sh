#!/usr/bin/env bash
echo "init_db..."
cd ./cmd/db-migrate || exit
go run main.go drop
go run main.go create
go run main.go migrate
