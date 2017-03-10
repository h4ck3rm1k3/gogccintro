package node_types

import (
	_ "reflect"
	_ "regexp"
	"database/sql"
	_ "fmt"
	"github.com/h4ck3rm1k3/gogccintro/models"
	//"github.com/h4ck3rm1k3/gogccintro/node_types"
)

// not generated
func (to *NodeBase) MappingGccTuParserNode(from *models.GccTuParserNode){
}

func (t*NodeTypeIntegerType) MappingAttrsType(node_id string){}

func foo(to NodeInterface, from sql.NullInt64) {}

//func (t*NodeTypeIntegerCst) MappingReference(from *models.GccTuParserNode){}



func (to *NodeTypeIntegerType) MappingGccTuParserNode(from *models.GccTuParserNode){
	to.Base.MappingGccTuParserNode(from)
}
func MappingAttrsType(to NodeInterface, from string) { to.Load(NodeIdFromString(from)) }
func (to *NodeTypeIntegerCst) MappingGccTuParserNode(from *models.GccTuParserNode){
	MappingAttrsType(to.AttrsType,from.AttrsType)
	to.Base.MappingGccTuParserNode(from)
}
func (to *NodeTypeIdentifierNode) MappingGccTuParserNode(from *models.GccTuParserNode){
	to.Base.MappingGccTuParserNode(from)
}
func MappingRefsRetn(to NodeInterface, from sql.NullInt64) { to.Load(from) }
func MappingRefsPrms(to NodeInterface, from sql.NullInt64) { to.Load(from) }
func MappingRefsSize(to NodeInterface, from sql.NullInt64) { to.Load(from) }
func (to *NodeTypeFunctionType) MappingGccTuParserNode(from *models.GccTuParserNode){
	MappingRefsRetn(to.RefsRetn,from.RefsRetn)
	MappingRefsPrms(to.RefsPrms,from.RefsPrms)
	MappingRefsSize(to.RefsSize,from.RefsSize)
	to.Base.MappingGccTuParserNode(from)
}
func MappingRefsBpos(to NodeInterface, from sql.NullInt64) { to.Load(from) }
func (to *NodeTypeFieldDecl) MappingGccTuParserNode(from *models.GccTuParserNode){
	to.Base.MappingGccTuParserNode(from)
	MappingRefsBpos(to.RefsBpos,from.RefsBpos)
}
func MappingRefsPtd(to NodeInterface, from sql.NullInt64) { to.Load(from) }
func (to *NodeTypePointerType) MappingGccTuParserNode(from *models.GccTuParserNode){
	to.Base.MappingGccTuParserNode(from)
	MappingRefsPtd(to.RefsPtd,from.RefsPtd)
}
func (to *NodeTypeTreeList) MappingGccTuParserNode(from *models.GccTuParserNode){
	to.Base.MappingGccTuParserNode(from)
}
func (to *NodeTypeTypeDecl) MappingGccTuParserNode(from *models.GccTuParserNode){
	to.Base.MappingGccTuParserNode(from)
}
func MappingRefsName(to NodeInterface, from sql.NullInt64) { to.Load(from) }
func MappingRefsFlds(to NodeInterface, from sql.NullInt64) { to.Load(from) }
func MappingRefsUnql(to NodeInterface, from sql.NullInt64) { to.Load(from) }
func (to *NodeTypeUnionType) MappingGccTuParserNode(from *models.GccTuParserNode){
	MappingRefsUnql(to.RefsUnql,from.RefsUnql)
	MappingRefsName(to.RefsName,from.RefsName)
	to.Base.MappingGccTuParserNode(from)
	MappingRefsSize(to.RefsSize,from.RefsSize)
	MappingRefsFlds(to.RefsFlds,from.RefsFlds)
}
func (to *NodeTypeRecordType) MappingGccTuParserNode(from *models.GccTuParserNode){
	MappingRefsSize(to.RefsSize,from.RefsSize)
	to.Base.MappingGccTuParserNode(from)
	MappingRefsFlds(to.RefsFlds,from.RefsFlds)
}
func (to *NodeTypeVoidType) MappingGccTuParserNode(from *models.GccTuParserNode){
	to.Base.MappingGccTuParserNode(from)
}
func MappingRefsType(to NodeInterface, from sql.NullInt64) { to.Load(from) }
func (to *NodeTypeFunctionDecl) MappingGccTuParserNode(from *models.GccTuParserNode){
	MappingRefsName(to.RefsName,from.RefsName)
	to.Base.MappingGccTuParserNode(from)
	MappingRefsType(to.RefsType,from.RefsType)
}
func MappingRefsDomn(to NodeInterface, from sql.NullInt64) { to.Load(from) }
func MappingRefsElts(to NodeInterface, from sql.NullInt64) { to.Load(from) }
func (to *NodeTypeArrayType) MappingGccTuParserNode(from *models.GccTuParserNode){
	to.Base.MappingGccTuParserNode(from)
	MappingRefsDomn(to.RefsDomn,from.RefsDomn)
	MappingRefsSize(to.RefsSize,from.RefsSize)
	MappingRefsElts(to.RefsElts,from.RefsElts)
}
