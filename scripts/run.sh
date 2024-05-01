#!/bin/bash

#starting mongodb
echo "starting mongodb server..."
systemctl start mongod

echo "build and runngin go file..."
#build and run go file
go run ./cmd/api/main.go