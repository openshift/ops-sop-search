FROM golang:latest

RUN apt -y install git

# Authorize SSH Host
 RUN mkdir -p /.ssh && \
     chmod 0777 /.ssh && \
     ssh-keyscan github.com > /.ssh/known_hosts

# RUN eval `ssh-agent -s` \
#     ssh-add /root/.ssh/id_rsa

WORKDIR /build

RUN chmod 0777 /build

COPY . .

COPY deploy/ssh_config.template /etc/ssh/ssh_config

RUN go mod download 

RUN go build cmd/main.go

#RUN useradd -r -m -u 1000 appuser

#USER appuser

RUN chmod +x ./main

ENTRYPOINT ./main