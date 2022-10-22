# Using the minimalist OS image for Golang enviroment.
FROM golang:1.19-alpine

# Setting our systems working directory.
WORKDIR /app

# Copying all files from present directory to system /app directory as we already set our working directory.
COPY . .

# Running commands to download module packages and building a exectuable of our project.
RUN go mod download
RUN go build cmd/webapp/*.go

# Exposing our connection at Port 8080
EXPOSE 8080

# A command for running the executable file.
CMD ./main
