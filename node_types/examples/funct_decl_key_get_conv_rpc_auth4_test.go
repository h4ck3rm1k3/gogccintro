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

type NodeTypeRange struct{}

type FieldType struct {
	// sample values
	// used by type
	// generic field vs field in type
	// domain of field (cardinality)
	// range of field (cardinality)
	// domain to range (cardinality matrix!)
	Implementation FieldTypeGeneric
	NodeTypes map[string] * NodeTypeRange
}


type NodeTypeInstance struct {
	fields_outgoing map[string] NodeFieldInstance // outgoing fields
	fields_incoming [] NodeFieldInstance // array of incoming objects
}

type NodeType struct {
	Name string
	FieldsOutgoing map[string] *FieldType // outgoing fields
	FieldsIncoming map[string] *FieldType // incoming fields
	Instances map[int] *NodeTypeInstance
	Implementation NodeTypeGeneric // interface to implementation hooks
	// cardinality of this node type
	// referenced by other nodes, contained by one?
	// references other nodes
}

func (t * NodeType) InitNodeType(){
	t.FieldsOutgoing = make(map[string] *FieldType) // outgoing fields
	t.FieldsIncoming = make(map[string] *FieldType) // incoming fields
	t.Instances = make(map[int] *NodeTypeInstance)
}

func (t * FieldType)	ReferenceNode(n * models.GccTuParserNode, name string, o * models.GccTuParserNode){

	if _, ok := t.NodeTypes[o.NodeType]; ok {
		//t.NodeTypes map[string] NodeTypeRange
	}else {
		fmt.Printf("%s->%s->%s\n", n.NodeType,name,o.NodeType,)
		t.NodeTypes[o.NodeType]=&NodeTypeRange{}
	}
	
}
// outgoing
func (t * NodeType)	ReferenceNode(n * models.GccTuParserNode, name string, o * models.GccTuParserNode){
	if o1, ok := t.FieldsOutgoing[name]; ok {
		o1.ReferenceNode(n, name, o )
	} else {

		fmt.Printf("%s->%s\n", n.NodeType,name)
		o2:=&FieldType{
			NodeTypes: make(map[string] *NodeTypeRange),
		}
		
		t.FieldsOutgoing[name]=o2
		o2.ReferenceNode(n, name, o )
	}
}

	// outgoing
func (t * NodeType)	ReferencedNode(n * models.GccTuParserNode, name string, o * models.GccTuParserNode){}
func (t* NodeType) StartNode(v * models.GccTuParserNode)(NodeInstanceGeneric){
	return nil
}
func (t* NodeType) EndGraph(){
	for i,x := range t.FieldsOutgoing {
		fmt.Printf("%s->%s->%s\n", t.Name, i,x)

		for j,_ := range x.NodeTypes {
			fmt.Printf("FINAL %s->%s->%s\n", t.Name, i, j)
		}
			
	}
}

type NodeInstanceGeneric interface {}

type NodeTypeGeneric interface {
	StartNode(v * models.GccTuParserNode)(NodeInstanceGeneric)
	// incoming
	ReferenceNode(n * models.GccTuParserNode, name string, o * models.GccTuParserNode)
	// outgoing
	ReferencedNode(n * models.GccTuParserNode, name string, o * models.GccTuParserNode)
	EndGraph()
}

type TReceiver struct {
	node_types map[string] NodeTypeGeneric
	node_types_generic map[string] NodeTypeGeneric
	ids   map[int] NodeInstanceGeneric// instances
}

var node_types=[]string {
	"integer_type",
	"type_decl",
	"array_type",
	"identifier_node",
	"pointer_type",
	"integer_cst",
	"union_type",
	"record_type",
	"field_decl",
	"tree_list",
	"function_decl",
	"void_type",
	"function_type",
}

func (r* TReceiver) StartGraph(){
	r.node_types_generic = make(map[string] NodeTypeGeneric)
	r.ids = make(map[int] NodeInstanceGeneric)
	// implementation types
	r.node_types =
		map[string] NodeTypeGeneric {
		// "integer_type":&NodeTypeIntegerType{},
		// "type_decl":&NodeTypeTypeDecl{},
		// "array_type":&NodeTypeArrayType{},
		"identifier_node":&NodeTypeIdentifierNode{
			names: make(map[string] *NodeInstanceIdentifierNode),// instances
		},
		// "pointer_type":&NodeTypePointerType{},
		// "integer_cst":&NodeTypeIntegerCst{},
		// "union_type":&NodeTypeUnionType{},
		// "record_type":&NodeTypeRecordType{},
		// "field_decl":&NodeTypeFieldDecl{},
		// "tree_list":&NodeTypeTreeList{},
		// "function_decl":&NodeTypeFunctionDecl{},
		// "void_type":&NodeTypeVoidType{},
		// "function_type":&NodeTypeFunctionType{},
	}
	for _,x := range node_types {
		//x=r.node_types
		if o, ok := r.node_types[x]; ok {
			o2:=&NodeType{
				Name:x,
				Implementation:o,
			}
			o2.InitNodeType()
			r.node_types_generic[x]=o2
		} else {
			o2:=&NodeType{
				Name:x,
			}
			o2.InitNodeType()
			r.node_types_generic[x]=o2
		}	
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

func (t* NodeTypeIdentifierNode) EndGraph(){}
//func GenericReferenceNode(n * models.GccTuParserNode, name string, o * models.GccTuParserNode){}
//func GenericReferencedNode(n * models.GccTuParserNode, name string, o * models.GccTuParserNode){}
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
	
	if o, ok := r.node_types_generic[v.NodeType]; ok {
		//fmt.Printf("node type %s %s\n",v.NodeType, val)
		t := o.StartNode(v)
		r.ids[v.NodeID]=t
	}else {
		//c := fmt.Sprintf("NodeType%s",CamelCase(v.NodeType))		
		fmt.Printf("node type not %s\n",v.NodeType)
		//fmt.Printf("type %s struct {}\n",c)
		//fmt.Printf("\"%s\":%s,\n",v.NodeType,c)		
		//r.node_types[v.NodeType]=NodeType{}
	}
	
}

func (r *TReceiver)	ReferenceNode(n * models.GccTuParserNode, name string, o * models.GccTuParserNode){
	//fmt.Printf("\tField %d/%s -> %s -> %d/%s\n", n.NodeID, n.NodeType, name, o.NodeID, o.NodeType)

	// look up the o.NodeType and
	//r.node_types[n.NodeType].ReferenceNode(n,name,o)
	//r.node_types[o.NodeType].ReferencedNode(n,name,o)

	// now lets peg
	if _, ok := r.node_types_generic[n.NodeType]; ok {
		r.node_types_generic[n.NodeType].ReferenceNode(n,name,o)
	} else	{
		///fmt.Printf("New NodeType %s\n", n.NodeType)
		c := fmt.Sprintf("NodeType%s",CamelCase(n.NodeType))		
		fmt.Printf("\"%s\":%s,\n",n.NodeType,c)		
	}
	
	if _, ok := r.node_types_generic[o.NodeType]; ok {
		r.node_types_generic[o.NodeType].ReferencedNode(n,name,o)
	} else	{
		//fmt.Printf("New NodeType %s\n", o.NodeType)
		c := fmt.Sprintf("NodeType%s",CamelCase(o.NodeType))		
		fmt.Printf("\"%s\":%s,\n",o.NodeType,c)				
	}
}

func (r *TReceiver)	ReferenceAttribute(n * models.GccTuParserNode, name string, value string){
	//fmt.Printf("\tField %d %s -> %s : %v\n", n.NodeID, n.NodeType,name,value)
}

func (r *TReceiver)	EndNode(n * models.GccTuParserNode){}
func (r* TReceiver)	EndGraph(){

	//fmt.Printf("End%v\n", r )
	//fmt.Printf("End%v\n", r.node_types )
	//for i,x := range r.node_types {		fmt.Printf("End: %v -> %v\n", i, x )	}

	for i,x := range r.node_types_generic {
		fmt.Printf("End: %v -> %v\n", i, x )
		x.EndGraph()
	}
	
}

func TestLoadType(*testing.T){
	//fmt.Printf("test load")
	const filename = "funct_decl_key_get_conv_rpc_auth.json"
	treemap := tree.NewTreeMapFromFile(filename)
	r := &TReceiver{}
	treemap.ResolveReferences(r)

}
