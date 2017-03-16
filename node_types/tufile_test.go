package node_types

import (
	"testing"
	"fmt"
	"bytes"
	"github.com/h4ck3rm1k3/gogccintro/tree"
	"github.com/h4ck3rm1k3/gogccintro/models"
		"encoding/json"
)

func (r* TUFile) EndNode(v * models.GccTuParserNode){}
func (r* TUFile) EndGraph(){

}
func (r* TUFile) StartGraph(tree * tree.TreeMap){
	r.Tree= tree
	fmt.Printf("\n")
}
func (t * TUFile) ReferenceNode(n * models.GccTuParserNode, name string, o * models.GccTuParserNode){}

func (t * TUFile) ReferencedNode(n * models.GccTuParserNode, name string, o * models.GccTuParserNode){

}

func (t * TUFile) ReferenceAttribute(n * models.GccTuParserNode, name string, val string){}
func (t * TUFile) StartNode(v * models.GccTuParserNode){
	//t.ids[v.NodeID]=t.NodeFactory.StartNode(v)
	i:=t.DispatchType(v)
	if i != nil {
		//fmt.Printf("Created %v from %v type:%s\n",i, v, v.NodeType)
		var buffer bytes.Buffer	
		body, _ := json.MarshalIndent(i,"\t","\t")
		buffer.Write(body)
		buffer.WriteString("\n")
		fmt.Printf("%s",buffer.String())
		//outf.WriteString(buffer.Striang())

		
	} else {
		//fmt.Printf("Created null from %v\n",v)
	}
		
}


func TestTUFile(*testing.T){
	t :=CreateTUFile ()
	//fmt.Printf("Test %s",t)

	const filename = "examples/funct_decl_key_get_conv_rpc_auth.json"
	treemap := tree.NewTreeMapFromFile(filename)
	treemap.ResolveReferences(t)
	//treemap.ResolveReferences(t)// second pass?

}
