package node_types
import (
	"github.com/h4ck3rm1k3/gogccintro/models"
	"github.com/h4ck3rm1k3/gogccintro/tree"
	"database/sql"
	//"encoding/json"
	"fmt"
)
// func (t * TUFile) LookupId(id int64) NodeInterface {
// 	return t.Ids[id]
// }

type TuFileInterface interface {
	LookupId(int) NodeInterface
}

func (t* TUFile) LookupGccNode(id int64) * models.GccTuParserNode {
	//fmt.Printf("lookup %s\n", id)
	if v,ok := t.Tree.Nodes[int(id)]; ok {
		return v
	} else {
		fmt.Printf("did not find id %s\n", id)
	}
	panic("Could not find node")
	return nil
}

type TUFile struct {
	SourceFileID int
	Ids   map[int] NodeInterface// resulting data
	Tree * tree.TreeMap
	NodeTypeIntegerTypeMap map[int] * NodeTypeIntegerType
	NodeTypeTreeListMap map[int] * NodeTypeTreeList
	NodeTypeIntegerCstMap map[int] * NodeTypeIntegerCst
	NodeTypeIdentifierNodeMap map[int] * NodeTypeIdentifierNode
	NodeTypeArrayTypeMap map[int] * NodeTypeArrayType
	NodeTypeTypeDeclMap map[int] * NodeTypeTypeDecl
	NodeTypePointerTypeMap map[int] * NodeTypePointerType
	NodeTypeFieldDeclMap map[int] * NodeTypeFieldDecl
	NodeTypeUnionTypeMap map[int] * NodeTypeUnionType
	NodeTypeVoidTypeMap map[int] * NodeTypeVoidType
	NodeTypeFunctionDeclMap map[int] *NodeTypeFunctionDecl
	NodeTypeFunctionTypeMap map[int] *NodeTypeFunctionType
	NodeTypeRecordTypeMap map[int] *NodeTypeRecordType

	//NodeTypeParamListMap map[int] *NodeTypeParamList
	
	// interfaces
	TypeInterfaceMap map[int] TypeInterface
	NameInterfaceMap map[int] NameInterface
		
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
		//File : t,
		NodeTypeName: NodeType{
			TypeName: from.NodeType,
			NodeType: t.LookupNodeType(from.NodeType),
		},
	}
}

//////////////////// generated following


func (t * TUFile) CreateNodeTypeTreeList(from *models.GccTuParserNode ) *NodeTypeTreeList {
	if from == nil {return nil }
	r:= &NodeTypeTreeList{
		Base : t.CreateBase(from),
		RefsChan: t.CreateRefNodeTypeTreeList(from.RefsChan),
	}
	t.NodeTypeTreeListMap[int(from.NodeID)]=r
	return r
}

func (t * TUFile) CreateNodeTypeFunctionDecl(from *models.GccTuParserNode ) *NodeTypeFunctionDecl {
	if from == nil {return nil }
	rt := t.CreateRefNodeTypeFunctionType(from.RefsType)
	//fmt.Printf("created func decl (type)1 %v\n",rt)
	r:= &NodeTypeFunctionDecl{
		RefsType: rt,
		RefsName: t.CreateRefNodeTypeIdentifierNode(from.RefsName),
		RefsArgs: t.CreateRefNodeTypeArgList(from.RefsArgs),
		Base : t.CreateBase(from),
	}
	//fmt.Printf("created func decl %s\n",r)
	//fmt.Printf("created func decl (type) %s\n",r.RefsType)
	t.NodeTypeFunctionDeclMap[int(from.NodeID)]=r
	return r
	
}

func (t * TUFile) CreateNodeTypeTypeDecl(from *models.GccTuParserNode ) *NodeTypeTypeDecl {
	if from == nil {return nil }
	r:= &NodeTypeTypeDecl{
		Base : t.CreateBase(from),
		RefsName : t.CreateRefNameInterface(from.RefsName),
	}
	r.RefsName.Named(r)
	t.NodeTypeTypeDeclMap[int(from.NodeID)]=r
	return r
}

func (t * TUFile) CreateNodeTypeIntegerCst(from *models.GccTuParserNode ) *NodeTypeIntegerCst {
	if from == nil {return nil }
	r:= &NodeTypeIntegerCst{
		Base : t.CreateBase(from),
		AttrsType: t.CreateRefNodeTypeIntegerType(NodeIdFromString(from.AttrsType)),
	}
	t.NodeTypeIntegerCstMap[int(from.NodeID)]=r
	return r
}

func (t * TUFile) CreateNodeTypeIdentifierNode(from *models.GccTuParserNode ) *NodeTypeIdentifierNode {
	if from == nil {return nil }
	r:= &NodeTypeIdentifierNode{
		Base : t.CreateBase(from),
		StringVal : from.AttrsString,
	}
	t.NodeTypeIdentifierNodeMap[int(from.NodeID)]=r
	return r
}

func (t * TUFile) CreateNodeTypeRecordType(from *models.GccTuParserNode ) *NodeTypeRecordType {
	if from == nil {return nil }
	r:= &NodeTypeRecordType{
		RefsFlds: t.CreateRefNodeTypeFieldDecl(from.RefsFlds),
		RefsSize: t.CreateRefNodeTypeIntegerCst(from.RefsSize),
		Base : t.CreateBase(from),
	}
	t.NodeTypeRecordTypeMap[int(from.NodeID)]=r
	return r
}

func (t * TUFile) CreateNodeTypeFieldDecl(from *models.GccTuParserNode ) *NodeTypeFieldDecl {
	if from == nil {return nil }
	r:= &NodeTypeFieldDecl{
		RefsBpos: t.CreateRefNodeTypeIntegerCst(from.RefsBpos),
		Base : t.CreateBase(from),
	}
	t.NodeTypeFieldDeclMap[int(from.NodeID)]=r
	return r
}

func (t * TUFile) CreateNodeTypeIntegerType(from *models.GccTuParserNode ) *NodeTypeIntegerType {
	if from == nil {return nil }
	r:= &NodeTypeIntegerType{
		Base : t.CreateBase(from),
	}
	t.NodeTypeIntegerTypeMap[int(from.NodeID)]=r
	return r
}

func (t * TUFile) CreateNodeTypeUnionType(from *models.GccTuParserNode ) *NodeTypeUnionType {
	if from == nil {return nil }
	r:= &NodeTypeUnionType{
		RefsFlds: t.CreateRefNodeTypeFieldDecl(from.RefsFlds),
		RefsSize: t.CreateRefNodeTypeIntegerCst(from.RefsSize),
		Base : t.CreateBase(from),
		RefsName: t.CreateRefNameInterface(from.RefsName),
		RefsUnql: t.CreateRefNodeTypeUnionType(from.RefsUnql),
	}
	r.RefsName.Named(r)
	t.NodeTypeUnionTypeMap[int(from.NodeID)]=r
	return r
}

func (t * TUFile) CreateNodeTypeVoidType(from *models.GccTuParserNode ) *NodeTypeVoidType {
	if from == nil {return nil }
	r:= &NodeTypeVoidType{
		Base : t.CreateBase(from),
	}
	t.NodeTypeVoidTypeMap[int(from.NodeID)]=r
	return r
}

func (t * TUFile) CreateNodeTypeArrayType(from *models.GccTuParserNode ) *NodeTypeArrayType {
	if from == nil {return nil }
	r:= &NodeTypeArrayType{
		RefsDomn: t.CreateRefTypeInterface(from.RefsDomn),
		RefsSize: t.CreateRefNodeTypeIntegerCst(from.RefsSize),
		Base : t.CreateBase(from),
		RefsElts: t.CreateRefTypeInterface(from.RefsElts),
	}
	t.NodeTypeArrayTypeMap[int(from.NodeID)]=r
	return r
}

func (t * TUFile) CreateNodeTypeFunctionType(from *models.GccTuParserNode ) *NodeTypeFunctionType {
	// fmt.Printf("create funct tpe %s %v %v %v\n", from,
	// 	from.RefsRetn,
	// 	from.RefsPrms,
	// 	from.RefsSize,
	// )
	if from == nil {return nil }
	r:= &NodeTypeFunctionType{
		RefsRetn: t.CreateRefTypeInterface(from.RefsRetn),
		RefsPrms: t.CreateRefNodeTypeParamList(from.RefsPrms),
		RefsSize: t.CreateRefNodeTypeIntegerCst(from.RefsSize),
		Base : t.CreateBase(from),
	}
	t.NodeTypeFunctionTypeMap[int(from.NodeID)]=r

	//fmt.Printf("created funct tpe %s %v %v %v\n", r)
	return r
}

func (t * TUFile) CreateNodeTypePointerType(from *models.GccTuParserNode ) *NodeTypePointerType {
	if from == nil {return nil }
	r := &NodeTypePointerType{
		Base : t.CreateBase(from),
		RefsPtd: t.CreateRefTypeInterface(from.RefsPtd),
	}
	t.NodeTypePointerTypeMap[int(from.NodeID)]=r
	return r
}

/// Refs


func (t * TUFile) CreateRefNodeTypeIntegerType(id sql.NullInt64 ) *NodeTypeIntegerType {
	if id.Valid {
		if node,ok := t.NodeTypeIntegerTypeMap[int(id.Int64)]; ok {
			return node
		} else {
			return t.CreateNodeTypeIntegerType(t.LookupGccNode(id.Int64))
		}	
	} else {
		// not set
		return nil
	}
}

func (t * TUFile) CreateRefNodeTypeTreeList(id sql.NullInt64 ) *NodeTypeTreeList {
	if id.Valid {
		if node,ok := t.NodeTypeTreeListMap[int(id.Int64)]; ok {
			return node
		} else {
			return t.CreateNodeTypeTreeList(t.LookupGccNode(id.Int64))
		}	
	} else {
		return nil
	}
}

func (t * TUFile) CreateRefNodeTypeIntegerCst(id sql.NullInt64 ) *NodeTypeIntegerCst {
	if id.Valid {
		if node,ok := t.NodeTypeIntegerCstMap[int(id.Int64)]; ok {
			return node
		} else {
			return t.CreateNodeTypeIntegerCst(t.LookupGccNode(id.Int64))
		}	
	} else {
		return nil
	}
}

func (t * TUFile) CreateRefNodeTypeIdentifierNode(id sql.NullInt64 ) *NodeTypeIdentifierNode {
	if id.Valid {
		if node,ok := t.NodeTypeIdentifierNodeMap[int(id.Int64)]; ok {
			return node
		} else {
			return t.CreateNodeTypeIdentifierNode(t.LookupGccNode(id.Int64))
		}	
	} else {
		return nil
	}
}

func (t * TUFile) CreateRefNodeTypeArrayType(id sql.NullInt64 ) *NodeTypeArrayType {
	if id.Valid {
		if node,ok := t.NodeTypeArrayTypeMap[int(id.Int64)]; ok {
			return node
		} else {
			return t.CreateNodeTypeArrayType(t.LookupGccNode(id.Int64))
		}	
	} else {
		return nil
	}
}

func (t * TUFile) CreateRefNodeTypeTypeDecl(id sql.NullInt64 ) *NodeTypeTypeDecl {
	if id.Valid {
		if node,ok := t.NodeTypeTypeDeclMap[int(id.Int64)]; ok {
			return node
		} else {
			return t.CreateNodeTypeTypeDecl(t.LookupGccNode(id.Int64))
		}	
	} else {
		return nil
	}
}

func (t * TUFile) CreateRefNodeTypePointerType(id sql.NullInt64 ) *NodeTypePointerType {
	if id.Valid {
		if node,ok := t.NodeTypePointerTypeMap[int(id.Int64)]; ok {
			return node
		} else {
			return t.CreateNodeTypePointerType(t.LookupGccNode(id.Int64))
		}	
	} else {
		return nil
	}

}


func (t * TUFile) CreateRefNodeTypeFieldDecl(id sql.NullInt64 ) *NodeTypeFieldDecl {
	if id.Valid {
		if node,ok := t.NodeTypeFieldDeclMap[int(id.Int64)]; ok {
			return node
		} else {
			return t.CreateNodeTypeFieldDecl(t.LookupGccNode(id.Int64))
		}	
	} else {
		return nil
	}

}


func (t * TUFile) CreateRefNodeTypeUnionType(id sql.NullInt64 ) *NodeTypeUnionType {
	if id.Valid {
		if node,ok := t.NodeTypeUnionTypeMap[int(id.Int64)]; ok {
			return node
		} else {
			return t.CreateNodeTypeUnionType(t.LookupGccNode(id.Int64))
		}	
	} else {
		return nil
	}
}

func (t * TUFile) CreateRefNodeTypeVoidType(id sql.NullInt64 ) *NodeTypeVoidType {
	if id.Valid {
		if node,ok := t.NodeTypeVoidTypeMap[int(id.Int64)]; ok {
			return node
		} else {
			return t.CreateNodeTypeVoidType(t.LookupGccNode(id.Int64))
		}	
	} else { return nil }
}

func (t * TUFile) CreateRefNodeTypeFunctionType(id sql.NullInt64 ) *NodeTypeFunctionType {
	//fmt.Printf("create reffunct type %v\n", id)
	if id.Valid {
		if node,ok := t.NodeTypeFunctionTypeMap[int(id.Int64)]; ok {
			return node
		} else {
			v := t.LookupGccNode(id.Int64)
			//fmt.Printf("func type lookup node %v\n", v)
			r := t.CreateNodeTypeFunctionType( v)
			//fmt.Printf("create reffunct type2 %v\n", r)
			return r
		}	
	} else {
		return nil }}

func (t * TUFile) CreateRefNodeTypeFunctionDecl(id sql.NullInt64 ) *NodeTypeFunctionDecl {
	if id.Valid {
		if node,ok := t.NodeTypeFunctionDeclMap[int(id.Int64)]; ok {
			return node
		} else {
			return t.CreateNodeTypeFunctionDecl(t.LookupGccNode(id.Int64))
		}	
	} else { return nil }}

func (t * TUFile) CreateRefNodeTypeRecordType(id sql.NullInt64 ) *NodeTypeRecordType {
	if id.Valid {
		if node,ok := t.NodeTypeRecordTypeMap[int(id.Int64)]; ok {
			return node
		} else {
			return t.CreateNodeTypeRecordType(t.LookupGccNode(id.Int64))
		}	
	} else { return nil }}

func CreateTUFile() *TUFile {
	return & TUFile {
		NodeTypeIntegerTypeMap : make(map[int] * NodeTypeIntegerType),
		NodeTypeTreeListMap : make(map[int] * NodeTypeTreeList),
		NodeTypeIntegerCstMap : make(map[int] * NodeTypeIntegerCst),
		NodeTypeRecordTypeMap : make(map[int] * NodeTypeRecordType),
		NodeTypeFunctionDeclMap : make(map[int] * NodeTypeFunctionDecl),
		NodeTypeFunctionTypeMap : make(map[int] * NodeTypeFunctionType),
		NodeTypeUnionTypeMap : make(map[int] * NodeTypeUnionType),
		NodeTypeVoidTypeMap : make(map[int] * NodeTypeVoidType),
		NodeTypeFieldDeclMap : make(map[int] * NodeTypeFieldDecl),
		NodeTypePointerTypeMap : make(map[int] * NodeTypePointerType),
		NodeTypeTypeDeclMap : make(map[int] * NodeTypeTypeDecl),
		NodeTypeArrayTypeMap : make(map[int] * NodeTypeArrayType),
		NodeTypeIdentifierNodeMap : make(map[int] * NodeTypeIdentifierNode),
		//NodeTypeParamListMap: make(map[int] *NodeTypeParamList),
		//
		TypeInterfaceMap :make( map[int] TypeInterface),
		NameInterfaceMap :make( map[int] NameInterface),

		Ids: make(map[int]NodeInterface), // result of create
		
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
		if node,ok := t.TypeInterfaceMap[int(id.Int64)]; ok {
			return node
		} else {
			v := t.LookupGccNode(id.Int64)
			if v == nil {
				panic("null")
				return nil
			}
			//switch on the type
			nt := v.NodeType
			switch (nt) {
			case "pointer_type":
				return t.CreateNodeTypePointerType(v)
				break
			case "void_type":
				return t.CreateNodeTypeVoidType(v)
				break
			case "function_type":
				r:= t.CreateNodeTypeFunctionType(v)
//				fmt.Printf("%s",r)
				return r
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
	return t.StringVal
}

func (t*NodeTypeTypeDecl) String() string {
	return t.RefsName.String()
}

// generics creating interfaces
func (t * TUFile) CreateRefNameInterface(id sql.NullInt64) NameInterface {
	if id.Valid {
		if node,ok := t.NameInterfaceMap[int(id.Int64)]; ok {
			return node
		} else {
			v := t.LookupGccNode(id.Int64)
			if v == nil {
				panic("null")
				return nil
			}
			var r NameInterface
			//switch on the type
				switch (v.NodeType) {
				case "identifier_node":
					r= t.CreateNodeTypeIdentifierNode(v)
					break
				case "type_decl":
					r= t.CreateNodeTypeTypeDecl(v)
					break
				default:
					fmt.Errorf("error: %s\n", v.NodeType)
					panic("null")
				}
			return r
			//return CreateNodeTypeFunctionDecl(t.LookupGccNode(id.Int64))
		}	
	}
	panic("null")
	return nil
}


// param list
func (t * TUFile) CreateRefNodeTypeParamList(id sql.NullInt64) *NodeTypeParamList {
	if id.Valid {
		//if node,ok := t.NodeTypeParamListMap[int(id.Int64)]; ok {
		//	return node
		//} else {
		return t.CreateNodeTypeParamList(t.LookupGccNode(id.Int64))
		//}	
	} else { return nil }}

func (t * TUFile) CreateRefNodeTypeArgList(id sql.NullInt64) *NodeTypeArgList {
	if id.Valid {
		//if node,ok := t.NodeTypeArgListMap[int(id.Int64)]; ok {
		//	return node
		//} else {
		return t.CreateNodeTypeArgList(t.LookupGccNode(id.Int64))
		//}	
	} else { return nil }}


func (t * TUFile) CreateNodeTypeParamList(from *models.GccTuParserNode ) *NodeTypeParamList {
	return &NodeTypeParamList{
		Base : t.CreateBase(from),
		RefsChan : t.CreateRefNodeTypeParamList(from.RefsChan),
		RefsValu : t.CreateRefNodeGeneric(from.RefsValu),
	}
}

func (t * TUFile) CreateNodeTypeArgList(from *models.GccTuParserNode ) *NodeTypeArgList {
	return &NodeTypeArgList{
		Base : t.CreateBase(from),
		RefsChan : t.CreateRefNodeTypeArgList(from.RefsChan),
		RefsValu : t.CreateRefNodeGeneric(from.RefsValu),
	}
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

func (t * TUFile) CreateRefNodeGeneric(id sql.NullInt64) NodeInterface {
	if id.Valid {
		//if node,ok := t.NodeInterfaceMap[int(id.Int64)]; ok {
		//	return node
		//} else {
			v := t.LookupGccNode(id.Int64)
			if v == nil {
					panic("null")
				return nil }

			return t.DispatchType(v)
//		}	
	}
	panic("null")
	return nil
}

func (t * TUFile) handle_type(v NodeInterface) NodeInterface{

	t.Ids[v.NodeId()]=v
	
	// is the type added
	return v
}
