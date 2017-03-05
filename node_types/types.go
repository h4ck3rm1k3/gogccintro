package node_types

type TypeInterface interface {}

type NodeTypeArrayType struct {
	RefsSize NodeTypeIntegerCst
	RefsDomn TypeInterface
	RefsElts TypeInterface
}

type NodeTypeFieldDecl struct {
	RefsBpos NodeTypeIntegerCst
}

type NodeTypeFunctionDecl struct {

	// the type of the function decl is always a function type
	RefsType * NodeTypeFunctionType

	// the identifier name of the function
	RefsName * NodeTypeIdentifierNode
}

type NodeTypeList struct {
	// list of node types
}

type NodeTypeFunctionType struct {
	RefsPrms NodeTypeList
	RefsRetn TypeInterface
	RefsSize NodeTypeIntegerCst
}

type NodeTypeIdentifierNode struct {
	Names string
}

type MinMaxMixin struct {
	RefsMax NodeTypeIntegerCst
	RefsMin NodeTypeIntegerCst
}

type TypeMixin struct {
	RefsUnql TypeInterface 
}

type NameInterface interface {
	// can be an identifier node or type decl
}

type NamedMixin struct {
	RefsName NameInterface
}

type NodeTypeIntegerCst struct {
	AttrsType NodeTypeIntegerType
}
type NodeTypeIntegerType struct {}

type NodeTypePointerType struct {
	// what is pointed to
	RefsPtd TypeInterface 
}

type NodeTypeRecordType struct {
	RefsFlds NodeTypeFieldDecl
	RefsSize NodeTypeIntegerCst
}

// type NodeTypeTreeList struct {}
type NodeTypeTypeDecl struct {
	
}

type NodeTypeUnionType struct {
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

