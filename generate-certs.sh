#! /bin/bash

mkdir certs
base=./certs

# CAs
step certificate create root-ca $base/root-ca.crt $base/root-ca.key --profile root-ca --insecure --no-password
step certificate create server-ca $base/server-ca.crt $base/server-ca.key --profile intermediate-ca --insecure --no-password \
  --ca $base/root-ca.crt --ca-key $base/root-ca.key
step certificate create client-ca $base/client-ca.crt $base/client-ca.key --profile intermediate-ca --insecure --no-password \
  --ca $base/root-ca.crt --ca-key $base/root-ca.key

# Server cert
step certificate create mtls.fbi.com $base/server.crt $base/server.key --profile leaf \
  --ca $base/server-ca.crt --ca-key $base/server-ca.key --insecure --no-password --san localhost --san mtls.fbi.com

# Client cert 1
step certificate create stephan $base/client-stephan.crt $base/client-stephan.key --profile leaf \
  --ca $base/client-ca.crt --ca-key $base/client-ca.key --insecure --no-password

# Client cert 2
step certificate create nick $base/client-nick.crt $base/client-nick.key --profile leaf \
  --ca $base/client-ca.crt --ca-key $base/client-ca.key --insecure --no-password
