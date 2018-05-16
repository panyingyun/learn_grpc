# learn_grpc
learn_grpc

Learn protocol buffer v3(Golang)

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


Learn rpc over http
1、rpc_http and rpc_tcp and rpcx_direct 

https://scguoi.github.io/DivisionByZero/2016/11/15/GO%E8%AF%AD%E8%A8%80RPC%E6%96%B9%E6%A1%88%E8%B0%83%E7%A0%94.html

2、rpcx_direct rpcx_etcdv3

https://github.com/rpcx-ecosystem/rpcx-examples3

go get -u -v -tags "etcd" github.com/smallnest/rpcx/...


