package node_types_examples
import (
	"testing"
	//"strconv"
	//"reflect"
	"fmt"
	"encoding/json"
	"github.com/h4ck3rm1k3/gogccintro/tree"
	"io/ioutil"
)

func TestLoad2(*testing.T){
	const filename = "funct_decl_key_get_conv_rpc_auth.json"
	
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("err %v\n", err)
		panic(err)
	}
	treemap := tree.NewTreeMap(0)
	err = json.Unmarshal(content, &treemap.Nodes)
	if err != nil {
		fmt.Printf("err %v\n", err)
		panic(err)
	}

	// first collect all identifiers, then what references them.
	
	//fmt.Printf("map %v\n", treemap.Nodes)
	for _,v := range treemap.Nodes {

		if v.NodeType == "identifier_node" {
			s:=treemap.FindName(v)
			fmt.Printf("node id %d %s name:%s\n", v.NodeID,v.NodeType,s)
		}	
	}
}
