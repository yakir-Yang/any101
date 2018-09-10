#! /bin/bash

go run main.go

sleep 1

echo "==> $ curl -X GET http://localhost:8001/test.go"
curl -X GET http://localhost:8001/test.go
