#!/bin/bash

go mod download
go build -o reservation.webapp cmd/webapp/*.go
if [[ $? == 0 ]]; then
  echo "To exit: Control + C"
  ./reservation.webapp
else
    echo "Failed to Execute"
fi