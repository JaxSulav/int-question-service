
RUN apt install -y protobuf-compiler
RUN export PATH="$PATH:$(go env GOPATH)/bin"
RUN go mod tidy -go=1.16 && go mod tidy -go=1.17
RUN go install \
        github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
        github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
        google.golang.org/protobuf/cmd/protoc-gen-go \
        google.golang.org/grpc/cmd/protoc-gen-go-grpc

RUN cd protos && buf generate --error-format=json 