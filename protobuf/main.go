package main

import (
	"fmt"
	"os"

	"github.com/22fortisetliber/lab/protobuf/pb"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func main() {
	data, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("err on reading file", err)
		return
	}
	fmt.Println(data)
	bpb := &pb.Book{}
	err = protojson.Unmarshal(data, bpb)
	if err != nil {
		fmt.Println("failed to unmarshal", err)
		return
	}
	bytes, err := proto.Marshal(bpb)
	if err != nil {
		fmt.Println("failed to marshal", err)
		return
	}
	fmt.Println(bytes)
}
