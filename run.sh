#!/bin/bash

echo "You need to run and setup Postgresql db locally for this"

go mod download
go build cmd/webapp/*.go
if [[ $? == 0 ]]; then
  echo "To exit: Control + C"
  ./main
else
    echo "Failed to Execute"
fi