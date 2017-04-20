package main

import (
	"testing"
	//"fmt"
	//"os"	
	//"io/ioutil"
	"github.com/gonum/graph/simple"
	//"../"
)


func TestReport(*testing.T){
	g := simple.NewDirectedGraph(0, 3000)

	// for id := 1; id < 15; id ++ {
	// 	//fn := simple.Node(id)
	// 	n0 := simple.Node(id)
	// 	g.AddNode(n0)
	// }
	
	for id := 1; id < 15; id ++ {
		fn := simple.Node(id)
		//g.AddNode(fn)
		for id2 := 1; id2 < 10; id2 ++ {
			tn := simple.Node(id2)
			//g.AddNode(tn)
			if id2 != id  {
				for an := 1; an < 5; an ++ {
					e:=simple.Edge{
						F: fn,
						T: tn,
						W: float64(an),
					}
					g.SetEdge(e)
				
				} //  for 
				
			} // if 
		}// for
	}// for
	Report(g)	
}// test report


