package tests

import (
	//"reflect"
	"fmt"
	"os"
	"io"
	"github.com/golang/protobuf/proto"
	///
//	"github.com/h4ck3rm1k3/gogccintro/fakego/ast"
	"github.com/h4ck3rm1k3/gogccintro/fakego/ast/proto"
	//"go/parser"
//	"github.com/h4ck3rm1k3/gogccintro/fakego/token"
)


func ReadBuf(){

	p:= &astproto.File{}
	
	file, err := os.Open("example.protobuf")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	stat, _ := file.Stat()

	buff := make([]byte, stat.Size())

	io.ReadFull(file, buff)
	file.Close()
	proto.Unmarshal(buff, p)

	fmt.Println(p)
}
