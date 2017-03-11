package node_types

import (
	//"reflect"
	//"fmt"
	"database/sql"
	//"github.com/h4ck3rm1k3/gogccintro/models"
	//"github.com/h4ck3rm1k3/gogccintro/tree"
)



/*
a name scope
*/
type NameScope struct {
	Names map[string] * NodeTypeIdentifierNode
}

type IntegerConstants struct {
	Ints map[string] * NodeTypeIntegerCst
}

type TypeCollection struct {
	Names map[string] * TypeInterface  // name lookup
	Ids   map[int] * TypeInterface  // id lookup
}

/*
collection of
 * integer constants
 * identifiers (except for field decls and other local its)
 * types ( how do we index them?)
*/

type GlobalScope struct {
	Names NameScope
	Integers IntegerConstants
	Types TypeCollection
}

type NodeInterface interface {
	Load(v sql.NullInt64)
}

type NamedObjectInterface interface {}

type NameInterface interface {
	// can be an identifier node or type decl
	String() string // all have a size
	Load(v sql.NullInt64)
}
type TypeInterface interface {
	// interface for types
	RefsSize() * NodeTypeIntegerCst // all have a size
	Load(v sql.NullInt64)
}
func NodeIdFromString(in string) (sql.NullInt64) {
	return sql.NullInt64{Valid:false}
}

type MinMaxMixin struct {
	RefsMax *NodeTypeIntegerCst
	RefsMin *NodeTypeIntegerCst
}

type TypeMixin struct {
	RefsUnql TypeInterface 
}

type NamedMixin struct {
	RefsName NameInterface
}

type NodeType struct {
	NodeType string
}

/*
node interface is the base interface into all objects looked up by nodeid
*/

type NodeBase struct {
	NodeID int
	File * TUFile 
	//NodeType * NodeType
	NodeType NodeTypeGeneric
}

type NodeTypeIdentifierNode struct {
	Base NodeBase
	Name string
	Named  NamedObjectInterface // what objects have this name?, can be multiple because names can be local
	Scope  NameScope
}


type NodeTypeArrayType struct {
	Base NodeBase
	RefsSize * NodeTypeIntegerCst
	RefsDomn TypeInterface
	RefsElts TypeInterface
}

type NodeTypeFieldDecl struct {
	Base NodeBase
	RefsBpos * NodeTypeIntegerCst
}

type NodeTypeFunctionDecl struct {
	Base NodeBase
	// the type of the function decl is always a function type
	RefsType * NodeTypeFunctionType `node: "contained,single"`

	// the identifier name of the function
	RefsName * NodeTypeIdentifierNode `node: "contained,single"`
}

// just recuse into the type given
type Recurse struct {
	Base NodeBase
}

type NodeTypeParamList struct {
	// list of node types
	Base NodeBase
	RefsChain Recurse
	RefsValu TypeInterface
}

type NodeTypeFunctionType struct {
	Base NodeBase
	RefsPrms * NodeTypeParamList `node: "contained,recurse"`
	RefsRetn TypeInterface `node: "reference"`
	RefsSize * NodeTypeIntegerCst `node: "reference"`
}

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
	RefsPtd TypeInterface `node: "reference"`
}

type NodeTypeRecordType struct {
	Base NodeBase
	RefsFlds * NodeTypeFieldDecl `node: "contained,recurse"`
	RefsSize * NodeTypeIntegerCst 
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
	RefsSize * NodeTypeIntegerCst `node: "reference"`
	RefsUnql * NodeTypeUnionType `node: "reference,backwards"`
	RefsFlds * NodeTypeFieldDecl `node: "contained,recurse"`
	RefsName NameInterface `node: "name"`
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
