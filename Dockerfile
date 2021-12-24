FROM golang:1.17-bullseye

RUN cat /etc/issue
RUN apt-get update

ADD . /usr/src/app

WORKDIR /usr/src/app
RUN apt-get -y install vim
RUN apt-get install -y protobuf-compiler

RUN export PATH="$PATH:$(go env GOPATH)/bin"

RUN go mod tidy -go=1.17
RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway && go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 && go install google.golang.org/protobuf/cmd/protoc-gen-go && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

RUN apt install -y build-essential

EXPOSE 8000
EXPOSE 50051