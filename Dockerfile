# Using the minimalist OS image for Golang enviroment.
FROM golang:1.16-alpine

# Setting our systems working directory.
WORKDIR /app

# Copying all files from present directory to system /app directory as we already set our working directory.
COPY . .

# Running commands to download module packages and building a exectuable of our project.
RUN go mod download
RUN go get github.com/jackc/pgx/v5
RUN go build cmd/webapp/*.go

# Exposing our connection at Port 8080
#ARG PORT

# A command for running the executable file.
CMD ./main
