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

func CamelCase(s string) string {
	var r []byte
	var i = 1

	// first letterx
	c2 := s[0]
	c2 ^= ' ' // Make it a capital letter.
	r = append(r, c2)
	
	for ; i < len(s); i++ {
		c := s[i]
		if c == '_' {
			c2 := s[i+1]
			c2 ^= ' ' // Make it a capital letter.
			i = i +1
			r = append(r, c2)
		} else {
			r = append(r, c)
		}
	}
	return string(r[:])
}

type Stats struct {

}

// generic instance of a type
type NodeFieldInstance interface {
}

type NodeFieldInstanceString struct {
}

type NodeFieldInstanceRef struct {
}


type FieldTypeGeneric interface {
}

type FieldType struct {
	// sample values
	// used by type
	// generic field vs field in type
	// domain of field (cardinality)
	// range of field (cardinality)
	// domain to range (cardinality matrix!)
}


type NodeTypeInstance struct {
	fields_outgoing map[string] NodeFieldInstance // outgoing fields
	fields_incoming [] NodeFieldInstance // array of incoming objects
}

type NodeType struct {
	fields_outgoing map[string] *FieldTypeGeneric // outgoing fields
	fields_incoming map[string] *FieldTypeGeneric // incoming fields
	instances map[int] *NodeTypeInstance
	// cardinality of this node type
	// referenced by other nodes, contained by one?
	// references other nodes
}

type NodeInstanceGeneric interface {
	
}

type NodeTypeGeneric interface {
	StartNode(v * models.GccTuParserNode)(NodeInstanceGeneric)

	// incoming
	ReferenceNode(n * models.GccTuParserNode, name string, o * models.GccTuParserNode)

	// outgoing
	ReferencedNode(n * models.GccTuParserNode, name string, o * models.GccTuParserNode)
		
}

type TReceiver struct {
	node_types map[string] NodeTypeGeneric
	node_types_generic map[string] NodeTypeGeneric
	ids   map[int] NodeInstanceGeneric// instances
}

func (r* TReceiver) StartGraph(){
	r.node_types_generic = make(map[string] NodeTypeGeneric)
	r.ids = make(map[int] NodeInstanceGeneric)
	
	r.node_types =
		map[string] NodeTypeGeneric {
		"integer_type":&NodeTypeIntegerType{},
		"type_decl":&NodeTypeTypeDecl{},
		"array_type":&NodeTypeArrayType{},
		"identifier_node":&NodeTypeIdentifierNode{
			names: make(map[string] *NodeInstanceIdentifierNode),// instances
		},
		"pointer_type":&NodeTypePointerType{},
		"integer_cst":&NodeTypeIntegerCst{},
		"union_type":&NodeTypeUnionType{},
		"record_type":&NodeTypeRecordType{},
		"field_decl":&NodeTypeFieldDecl{},
		"tree_list":&NodeTypeTreeList{},
		"function_decl":&NodeTypeFunctionDecl{},
		"void_type":&NodeTypeVoidType{},
		"function_type":&NodeTypeFunctionType{},
	}	
}

type NamedObjectInterface interface {}
type NodeInstanceIdentifierNode struct {
	String string
	NodeID int
	Named  NamedObjectInterface // what object is named?
}

///////////////////////////////

type NodeTypeArrayType struct {}
type NodeTypeFieldDecl struct {}
type NodeTypeFunctionDecl struct {}
type NodeTypeFunctionType struct {}
type NodeTypeIdentifierNode struct {
	names map[string] *NodeInstanceIdentifierNode// instances
}
type NodeTypeIntegerCst struct {}
type NodeTypeIntegerType struct {}
type NodeTypePointerType struct {}
type NodeTypeRecordType struct {}
type NodeTypeTreeList struct {}
type NodeTypeTypeDecl struct {}
type NodeTypeUnionType struct {}
type NodeTypeVoidType struct {}

///////////////////////////////////


func (t* NodeTypeArrayType) ReferenceNode(n * models.GccTuParserNode, name string, o * models.GccTuParserNode){}
func (t* NodeTypeFieldDecl) ReferenceNode(n * models.GccTuParserNode, name string, o * models.GccTuParserNode){}
func (t* NodeTypeFunctionDecl) ReferenceNode(n * models.GccTuParserNode, name string, o * models.GccTuParserNode){}
func (t* NodeTypeFunctionType) ReferenceNode(n * models.GccTuParserNode, name string, o * models.GccTuParserNode){}
func (t* NodeTypeIdentifierNode) ReferenceNode(n * models.GccTuParserNode, name string, o * models.GccTuParserNode){}
func (t* NodeTypeIntegerCst) ReferenceNode(n * models.GccTuParserNode, name string, o * models.GccTuParserNode){}
func (t* NodeTypeIntegerType) ReferenceNode(n * models.GccTuParserNode, name string, o * models.GccTuParserNode){}
func (t* NodeTypePointerType) ReferenceNode(n * models.GccTuParserNode, name string, o * models.GccTuParserNode){}
func (t* NodeTypeRecordType) ReferenceNode(n * models.GccTuParserNode, name string, o * models.GccTuParserNode){}
func (t* NodeTypeTreeList) ReferenceNode(n * models.GccTuParserNode, name string, o * models.GccTuParserNode){}
func (t* NodeTypeTypeDecl) ReferenceNode(n * models.GccTuParserNode, name string, o * models.GccTuParserNode){}
func (t* NodeTypeUnionType) ReferenceNode(n * models.GccTuParserNode, name string, o * models.GccTuParserNode){}
func (t* NodeTypeVoidType) ReferenceNode(n * models.GccTuParserNode, name string, o * models.GccTuParserNode){}

func (t* NodeTypeArrayType) ReferencedNode(n * models.GccTuParserNode, name string, o * models.GccTuParserNode){}
func (t* NodeTypeFieldDecl) ReferencedNode(n * models.GccTuParserNode, name string, o * models.GccTuParserNode){}
func (t* NodeTypeFunctionDecl) ReferencedNode(n * models.GccTuParserNode, name string, o * models.GccTuParserNode){}
func (t* NodeTypeFunctionType) ReferencedNode(n * models.GccTuParserNode, name string, o * models.GccTuParserNode){}
func (t* NodeTypeIdentifierNode) ReferencedNode(n * models.GccTuParserNode, name string, o * models.GccTuParserNode){}
func (t* NodeTypeIntegerCst) ReferencedNode(n * models.GccTuParserNode, name string, o * models.GccTuParserNode){}
func (t* NodeTypeIntegerType) ReferencedNode(n * models.GccTuParserNode, name string, o * models.GccTuParserNode){}
func (t* NodeTypePointerType) ReferencedNode(n * models.GccTuParserNode, name string, o * models.GccTuParserNode){}
func (t* NodeTypeRecordType) ReferencedNode(n * models.GccTuParserNode, name string, o * models.GccTuParserNode){}
func (t* NodeTypeTreeList) ReferencedNode(n * models.GccTuParserNode, name string, o * models.GccTuParserNode){}
func (t* NodeTypeTypeDecl) ReferencedNode(n * models.GccTuParserNode, name string, o * models.GccTuParserNode){}
func (t* NodeTypeUnionType) ReferencedNode(n * models.GccTuParserNode, name string, o * models.GccTuParserNode){}
func (t* NodeTypeVoidType) ReferencedNode(n * models.GccTuParserNode, name string, o * models.GccTuParserNode){}



func (t* NodeTypeArrayType) StartNode(v * models.GccTuParserNode)(NodeInstanceGeneric){
	return nil
}
func (t* NodeTypeFieldDecl) StartNode(v * models.GccTuParserNode)(NodeInstanceGeneric){
	return nil
}
func (t* NodeTypeFunctionDecl) StartNode(v * models.GccTuParserNode)(NodeInstanceGeneric){
	return nil
}
func (t* NodeTypeFunctionType) StartNode(v * models.GccTuParserNode)(NodeInstanceGeneric){
	return nil
}

func (t* NodeTypeIdentifierNode) StartNode(v * models.GccTuParserNode)(NodeInstanceGeneric){
	o:=NodeInstanceIdentifierNode{
		String :v.AttrsString,
		NodeID: v.NodeID,
	}
	t.names[v.AttrsString]=&o
	return &o
}

func (t* NodeTypeIntegerCst) StartNode(v * models.GccTuParserNode)(NodeInstanceGeneric){
	return nil
}
func (t* NodeTypeIntegerType) StartNode(v * models.GccTuParserNode)(NodeInstanceGeneric){
	return nil
}
func (t* NodeTypePointerType) StartNode(v * models.GccTuParserNode)(NodeInstanceGeneric){
	return nil
}
func (t* NodeTypeRecordType) StartNode(v * models.GccTuParserNode)(NodeInstanceGeneric){
	return nil
}
func (t* NodeTypeTreeList) StartNode(v * models.GccTuParserNode)(NodeInstanceGeneric){
	return nil
}
func (t* NodeTypeTypeDecl) StartNode(v * models.GccTuParserNode)(NodeInstanceGeneric){
	return nil
}
func (t* NodeTypeUnionType) StartNode(v * models.GccTuParserNode)(NodeInstanceGeneric){
	return nil
}
func (t* NodeTypeVoidType) StartNode(v * models.GccTuParserNode)(NodeInstanceGeneric){
	return nil
}

func (r* TReceiver) StartNode(v * models.GccTuParserNode){
	//fmt.Printf("node id %d %s\n", v.NodeID,v.NodeType)

	o := r.node_types[v.NodeType].StartNode(v)

	r.ids[v.NodeID]=o
	
	if _, ok := r.node_types_generic[v.NodeType]; ok {
		//fmt.Printf("node type %s %s\n",v.NodeType, val)
		
	}else {
		c := fmt.Sprintf("NodeType%s",CamelCase(v.NodeType))		
		//fmt.Printf("node type not %s %s -> %s\n",v.NodeType, val, c)
		//fmt.Printf("type %s struct {}\n",c)
		fmt.Printf("\"%s\":%s,\n",v.NodeType,c)		
		//r.node_types[v.NodeType]=NodeType{}
	}
	
}

func (r *TReceiver)	ReferenceNode(n * models.GccTuParserNode, name string, o * models.GccTuParserNode){
	fmt.Printf("\tField %d/%s -> %s -> %d/%s\n", n.NodeID, n.NodeType, name, o.NodeID, o.NodeType)

	// look up the o.NodeType and
	r.node_types[n.NodeType].ReferenceNode(n,name,o)
	r.node_types[o.NodeType].ReferencedNode(n,name,o)
}

func (r *TReceiver)	ReferenceAttribute(n * models.GccTuParserNode, name string, value string){
	fmt.Printf("\tField %d %s -> %s : %v\n", n.NodeID, n.NodeType,name,value)
}

func (r *TReceiver)	EndNode(n * models.GccTuParserNode){}
func (r* TReceiver)	EndGraph(){

	fmt.Printf("End%v\n", r )
	fmt.Printf("End%v\n", r.node_types )
	for i,x := range r.node_types {
		fmt.Printf("End: %v -> %v\n", i, x )
	}

	
}

func TestLoadType(*testing.T){
	//fmt.Printf("test load")
	const filename = "funct_decl_key_get_conv_rpc_auth.json"
	treemap := tree.NewTreeMapFromFile(filename)
	r := &TReceiver{}
	treemap.ResolveReferences(r)

}
