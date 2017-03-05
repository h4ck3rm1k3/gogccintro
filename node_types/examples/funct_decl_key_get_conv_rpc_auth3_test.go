package node_types_examples
import (
	"testing"
	//"strconv"
	//"reflect"
	"fmt"
//	"encoding/json"
	"github.com/h4ck3rm1k3/gogccintro/tree"
	"github.com/h4ck3rm1k3/gogccintro/models"
	//"io/ioutil"
	//"os"
)

type AReceiver struct {

}

func (r* AReceiver) StartGraph(tree *tree.TreeMap){}
func (r* AReceiver)	StartNode(v * models.GccTuParserNode){
	fmt.Printf("node id %d %s\n", v.NodeID,v.NodeType)
}

func (r *AReceiver)	ReferenceNode(n * models.GccTuParserNode, name string, o * models.GccTuParserNode){
	fmt.Printf("\tField %d/%s -> %s -> %d/%s\n", n.NodeID, n.NodeType, name, o.NodeID, o.NodeType)
}

func (r *AReceiver)	ReferenceAttribute(n * models.GccTuParserNode, name string, value string){
	fmt.Printf("\tField %d %s -> %s : %v\n", n.NodeID, n.NodeType,name,value)
}

func (r *AReceiver)	EndNode(n * models.GccTuParserNode){}
func (r* AReceiver)	EndGraph(){}

func TestLoad3(*testing.T){

	const filename = "funct_decl_key_get_conv_rpc_auth.json"
	treemap := tree.NewTreeMapFromFile(filename)
	r := &AReceiver{}
	treemap.ResolveReferences(r)

}
