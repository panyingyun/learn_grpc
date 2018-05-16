//go:generate protoc --go_out=. ./packet/test.proto

package main

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/panyingyun/learn_grpc/protocolexp/packet"
)

//https://segmentfault.com/a/1190000009277748

func main() {
	// 填充数据
	user := packet.UserInfo{
		Message: "panyingyun",
		Length:  1000,
		Cnt:     32,
	}

	// 对数据进行序列化
	data, err := proto.Marshal(&user)
	if err != nil {
		fmt.Errorf("Mashal data error: %v", err)
	}
	fmt.Println("Marshal Data = ", data)

	// 对已经序列化的数据进行反序列化
	var target packet.UserInfo
	err = proto.Unmarshal(data, &target)
	if err != nil {
		fmt.Errorf("UnMashal data error: %v", err)
	}
	fmt.Println(target.Message)
	fmt.Println(target.Length)
	fmt.Println(target.Cnt)
}
