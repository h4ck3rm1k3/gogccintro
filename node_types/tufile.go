package node_types
import (
	"github.com/h4ck3rm1k3/gogccintro/models"
	"github.com/h4ck3rm1k3/gogccintro/tree"
	"database/sql"
	
	"fmt"
)
// func (t * TUFile) LookupId(id int64) NodeInterface {
// 	return t.Ids[id]
// }

type TuFileInterface interface {
	LookupId(int) NodeInterface
}

func (t* TUFile) LookupGccNode(id int64) * models.GccTuParserNode {
	if v,ok := t.Tree.Nodes[int(id)]; ok { ///TODO
		return v
	} else {
		fmt.Printf("did not find id %s\n", id)
	}
	panic("Could not find node")
	return nil
}

type TUFile struct {
	SourceFileID int
	//Ids   map[int64] * models.GccTuParserNode// raw data
	Tree * tree.TreeMap
	NodeTypeIntegerTypeMap map[int64] * NodeTypeIntegerType
	NodeTypeTreeListMap map[int64] * NodeTypeTreeList
	NodeTypeIntegerCstMap map[int64] * NodeTypeIntegerCst
	NodeTypeIdentifierNodeMap map[int64] * NodeTypeIdentifierNode
	NodeTypeArrayTypeMap map[int64] * NodeTypeArrayType
	NodeTypeTypeDeclMap map[int64] * NodeTypeTypeDecl
	NodeTypePointerTypeMap map[int64] * NodeTypePointerType
	NodeTypeFieldDeclMap map[int64] * NodeTypeFieldDecl
	NodeTypeUnionTypeMap map[int64] * NodeTypeUnionType
	NodeTypeVoidTypeMap map[int64] * NodeTypeVoidType
	NodeTypeFunctionDeclMap map[int64] *NodeTypeFunctionDecl
	NodeTypeFunctionTypeMap map[int64] *NodeTypeFunctionType
	NodeTypeRecordTypeMap map[int64] *NodeTypeRecordType

	NodeTypeParamListMap map[int64] *NodeTypeParamList
	// interfaces
	TypeInterfaceMap map[int64] TypeInterface
	NameInterfaceMap map[int64] NameInterface
		
}


func (t * TUFile) LookupNodeType(name string) NodeTypeGeneric{
	return NodePrototypes[name]
}

func (t * TUFile) CreateBase(from *models.GccTuParserNode) NodeBase{
	if from == nil {
		panic("null base")
	} 
	if t.SourceFileID!=from.SourceFileID{
		fmt.Errorf("file %d!=%d", t.SourceFileID,from.SourceFileID)
	}
	
	return NodeBase{
		NodeID : from.NodeID,
		File : t,
		NodeType: t.LookupNodeType(from.NodeType),
	}
}

//////////////////// generated following


func (t * TUFile) CreateNodeTypeTreeList(from *models.GccTuParserNode ) *NodeTypeTreeList {
	return &NodeTypeTreeList{
		Base : t.CreateBase(from),
	}
}

func (t * TUFile) CreateNodeTypeFunctionDecl(from *models.GccTuParserNode ) *NodeTypeFunctionDecl {
	return &NodeTypeFunctionDecl{
		RefsType: t.CreateRefNodeTypeFunctionType(from.RefsType),
		RefsName: t.CreateRefNodeTypeIdentifierNode(from.RefsName),
		Base : t.CreateBase(from),
	}
}
func (t * TUFile) CreateNodeTypeTypeDecl(from *models.GccTuParserNode ) *NodeTypeTypeDecl {
	return &NodeTypeTypeDecl{
		Base : t.CreateBase(from),
		Name : t.CreateRefNameInterface(from.RefsName),
	}
}
func (t * TUFile) CreateNodeTypeIntegerCst(from *models.GccTuParserNode ) *NodeTypeIntegerCst {
	return &NodeTypeIntegerCst{
		Base : t.CreateBase(from),
		AttrsType: t.CreateRefNodeTypeIntegerType(NodeIdFromString(from.AttrsType)),
	}
}
func (t * TUFile) CreateNodeTypeIdentifierNode(from *models.GccTuParserNode ) *NodeTypeIdentifierNode {
	return &NodeTypeIdentifierNode{
		Base : t.CreateBase(from),
	}
}
func (t * TUFile) CreateNodeTypeRecordType(from *models.GccTuParserNode ) *NodeTypeRecordType {
	return &NodeTypeRecordType{
		RefsFlds: t.CreateRefNodeTypeFieldDecl(from.RefsFlds),
		RefsSize: t.CreateRefNodeTypeIntegerCst(from.RefsSize),
		Base : t.CreateBase(from),
	}
}

func (t * TUFile) CreateNodeTypeFieldDecl(from *models.GccTuParserNode ) *NodeTypeFieldDecl {
	return &NodeTypeFieldDecl{
		RefsBpos: t.CreateRefNodeTypeIntegerCst(from.RefsBpos),
		Base : t.CreateBase(from),
	}
}

func (t * TUFile) CreateNodeTypeIntegerType(from *models.GccTuParserNode ) *NodeTypeIntegerType {
	return &NodeTypeIntegerType{
		Base : t.CreateBase(from),
	}
}
func (t * TUFile) CreateNodeTypeUnionType(from *models.GccTuParserNode ) *NodeTypeUnionType {
	return &NodeTypeUnionType{
		RefsFlds: t.CreateRefNodeTypeFieldDecl(from.RefsFlds),
		RefsSize: t.CreateRefNodeTypeIntegerCst(from.RefsSize),
		Base : t.CreateBase(from),
		RefsName: t.CreateRefNameInterface(from.RefsName),
		RefsUnql: t.CreateRefNodeTypeUnionType(from.RefsUnql),
	}
}
func (t * TUFile) CreateNodeTypeVoidType(from *models.GccTuParserNode ) *NodeTypeVoidType {
	return &NodeTypeVoidType{
		Base : t.CreateBase(from),
	}
}
func (t * TUFile) CreateNodeTypeArrayType(from *models.GccTuParserNode ) *NodeTypeArrayType {
	return &NodeTypeArrayType{
		RefsDomn: t.CreateRefTypeInterface(from.RefsDomn),
		RefsSize: t.CreateRefNodeTypeIntegerCst(from.RefsSize),
		Base : t.CreateBase(from),
		RefsElts: t.CreateRefTypeInterface(from.RefsElts),
	}
}
func (t * TUFile) CreateNodeTypeFunctionType(from *models.GccTuParserNode ) *NodeTypeFunctionType {
	return &NodeTypeFunctionType{
		RefsRetn: t.CreateRefTypeInterface(from.RefsRetn),
		RefsPrms: t.CreateRefNodeTypeParamList(from.RefsPrms),
		RefsSize: t.CreateRefNodeTypeIntegerCst(from.RefsSize),
		Base : t.CreateBase(from),
	}
}

func (t * TUFile) CreateNodeTypePointerType(from *models.GccTuParserNode ) *NodeTypePointerType {
	return &NodeTypePointerType{
		Base : t.CreateBase(from),
		RefsPtd: t.CreateRefTypeInterface(from.RefsPtd),
	}
}

/// Refs


func (t * TUFile) CreateRefNodeTypeIntegerType(id sql.NullInt64 ) *NodeTypeIntegerType {
	if id.Valid {
		if node,ok := t.NodeTypeIntegerTypeMap[id.Int64]; ok {
			return node
		} else {
			return CreateNodeTypeIntegerType(t.LookupGccNode(id.Int64))
		}	
	} else {
		// not set
		return nil
	}
}

func (t * TUFile) CreateRefNodeTypeTreeList(id sql.NullInt64 ) *NodeTypeTreeList {
	if id.Valid {
		if node,ok := t.NodeTypeTreeListMap[id.Int64]; ok {
			return node
		} else {
			return CreateNodeTypeTreeList(t.LookupGccNode(id.Int64))
		}	
	} else {
		return nil
	}
}

func (t * TUFile) CreateRefNodeTypeIntegerCst(id sql.NullInt64 ) *NodeTypeIntegerCst {
	if id.Valid {
		if node,ok := t.NodeTypeIntegerCstMap[id.Int64]; ok {
			return node
		} else {
			return CreateNodeTypeIntegerCst(t.LookupGccNode(id.Int64))
		}	
	} else {
		return nil
	}
}

func (t * TUFile) CreateRefNodeTypeIdentifierNode(id sql.NullInt64 ) *NodeTypeIdentifierNode {
	if id.Valid {
		if node,ok := t.NodeTypeIdentifierNodeMap[id.Int64]; ok {
			return node
		} else {
			return CreateNodeTypeIdentifierNode(t.LookupGccNode(id.Int64))
		}	
	} else {
		return nil
	}
}

func (t * TUFile) CreateRefNodeTypeArrayType(id sql.NullInt64 ) *NodeTypeArrayType {
	if id.Valid {
		if node,ok := t.NodeTypeArrayTypeMap[id.Int64]; ok {
			return node
		} else {
			return CreateNodeTypeArrayType(t.LookupGccNode(id.Int64))
		}	
	} else {
		return nil
	}
}

func (t * TUFile) CreateRefNodeTypeTypeDecl(id sql.NullInt64 ) *NodeTypeTypeDecl {
	if id.Valid {
		if node,ok := t.NodeTypeTypeDeclMap[id.Int64]; ok {
			return node
		} else {
			return CreateNodeTypeTypeDecl(t.LookupGccNode(id.Int64))
		}	
	} else {
		return nil
	}

}


func (t * TUFile) CreateRefNodeTypePointerType(id sql.NullInt64 ) *NodeTypePointerType {
	if id.Valid {
		if node,ok := t.NodeTypePointerTypeMap[id.Int64]; ok {
			return node
		} else {
			return CreateNodeTypePointerType(t.LookupGccNode(id.Int64))
		}	
	} else {
		return nil
	}

}


func (t * TUFile) CreateRefNodeTypeFieldDecl(id sql.NullInt64 ) *NodeTypeFieldDecl {
	if id.Valid {
		if node,ok := t.NodeTypeFieldDeclMap[id.Int64]; ok {
			return node
		} else {
			return CreateNodeTypeFieldDecl(t.LookupGccNode(id.Int64))
		}	
	} else {
		return nil
	}

}


func (t * TUFile) CreateRefNodeTypeUnionType(id sql.NullInt64 ) *NodeTypeUnionType {
	if id.Valid {
		if node,ok := t.NodeTypeUnionTypeMap[id.Int64]; ok {
			return node
		} else {
			return CreateNodeTypeUnionType(t.LookupGccNode(id.Int64))
		}	
	} else {
		return nil
	}
}

func (t * TUFile) CreateRefNodeTypeVoidType(id sql.NullInt64 ) *NodeTypeVoidType {
	if id.Valid {
		if node,ok := t.NodeTypeVoidTypeMap[id.Int64]; ok {
			return node
		} else {
			return CreateNodeTypeVoidType(t.LookupGccNode(id.Int64))
		}	
	} else { return nil }}

func (t * TUFile) CreateRefNodeTypeFunctionType(id sql.NullInt64 ) *NodeTypeFunctionType {
	if id.Valid {
		if node,ok := t.NodeTypeFunctionTypeMap[id.Int64]; ok {
			return node
		} else {
			return CreateNodeTypeFunctionType(t.LookupGccNode(id.Int64))
		}	
	} else { return nil }}

func (t * TUFile) CreateRefNodeTypeFunctionDecl(id sql.NullInt64 ) *NodeTypeFunctionDecl {
	if id.Valid {
		if node,ok := t.NodeTypeFunctionDeclMap[id.Int64]; ok {
			return node
		} else {
			return CreateNodeTypeFunctionDecl(t.LookupGccNode(id.Int64))
		}	
	} else { return nil }}

func (t * TUFile) CreateRefNodeTypeRecordType(id sql.NullInt64 ) *NodeTypeRecordType {
	if id.Valid {
		if node,ok := t.NodeTypeRecordTypeMap[id.Int64]; ok {
			return node
		} else {
			return CreateNodeTypeRecordType(t.LookupGccNode(id.Int64))
		}	
	} else { return nil }}

func CreateTUFile() *TUFile {
	return & TUFile {
		NodeTypeIntegerTypeMap : make(map[int64] * NodeTypeIntegerType),
		NodeTypeTreeListMap : make(map[int64] * NodeTypeTreeList),
		NodeTypeIntegerCstMap : make(map[int64] * NodeTypeIntegerCst),
		NodeTypeRecordTypeMap : make(map[int64] * NodeTypeRecordType),
		NodeTypeFunctionDeclMap : make(map[int64] * NodeTypeFunctionDecl),
		NodeTypeFunctionTypeMap : make(map[int64] * NodeTypeFunctionType),
		NodeTypeUnionTypeMap : make(map[int64] * NodeTypeUnionType),
		NodeTypeVoidTypeMap : make(map[int64] * NodeTypeVoidType),
		NodeTypeFieldDeclMap : make(map[int64] * NodeTypeFieldDecl),
		NodeTypePointerTypeMap : make(map[int64] * NodeTypePointerType),
		NodeTypeTypeDeclMap : make(map[int64] * NodeTypeTypeDecl),
		NodeTypeArrayTypeMap : make(map[int64] * NodeTypeArrayType),
		NodeTypeIdentifierNodeMap : make(map[int64] * NodeTypeIdentifierNode),
		NodeTypeParamListMap: make(map[int64] *NodeTypeParamList),
		//
		TypeInterfaceMap :make( map[int64] TypeInterface),
		NameInterfaceMap :make( map[int64] NameInterface),
		
	}
}

func(t *NodeTypePointerType) GetRefsSize() * NodeTypeIntegerCst {
	return t.RefsSize
}

func(t *NodeTypeVoidType) GetRefsSize() * NodeTypeIntegerCst {
	return t.RefsSize
}

func(t *NodeTypeFunctionType) GetRefsSize() * NodeTypeIntegerCst {
	return t.RefsSize
}

func(t *NodeTypeIntegerType) GetRefsSize() * NodeTypeIntegerCst {
	return t.RefsSize
}

func(t *NodeTypeArrayType) GetRefsSize() * NodeTypeIntegerCst {
	return t.RefsSize
}

func(t *NodeTypeUnionType) GetRefsSize() * NodeTypeIntegerCst {
	return t.RefsSize
}

func(t *NodeTypeRecordType) GetRefsSize() * NodeTypeIntegerCst {
	return t.RefsSize
}

// generics creating interfaces
func (t * TUFile) CreateRefTypeInterface(id sql.NullInt64 ) TypeInterface {
	if id.Valid {
		if node,ok := t.TypeInterfaceMap[id.Int64]; ok {
			return node
		} else {
			v := t.LookupGccNode(id.Int64)
			if v == nil {
				panic("null")
				return nil
			}
			//switch on the type
			fmt.Printf("TODO switch on type %s",v)
			nt := v.NodeType
			switch (nt) {
			case "pointer_type":
				return t.CreateNodeTypePointerType(v)
				break
			case "void_type":
				return t.CreateNodeTypeVoidType(v)
				break
			case "function_type":
				return t.CreateNodeTypeFunctionType(v)
				break
			case "integer_type":
				return t.CreateNodeTypeIntegerType(v)
				break
			case "array_type":
				return t.CreateNodeTypeArrayType(v)
				break
			case "union_type":
				return t.CreateNodeTypeUnionType(v)
				break
			case "record_type":
				return t.CreateNodeTypeRecordType(v)
				break

				default :
				fmt.Errorf("unhandled node type%s", v.NodeType)
				panic("null")
				return nil
				break
			}
				

			//return CreateNodeTypeFunctionDecl(t.LookupGccNode(id.Int64))
		}	
	}
	panic("null")
	return nil
}

func (t*NodeTypeIdentifierNode) String() string {
	return t.Name
}

func (t*NodeTypeTypeDecl) String() string {
	return t.Name.String()
}

// generics creating interfaces
func (t * TUFile) CreateRefNameInterface(id sql.NullInt64) NameInterface {
	if id.Valid {
		if node,ok := t.NameInterfaceMap[id.Int64]; ok {
			return node
		} else {
			v := t.LookupGccNode(id.Int64)
			if v == nil {
					panic("null")
				return nil }
			//switch on the type
			//fmt.Printf("TODO switch on type %s",v)
				switch (v.NodeType) {
				case "identifier_node":
					return t.CreateNodeTypeIdentifierNode(v)
					break
				case "type_decl":
					return t.CreateNodeTypeTypeDecl(v)
					break
				default:
					fmt.Errorf("error: %s\n", v.NodeType)
					panic("null")
					return nil
				}

			//return CreateNodeTypeFunctionDecl(t.LookupGccNode(id.Int64))
		}	
	}
	panic("null")
	return nil
}


// param list
func (t * TUFile) CreateRefNodeTypeParamList(id sql.NullInt64) *NodeTypeParamList {
	if id.Valid {
		if node,ok := t.NodeTypeParamListMap[id.Int64]; ok {
			return node
		} else {
			return t.CreateNodeTypeParamList(t.LookupGccNode(id.Int64))
		}	
	} else { return nil }}

func (t * TUFile) CreateNodeTypeParamList(from *models.GccTuParserNode ) *NodeTypeParamList {
	return &NodeTypeParamList{
		Base : t.CreateBase(from),
	}
}
