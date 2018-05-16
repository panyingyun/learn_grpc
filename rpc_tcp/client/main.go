package main

import (
	"errors"
	"fmt"
	"net/rpc"
	"os"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	var client, err = rpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Println("连接不到服务器：", err)
		os.Exit(1)
	}
	var args = Args{44, 33}
	var result int
	fmt.Println("开始调用！")
	err = client.Call("Arith.Multiply", args, &result)
	if err != nil {
		fmt.Println("调用失败！", err)
	}
	fmt.Println("Arith: ", args.A, " * ", args.B, " = ", result)
}
