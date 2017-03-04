package node_types_examples
import (
	"testing"
	//"strconv"
	"log"
	"fmt"
//	"encoding/json"
	"github.com/h4ck3rm1k3/gogccintro/tree"
	"github.com/h4ck3rm1k3/gogccintro/models"
	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/quad"
//	"github.com/cayleygraph/cayley/store"
	//"io/ioutil"
	//"os"
)

type AReceiverCayley struct {
	store * cayley.Handle
}

func (r * AReceiverCayley) StartGraph(){
	//fmt.Printf("start graph %v\n", r)
	store, err := cayley.NewMemoryGraph()
	r.store =store
	if err != nil {
		log.Fatalln(err)
	}
	//fmt.Printf("created graph %s\n", r.store)
}
func (r * AReceiverCayley)	StartNode(v * models.GccTuParserNode){
	//fmt.Printf("start node %v\n", r)
	//fmt.Printf("in graph %s\n", r.store)
	//fmt.Printf("node id %d %s\n", v.NodeID,v.NodeType)
	if r.store != nil{
		r.store.AddQuad(quad.Make(fmt.Sprintf("%d",v.NodeID), "node_type", v.NodeType, nil))
	} else {
		fmt.Printf("start node no store\n")
	}
}

func (r * AReceiverCayley)	ReferenceNode(n * models.GccTuParserNode, name string, o * models.GccTuParserNode){
	//fmt.Printf("\tField %d/%s -> %s -> %d/%s\n", n.NodeID, n.NodeType, name, o.NodeID, o.NodeType)
	if r.store != nil{
		r.store.AddQuad(quad.Make(fmt.Sprintf("%d",n.NodeID), name, fmt.Sprintf("%d",o.NodeID), nil))
	} else {
		fmt.Printf("start node no store\n")
	}
}

func (r * AReceiverCayley)	ReferenceAttribute(n * models.GccTuParserNode, name string, value string){
	//fmt.Printf("\tField %d %s -> %s : %v\n", n.NodeID, n.NodeType,name,value)
	if r.store != nil{
		r.store.AddQuad(quad.Make(n.NodeID, name, value, nil))
	} else {
		fmt.Printf("start node no store\n")
	}
}

func (r * AReceiverCayley)	EndNode(n * models.GccTuParserNode){
	//fmt.Printf("end node %v\n", r)
}
func (r * AReceiverCayley)	EndGraph(){
	//fmt.Printf("end graph %v\n", r)
	it := r.store.QuadsAllIterator()
	for it.Next() {
		qu := r.store.Quad(it.Result())
		fmt.Println(qu.Subject.Native(), qu.Predicate.Native(), qu.Object.Native())
		
	}
}

func TestLoadCayley(*testing.T){

	const filename = "funct_decl_key_get_conv_rpc_auth.json"
	treemap := tree.NewTreeMapFromFile(filename)
	r := &AReceiverCayley{}
	treemap.ResolveReferences(r)

}
