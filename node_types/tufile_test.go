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
func (t * TUFile) ReferencedNode(n * models.GccTuParserNode, name string, o * models.GccTuParserNode){}
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

func (t * TUFile) handle_type(v NodeInterface) NodeInterface{
	return v
}

func (t * TUFile) DispatchType(v * models.GccTuParserNode) NodeInterface{
	if v == nil {
		return nil
	}
	switch (v.NodeType) {
	case "identifier_node":
		return t.handle_type(t.CreateNodeTypeIdentifierNode(v))
		break
	case "pointer_type":
		return t.handle_type(t.CreateNodeTypePointerType(v))
		break
	case "function_decl":
		return t.handle_type(t.CreateNodeTypeFunctionDecl(v))
		break
	case "void_type":
		return t.handle_type(t.CreateNodeTypeVoidType(v))
		break
	case "function_type":
		return t.handle_type(t.CreateNodeTypeFunctionType(v))
		break
	case "integer_type":
		return t.handle_type(t.CreateNodeTypeIntegerType(v))
		break
	case "type_decl":
		return t.handle_type(t.CreateNodeTypeTypeDecl(v))
		break
	case "array_type":
		return t.handle_type(t.CreateNodeTypeArrayType(v))
		break
	case "integer_cst":
		return t.handle_type(t.CreateNodeTypeIntegerCst(v))
		break
	case "union_type":
		return t.handle_type(t.CreateNodeTypeUnionType(v))
		break
	case "record_type":
		return t.handle_type(t.CreateNodeTypeRecordType(v))
		break
	case "field_decl":
		return t.handle_type(t.CreateNodeTypeFieldDecl(v))
		break
	case "tree_list":
		return t.handle_type(t.CreateNodeTypeTreeList(v))
		break
	default:
		fmt.Errorf("error: %s\n", v.NodeType)
		return nil
	}
	return nil
}

func TestTUFile(*testing.T){
	t :=CreateTUFile ()
	//fmt.Printf("Test %s",t)

	const filename = "examples/funct_decl_key_get_conv_rpc_auth.json"
	treemap := tree.NewTreeMapFromFile(filename)
	treemap.ResolveReferences(t)

}
