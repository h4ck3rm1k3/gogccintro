package main

import (
	"fmt"
	//"os"	
	//"io/ioutil"
	//"github.com/golang/protobuf/proto"
	"github.com/gonum/graph/simple"
	"../"
)

func Report(g * simple.DirectedGraph) {
	fmt.Printf("Report\n")

	//g.from
	for _,x := range(g.Nodes()) {// all the nodes

		in_nodes := g.To(x)

		// if 0 in nodes then it is a starting point
		if len(in_nodes) > 0 {
			// only one node references this
			// the types of fields that are containers are
			//  prms, chan, valu, flds

			// others that are random:
			// ptd, min, type , max, retn

    //   1 flds
    //   1 retn
    //   4 mngl
    //   7 ptd
    //   7 valu
    //  14 min
    //  16 max
    // 123 type
    // 297 prms
    // 360 chan
			// this should be unique : mngl
		
			//fmt.Printf("%d : %d\n",x, len(in_nodes))
		
			for _,f := range(in_nodes) { // all the edges to Nodes
				//d := g.Edge(x,f)
				w,ok := g.Weight(f,x)
				w2 := int32(w)
				
				if w2 != int32(astproto.Field_name_field)  {
					if ok {
						fmt.Printf("edge to %d from %d edge Found %f:%s max:%d\n",
							x.ID(),
							f.ID(),
							w,astproto.Field_name[w2],
							len(in_nodes),
						)
					} else {
						//	fmt.Printf("edge to %d from %d edge FAIL noweight\n",x,f)
						panic("fail")
					}
				}
				//func (g *DirectedGraph) To(n graph.Node) []graph.Node {
			}
		}
	}
}
