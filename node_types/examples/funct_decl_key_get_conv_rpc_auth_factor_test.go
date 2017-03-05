package node_types_examples
import (
	"testing"
	//"strconv"
	//"reflect"
	//"fmt"
	//	"encoding/json"
	"github.com/h4ck3rm1k3/gogccintro/tree"
	"github.com/h4ck3rm1k3/gogccintro/node_types"
	"github.com/h4ck3rm1k3/gogccintro/models"
	//"io/ioutil"
	//"os"
)

type receiver struct {
	ids   map[int] NodeInstanceGeneric// instances
	NodeFactory node_types.NodeFactory
}
func (r* receiver) EndNode(v * models.GccTuParserNode){}
func (r* receiver) EndGraph(){}
func (r* receiver) StartGraph(tree * tree.TreeMap){
	r.ids = make(map[int] NodeInstanceGeneric)
	r.NodeFactory.StartGraph(tree)
}
func (t * receiver) ReferenceNode(n * models.GccTuParserNode, name string, o * models.GccTuParserNode){}
func (t * receiver) ReferencedNode(n * models.GccTuParserNode, name string, o * models.GccTuParserNode){}
func (t * receiver) ReferenceAttribute(n * models.GccTuParserNode, name string, val string){}
func (t * receiver) StartNode(v * models.GccTuParserNode){
	t.ids[v.NodeID]=t.NodeFactory.StartNode(v)
}

func TestFactory(*testing.T){
	//fmt.Printf("test load")
	const filename = "funct_decl_key_get_conv_rpc_auth.json"
	treemap := tree.NewTreeMapFromFile(filename)
	r := &receiver{
		NodeFactory:node_types.NodeFactory{
			Tree: treemap,
		},
	}
	treemap.ResolveReferences(r)
}
