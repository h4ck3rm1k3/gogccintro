package node_types

import (
	//"reflect"
	//"fmt"
	//"github.com/h4ck3rm1k3/gogccintro/models"
	//"github.com/h4ck3rm1k3/gogccintro/tree"
)
type NamedObjectInterface interface {}
type NodeInterface interface {}
type NameInterface interface {
	// can be an identifier node or type decl
}

type MinMaxMixin struct {

	RefsMax NodeTypeIntegerCst
	RefsMin NodeTypeIntegerCst
}

type TypeMixin struct {

	RefsUnql TypeInterface 
}


type NamedMixin struct {
	RefsName NameInterface
}

type NodeBase struct {
	NodeID int
	FileID int
	NodeType string
}

type TypeInterface struct {

	// interface for types
	RefsSize NodeTypeIntegerCst // all have a size
}

type NodeTypeArrayType struct {
	Base NodeBase
	RefsSize NodeTypeIntegerCst
	RefsDomn TypeInterface
	RefsElts TypeInterface
}

type NodeTypeFieldDecl struct {
	Base NodeBase
	RefsBpos NodeTypeIntegerCst
}

type NodeTypeFunctionDecl struct {
	Base NodeBase
	// the type of the function decl is always a function type
	RefsType NodeTypeFunctionType

	// the identifier name of the function
	RefsName NodeTypeIdentifierNode
}

type NodeTypeList struct {
	// list of node types
}

type NodeTypeFunctionType struct {
	Base NodeBase
	RefsPrms NodeTypeList
	RefsRetn TypeInterface
	RefsSize NodeTypeIntegerCst
}



type NodeTypeIdentifierNode struct {
	Base NodeBase
	Name string
	Named  NamedObjectInterface // what objects have this name?, can be multiple because names can be local
}

// func CreateNodeTypeIdentifierNode(NodeID int,Name string) *NodeTypeIdentifierNode{
// 	return &NodeTypeIdentifierNode{
// 		Name : Name,
// 		Base : NodeBase{
// 			NodeID:NodeID,
// 		},
// 	}
// }

	// local 
	//     field_decl
	
	// not local
	//     union_type, integer_type, type_decl, function_decl


type NodeTypeIntegerCst struct {
	Base NodeBase
	AttrsType * NodeTypeIntegerType

}
type NodeTypeIntegerType struct {
	Base NodeBase
}

type NodeTypePointerType struct {
	Base NodeBase
	// what is pointed to
	RefsPtd TypeInterface 
}

type NodeTypeRecordType struct {
	Base NodeBase
	RefsFlds NodeTypeFieldDecl
	RefsSize NodeTypeIntegerCst
}

type NodeTypeTreeList struct {
	Base NodeBase
}

type NodeTypeVoidType struct {
	Base NodeBase
}

type NodeTypeTypeDecl struct {
	Base NodeBase
}

type NodeTypeUnionType struct {
	Base NodeBase
	RefsSize NodeTypeIntegerCst
	RefsUnql * NodeTypeUnionType
	RefsFlds NodeTypeFieldDecl
	RefsName NameInterface
}

var NodeTypeNames=[]string {
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

var NodeTypeMap =
	map[string]string {
	"integer_type":"NodeTypeIntegerType",
	"type_decl":"NodeTypeTypeDecl",
	"array_type":"NodeTypeArrayType",
	"identifier_node":"NodeTypeIdentifierNode",
	"pointer_type":"NodeTypePointerType",
	"integer_cst":"NodeTypeIntegerCst",
	"union_type":"NodeTypeUnionType",
	"record_type":"NodeTypeRecordType",
	"field_decl":"NodeTypeFieldDecl",
	"tree_list":"NodeTypeTreeList",
	"function_decl":"NodeTypeFunctionDecl",
	"void_type":"NodeTypeVoidType",
	"function_type":"NodeTypeFunctionType",
}
