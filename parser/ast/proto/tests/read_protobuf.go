package main

import (
	"fmt"
	"os"
	
	"io/ioutil"
	"github.com/golang/protobuf/proto"

	"github.com/gonum/graph/simple"
	
	//"github.com/h4ck3rm1k3/gogccintro/parser/ast/proto"
	"../"
)

type GraphVistor struct {
	//
	Out * simple.DirectedGraph
}

func (v *GraphVistor) RecAttr(In * astproto.Attr){

}

func (v *GraphVistor) RecNode(In * astproto.Node){
}

func (v *GraphVistor)RecFile(In * astproto.File){
	for i,j := range In.GetNodes() {
		//v.RecNode(j)
		id := j.GetNodeID()
		t  := j.GetNodeType()
		
		for i2,j2 := range j.GetAttrs() {
			//	v.RecAttr(j)
			nt := j2.GetNodeType()
			an := j2.GetAttrName()
			
			fmt.Printf("off:%d / id:%d type:%s / foff:%d ft:%#v fn:%s\n",
				i,
				id,
				t,
				i2,
				nt,
				an)
		}
		
	}
}

func main() {
	if len(os.Args) < 2 {
		name := os.Args[0]
		fmt.Printf("Usage: %v \"Test File\"\n", name)
		os.Exit(1)
	}
	filename := os.Args[1]
	data, err := ioutil.ReadFile(filename)

	if err != nil {panic(err)}
	newTest := &astproto.File{}
	err = proto.Unmarshal(data, newTest)

	if err != nil {	panic(fmt.Sprintf("unmarshaling error: %s", err))}
	
	// create a graph
	g  := &simple.DirectedGraph{}

	v:=GraphVistor{	Out:g }
	v.RecFile(newTest)
	//n0 := Node(g.NewNodeID())
	//g.AddNode(n0)
	
	
}
