package node_types


import (
	//"reflect"
	"fmt"
	"github.com/h4ck3rm1k3/gogccintro/models"
	"github.com/h4ck3rm1k3/gogccintro/tree"

)

type NodeTypeGeneric interface {
	Create(v * models.GccTuParserNode) NodeInterface
}

//func CreateNodeTypeIntegerType(v * models.GccTuParserNode) *NodeTypeIntegerType {
//	return &NodeTypeIntegerType{}
//}
type NodeTypeTypeDeclFactory struct {}
type NodeTypeArrayTypeFactory struct {}
type NodeTypePointerTypeFactory struct {}
type NodeTypeIntegerCstFactory struct {}
type NodeTypeRecordTypeFactory struct {}
type NodeTypeTreeListFactory struct {}
type NodeTypeFunctionDeclFactory struct {}
type NodeTypeIntegerTypeFactory struct {}
type NodeTypeIdentifierNodeFactory struct {}
type NodeTypeUnionTypeFactory struct {}
type NodeTypeFieldDeclFactory struct {}
type NodeTypeVoidTypeFactory struct {}
type NodeTypeFunctionTypeFactory struct {}


// define prototypes for each node type
var NodePrototypes =
	map[string] NodeTypeGeneric {
	"identifier_node": &NodeTypeIdentifierNodeFactory{},
	"pointer_type": &NodeTypePointerTypeFactory{},
	"function_decl": &NodeTypeFunctionDeclFactory{},
	"void_type": &NodeTypeVoidTypeFactory{},
	"function_type": &NodeTypeFunctionTypeFactory{},
	"integer_type": &NodeTypeIntegerTypeFactory{},
	"type_decl": &NodeTypeTypeDeclFactory{},
	"array_type": &NodeTypeArrayTypeFactory{},
	"integer_cst": &NodeTypeIntegerCstFactory{},
	"union_type": &NodeTypeUnionTypeFactory{},
	"record_type": &NodeTypeRecordTypeFactory{},
	"field_decl": &NodeTypeFieldDeclFactory{},
	"tree_list": &NodeTypeTreeListFactory{},
	
}

func CreateNodeTypeUnionType(v * models.GccTuParserNode) *NodeTypeUnionType { return &NodeTypeUnionType{}}
func (t * NodeTypeUnionTypeFactory)Create(v * models.GccTuParserNode) NodeInterface { return CreateNodeTypeUnionType(v)}
func CreateNodeTypeRecordType(v * models.GccTuParserNode) *NodeTypeRecordType { return &NodeTypeRecordType{}}
func (t * NodeTypeRecordTypeFactory)Create(v * models.GccTuParserNode) NodeInterface { return CreateNodeTypeRecordType(v)}
func CreateNodeTypeFieldDecl(v * models.GccTuParserNode) *NodeTypeFieldDecl { return &NodeTypeFieldDecl{}}
func (t * NodeTypeFieldDeclFactory)Create(v * models.GccTuParserNode) NodeInterface { return CreateNodeTypeFieldDecl(v)}
func CreateNodeTypeTreeList(v * models.GccTuParserNode) *NodeTypeTreeList { return &NodeTypeTreeList{}}
func (t * NodeTypeTreeListFactory)Create(v * models.GccTuParserNode) NodeInterface { return CreateNodeTypeTreeList(v)}
func CreateNodeTypeFunctionDecl(v * models.GccTuParserNode) *NodeTypeFunctionDecl { return &NodeTypeFunctionDecl{}}
func (t * NodeTypeFunctionDeclFactory)Create(v * models.GccTuParserNode) NodeInterface { return CreateNodeTypeFunctionDecl(v)}
func CreateNodeTypeIntegerType(v * models.GccTuParserNode) *NodeTypeIntegerType { return &NodeTypeIntegerType{}}
func (t * NodeTypeIntegerTypeFactory)Create(v * models.GccTuParserNode) NodeInterface { return CreateNodeTypeIntegerType(v)}
func CreateNodeTypeArrayType(v * models.GccTuParserNode) *NodeTypeArrayType { return &NodeTypeArrayType{}}
func (t * NodeTypeArrayTypeFactory)Create(v * models.GccTuParserNode) NodeInterface { return CreateNodeTypeArrayType(v)}
func CreateNodeTypeIdentifierNode(v * models.GccTuParserNode) *NodeTypeIdentifierNode { return &NodeTypeIdentifierNode{}}
func (t * NodeTypeIdentifierNodeFactory)Create(v * models.GccTuParserNode) NodeInterface { return CreateNodeTypeIdentifierNode(v)}
func CreateNodeTypeFunctionType(v * models.GccTuParserNode) *NodeTypeFunctionType { return &NodeTypeFunctionType{}}
func (t * NodeTypeFunctionTypeFactory)Create(v * models.GccTuParserNode) NodeInterface { return CreateNodeTypeFunctionType(v)}
func CreateNodeTypeVoidType(v * models.GccTuParserNode) *NodeTypeVoidType { return &NodeTypeVoidType{}}
func (t * NodeTypeVoidTypeFactory)Create(v * models.GccTuParserNode) NodeInterface { return CreateNodeTypeVoidType(v)}
func CreateNodeTypeTypeDecl(v * models.GccTuParserNode) *NodeTypeTypeDecl { return &NodeTypeTypeDecl{}}
func (t * NodeTypeTypeDeclFactory)Create(v * models.GccTuParserNode) NodeInterface { return CreateNodeTypeTypeDecl(v)}
func CreateNodeTypePointerType(v * models.GccTuParserNode) *NodeTypePointerType { return &NodeTypePointerType{}}
func (t * NodeTypePointerTypeFactory)Create(v * models.GccTuParserNode) NodeInterface { return CreateNodeTypePointerType(v)}
func CreateNodeTypeIntegerCst(v * models.GccTuParserNode) *NodeTypeIntegerCst { return &NodeTypeIntegerCst{}}
func (t * NodeTypeIntegerCstFactory)Create(v * models.GccTuParserNode) NodeInterface { return CreateNodeTypeIntegerCst(v)}

type NodeFactory struct {
	Tree * tree.TreeMap
	Types * TypesMap
	TypeNames map[string] bool
}

func GenerateCode(){

	fmt.Printf("var NodePrototypes =map[string] NodeTypeGeneric {\n")
	for n,x := range NodeTypeMap {
		fmt.Printf("\"%s\": & %sFactory{},\n", n,x)
	}
	fmt.Printf("}\n")
	
	for _,x := range NodeTypeMap {
		fmt.Printf("type %sFactory struct {}\n", x)
		fmt.Printf("func Create%s(v * models.GccTuParserNode) *%s { return &%s{}}\n",x,x,x)
		fmt.Printf("func (t * %sFactory)Create(v * models.GccTuParserNode) NodeInterface { return Create%s(v)}\n",x,x)
			
		//"integer_type":NodeTypeIntegerTypeFactory{},
	}
}

func (t * NodeFactory)StartGraph(tree * tree.TreeMap) {
	t.Types = CreateTypesMap()
	t.TypeNames= make(map[string] bool)
}

func (t * NodeFactory)EndGraph() {
	t.Types.Report()
}

func (t * NodeFactory)StartNode(v * models.GccTuParserNode)(NodeInterface) {

	//fmt.Printf("-------------------------------\n")
	//x := CreateNodeTypeIntegerType
	//	fmt.Printf("%v\n",x)
//	c := x(v)
//	fmt.Printf("%v\n",c)
	
	//objValue := reflect.ValueOf(v).Elem()

	if _, ok := t.TypeNames[v.NodeType]; ok {
	}else {
		t.TypeNames[v.NodeType]=true
		if o, ok := NodePrototypes[v.NodeType]; ok {
			o2:= o.Create(v)
			t.Types.MapType(v,o2)
			// 	//fmt.Printf("\tNew Object for type: %s\n",v.NodeType)
			// 	fmt.Printf("\tObject type: %v\n",o)
			// 	n1:=o(v)
			// 	fmt.Printf("\tNew: %v\n",n1)
		}
	}
	return nil
}
