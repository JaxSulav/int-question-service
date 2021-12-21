protoc -I=./ --go_out=. --go_opt=paths=import --go-grpc_out=. --go-grpc_opt=paths=import --proto_path=protos/ \
  protos/main.proto \
  protos/datetime.proto \
  protos/question.proto \
  protos/questionset.proto \
  protos/set.proto \
  protos/type.proto
