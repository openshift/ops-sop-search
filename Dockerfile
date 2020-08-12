FROM golang:latest

RUN apt -y install git

# Authorize SSH Host
 RUN mkdir -p /.ssh && \
     chmod 0777 /.ssh && \
     ssh-keyscan github.com > /.ssh/known_hosts

WORKDIR /build

RUN chmod 0777 /build

COPY . .

RUN go mod download 

RUN go build cmd/main.go

RUN chmod +x ./main

ENTRYPOINT ./main