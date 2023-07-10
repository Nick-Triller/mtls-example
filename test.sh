#! /bin/bash

# Kill process group including server that runs in background
trap "trap - SIGTERM && kill -- -$$" SIGINT SIGTERM EXIT

# Start server in background
go run cmd/server.go &

# Wait for server to start
sleep 1

# Run client
go run cmd/client.go --clientCertFile ./certs/client-stephan.crt --clientKeyFile ./certs/client-stephan.key
go run cmd/client.go --clientCertFile ./certs/client-nick.crt --clientKeyFile ./certs/client-nick.key
