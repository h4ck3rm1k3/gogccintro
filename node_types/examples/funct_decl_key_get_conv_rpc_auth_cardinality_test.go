package node_types_examples
/*
prints out the number of times a given object is referenced and the type of it,
this helps understand the cardinality of objects, if they are contained in total by the parent or if they need to exist outside and be looked up multiple times.
*/

import (
	"testing"
	//"strconv"
	//"reflect"
	"fmt"
	//	"encoding/json"
	"github.com/h4ck3rm1k3/gogccintro/tree"
	"github.com/h4ck3rm1k3/gogccintro/node_types"
	"github.com/h4ck3rm1k3/gogccintro/models"
	//"io/ioutil"
	//"os"
)

type counter struct {
	name string
	count int
}
func (c * counter) do_count(){
	c.count++
}
type receiver struct {
	ids   map[int] *counter
	NodeFactory node_types.NodeFactory
}
func (r* receiver) EndNode(v * models.GccTuParserNode){}
func (r* receiver) EndGraph(){}
func (r* receiver) StartGraph(tree * tree.TreeMap){
	r.ids = make(map[int] *counter)
}
func (r * receiver) ReferenceNode(n * models.GccTuParserNode, name string, o * models.GccTuParserNode){
	if o1, ok := r.ids[o.NodeID]; ok {
		o1.do_count();
	}else {
		r.ids[o.NodeID]=&counter{count:1,name:o.NodeType}
	}
}
func (t * receiver) ReferencedNode(n * models.GccTuParserNode, name string, o * models.GccTuParserNode){
	//r.ids[o.NodeId]++;
}
func (t * receiver) ReferenceAttribute(n * models.GccTuParserNode, name string, val string){}
func (t * receiver) StartNode(v * models.GccTuParserNode){
	if _, ok := t.ids[v.NodeID]; ok {
		
	} else {
		t.ids[v.NodeID]=&counter{count:0,name:v.NodeType}
	}
}

func TestCardinality(*testing.T){
	//fmt.Printf("test load")
	const filename = "funct_decl_key_get_conv_rpc_auth.json"
	treemap := tree.NewTreeMapFromFile(filename)
	r := &receiver{
		NodeFactory:node_types.NodeFactory{
			Tree: treemap,
		},
	}
	treemap.ResolveReferences(r)
	for x,y := range r.ids {
		fmt.Printf("%d -> %s:%d\n", x,y.name,y.count)
	}
}
