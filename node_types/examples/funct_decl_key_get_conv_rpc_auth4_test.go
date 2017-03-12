package node_types_examples
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

type Stats struct {}

// generic instance of a type
type NodeFieldInstance interface {}
type NodeFieldInstanceString struct {}
type NodeFieldInstanceRef struct {}
type FieldTypeGeneric interface {}
type NodeTypeRange struct{}
type NodeTypeDomain struct{}

type FieldType struct {
	Name string
	// sample values
	// used by type
	// generic field vs field in type
	// domain of field (cardinality)
	// range of field (cardinality)
	// domain to range (cardinality matrix!)
	Implementation FieldTypeGeneric
	RangeTypes map[string] * NodeTypeRange
	DomainTypes map[string] * NodeTypeDomain
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

// out going
func (t * FieldType) ReferenceNode(n * models.GccTuParserNode, name string, o * models.GccTuParserNode){
	if _, ok := t.RangeTypes[o.NodeType]; ok {
		//t.NodeTypes map[string] NodeTypeRange
	}else {
		//fmt.Printf("%s->%s->%s\n", n.NodeType,name,o.NodeType,)
		t.RangeTypes[o.NodeType]=&NodeTypeRange{}
	}
}

func (t * NodeType)	ReferenceNode(n * models.GccTuParserNode, name string, o * models.GccTuParserNode){
	if o1, ok := t.FieldsOutgoing[name]; ok {
		o1.ReferenceNode(n, name, o )
	} else {
		//fmt.Printf("%s->%s\n", n.NodeType,name)
		o2:=&FieldType{
			Name:name,
			RangeTypes: make(map[string] *NodeTypeRange),
			DomainTypes: make(map[string] *NodeTypeDomain),
		}
		t.FieldsOutgoing[name]=o2
		o2.ReferenceNode(n, name, o )
	}
}


// incoming

// called on the o.nodetype, so the incoming is from n
func (t * FieldType) ReferencedNode(n * models.GccTuParserNode, name string, o * models.GccTuParserNode){
	if _, ok := t.DomainTypes[n.NodeType]; ok {
	}else {
		//fmt.Printf("field referenced %s<-%s<-%s\n", o.NodeType,name,n.NodeType,)
		t.DomainTypes[n.NodeType]=&NodeTypeDomain{}
	}
}

func (t * NodeType)	ReferencedNode(n * models.GccTuParserNode, name string, o * models.GccTuParserNode){
	//fmt.Printf("starting referenced node %s<-%s in %v\n", o.NodeType,name,t)
	if o1, ok := t.FieldsIncoming[name]; ok {
		o1.ReferencedNode(n, name, o )
	} else {
		//fmt.Printf("adding new referenced %s<-%s\n", n.NodeType,name)
		o2:=&FieldType{
			Name:name,
			RangeTypes: make(map[string] *NodeTypeRange),
			DomainTypes: make(map[string] *NodeTypeDomain),
		}
		t.FieldsIncoming[name]=o2
		o2.ReferencedNode(n, name, o )
	}
}

func (t* NodeType) StartNode(v * models.GccTuParserNode)(NodeInstanceGeneric){
	return nil
}

func (t* NodeType) EndGraph(){

	fmt.Printf("REFS\n")
	for i,x := range t.FieldsOutgoing {
		//fmt.Printf("%s->%s->%s\n", t.Name, i,x)
		for j,_ := range x.RangeTypes {
			c := fmt.Sprintf("NodeType%s",CamelCase(j))
			fmt.Printf("\t%s %s\n", i, c)
		}
	}
	
	fmt.Printf("IN\n")
	for i,x := range t.FieldsIncoming {

		//fmt.Printf("%s<-%s<-%s\n", t.Name, i,x)
		for j,_ := range x.DomainTypes {
			c := fmt.Sprintf("NodeType%s",CamelCase(j))
			fmt.Printf("\t%s %s\n", i, c)
		}
	}
}

type NodeInstanceGeneric interface {}

type NodeTypeGeneric interface {
	StartNode(v * models.GccTuParserNode)(NodeInstanceGeneric)
    ReferenceNode(n * models.GccTuParserNode, name string, o * models.GccTuParserNode)
    ReferencedNode(n * models.GccTuParserNode, name string, o * models.GccTuParserNode)
	EndGraph()
}

type TReceiver struct {
	node_types map[string] NodeTypeGeneric
	node_types_generic map[string] NodeTypeGeneric
	ids   map[int] NodeInstanceGeneric// instances
	NodeFactory node_types.NodeFactory
}

func (r* TReceiver) StartGraph(tree * tree.TreeMap){
	r.node_types_generic = make(map[string] NodeTypeGeneric)
	r.ids = make(map[int] NodeInstanceGeneric)
	r.NodeFactory.StartGraph(tree)
	// implementation types
	r.node_types =
		map[string] NodeTypeGeneric {
		// "integer_type":&NodeTypeIntegerType{},
		// "type_decl":&NodeTypeTypeDecl{},
		// "array_type":&NodeTypeArrayType{},
		"identifier_node":&NodeTypeIdentifierNode{
			Names: make(map[string] *node_types.NodeTypeIdentifierNode),// instances
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
	for _,x := range node_types.NodeTypeNames {
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

func (t * NodeTypeIdentifierNode) ReferenceNode(n * models.GccTuParserNode, name string, o * models.GccTuParserNode){}
func (t * NodeTypeIdentifierNode) ReferencedNode(n * models.GccTuParserNode, name string, o * models.GccTuParserNode){}
	

///////////////////////////////

type NodeTypeIdentifierNode struct {
	Names map[string] *node_types.NodeTypeIdentifierNode// instances
}

// type NodeTypeVoidType struct {}

///////////////////////////////////

func (t* NodeTypeIdentifierNode) EndGraph(){}

func (t* NodeTypeIdentifierNode) StartNode(v * models.GccTuParserNode)(NodeInstanceGeneric){
	//o:=node_types.CreateNodeTypeIdentifierNode(v)
	//t.Names[v.AttrsString]=o
	return nil
}

func (r* TReceiver) StartNode(v * models.GccTuParserNode){

	//r.NodeFactory.StartNode(v)
	
	//fmt.Printf("node id %d %s\n", v.NodeID,v.NodeType)
	if o, ok := r.node_types_generic[v.NodeType]; ok {
		//fmt.Printf("node type %s %s\n",v.NodeType, val)
		t := o.StartNode(v)
		r.ids[v.NodeID]=t
	}else {
		//c := fmt.Sprintf("NodeType%s",CamelCase(v.NodeType))
		//fmt.Printf("node type not %s\n",v.NodeType)
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
		//c := fmt.Sprintf("NodeType%s",CamelCase(n.NodeType))
		//fmt.Printf("reference \"%s\":%s,\n",n.NodeType,c)
	}

	// this is the referenced node, incoming to the o object 
	if _, ok := r.node_types_generic[o.NodeType]; ok {
		//fmt.Printf("Referenced %s\n",o.NodeType)
		r.node_types_generic[o.NodeType].ReferencedNode(n,name,o)
	} else	{
		//fmt.Printf("New NodeType %s\n", o.NodeType)
		//c := fmt.Sprintf("NodeType%s",CamelCase(o.NodeType))
		//fmt.Printf("refrenced \"%s\":%s,\n",o.NodeType,c)
	}
}

func (r *TReceiver)	ReferenceAttribute(n * models.GccTuParserNode, name string, value string){
	//fmt.Printf("\tField %d %s -> %s : %v\n", n.NodeID, n.NodeType,name,value)
}

func (r *TReceiver)	EndNode(n * models.GccTuParserNode){}
func (r* TReceiver)	EndGraph(){
	fmt.Printf("Prototypes %v\n", node_types.NodePrototypes )
	//fmt.Printf("End%v\n", r.node_types )
    //for i,x := range r.node_types {		fmt.Printf("End: %v -> %v\n", i, x )	}
	for i,x := range r.node_types_generic {
		fmt.Printf("End: %v\n", i )
		x.EndGraph()
	}
}

func TestLoadType(*testing.T){
	//fmt.Printf("test load")
	const filename = "funct_decl_key_get_conv_rpc_auth.json"
	treemap := tree.NewTreeMapFromFile(filename)
	r := &TReceiver{
		NodeFactory:node_types.NodeFactory{
			Tree: treemap,
		},
	}
	treemap.ResolveReferences(r)
}
