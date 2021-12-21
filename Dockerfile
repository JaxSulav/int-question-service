
RUN apt install -y protobuf-compiler

RUN go mod tidy
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u google.golang.org/grpc
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
RUN export PATH="$PATH:$(go env GOPATH)/bin"

RUN ./generatepb.sh