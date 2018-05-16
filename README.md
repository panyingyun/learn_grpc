# learn_grpc
learn_grpc

Study protocol buffer v3(Golang)

1、 download compiler protoc
https://github.com/google/protobuf/releases/download/v3.5.1/protoc-3.5.1-win32.zip

2、install protocol buffer package
go get github.com/golang/protobuf/proto

3、install plug-in for protoc compiler
go get github.com/golang/protobuf/protoc-gen-go

4、compiler .proto
go generate

or 

protoc --go_out=. ./packet/test.proto

5、Others gogoprotobuf a fast protocol buffer decode and encoder package 
https://segmentfault.com/a/1190000009277748


