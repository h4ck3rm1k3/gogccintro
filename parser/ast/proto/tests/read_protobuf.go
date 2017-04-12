package main

import (
	"fmt"
	"os"	
	"io/ioutil"
	"github.com/golang/protobuf/proto"
	"github.com/gonum/graph/simple"
	"../"
)


func ReadProto(filename string ) (* astproto.File) {

	data, err := ioutil.ReadFile(filename)

	if err != nil {panic(err)}
	newTest := &astproto.File{}
	err = proto.Unmarshal(data, newTest)

	if err != nil {	panic(fmt.Sprintf("unmarshaling error: %s", err))}
	return newTest
}
type Graph interface {
	RecFile(* astproto.File)
}
func main() {
	if len(os.Args) < 2 {
		name := os.Args[0]
		fmt.Printf("Usage: %v \"Test File\"\n", name)
		os.Exit(1)
	}

	filename := os.Args[1]
	otype := os.Args[2]

	fmt.Printf("type: %s\n", otype)
	
	newTest := ReadProto(filename)
	//g  := &simple.DirectedGraph{}
	g := simple.NewDirectedGraph(0, 3000)

	var v Graph
	if otype == "print" {
		// create a graph
		v=&GraphVistor{	Out:g }
	} else if otype == "create" {
		v=&GraphCreator{	Out:g }
	}else {
		panic("unknown")
	}
	
	v.RecFile(newTest)	
	//n0 := Node(g.NewNodeID())
	//g.AddNode(n0)
	
	
}
