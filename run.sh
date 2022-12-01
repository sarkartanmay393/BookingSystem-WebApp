#!/bin/bash

echo "You need to run and setup Postgresql db locally for this"
echo "To automate this, you can use docker-compose"
useCompose=false
choice="Y"
echo "Do you want to use docker-compose? (Y/n)"
read choice
if [[ choice == "Y" ]]; then
    useCompose=true
fi
sleep 3

if [[ useCompose == true ]]; then
    echo "Starting docker-compose"
    docker-compose up -d
    sleep 3
fi

#go mod download
#go build cmd/webapp/*.go
#if [[ $? == 0 ]]; then
#  echo "To exit: Control + C"
#  ./main
#else
#    echo "Failed to Execute"
#fi