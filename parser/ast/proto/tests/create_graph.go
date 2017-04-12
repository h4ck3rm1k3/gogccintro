package main

import (
	//"fmt"
	//"os"	
	//"io/ioutil"
	//"github.com/golang/protobuf/proto"
	"github.com/gonum/graph/simple"
	"../"
	//"github.com/gonum/graph"
)

type GraphCreator struct {
	Out * simple.DirectedGraph
}

func (v *GraphCreator)Edge(from int32, to int32, edge int){
	v.Out.SetEdge(
		simple.Edge{
			F: simple.Node(from),
			T: simple.Node(to),
			W: float64(edge)})
}

func (v *GraphCreator)RecFile(In * astproto.File){
	for _,j := range In.GetNodes() {
		//v.RecNode(j)
		id := j.GetNodeID()
		//t  := j.GetNodeType()
		
		// fmt.Printf("id:%d type:%s :\n",
		// 	id,
		// 	t)
			
		for _,j2 := range j.GetAttrs() {
			//	v.RecAttr(j)
			nt := j2.GetNodeType()
			an := j2.GetAttrName()
			
			// fmt.Printf("\to%d f%d fn:%-20s FN:%d\t",
			// 	i2,
			// 	nt,
			// 	an,
			// 	int(an))

			//GetNodeType() TuNodeType {
			switch(nt){
				
			case astproto.TuNodeType_NTSourceAttr:
				//fmt.Printf("%s:%d\n",j2.SourceAttr.GetFileName(),j2.SourceAttr.GetLineNumber())
		
			case astproto.TuNodeType_NTNumberedNodeRef:
				v.Edge(id, j2.NodeAttr.GetNodeId(), int(an))
				
			case astproto.TuNodeType_NTNodeAttr:
				v.Edge(id, j2.NodeAttr.GetNodeId(), int(an))
				
			case astproto.TuNodeType_NTIntAttr:
				//fmt.Printf("Int:%s\n",j2.IntAttr.GetValue()) // : &IntAttr{ Value: &val },
				
			case astproto.TuNodeType_NTStringAttr:
				//fmt.Printf("String:%s\n",j2.GetStringAttr()) // : &val
				
			case astproto.TuNodeType_NTTagAttr:
				//fmt.Printf("Tag:%#v\n",j2.GetTagAttr()) // : &val }	)
				
			case astproto.TuNodeType_NTAccsAttr:
				//fmt.Printf("Acc:%#v\n",j2.GetAccess()) // : &val
				
			case astproto.TuNodeType_NTSpecValue:
				//fmt.Printf("Spec:%#v\n",j2.GetSpecAttr()) // : &spec,
				
			case astproto.TuNodeType_NTQualAttr:
				//fmt.Printf("Qual:%#v\n",j2.GetQualAttr()) // : &val 	}	)
				
			case astproto.TuNodeType_NTSignAttr:
				//fmt.Printf("Sign:%#v\n",j2.GetSignAttr()) // : &val	})
				
			case astproto.TuNodeType_NTNoteAttr:
				//fmt.Printf("Note:%#v\n",j2.GetNoteAttr()) // : &val })
				
			case astproto.TuNodeType_NTLinkAttr:
				//fmt.Printf("Link:%#v\n",j2.GetLinkAttr()) // : &val	})			
				
			
			default:
				panic("unknown nt")
			}
		}
		
	}
}

