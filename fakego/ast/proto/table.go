package astproto
import (
	"fmt"
	"github.com/h4ck3rm1k3/gogccintro/fakego/ast"
	//"github.com/h4ck3rm1k3/gogccintro/fakego/ast/proto"
	//"github.com/h4ck3rm1k3/gogccintro/fakego/token"
)

type AddressTable struct {
	Idsold          map[string] interface{}
	Idsnew          map[string] interface{}
	OldNew          map[interface{}] interface{}


	Scopes          map[string]*ast.Scope          `protobuf:"bytes,1,rep,name=Scopes,json=scopes" json:"Scopes,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Objects         map[string]*ast.Object         `protobuf:"bytes,2,rep,name=Objects,json=objects" json:"Objects,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	ArrayTypes      map[string]*ast.ArrayType      `protobuf:"bytes,3,rep,name=ArrayTypes,json=arrayTypes" json:"ArrayTypes,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	AssignStmts     map[string]*ast.AssignStmt     `protobuf:"bytes,4,rep,name=AssignStmts,json=assignStmts" json:"AssignStmts,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	BadDecls        map[string]*ast.BadDecl        `protobuf:"bytes,5,rep,name=BadDecls,json=badDecls" json:"BadDecls,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	BadExprs        map[string]*ast.BadExpr        `protobuf:"bytes,6,rep,name=BadExprs,json=badExprs" json:"BadExprs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	BadStmts        map[string]*ast.BadStmt        `protobuf:"bytes,7,rep,name=BadStmts,json=badStmts" json:"BadStmts,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	BasicLits       map[string]*ast.BasicLit       `protobuf:"bytes,8,rep,name=BasicLits,json=basicLits" json:"BasicLits,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	BinaryExprs     map[string]*ast.BinaryExpr     `protobuf:"bytes,9,rep,name=BinaryExprs,json=binaryExprs" json:"BinaryExprs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	BlockStmts      map[string]*ast.BlockStmt      `protobuf:"bytes,10,rep,name=BlockStmts,json=blockStmts" json:"BlockStmts,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	BranchStmts     map[string]*ast.BranchStmt     `protobuf:"bytes,11,rep,name=BranchStmts,json=branchStmts" json:"BranchStmts,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	CallExprs       map[string]*ast.CallExpr       `protobuf:"bytes,12,rep,name=CallExprs,json=callExprs" json:"CallExprs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	CaseClauses     map[string]*ast.CaseClause     `protobuf:"bytes,13,rep,name=CaseClauses,json=caseClauses" json:"CaseClauses,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	ChanDirs        map[string]*ast.ChanDir        `protobuf:"bytes,14,rep,name=ChanDirs,json=chanDirs" json:"ChanDirs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	ChanTypes       map[string]*ast.ChanType       `protobuf:"bytes,15,rep,name=ChanTypes,json=chanTypes" json:"ChanTypes,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	CommClauses     map[string]*ast.CommClause     `protobuf:"bytes,16,rep,name=CommClauses,json=commClauses" json:"CommClauses,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	CommentGroups   map[string]*ast.CommentGroup   `protobuf:"bytes,17,rep,name=CommentGroups,json=commentGroups" json:"CommentGroups,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Comments        map[string]*ast.Comment        `protobuf:"bytes,18,rep,name=Comments,json=comments" json:"Comments,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	CompositeLits   map[string]*ast.CompositeLit   `protobuf:"bytes,19,rep,name=CompositeLits,json=compositeLits" json:"CompositeLits,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	DeclStmts       map[string]*ast.DeclStmt       `protobuf:"bytes,20,rep,name=DeclStmts,json=declStmts" json:"DeclStmts,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	DeferStmts      map[string]*ast.DeferStmt      `protobuf:"bytes,21,rep,name=DeferStmts,json=deferStmts" json:"DeferStmts,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Ellipsiss       map[string]*ast.Ellipsis       `protobuf:"bytes,22,rep,name=Ellipsiss,json=ellipsiss" json:"Ellipsiss,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	EmptyStmts      map[string]*ast.EmptyStmt      `protobuf:"bytes,23,rep,name=EmptyStmts,json=emptyStmts" json:"EmptyStmts,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	ExprStmts       map[string]*ast.ExprStmt       `protobuf:"bytes,24,rep,name=ExprStmts,json=exprStmts" json:"ExprStmts,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	FieldLists      map[string]*ast.FieldList      `protobuf:"bytes,25,rep,name=FieldLists,json=fieldLists" json:"FieldLists,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Fields          map[string]*ast.Field          `protobuf:"bytes,26,rep,name=Fields,json=fields" json:"Fields,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Files           map[string]*ast.File           `protobuf:"bytes,27,rep,name=Files,json=files" json:"Files,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	ForStmts        map[string]*ast.ForStmt        `protobuf:"bytes,28,rep,name=ForStmts,json=forStmts" json:"ForStmts,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	FuncDecls       map[string]*ast.FuncDecl       `protobuf:"bytes,29,rep,name=FuncDecls,json=funcDecls" json:"FuncDecls,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	FuncLits        map[string]*ast.FuncLit        `protobuf:"bytes,30,rep,name=FuncLits,json=funcLits" json:"FuncLits,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	FuncTypes       map[string]*ast.FuncType       `protobuf:"bytes,31,rep,name=FuncTypes,json=funcTypes" json:"FuncTypes,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	GenDecls        map[string]*ast.GenDecl        `protobuf:"bytes,32,rep,name=GenDecls,json=genDecls" json:"GenDecls,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	GoStmts         map[string]*ast.GoStmt         `protobuf:"bytes,33,rep,name=GoStmts,json=goStmts" json:"GoStmts,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Idents          map[string]*ast.Ident          `protobuf:"bytes,34,rep,name=Idents,json=idents" json:"Idents,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	IfStmts         map[string]*ast.IfStmt         `protobuf:"bytes,35,rep,name=IfStmts,json=ifStmts" json:"IfStmts,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	ImportSpecs     map[string]*ast.ImportSpec     `protobuf:"bytes,36,rep,name=ImportSpecs,json=importSpecs" json:"ImportSpecs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	IncDecStmts     map[string]*ast.IncDecStmt     `protobuf:"bytes,37,rep,name=IncDecStmts,json=incDecStmts" json:"IncDecStmts,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	IndexExprs      map[string]*ast.IndexExpr      `protobuf:"bytes,38,rep,name=IndexExprs,json=indexExprs" json:"IndexExprs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	InterfaceTypes  map[string]*ast.InterfaceType  `protobuf:"bytes,39,rep,name=InterfaceTypes,json=interfaceTypes" json:"InterfaceTypes,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	IsExporteds     map[string]*ast.IsExported     `protobuf:"bytes,40,rep,name=IsExporteds,json=isExporteds" json:"IsExporteds,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	KeyValueExprs   map[string]*ast.KeyValueExpr   `protobuf:"bytes,41,rep,name=KeyValueExprs,json=keyValueExprs" json:"KeyValueExprs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	LabeledStmts    map[string]*ast.LabeledStmt    `protobuf:"bytes,42,rep,name=LabeledStmts,json=labeledStmts" json:"LabeledStmts,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	MapTypes        map[string]*ast.MapType        `protobuf:"bytes,43,rep,name=MapTypes,json=mapTypes" json:"MapTypes,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	NewIdents       map[string]*ast.NewIdent       `protobuf:"bytes,44,rep,name=NewIdents,json=newIdents" json:"NewIdents,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Nodes           map[string]*ast.Node           `protobuf:"bytes,45,rep,name=Nodes,json=nodes" json:"Nodes,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Packages        map[string]*ast.Package        `protobuf:"bytes,46,rep,name=Packages,json=packages" json:"Packages,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	ParenExprs      map[string]*ast.ParenExpr      `protobuf:"bytes,47,rep,name=ParenExprs,json=parenExprs" json:"ParenExprs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	RangeStmts      map[string]*ast.RangeStmt      `protobuf:"bytes,48,rep,name=RangeStmts,json=rangeStmts" json:"RangeStmts,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	ReturnStmts     map[string]*ast.ReturnStmt     `protobuf:"bytes,49,rep,name=ReturnStmts,json=returnStmts" json:"ReturnStmts,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	SelectStmts     map[string]*ast.SelectStmt     `protobuf:"bytes,50,rep,name=SelectStmts,json=selectStmts" json:"SelectStmts,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	SelectorExprs   map[string]*ast.SelectorExpr   `protobuf:"bytes,51,rep,name=SelectorExprs,json=selectorExprs" json:"SelectorExprs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	SendStmts       map[string]*ast.SendStmt       `protobuf:"bytes,52,rep,name=SendStmts,json=sendStmts" json:"SendStmts,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	SliceExprs      map[string]*ast.SliceExpr      `protobuf:"bytes,53,rep,name=SliceExprs,json=sliceExprs" json:"SliceExprs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	StarExprs       map[string]*ast.StarExpr       `protobuf:"bytes,54,rep,name=StarExprs,json=starExprs" json:"StarExprs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	StructTypes     map[string]*ast.StructType     `protobuf:"bytes,55,rep,name=StructTypes,json=structTypes" json:"StructTypes,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	SwitchStmts     map[string]*ast.SwitchStmt     `protobuf:"bytes,56,rep,name=SwitchStmts,json=switchStmts" json:"SwitchStmts,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	TypeAssertExprs map[string]*ast.TypeAssertExpr `protobuf:"bytes,57,rep,name=TypeAssertExprs,json=typeAssertExprs" json:"TypeAssertExprs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	TypeSpecs       map[string]*ast.TypeSpec       `protobuf:"bytes,58,rep,name=TypeSpecs,json=typeSpecs" json:"TypeSpecs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	TypeSwitchStmts map[string]*ast.TypeSwitchStmt `protobuf:"bytes,59,rep,name=TypeSwitchStmts,json=typeSwitchStmts" json:"TypeSwitchStmts,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	UnaryExprs      map[string]*ast.UnaryExpr      `protobuf:"bytes,60,rep,name=UnaryExprs,json=unaryExprs" json:"UnaryExprs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	ValueSpecs      map[string]*ast.ValueSpec      `protobuf:"bytes,61,rep,name=ValueSpecs,json=valueSpecs" json:"ValueSpecs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// interfaces
	Specs            map[string]*ast.Spec `protobuf:"bytes,62,rep,name=Specs,json=specs" json:"Specs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Decls            map[string]*ast.Decl `protobuf:"bytes,63,rep,name=Decls,json=decls" json:"Decls,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Exprs            map[string]*ast.Expr `protobuf:"bytes,64,rep,name=Exprs,json=exprs" json:"Exprs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Stmts            map[string]*ast.Stmt `protobuf:"bytes,65,rep,name=Stmts,json=stmts" json:"Stmts,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_unrecognized []byte           `json:"-"`
}

//func (t* AddressTable) PtrmapArrayType(id string) (* ArrayType){ return t.ArrayTypes[id] }
func (t* AddressTable) PtrmapAssignStmt(id string) (* ast.AssignStmt){ return t.AssignStmts[id] }
//func (t* AddressTable) PtrmapBadDecl(id string) (* ast.BadDecl){ return t.BadDecls[id] }
func (t* AddressTable) PtrmapBadExpr(id string) (* ast.BadExpr){ return t.BadExprs[id] }
func (t* AddressTable) PtrmapBadStmt(id string) (* ast.BadStmt){ return t.BadStmts[id] }
func (t* AddressTable) PtrmapBasicLit(id string) (* ast.BasicLit){ return t.BasicLits[id] }
func (t* AddressTable) PtrmapBinaryExpr(id string) (* ast.BinaryExpr){ return t.BinaryExprs[id] }
func (t* AddressTable) PtrmapBlockStmt(id string) (* ast.BlockStmt){ return t.BlockStmts[id] }
func (t* AddressTable) PtrmapBranchStmt(id string) (* ast.BranchStmt){ return t.BranchStmts[id] }
func (t* AddressTable) PtrmapCallExpr(id string) (* ast.CallExpr){ return t.CallExprs[id] }
func (t* AddressTable) PtrmapCaseClause(id string) (* ast.CaseClause){ return t.CaseClauses[id] }
//func (t* AddressTable) PtrmapChanDir(id string) (* ast.ChanDir){ return t.ChanDirs[id] }
//func (t* AddressTable) PtrmapChanType(id string) (* ast.ChanType){ return t.ChanTypes[id] }
func (t* AddressTable) PtrmapCommClause(id string) (* ast.CommClause){ return t.CommClauses[id] }
//func (t* AddressTable) PtrmapComment(id string) (* ast.Comment){ return t.Comments[id] }
func (t* AddressTable) PtrmapCommentGroup(id string) (* ast.CommentGroup){ return t.CommentGroups[id] }
func (t* AddressTable) PtrmapCompositeLit(id string) (* ast.CompositeLit){ return t.CompositeLits[id] }
//func (t* AddressTable) PtrmapDeclStmt(id string) (* ast.DeclStmt){ return t.DeclStmts[id] }
func (t* AddressTable) PtrmapDeferStmt(id string) (* ast.DeferStmt){ return t.DeferStmts[id] }
func (t* AddressTable) PtrmapEllipsis(id string) (* ast.Ellipsis){ return t.Ellipsiss[id] }
func (t* AddressTable) PtrmapEmptyStmt(id string) (* ast.EmptyStmt){ return t.EmptyStmts[id] }
//func (t* AddressTable) PtrmapExprStmt(id string) (* ast.ExprStmt){ return t.ExprStmts[id] }
func (t* AddressTable) PtrmapFile(id string) (* ast.File){ return t.Files[id] }
func (t* AddressTable) PtrmapForStmt(id string) (* ast.ForStmt){ return t.ForStmts[id] }
func (t* AddressTable) PtrmapFuncDecl(id string) (* ast.FuncDecl){ return t.FuncDecls[id] }
func (t* AddressTable) PtrmapFuncLit(id string) (* ast.FuncLit){ return t.FuncLits[id] }
func (t* AddressTable) PtrmapFuncType(id string) (* ast.FuncType){ return t.FuncTypes[id] }
func (t* AddressTable) PtrmapGenDecl(id string) (* ast.GenDecl){ return t.GenDecls[id] }
func (t* AddressTable) PtrmapGoStmt(id string) (* ast.GoStmt){ return t.GoStmts[id] }
func (t* AddressTable) PtrmapIdent(id string) (* ast.Ident){ return t.Idents[id] }
func (t* AddressTable) PtrmapIfStmt(id string) (* ast.IfStmt){ return t.IfStmts[id] }
func (t* AddressTable) PtrmapImportSpec(id string) (* ast.ImportSpec){ return t.ImportSpecs[id] }
func (t* AddressTable) PtrmapIncDecStmt(id string) (* ast.IncDecStmt){ return t.IncDecStmts[id] }
func (t* AddressTable) PtrmapIndexExpr(id string) (* ast.IndexExpr){ return t.IndexExprs[id] }
func (t* AddressTable) PtrmapInterfaceType(id string) (* ast.InterfaceType){ return t.InterfaceTypes[id] }
//func (t* AddressTable) PtrmapIsExported(id string) (* ast.IsExported){ return t.IsExporteds[id] }
func (t* AddressTable) PtrmapKeyValueExpr(id string) (* ast.KeyValueExpr){ return t.KeyValueExprs[id] }
func (t* AddressTable) PtrmapLabeledStmt(id string) (* ast.LabeledStmt){ return t.LabeledStmts[id] }
func (t* AddressTable) PtrmapMapType(id string) (* ast.MapType){ return t.MapTypes[id] }
//func (t* AddressTable) PtrmapNewIdent(id string) (* ast.NewIdent){ return t.NewIdents[id] }
//func (t* AddressTable) PtrmapNode(id string) (* ast.Node){ return t.Nodes[id] }
func (t* AddressTable) PtrmapPackage(id string) (* ast.Package){ return t.Packages[id] }
func (t* AddressTable) PtrmapParenExpr(id string) (* ast.ParenExpr){ return t.ParenExprs[id] }
func (t* AddressTable) PtrmapRangeStmt(id string) (* ast.RangeStmt){ return t.RangeStmts[id] }
//func (t* AddressTable) PtrmapReturnStmt(id string) (* ast.ReturnStmt){ return t.ReturnStmts[id] }
func (t* AddressTable) PtrmapSelectStmt(id string) (* ast.SelectStmt){ return t.SelectStmts[id] }
func (t* AddressTable) PtrmapSelectorExpr(id string) (* ast.SelectorExpr){ return t.SelectorExprs[id] }
func (t* AddressTable) PtrmapSendStmt(id string) (* ast.SendStmt){ return t.SendStmts[id] }
func (t* AddressTable) PtrmapSliceExpr(id string) (* ast.SliceExpr){ return t.SliceExprs[id] }
//func (t* AddressTable) PtrmapStarExpr(id string) (* ast.StarExpr){ return t.StarExprs[id] }
func (t* AddressTable) PtrmapStructType(id string) (* ast.StructType){ return t.StructTypes[id] }
func (t* AddressTable) PtrmapSwitchStmt(id string) (* ast.SwitchStmt){ return t.SwitchStmts[id] }
func (t* AddressTable) PtrmapTypeAssertExpr(id string) (* ast.TypeAssertExpr){ return t.TypeAssertExprs[id] }


func (t* AddressTable) PtrmapTypeSwitchStmt(id string) (* ast.TypeSwitchStmt){ return t.TypeSwitchStmts[id] }
func (t* AddressTable) PtrmapUnaryExpr(id string) (* ast.UnaryExpr){ return t.UnaryExprs[id] }
func (t* AddressTable) PtrmapValueSpec(id string) (* ast.ValueSpec){ return t.ValueSpecs[id] }


func CreateTable() (*AddressTable) {
	return &AddressTable{
		Idsold:          make(map[string] interface{}),
		Idsnew :         make(map[string] interface{}),
		OldNew  :        make(map[interface{}] interface{}),

		Scopes:make( map[string] *ast.Scope),
		Objects:make( map[string] *ast.Object),
		ArrayTypes : make(map[string]*ast.ArrayType),
		AssignStmts : make(map[string]*ast.AssignStmt),
		BadDecls : make(map[string]*ast.BadDecl),
		BadExprs : make(map[string]*ast.BadExpr),
		BadStmts : make(map[string]*ast.BadStmt),
		BasicLits : make(map[string]*ast.BasicLit),
BinaryExprs : make(map[string]*ast.BinaryExpr),
BlockStmts : make(map[string]*ast.BlockStmt),
BranchStmts : make(map[string]*ast.BranchStmt),
CallExprs : make(map[string]*ast.CallExpr),
     CaseClauses : make(map[string]*ast.CaseClause),
     ChanDirs : make(map[string]*ast.ChanDir),
     ChanTypes : make(map[string]*ast.ChanType),
     CommClauses : make(map[string]*ast.CommClause),
     CommentGroups : make(map[string]*ast.CommentGroup),
     Comments : make(map[string]*ast.Comment),
     CompositeLits : make(map[string]*ast.CompositeLit),
     DeclStmts : make(map[string]*ast.DeclStmt),

     DeferStmts : make(map[string]*ast.DeferStmt),
     Ellipsiss : make(map[string]*ast.Ellipsis),
     EmptyStmts : make(map[string]*ast.EmptyStmt),
     ExprStmts : make(map[string]*ast.ExprStmt),

     FieldLists : make(map[string]*ast.FieldList),
     Fields : make(map[string]*ast.Field),
     Files : make(map[string]*ast.File),
     ForStmts : make(map[string]*ast.ForStmt),
     FuncDecls : make(map[string]*ast.FuncDecl),
     FuncLits : make(map[string]*ast.FuncLit),
     FuncTypes : make(map[string]*ast.FuncType),
     GenDecls : make(map[string]*ast.GenDecl),
     GoStmts : make(map[string]*ast.GoStmt),
     Idents : make(map[string]*ast.Ident),
     IfStmts : make(map[string]*ast.IfStmt),
     ImportSpecs : make(map[string]*ast.ImportSpec),
     IncDecStmts : make(map[string]*ast.IncDecStmt),
     IndexExprs : make(map[string]*ast.IndexExpr),
     InterfaceTypes : make(map[string]*ast.InterfaceType),
     IsExporteds : make(map[string]*ast.IsExported),
     KeyValueExprs : make(map[string]*ast.KeyValueExpr),
     LabeledStmts : make(map[string]*ast.LabeledStmt),
     MapTypes : make(map[string]*ast.MapType),
     NewIdents : make(map[string]*ast.NewIdent),
     Nodes : make(map[string]*ast.Node),
     Packages : make(map[string]*ast.Package),
     ParenExprs : make(map[string]*ast.ParenExpr),
     RangeStmts : make(map[string]*ast.RangeStmt),
     ReturnStmts : make(map[string]*ast.ReturnStmt),
    SelectStmts : make(map[string]*ast.SelectStmt),
    SelectorExprs : make(map[string]*ast.SelectorExpr),
    SendStmts : make(map[string]*ast.SendStmt),
    SliceExprs : make(map[string]*ast.SliceExpr),

    StarExprs : make(map[string]*ast.StarExpr),

    StructTypes : make(map[string]*ast.StructType),
    SwitchStmts : make(map[string]*ast.SwitchStmt),
    TypeAssertExprs : make(map[string]*ast.TypeAssertExpr),
    TypeSpecs : make(map[string]*ast.TypeSpec),
    TypeSwitchStmts : make(map[string]*ast.TypeSwitchStmt),
    UnaryExprs : make(map[string]*ast.UnaryExpr),
    ValueSpecs : make(map[string]*ast.ValueSpec),

		// interfaces
		Stmts : make(map[string]*ast.Stmt),
		Exprs : make(map[string]*ast.Expr),
		Specs : make(map[string]*ast.Spec),
		Decls : make(map[string]*ast.Decl),
	}
}

// func (t* AddressTable) StrmapTypeSpec(id string, f * TypeSpec) (*TypeSpec){
//	t.TypeSpecs[id] =f;
//	f.Report();
//	return f
// }


func (t* Deferred) Report() (string){
	// needs access to table to resolve set
	return "TODO"
	//if t.Data != nil {
	//	return t.Data.Report()
	//} else {
	// i := t.Set
	// switch v:= i.(type) {
	//	case *map[string]*Field :
	//		if val, ok := (*v)[t.Id]; ok {
	//			if val == nil {
	//				fmt.Printf("did not find %s %s",v,t)
	//			} else {
	//				d := (*v)[t.Id]
	//				return d.Report()
	//			}
	//		}
	//	case *map[string]*FieldList :
	//		if val, ok := (*v)[t.Id]; ok {
	//			if val == nil {
	//				fmt.Printf("did not find %s %s",v,t)
	//			} else {
	//				d := (*v)[t.Id]
	//				return d.Report()
	//			}
	//		}
	//	case *map[string]*TypeSpec :

	//		if val, ok := (*v)[t.Id]; ok {
	//			if val == nil {
	//				fmt.Printf("did not find %s %s",v,t)
	//			} else {
	//				d := (*v)[t.Id]
	//				return d.Report()
	//			}
	//		}


	//	default:
	//		//t.Data = v[t.Id]
	//		fmt.Printf("unknown type %s %s",v,t)
	//		panic("unknown type")
	//		return "unkown"
	//	}
	// //}
	// return "huh?"
}


func (t* AddressTable) PtrmapTypeSpec(id string) (ast.Foo2){

	if x, ok := t.Idsold[id]; ok {
		switch v:=x.(type) {
		default:
			fmt.Printf("default %t\n",v)
		}

	}
	if val, ok := t.TypeSpecs[id]; ok {
		//t := NodeType_TYPESPEC
		//return &Foo2{
		//	Type: &t,
		//	Typespec : val,
		//}
		return val
	}

	// create deferred
	// t := NodeType_TYPESPEC
	// t2 := NodeType_DEFERRED
	// return &Foo2{
	//	Type: &t2,
	//	Deferred:&Deferred{
	//		Id:&id,
	//		Type: &t,
	//	},
	// }
	return &ast.Deferred{
		Id: id,
		Set: &t.TypeSpecs,
	}
	//panic("todo"); return nil;
}

func (t* AddressTable) PtrmapTypeSpec2(id string) (ast.Foo2){
	panic("todo"); return nil;
	// if val, ok := t.TypeSpecs[id]; ok {
	//	if val == nil {
	//	} else {

	//	}
	// } else {
	//	//return &Deferred{ Id:id, Set: &t.TypeSpecs  }
	//	t := NodeType_TYPESPEC
	//	t2 := NodeType_DEFERRED
	//	return &Foo2{
	//		Type: &t2,
	//		Deferred:&Deferred{
	//			Id:&id,
	//			Type: &t,
	//		},
	//	}
	// }
}

func (t* AddressTable) PtrmapField(id string) (ast.Foo2){

	// t3 := NodeType_FIELD
	// //return t.Fields[id]
	if val, ok := t.Fields[id]; ok {
		return val
	}
	//panic("todo"); return nil;
	return &ast.FieldDeferred {
		Id : id,
	}
	//	if val == nil {
	//		//return &Deferred{ Id:id, Set: &t.Fields  }

	//		t2 := NodeType_DEFERRED
	//		return &Foo2{
	//			Type: &t2,
	//			Deferred:&Deferred{
	//			Id:&id,
	//			Type: &t3,
	//			},
	//		}

	//	} else {
	//		//return val
	//		return &Foo2{
	//			Type: &t3,
	//			Field : val,
	//		}
	//	}

	// } else {
	//	//return &Deferred{ Id:id, Set: &t.Fields  }
	//	t := NodeType_FIELD
	//	t2 := NodeType_DEFERRED
	//	return &Foo2{
	//		Type: &t2,
	//		Deferred:&Deferred{
	//			Id:&id,
	//			Type: &t,
	//		},
	//	}

	// }
}
func (t* AddressTable) PtrmapFieldList(id string) (ast.Foo2){
	panic("todo"); return nil;
	// return t.FieldLists[id]
	// t3 := NodeType_FIELDLIST
	// if val, ok := t.FieldLists[id]; ok {
	//	if val != nil {
	//		return Foo2{
	//			Type: &t3,
	//			Fieldlist : val,
	//		}
	//	}
	// }
	// t2 := NodeType_DEFERRED
	// return Foo2{
	//	Type: &t2,
	//	Deferred:&Deferred{
	//		Id:&id,
	//		Type: &t3,
	//	},
	// }

}

func (t* AddressTable) PtrmapObject(id string) (* ast.Object){
	// todo
	if val, ok := t.Objects[id]; ok {
		return val
	}
	return &ast.Object {
		Deferred:&ast.Deferred{
			Id:id,
			Set: &t.Objects,
		},
	}
	//panic("todo"); return nil;
}

// func (t* AddressTable) PtrmapObject(id string) (*ast.Object){
//	t2 := NodeType_OBJECT
//	if val, ok := t.Objects[id]; ok {
//		if val != nil {
//			return val
//		}
//	}
//	return &Object{
//		Deferred:&Deferred{
//			Id:&id,
//			Type: &t2,
//		},
//	}
// }

// func (t* AddressTable) StrmapObject(id string, f * Object) (*Object){ t.Objects[id] =f; f.Report(); return f}
// func (t* AddressTable) StrmapValueSpec(id string, f * ValueSpec) (Foo2){
//	//t.ValueSpecs[id] =f; f.Report(); return f
//	s := t.ValueSpecs
//	t2 := NodeType_VALUESPEC
//	if val, ok := s[id]; ok {
//		if val != nil {
//			return Foo2{
//				Type: &t2,
//				Valuespec : val,
//			}
//		}
//	}


//	t3 := NodeType_DEFERRED
//	return Foo2{
//		Type: &t2,
//		Deferred:&Deferred{
//			Id:&id,
//			Type: &t3,
//		},
//	}

// }

// func (t* AddressTable) StrmapFuncDecl(id string, f * FuncDecl) (Foo2){
//	fmt.Printf("test")
//	t2 := NodeType_FUNCDECL
//	//t.FuncDecls[id] =f; f.Report(); return f
//	s := t.FuncDecls
//	if val, ok := s[id]; ok {
//		if val != nil {
//			return Foo2{
//				Type: &t2,
//				Funcdecl : val,
//			}
//		}
//	}
//	t3 := NodeType_DEFERRED
//	return Foo2{
//		Type: &t2,
//		Deferred:&Deferred{
//			Id:&id,
//			Type: &t3,
//		},
//	}

// }


// func (t* AddressTable) StrmapDecl(id string, f Decl) (Decl){ t.Decls[id] = &f; f.Report(); return f}
// func (t* AddressTable) StrmapExpr(id string, f Expr) (Expr){ t.Exprs[id] = &f; f.Report(); return f}
// func (t* AddressTable) StrmapSpec(id string, f Spec) (Spec){ t.Specs[id] = &f; f.Report(); return f}
// func (t* AddressTable) StrmapStmt(id string, f Stmt) (Stmt){ t.Stmts[id] = &f; f.Report(); return f}

func (t* AddressTable) PtrmapDecl(id string) (ast.Decl){
	//panic("todo"); return nil;
	return *t.Decls[id]
}
func (t* AddressTable) PtrmapExpr(id string) (ast.Expr){ return *t.Exprs[id] }
func (t* AddressTable) PtrmapStmt(id string) (ast.Stmt){ return *t.Stmts[id] }
func (t* AddressTable) PtrmapSpec(id string) (ast.Spec){ return *t.Specs[id] }



///func (t* AddressTable) StrmapStructType(id string, f * StructType) (*StructType){ t.StructTypes[id] =f; f.Report(); return f}
func (t* AddressTable) ConvertFoo3(f ast.Foo3) (*Foo3){
	// fmt.Printf("Foo3 %#v\n", f)
	if f == nil {
		return nil;
	}
	switch v:= f.(type) {
	case string:
		fmt.Printf("Foo3 string: %s\n", v)
		panic("wrong type")
	case int:
		var  v2 int64 = (int64(v))
		return &Foo3{
			Tokenpos : &v2,
		}
	default:
		fmt.Printf("other ConvertFoo3: %T\n", v)
	}

	panic("todo"); return nil;
}

// func (t* AddressTable) ConvertFieldList(f *ast.FieldList) (*FieldList){
//	panic("todo"); return nil;
// }

// func (t* AddressTable) ConvertStructType(f *ast.StructType) (*StructType){
//
//	return &StructType{
//		Struct : t.ConvertFoo3(f.Struct),
//		Fields : t.ConvertFieldList(f.Fields),
//		Incomplete : &f2,
//	}
// }

// func (t* AddressTable) StrmapStructType(id string, f * ast.StructType) (*ast.StructType){ t.StructTypes[id] = t.ConvertStructType(f); f.Report(); return f}


func (t* AddressTable) Converttoken_Pos(f ast.Foo3) (*Foo3){
	switch v:=f.(type) {
	case int64:
		return &Foo3{
			Tokenpos: &v,
		}
	case int:
		var f2 int64
		f2= int64(v)
		return &Foo3{
			Tokenpos: &f2,
		}
	default:
		fmt.Printf("pos %T\n",v)
	}
	panic("todo")

}

func (t* AddressTable) Converttoken_Pos2(f ast.Foo3) (*Foo2){
	panic("todo"); return nil;
}

func (t* AddressTable) Converttoken_Token(f ast.Foo3) (*Foo3){
	if f == nil { return nil }
	switch v:=f.(type) {
	default:
		fmt.Printf("type debug %T",v)

	}
	panic("todo"); return nil;
}

func (t* AddressTable) ConvertTypeExpr3(f ast.Foo3) (*Expr){
	panic("todo")
	panic("todo"); return nil;
}

func (t* AddressTable) ConvertTypeExpr(f ast.Expr) (*Expr){

	switch v:= f.(type) {
	case nil:
		fmt.Printf("found nil type expr %T",v)
		return nil;
	case *ast.IndexExpr:
		t2 := NodeType_INDEXEXPR
		return &Expr{
			Type: &t2,
			Indexexpr: t.ConvertIndexExpr(v),
		}

	case *ast.BinaryExpr:
		t2 := NodeType_BINARYEXPR
		return &Expr{
			Type: &t2,
			Binaryexpr: t.ConvertBinaryExpr(v),
		}


	case *ast.BasicLit:
		t2 := NodeType_BASICLIT
		return &Expr{
			Type: &t2,
			Basiclit: t.ConvertBasicLit(v),
		}

	case *ast.ArrayType:
		t2 := NodeType_ARRAYTYPE
		return &Expr{
			Type: &t2,
			Arraytype: t.ConvertArrayType(v),
		}

	case *ast.CallExpr:
		t2 := NodeType_CALLEXPR
		return &Expr{
			Type: &t2,
			Callexpr: t.ConvertCallExpr(v),
		}

	case *ast.StarExpr:
		t2 := NodeType_STAREXPR
		return &Expr{
			Type: &t2,
			Starexpr: t.ConvertStarExpr(v),
		}

	case *ast.StructType:
		t2 := NodeType_STRUCTTYPE
		return &Expr{
			Type: &t2,
			Structtype: t.ConvertStructType(v),
		}
	case *ast.InterfaceType:
		t2 := NodeType_INTERFACETYPE
		return &Expr{
			Type: &t2,
			Interfacetype: t.ConvertInterfaceType(v),
		}
	case *ast.FuncType:
		t2 := NodeType_FUNCTYPE
		return &Expr{
			Type: &t2,
			Functype: t.ConvertFuncType(v),
		}
	case *ast.SelectorExpr:
		t2 := NodeType_SELECTOREXPR
		return &Expr{
			Type: &t2,
			Selectorexpr: t.ConvertSelectorExpr(v),
		}
	case *ast.Ident:
		t2 := NodeType_IDENT
		return &Expr{
			Type: &t2,
			Ident: t.ConvertIdent(v),
		}

	case	*ast.CompositeLit:
		t2 := NodeType_COMPOSITELIT
		return &Expr{
			Type: &t2,
			Compositelit: t.ConvertCompositeLit(v),
		}


	/*
case *ast.AssignStmt:
	t2 := NodeType_ASSIGNSTMT
	return &Expr{
	Type: &t2,
	Assignstmt: t.ConvertAssignStmt(v),
	}
case *ast.BadDecl:
	t2 := NodeType_BADDECL
	return &Expr{
	Type: &t2,
	Baddecl: t.ConvertBadDecl(v),
	}
case *ast.BadExpr:
	t2 := NodeType_BADEXPR
	return &Expr{
	Type: &t2,
	Badexpr: t.ConvertBadExpr(v),
	}
case *ast.BadStmt:
	t2 := NodeType_BADSTMT
	return &Expr{
	Type: &t2,
	Badstmt: t.ConvertBadStmt(v),
	}
case *ast.BasicLit:
	t2 := NodeType_BASICLIT
	return &Expr{
	Type: &t2,
	Basiclit: t.ConvertBasicLit(v),
	}
case *ast.BinaryExpr:
	t2 := NodeType_BINARYEXPR
	return &Expr{
	Type: &t2,
	Binaryexpr: t.ConvertBinaryExpr(v),
	}
case *ast.BlockStmt:
	t2 := NodeType_BLOCKSTMT
	return &Expr{
	Type: &t2,
	Blockstmt: t.ConvertBlockStmt(v),
	}
case *ast.BranchStmt:
	t2 := NodeType_BRANCHSTMT
	return &Expr{
	Type: &t2,
	Branchstmt: t.ConvertBranchStmt(v),
	}
case *ast.CallExpr:
	t2 := NodeType_CALLEXPR
	return &Expr{
	Type: &t2,
	Callexpr: t.ConvertCallExpr(v),
	}
case *ast.CaseClause:
	t2 := NodeType_CASECLAUSE
	return &Expr{
	Type: &t2,
	Caseclause: t.ConvertCaseClause(v),
	}
case *ast.ChanDir:
	t2 := NodeType_CHANDIR
	return &Expr{
	Type: &t2,
	Chandir: t.ConvertChanDir(v),
	}
case *ast.ChanType:
	t2 := NodeType_CHANTYPE
	return &Expr{
	Type: &t2,
	Chantype: t.ConvertChanType(v),
	}
case *ast.CommClause:
	t2 := NodeType_COMMCLAUSE
	return &Expr{
	Type: &t2,
	Commclause: t.ConvertCommClause(v),
	}
case *ast.Comment:
	t2 := NodeType_COMMENT
	return &Expr{
	Type: &t2,
	Comment: t.ConvertComment(v),
	}
case *ast.CompositeLit:
	t2 := NodeType_COMPOSITELIT
	return &Expr{
	Type: &t2,
	Compositelit: t.ConvertCompositeLit(v),
	}
case *ast.DeclStmt:
	t2 := NodeType_DECLSTMT
	return &Expr{
	Type: &t2,
	Declstmt: t.ConvertDeclStmt(v),
	}
case *ast.DeferStmt:
	t2 := NodeType_DEFERSTMT
	return &Expr{
	Type: &t2,
	Deferstmt: t.ConvertDeferStmt(v),
	}
case *ast.Ellipsis:
	t2 := NodeType_ELLIPSIS
	return &Expr{
	Type: &t2,
	Ellipsis: t.ConvertEllipsis(v),
	}
case *ast.EmptyStmt:
	t2 := NodeType_EMPTYSTMT
	return &Expr{
	Type: &t2,
	Emptystmt: t.ConvertEmptyStmt(v),
	}
case *ast.ExprStmt:
	t2 := NodeType_EXPRSTMT
	return &Expr{
	Type: &t2,
	Exprstmt: t.ConvertExprStmt(v),
	}
case *ast.Field:
	t2 := NodeType_FIELD
	return &Expr{
	Type: &t2,
	Field: t.ConvertField(v),
	}
case *ast.FieldList:
	t2 := NodeType_FIELDLIST
	return &Expr{
	Type: &t2,
	Fieldlist: t.ConvertFieldList(v),
	}
case *ast.File:
	t2 := NodeType_FILE
	return &Expr{
	Type: &t2,
	File: t.ConvertFile(v),
	}
case *ast.ForStmt:
	t2 := NodeType_FORSTMT
	return &Expr{
	Type: &t2,
	Forstmt: t.ConvertForStmt(v),
	}
case *ast.FuncDecl:
	t2 := NodeType_FUNCDECL
	return &Expr{
	Type: &t2,
	Funcdecl: t.ConvertFuncDecl(v),
	}
case *ast.FuncLit:
	t2 := NodeType_FUNCLIT
	return &Expr{
	Type: &t2,
	Funclit: t.ConvertFuncLit(v),
	}
case *ast.FuncType:
	t2 := NodeType_FUNCTYPE
	return &Expr{
	Type: &t2,
	Functype: t.ConvertFuncType(v),
	}
case *ast.GenDecl:
	t2 := NodeType_GENDECL
	return &Expr{
	Type: &t2,
	Gendecl: t.ConvertGenDecl(v),
	}
case *ast.GoStmt:
	t2 := NodeType_GOSTMT
	return &Expr{
	Type: &t2,
	Gostmt: t.ConvertGoStmt(v),
	}
case *ast.Ident:
	t2 := NodeType_IDENT
	return &Expr{
	Type: &t2,
	Ident: t.ConvertIdent(v),
	}
case *ast.IfStmt:
	t2 := NodeType_IFSTMT
	return &Expr{
	Type: &t2,
	Ifstmt: t.ConvertIfStmt(v),
	}
case *ast.ImportSpec:
	t2 := NodeType_IMPORTSPEC
	return &Expr{
	Type: &t2,
	Importspec: t.ConvertImportSpec(v),
	}
case *ast.IncDecStmt:
	t2 := NodeType_INCDECSTMT
	return &Expr{
	Type: &t2,
	Incdecstmt: t.ConvertIncDecStmt(v),
	}
case *ast.IndexExpr:
	t2 := NodeType_INDEXEXPR
	return &Expr{
	Type: &t2,
	Indexexpr: t.ConvertIndexExpr(v),
	}
case *ast.InterfaceType:
	t2 := NodeType_INTERFACETYPE
	return &Expr{
	Type: &t2,
	Interfacetype: t.ConvertInterfaceType(v),
	}
case *ast.IsExported:
	t2 := NodeType_ISEXPORTED
	return &Expr{
	Type: &t2,
	Isexported: t.ConvertIsExported(v),
	}
case *ast.KeyValueExpr:
	t2 := NodeType_KEYVALUEEXPR
	return &Expr{
	Type: &t2,
	Keyvalueexpr: t.ConvertKeyValueExpr(v),
	}
case *ast.LabeledStmt:
	t2 := NodeType_LABELEDSTMT
	return &Expr{
	Type: &t2,
	Labeledstmt: t.ConvertLabeledStmt(v),
	}
*/
case *ast.MapType:
	t2 := NodeType_MAPTYPE
	return &Expr{
	Type: &t2,
	Maptype: t.ConvertMapType(v),
	}
/*
case *ast.NewIdent:
	t2 := NodeType_NEWIDENT
	return &Expr{
	Type: &t2,
	Newident: t.ConvertNewIdent(v),
	}
case *ast.Package:
	t2 := NodeType_PACKAGE
	return &Expr{
	Type: &t2,
	Package: t.ConvertPackage(v),
	}
case *ast.ParenExpr:
	t2 := NodeType_PARENEXPR
	return &Expr{
	Type: &t2,
	Parenexpr: t.ConvertParenExpr(v),
	}
case *ast.RangeStmt:
	t2 := NodeType_RANGESTMT
	return &Expr{
	Type: &t2,
	Rangestmt: t.ConvertRangeStmt(v),
	}
case *ast.ReturnStmt:
	t2 := NodeType_RETURNSTMT
	return &Expr{
	Type: &t2,
	Returnstmt: t.ConvertReturnStmt(v),
	}
case *ast.SelectStmt:
	t2 := NodeType_SELECTSTMT
	return &Expr{
	Type: &t2,
	Selectstmt: t.ConvertSelectStmt(v),
	}
case *ast.SelectorExpr:
	t2 := NodeType_SELECTOREXPR
	return &Expr{
	Type: &t2,
	Selectorexpr: t.ConvertSelectorExpr(v),
	}
case *ast.SendStmt:
	t2 := NodeType_SENDSTMT
	return &Expr{
	Type: &t2,
	Sendstmt: t.ConvertSendStmt(v),
	}
case *ast.SliceExpr:
	t2 := NodeType_SLICEEXPR
	return &Expr{
	Type: &t2,
	Sliceexpr: t.ConvertSliceExpr(v),
	}
case *ast.Spec:
	t2 := NodeType_SPEC
	return &Expr{
	Type: &t2,
	Spec: t.ConvertSpec(v),
	}
case *ast.StarExpr:
	t2 := NodeType_STAREXPR
	return &Expr{
	Type: &t2,
	Starexpr: t.ConvertStarExpr(v),
	}
case *ast.StructType:
	t2 := NodeType_STRUCTTYPE
	return &Expr{
	Type: &t2,
	Structtype: t.ConvertStructType(v),
	}
case *ast.SwitchStmt:
	t2 := NodeType_SWITCHSTMT
	return &Expr{
	Type: &t2,
	Switchstmt: t.ConvertSwitchStmt(v),
	}
case *ast.TypeAssertExpr:
	t2 := NodeType_TYPEASSERTEXPR
	return &Expr{
	Type: &t2,
	Typeassertexpr: t.ConvertTypeAssertExpr(v),
	}
case *ast.TypeSpec:
	t2 := NodeType_TYPESPEC
	return &Expr{
	Type: &t2,
	Typespec: t.ConvertTypeSpec(v),
	}
case *ast.TypeSwitchStmt:
	t2 := NodeType_TYPESWITCHSTMT
	return &Expr{
	Type: &t2,
	Typeswitchstmt: t.ConvertTypeSwitchStmt(v),
	}
case *ast.UnaryExpr:
	t2 := NodeType_UNARYEXPR
	return &Expr{
	Type: &t2,
	Unaryexpr: t.ConvertUnaryExpr(v),
	}
case *ast.ValueSpec:
	t2 := NodeType_VALUESPEC
	return &Expr{
	Type: &t2,
	Valuespec: t.ConvertValueSpec(v),
	}
case *ast.CommentGroup:
	t2 := NodeType_COMMENTGROUP
	return &Expr{
	Type: &t2,
	Commentgroup: t.ConvertCommentGroup(v),
	}
case *ast.Spec:
	t2 := NodeType_SPEC
	return &Expr{
	Type: &t2,
	Spec: t.ConvertSpec(v),
	}
case *ast.Object:
	t2 := NodeType_OBJECT
	return &Expr{
	Type: &t2,
	Object: t.ConvertObject(v),
	}
*/
	default:
		fmt.Printf("type debug %T",v)

	}
	panic("todo")
	//panic("todo"); return nil;
}

func (t* AddressTable) ConvertPointerFieldList(f * ast.FieldList) (*FieldList){
	return t.ConvertFieldList(f)
}


//func (t* AddressTable) ConvertTypebool(f ast.Boo) (*bool){
//	panic("todo")
//	panic("todo"); return nil;
//}

func (t* AddressTable) ConvertArrayExpr(f []ast.Expr) ([]*Expr){

	v2 := make([]*Expr, len(f))
	for i,j := range(f) {
		x:=t.ConvertExpr(j)
		//fmt.Printf("created %d %#v\n", i,*x)
		v2[i]=x
	}
	//fmt.Printf("created %#v\n", f)
	return v2

	//panic("todo"); return nil;
}

func (t* AddressTable) ConvertPointerBasicLit(f * ast.BasicLit) (*BasicLit){
	return t.ConvertBasicLit(f)
}

func (t* AddressTable) ConvertTypeStmt(f ast.Stmt) (*Stmt){
	//panic("todo"); return nil;
	switch v:= f.(type) {
	case *ast.AssignStmt:
		t2 := NodeType_ASSIGNSTMT
		return &Stmt{
			Type: &t2,
			Assignstmt: t.ConvertAssignStmt(v),
		}
	case nil:
		return nil
	default:
		fmt.Printf("stmt unknown %T",v)

	}
	panic("todo")

}

func (t* AddressTable) ConvertPointerBlockStmt(f * ast.BlockStmt) (*BlockStmt){
	panic("todo"); return nil;
}


func (t* AddressTable) ConvertMapType(f *ast.MapType) (*MapType){
//o
	return &MapType{
		Map : t.Converttoken_Pos( f.Map),
		Key : t.ConvertExpr( f.Key),
		Value : t.ConvertExpr( f.Value),
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertField(f *ast.Field) (*Field){

	return &Field{
		//Doc : t.ConvertPointerCommentGroup( f.Doc),
		Names : t.ConvertArrayStarIdent( f.Names),
		Type : t.ConvertTypeExpr( f.Type),
		Tag : t.ConvertBasicLit( f.Tag),
		//Comment : t.ConvertPointerCommentGroup( f.Comment),
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertSliceExpr(f *ast.SliceExpr) (*SliceExpr){

	return &SliceExpr{
		X : t.ConvertTypeExpr( f.X),
		Lbrack : t.Converttoken_Pos( f.Lbrack),
		Low : t.ConvertTypeExpr( f.Low),
		High : t.ConvertTypeExpr( f.High),
		Max : t.ConvertTypeExpr( f.Max),
		Slice3 : & f.Slice3,
		Rbrack : t.Converttoken_Pos( f.Rbrack),
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertExprStmt(f *ast.ExprStmt) (*ExprStmt){

	return &ExprStmt{
		X : t.ConvertTypeExpr( f.X),
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertAssignStmt(f *ast.AssignStmt) (*AssignStmt){
	if f == nil {return nil }
	return &AssignStmt{
		Lhs : t.ConvertArrayExpr( f.Lhs),
		TokPos : t.Converttoken_Pos( f.TokPos),
		Tok : t.Converttoken_Token( f.Tok),
		Rhs : t.ConvertArrayExpr( f.Rhs),
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertInterfaceType(f *ast.InterfaceType) (*InterfaceType){

	return &InterfaceType{
		Interface : t.Converttoken_Pos( f.Interface),
		Methods : t.ConvertFieldList( f.Methods),
		Incomplete : & f.Incomplete,
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertIfStmt(f *ast.IfStmt) (*IfStmt){

	return &IfStmt{
		If : t.Converttoken_Pos( f.If),
		Init : t.ConvertTypeStmt( f.Init),
		Cond : t.ConvertTypeExpr( f.Cond),
		Body : t.ConvertBlockStmt( f.Body),
		Else : t.ConvertTypeStmt( f.Else),
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertBadExpr(f *ast.BadExpr) (*BadExpr){

	return &BadExpr{
		From : t.Converttoken_Pos2( f.From),
		To : t.Converttoken_Pos2( f.To),
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertChanDir(f *ast.ChanDir) (*ChanDir){

	return &ChanDir{
		//identifier &{NamePos:10175 Name:ChanDir Obj:0xc420148720}
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertBinaryExpr(f *ast.BinaryExpr) (*BinaryExpr){

	return &BinaryExpr{
		X : t.ConvertTypeExpr( f.X),
		OpPos : t.Converttoken_Pos( f.OpPos),
		Op : t.Converttoken_Token( f.Op),
		Y : t.ConvertTypeExpr( f.Y),
		//Incomplete: &f2,
	}
}

// func (t* AddressTable) ConvertBadDecl(f *ast.BadDecl) (*BadDecl){
//
//	return &BadDecl{
//		From : t.Converttoken_Pos( f.From),
//		To : t.Converttoken_Pos( f.To),
//		Incomplete: &f2,
//	}
// }

func (t* AddressTable) ConvertIncDecStmt(f *ast.IncDecStmt) (*IncDecStmt){

	return &IncDecStmt{
		X : t.ConvertTypeExpr( f.X),
		TokPos : t.Converttoken_Pos( f.TokPos),
		Tok : t.Converttoken_Token( f.Tok),
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertPointerCallExpr(f *ast.CallExpr) (*CallExpr){
	panic("todo"); return nil;
}

func (t* AddressTable) ConvertDeferStmt(f *ast.DeferStmt) (*DeferStmt){

	return &DeferStmt{
		Defer : t.Converttoken_Pos( f.Defer),
		Call : t.ConvertCallExpr( f.Call),
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertKeyValueExpr(f *ast.KeyValueExpr) (*KeyValueExpr){

	return &KeyValueExpr{
		Key : t.ConvertTypeExpr( f.Key),
		Colon : t.Converttoken_Pos( f.Colon),
		Value : t.ConvertTypeExpr( f.Value),
		//Incomplete: &f2,
	}
}

// func (t* AddressTable) ConvertSEND(f *ast.SEND) (*SEND){
//
//	return &SEND{
// // other type &{}
//		Incomplete: &f2,
//	}
// }

// func (t* AddressTable) ConvertisWhitespace(f *ast.isWhitespace) (*isWhitespace){
//
//	return &isWhitespace{
// // other type &{}
//		Incomplete: &f2,
//	}
// }

func (t* AddressTable) ConvertSelectorExpr(f *ast.SelectorExpr) (*SelectorExpr){

	return &SelectorExpr{
		X : t.ConvertTypeExpr( f.X),
		Sel : t.ConvertIdent( f.Sel),
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertCompositeLit(f *ast.CompositeLit) (*CompositeLit){

	return &CompositeLit{
		Type : t.ConvertTypeExpr( f.Type),
		Lbrace : t.Converttoken_Pos( f.Lbrace),
		Elts : t.ConvertArrayExpr( f.Elts),
		Rbrace : t.Converttoken_Pos( f.Rbrace),
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertNewIdent(f *ast.NewIdent) (*NewIdent){

	return &NewIdent{
// other type &{}
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertSwitchStmt(f *ast.SwitchStmt) (*SwitchStmt){

	return &SwitchStmt{
		Switch : t.Converttoken_Pos( f.Switch),
		Init : t.ConvertTypeStmt( f.Init),
		Tag : t.ConvertTypeExpr( f.Tag),
		Body : t.ConvertBlockStmt( f.Body),
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertStructType(f *ast.StructType) (*StructType){

	return &StructType{
		Struct : t.Converttoken_Pos( f.Struct),
		Fields : t.ConvertFieldList( f.Fields),
		Incomplete : &f.Incomplete,
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertArrayType(f *ast.ArrayType) (*ArrayType){

	return &ArrayType{
		Lbrack : t.Converttoken_Pos( f.Lbrack),
		//Len : t.ConvertTypeExpr( f.Len),
		Elt : t.ConvertTypeExpr( f.Elt),
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertArrayStmt(f []ast.Stmt) ([]*Stmt){
	//panic("todo"); return nil;
	v2 := make([]*Stmt, len(f))
	for i,j := range(f) {
		x:=t.ConvertStmt(&j)
		//fmt.Printf("created %d %#v\n", i,*x)
		v2[i]=x
	}
	//fmt.Printf("created %#v\n", f)
	return v2

}

func (t* AddressTable) ConvertCaseClause(f *ast.CaseClause) (*CaseClause){

	return &CaseClause{
		Case : t.Converttoken_Pos( f.Case),
		List : t.ConvertArrayExpr( f.List),
		Colon : t.Converttoken_Pos( f.Colon),
		Body : t.ConvertArrayStmt( f.Body),
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertExpr(f ast.Expr) (*Expr){

	return &Expr{
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertCallExpr(f *ast.CallExpr) (*CallExpr){

	return &CallExpr{
		Fun : t.ConvertTypeExpr( f.Fun),
		Lparen : t.Converttoken_Pos( f.Lparen),
		Args : t.ConvertArrayExpr( f.Args),
		Ellipsis : t.Converttoken_Pos( f.Ellipsis),
		Rparen : t.Converttoken_Pos( f.Rparen),
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertValueSpec(f *ast.ValueSpec) (*ValueSpec){
	if f == nil { return nil }
	return &ValueSpec{
		//Doc : t.ConvertPointerCommentGroup( f.Doc),
		Names : t.ConvertArrayStarIdent( f.Names),
		Type : t.ConvertTypeExpr( f.Type),
		Values : t.ConvertArrayExpr( f.Values),
		//Comment : t.ConvertPointerCommentGroup( f.Comment),
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertStmt(f *ast.Stmt) (*Stmt){

	return &Stmt{
		//Incomplete: &f2,
	}
}

// func (t* AddressTable) ConvertRECV(f *ast.RECV) (*RECV){
//
//	return &RECV{
// // other type &{}
//		//Incomplete: &f2,
//	}
// }

func (t* AddressTable) ConvertComment(f *ast.Comment) (*Comment){

	return &Comment{
		//Slash : t.Converttoken_Pos( f.Slash),
		//Text : t.ConvertTypestring( f.Text),
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertBadStmt(f *ast.BadStmt) (*BadStmt){

	return &BadStmt{
		From : t.Converttoken_Pos2( f.From),
		To : t.Converttoken_Pos2( f.To),
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertLabeledStmt(f *ast.LabeledStmt) (*LabeledStmt){

	return &LabeledStmt{
		Label : t.ConvertIdent( f.Label),
		Colon : t.Converttoken_Pos( f.Colon),
		Stmt : t.ConvertTypeStmt( f.Stmt),
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertBasicLit(f *ast.BasicLit) (*BasicLit){

	return &BasicLit{
// other type &{}
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertTypeSpec(f *ast.TypeSpec) (*TypeSpec){

	return &TypeSpec{
		//Doc : t.ConvertPointerCommentGroup( f.Doc),
		Name : t.ConvertIdent( f.Name),
		Type : t.ConvertTypeExpr( f.Type),
		//Comment : t.ConvertPointerCommentGroup( f.Comment),
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertDecl(f ast.Decl) (*Decl){

	switch v:= f.(type) {
	case *ast.ArrayType:
	t2 := NodeType_ARRAYTYPE
	return &Decl{
	Type: &t2,
	Arraytype: t.ConvertArrayType(v),
	}
case *ast.AssignStmt:
	t2 := NodeType_ASSIGNSTMT
	return &Decl{
	Type: &t2,
	Assignstmt: t.ConvertAssignStmt(v),
	}
case *ast.BadDecl:
	t2 := NodeType_BADDECL
	return &Decl{
	Type: &t2,
	Baddecl: t.ConvertBadDecl(v),
	}
case *ast.BadExpr:
	t2 := NodeType_BADEXPR
	return &Decl{
	Type: &t2,
	Badexpr: t.ConvertBadExpr(v),
	}
case *ast.BadStmt:
	t2 := NodeType_BADSTMT
	return &Decl{
	Type: &t2,
	Badstmt: t.ConvertBadStmt(v),
	}
case *ast.BasicLit:
	t2 := NodeType_BASICLIT
	return &Decl{
	Type: &t2,
	Basiclit: t.ConvertBasicLit(v),
	}
case *ast.BinaryExpr:
	t2 := NodeType_BINARYEXPR
	return &Decl{
	Type: &t2,
	Binaryexpr: t.ConvertBinaryExpr(v),
	}
case *ast.BlockStmt:
	t2 := NodeType_BLOCKSTMT
	return &Decl{
	Type: &t2,
	Blockstmt: t.ConvertBlockStmt(v),
	}
case *ast.BranchStmt:
	t2 := NodeType_BRANCHSTMT
	return &Decl{
	Type: &t2,
	Branchstmt: t.ConvertBranchStmt(v),
	}
case *ast.CallExpr:
	t2 := NodeType_CALLEXPR
	return &Decl{
	Type: &t2,
	Callexpr: t.ConvertCallExpr(v),
	}
case *ast.CaseClause:
	t2 := NodeType_CASECLAUSE
	return &Decl{
	Type: &t2,
	Caseclause: t.ConvertCaseClause(v),
	}
case *ast.ChanDir:
	t2 := NodeType_CHANDIR
	return &Decl{
	Type: &t2,
	Chandir: t.ConvertChanDir(v),
	}
case *ast.ChanType:
	t2 := NodeType_CHANTYPE
	return &Decl{
	Type: &t2,
	Chantype: t.ConvertChanType(v),
	}
case *ast.CommClause:
	t2 := NodeType_COMMCLAUSE
	return &Decl{
	Type: &t2,
	Commclause: t.ConvertCommClause(v),
	}
case *ast.Comment:
	t2 := NodeType_COMMENT
	return &Decl{
	Type: &t2,
	Comment: t.ConvertComment(v),
	}
case *ast.CompositeLit:
	t2 := NodeType_COMPOSITELIT
	return &Decl{
	Type: &t2,
	Compositelit: t.ConvertCompositeLit(v),
	}
case *ast.DeclStmt:
	t2 := NodeType_DECLSTMT
	return &Decl{
	Type: &t2,
	Declstmt: t.ConvertDeclStmt(v),
	}
case *ast.DeferStmt:
	t2 := NodeType_DEFERSTMT
	return &Decl{
	Type: &t2,
	Deferstmt: t.ConvertDeferStmt(v),
	}
case *ast.Ellipsis:
	t2 := NodeType_ELLIPSIS
	return &Decl{
	Type: &t2,
	Ellipsis: t.ConvertEllipsis(v),
	}
case *ast.EmptyStmt:
	t2 := NodeType_EMPTYSTMT
	return &Decl{
	Type: &t2,
	Emptystmt: t.ConvertEmptyStmt(v),
	}
case *ast.ExprStmt:
	t2 := NodeType_EXPRSTMT
	return &Decl{
	Type: &t2,
	Exprstmt: t.ConvertExprStmt(v),
	}
case *ast.Field:
	t2 := NodeType_FIELD
	return &Decl{
	Type: &t2,
	Field: t.ConvertField(v),
	}
case *ast.FieldList:
	t2 := NodeType_FIELDLIST
	return &Decl{
	Type: &t2,
	Fieldlist: t.ConvertFieldList(v),
	}
case *ast.File:
	t2 := NodeType_FILE
	return &Decl{
	Type: &t2,
	File: t.ConvertFile(v),
	}
case *ast.ForStmt:
	t2 := NodeType_FORSTMT
	return &Decl{
	Type: &t2,
	Forstmt: t.ConvertForStmt(v),
	}
case *ast.FuncDecl:
	t2 := NodeType_FUNCDECL
	return &Decl{
	Type: &t2,
	Funcdecl: t.ConvertFuncDecl(v),
	}
case *ast.FuncLit:
	t2 := NodeType_FUNCLIT
	return &Decl{
	Type: &t2,
	Funclit: t.ConvertFuncLit(v),
	}
case *ast.FuncType:
	t2 := NodeType_FUNCTYPE
	return &Decl{
	Type: &t2,
	Functype: t.ConvertFuncType(v),
	}
case *ast.GenDecl:
	t2 := NodeType_GENDECL
	return &Decl{
	Type: &t2,
	Gendecl: t.ConvertGenDecl(v),
	}
case *ast.GoStmt:
	t2 := NodeType_GOSTMT
	return &Decl{
	Type: &t2,
	Gostmt: t.ConvertGoStmt(v),
	}
case *ast.Ident:
	t2 := NodeType_IDENT
	return &Decl{
	Type: &t2,
	Ident: t.ConvertIdent(v),
	}
case *ast.IfStmt:
	t2 := NodeType_IFSTMT
	return &Decl{
	Type: &t2,
	Ifstmt: t.ConvertIfStmt(v),
	}
case *ast.ImportSpec:
	t2 := NodeType_IMPORTSPEC
	return &Decl{
	Type: &t2,
	Importspec: t.ConvertImportSpec(v),
	}
case *ast.IncDecStmt:
	t2 := NodeType_INCDECSTMT
	return &Decl{
	Type: &t2,
	Incdecstmt: t.ConvertIncDecStmt(v),
	}
case *ast.IndexExpr:
	t2 := NodeType_INDEXEXPR
	return &Decl{
	Type: &t2,
	Indexexpr: t.ConvertIndexExpr(v),
	}
case *ast.InterfaceType:
	t2 := NodeType_INTERFACETYPE
	return &Decl{
	Type: &t2,
	Interfacetype: t.ConvertInterfaceType(v),
	}
case *ast.IsExported:
	t2 := NodeType_ISEXPORTED
	return &Decl{
	Type: &t2,
	Isexported: t.ConvertIsExported(v),
	}
case *ast.KeyValueExpr:
	t2 := NodeType_KEYVALUEEXPR
	return &Decl{
	Type: &t2,
	Keyvalueexpr: t.ConvertKeyValueExpr(v),
	}
case *ast.LabeledStmt:
	t2 := NodeType_LABELEDSTMT
	return &Decl{
	Type: &t2,
	Labeledstmt: t.ConvertLabeledStmt(v),
	}
case *ast.MapType:
	t2 := NodeType_MAPTYPE
	return &Decl{
	Type: &t2,
	Maptype: t.ConvertMapType(v),
	}
case *ast.NewIdent:
	t2 := NodeType_NEWIDENT
	return &Decl{
	Type: &t2,
	Newident: t.ConvertNewIdent(v),
	}
case *ast.Package:
	t2 := NodeType_PACKAGE
	return &Decl{
	Type: &t2,
	Package: t.ConvertPackage(v),
	}
case *ast.ParenExpr:
	t2 := NodeType_PARENEXPR
	return &Decl{
	Type: &t2,
	Parenexpr: t.ConvertParenExpr(v),
	}
case *ast.RangeStmt:
	t2 := NodeType_RANGESTMT
	return &Decl{
	Type: &t2,
	Rangestmt: t.ConvertRangeStmt(v),
	}
case *ast.ReturnStmt:
	t2 := NodeType_RETURNSTMT
	return &Decl{
	Type: &t2,
	Returnstmt: t.ConvertReturnStmt(v),
	}
case *ast.SelectStmt:
	t2 := NodeType_SELECTSTMT
	return &Decl{
	Type: &t2,
	Selectstmt: t.ConvertSelectStmt(v),
	}
case *ast.SelectorExpr:
	t2 := NodeType_SELECTOREXPR
	return &Decl{
	Type: &t2,
	Selectorexpr: t.ConvertSelectorExpr(v),
	}
case *ast.SendStmt:
	t2 := NodeType_SENDSTMT
	return &Decl{
	Type: &t2,
	Sendstmt: t.ConvertSendStmt(v),
	}
case *ast.SliceExpr:
	t2 := NodeType_SLICEEXPR
	return &Decl{
	Type: &t2,
	Sliceexpr: t.ConvertSliceExpr(v),
	}
/*
case *ast.Spec:
	t2 := NodeType_SPEC
	return &Decl{
	Type: &t2,
	Spec: t.ConvertSpec(v),
	}*/
case *ast.StarExpr:
	t2 := NodeType_STAREXPR
	return &Decl{
	Type: &t2,
	Starexpr: t.ConvertStarExpr(v),
	}
case *ast.StructType:
	t2 := NodeType_STRUCTTYPE
	return &Decl{
	Type: &t2,
	Structtype: t.ConvertStructType(v),
	}
case *ast.SwitchStmt:
	t2 := NodeType_SWITCHSTMT
	return &Decl{
	Type: &t2,
	Switchstmt: t.ConvertSwitchStmt(v),
	}
case *ast.TypeAssertExpr:
	t2 := NodeType_TYPEASSERTEXPR
	return &Decl{
	Type: &t2,
	Typeassertexpr: t.ConvertTypeAssertExpr(v),
	}
case *ast.TypeSpec:
	t2 := NodeType_TYPESPEC
	return &Decl{
	Type: &t2,
	Typespec: t.ConvertTypeSpec(v),
	}
case *ast.TypeSwitchStmt:
	t2 := NodeType_TYPESWITCHSTMT
	return &Decl{
	Type: &t2,
	Typeswitchstmt: t.ConvertTypeSwitchStmt(v),
	}
case *ast.UnaryExpr:
	t2 := NodeType_UNARYEXPR
	return &Decl{
	Type: &t2,
	Unaryexpr: t.ConvertUnaryExpr(v),
	}
case *ast.ValueSpec:
	t2 := NodeType_VALUESPEC
	return &Decl{
	Type: &t2,
	Valuespec: t.ConvertValueSpec(v),
	}
case *ast.CommentGroup:
	t2 := NodeType_COMMENTGROUP
	return &Decl{
	Type: &t2,
	Commentgroup: t.ConvertCommentGroup(v),
	}
case *ast.Object:
	t2 := NodeType_OBJECT
	return &Decl{
	Type: &t2,
	Object: t.ConvertObject(v),
	}

	default:
		fmt.Printf("stmt unknown %T",v)
		
	}
	// return &Decl{
	// 	//Incomplete: &f2,
	// }

	panic("todo")
}

func (t* AddressTable) ConvertPointerFuncType(f *ast.FuncType) (*FuncType){
	panic("todo"); return nil;
}

func (t* AddressTable) ConvertPointerFuncType2(f ast.Foo2) (*Foo2){
	panic("todo"); return nil;
}

func (t* AddressTable) ConvertFuncDecl(f *ast.FuncDecl) (*FuncDecl){
	if f == nil {return nil}
	return &FuncDecl{
		//Doc : t.ConvertPointerCommentGroup( f.Doc),
		Recv : t.ConvertFieldList( f.Recv),
		Name : t.ConvertIdent( f.Name),
		Type : t.ConvertFuncType( f.Type),
		Body : t.ConvertBlockStmt( f.Body),
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertFuncLit(f *ast.FuncLit) (*FuncLit){

	return &FuncLit{
		Type : t.ConvertFuncType( f.Type),
		Body : t.ConvertBlockStmt( f.Body),
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertForStmt(f *ast.ForStmt) (*ForStmt){

	return &ForStmt{
		For : t.Converttoken_Pos( f.For),
		Init : t.ConvertTypeStmt( f.Init),
		Cond : t.ConvertTypeExpr( f.Cond),
		Post : t.ConvertTypeStmt( f.Post),
		Body : t.ConvertBlockStmt( f.Body),
		//Incomplete: &f2,
	}
}



func (t* AddressTable) ConvertSelectStmt(f *ast.SelectStmt) (*SelectStmt){

	return &SelectStmt{
		Select : t.Converttoken_Pos2( f.Select),
		Body : t.ConvertBlockStmt( f.Body),
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertSendStmt(f *ast.SendStmt) (*SendStmt){

	return &SendStmt{
		Chan : t.ConvertTypeExpr( f.Chan),
		Arrow : t.Converttoken_Pos2( f.Arrow),
		Value : t.ConvertTypeExpr( f.Value),
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertPackage(f *ast.Package) (*Package){

	return &Package{
		Name : t.ConvertTypestring( f.Name),
		Scope : t.ConvertPointerScope2( f.Scope),
		Imports : t.ConvertMapstringToPointerObject( f.Imports),
		Files : t.ConvertMapstringToPointerFile( f.Files),
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertPointerFieldList2(f ast.Foo2) (*Foo2){
	panic("todo"); return nil;
}

func (t* AddressTable) ConvertFieldList(f *ast.FieldList) (*FieldList){
	if f == nil {return nil}
	return &FieldList{
		Opening : t.Converttoken_Pos( f.Opening),
		List : t.ConvertArrayStarField( f.List),
		Closing : t.Converttoken_Pos( f.Closing),
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertImportSpec(f *ast.ImportSpec) (*ImportSpec){

	return &ImportSpec{
		//Doc : t.ConvertPointerCommentGroup( f.Doc),
		Name : t.ConvertIdent( f.Name),
		Path : t.ConvertBasicLit( f.Path),
		//Comment : t.ConvertPointerCommentGroup( f.Comment),
		//EndPos : t.Converttoken_Pos( f.EndPos),
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertReturnStmt(f *ast.ReturnStmt) (*ReturnStmt){

	return &ReturnStmt{
		//Return : t.Converttoken_Pos( f.Return),
		//Results : t.ConvertArrayExpr( f.Results),
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertMapstringToPointerObject(f map[string]*ast.Object) (map[string]*Object){

	r := make(map[string]*Object)
	
	for i,j := range(f) {
		o:=t.ConvertObject(j)
		r[i]=o
	}
	return r
	//panic("todo"); return nil;
}

func (t* AddressTable) ConvertMapstringToPointerFile(f map[string]*ast.File) (map[string]*File){
	panic("todo"); return nil;
}

func (t* AddressTable) ConvertArrayStarField(f []*ast.Field) ([]*Field){
	//panic("todo"); return nil;
	v2 := make([]*Field, len(f))
	for i,j := range(f) {
		x:=t.ConvertField(j)
		//fmt.Printf("created %d %#v\n", i,*x)
		v2[i]=x
	}
	//fmt.Printf("created %#v\n", f)
	return v2
}


func (t* AddressTable) ConvertPointerBlockStmt2(f *ast.BlockStmt) (*BlockStmt){
	panic("todo"); return nil;
}

func (t* AddressTable) ConvertPointerScope(f *ast.Scope) (*Scope){
	panic("todo"); return nil;
}

func (t* AddressTable) ConvertPointerScope2(f ast.Foo2) (*Foo2){
	panic("todo"); return nil;
}

func (t* AddressTable) ConvertTypestring(f string) (*string){
	panic("todo"); return nil;
}

func (t* AddressTable) ConvertBlockStmt(f *ast.BlockStmt) (*BlockStmt){

	return &BlockStmt{
// other type &{}
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertTypeSwitchStmt(f *ast.TypeSwitchStmt) (*TypeSwitchStmt){

	return &TypeSwitchStmt{
		Switch : t.Converttoken_Pos2( f.Switch),
		Init : t.ConvertTypeStmt( f.Init),
		Assign : t.ConvertTypeStmt( f.Assign),
		Body : t.ConvertBlockStmt( f.Body),
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertSpec(f *ast.Spec) (*Spec){

	return &Spec{
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertStarExpr(f *ast.StarExpr) (*StarExpr){

	return &StarExpr{
		//Star : t.Converttoken_Pos( f.Star),
		//X : t.ConvertTypeExpr( f.X),
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertGenDecl(f *ast.GenDecl) (*GenDecl){

	return &GenDecl{
		//Doc : t.ConvertPointerCommentGroup( f.Doc),
		TokPos : t.Converttoken_Pos( f.TokPos),
		Tok : t.Converttoken_Token( f.Tok),
		Lparen : t.Converttoken_Pos( f.Lparen),
		Specs : t.ConvertArraySpec( f.Specs),
		Rparen : t.Converttoken_Pos( f.Rparen),
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertRangeStmt(f *ast.RangeStmt) (*RangeStmt){

	return &RangeStmt{
		For : t.Converttoken_Pos( f.For),
		Key : t.ConvertTypeExpr( f.Key),
		Value : t.ConvertTypeExpr( f.Value),
		TokPos : t.Converttoken_Pos( f.TokPos),
		Tok : t.Converttoken_Token( f.Tok),
		X : t.ConvertTypeExpr( f.X),
		Body : t.ConvertBlockStmt( f.Body),
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertParenExpr(f *ast.ParenExpr) (*ParenExpr){

	return &ParenExpr{
		Lparen : t.Converttoken_Pos( f.Lparen),
		X : t.ConvertTypeExpr( f.X),
		Rparen : t.Converttoken_Pos( f.Rparen),
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertDeclStmt(f *ast.DeclStmt) (*DeclStmt){

	return &DeclStmt{
		//Decl : t.ConvertTypeDecl( f.Decl),
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertTypeAssertExpr(f *ast.TypeAssertExpr) (*TypeAssertExpr){

	return &TypeAssertExpr{
		X : t.ConvertTypeExpr( f.X),
		Lparen : t.Converttoken_Pos( f.Lparen),
		Type : t.ConvertTypeExpr( f.Type),
		Rparen : t.Converttoken_Pos( f.Rparen),
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertEmptyStmt(f *ast.EmptyStmt) (*EmptyStmt){

	return &EmptyStmt{
		Semicolon : t.Converttoken_Pos2( f.Semicolon),
		Implicit : &f.Implicit,
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertIndexExpr(f *ast.IndexExpr) (*IndexExpr){

	return &IndexExpr{
		X : t.ConvertTypeExpr( f.X),
		Lbrack : t.Converttoken_Pos( f.Lbrack),
		Index : t.ConvertTypeExpr( f.Index),
		Rbrack : t.Converttoken_Pos( f.Rbrack),
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertArrayDecl(f []ast.Decl)([]*Decl){
	
	v2 := make([]*Decl, len(f))
	for i,j := range(f) {
		x:=t.ConvertDecl(j)
		//fmt.Printf("created %d %#v\n", i,*x)
		v2[i]=x
	}
	//fmt.Printf("created %#v\n", f)
	return v2

}

func (t* AddressTable) ConvertArrayStarImportSpec (f []*ast.ImportSpec)([]*ImportSpec){
	//panic("todo"); return nil;
	v2 := make([]*ImportSpec, len(f))
	for i,j := range(f) {
		x:=t.ConvertImportSpec(j)
		//fmt.Printf("created %d %#v\n", i,*x)
		v2[i]=x
	}
	//fmt.Printf("created %#v\n", f)
	return v2

}
func (t* AddressTable) ConvertArrayStarCommentGroup(f []*ast.CommentGroup)([]*CommentGroup){
	//panic("todo"); return nil;
	v2 := make([]*CommentGroup, len(f))
	for i,j := range(f) {
		x:=t.ConvertCommentGroup(j)
		//fmt.Printf("created %d %#v\n", i,*x)
		v2[i]=x
	}
	//fmt.Printf("created %#v\n", f)
	return v2
	
}

func (t* AddressTable) ConvertFile(f *ast.File) (*File){

	return &File{
		//Doc : t.ConvertPointerCommentGroup( f.Doc),
		Package : t.Converttoken_Pos( f.Package),
		Name : t.ConvertIdent( f.Name),
		Decls : t.ConvertArrayDecl( f.Decls),
		Scope : t.ConvertScope( f.Scope),
		Imports : t.ConvertArrayStarImportSpec( f.Imports),
		Unresolved : t.ConvertArrayStarIdent( f.Unresolved),
		Comments : t.ConvertArrayStarCommentGroup( f.Comments),
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertBranchStmt(f *ast.BranchStmt) (*BranchStmt){

	return &BranchStmt{
		TokPos : t.Converttoken_Pos( f.TokPos),
		Tok : t.Converttoken_Token( f.Tok),
		Label : t.ConvertIdent( f.Label),
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertCommentGroup(f *ast.CommentGroup) (*CommentGroup){

	return &CommentGroup{
		//List : t.ConvertArrayStarComment( f.List),
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertArraySpec(f []ast.Spec) ([]*Spec){
	v2 := make([]*Spec, len(f))
	for i,j := range(f) {
		x:=t.ConvertSpec(&j)
		//fmt.Printf("created %d %#v\n", i,*x)
		v2[i]=x
	}
	//fmt.Printf("created %#v\n", f)
	return v2
}

func (t* AddressTable) ConvertCommClause(f *ast.CommClause) (*CommClause){

	return &CommClause{
		Case : t.Converttoken_Pos2( f.Case),
		Comm : t.ConvertTypeStmt( f.Comm),
		Colon : t.Converttoken_Pos2( f.Colon),
		Body : t.ConvertArrayStmt( f.Body),
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertIsExported(f *ast.IsExported) (*IsExported){

	return &IsExported{
// other type &{}
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertChanType(f *ast.ChanType) (*ChanType){

	return &ChanType{
		//Begin : t.Converttoken_Pos( f.Begin),
//		Arrow : t.Converttoken_Pos( f.Arrow),
//		Dir : t.ConvertTypeChanDir( f.Dir),
//		Value : t.ConvertTypeExpr( f.Value),
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertGoStmt(f *ast.GoStmt) (*GoStmt){

	return &GoStmt{
		Go : t.Converttoken_Pos2( f.Go),
		Call : t.ConvertCallExpr( f.Call),
		//Incomplete: &f2,
	}
}

// func (t* AddressTable) ConvertstripTrailingWhitespace(f *ast.stripTrailingWhitespace) (*stripTrailingWhitespace){
//
//	return &stripTrailingWhitespace{
// // other type &{}
//		//Incomplete: &f2,
//	}
// }

func (t* AddressTable) ConvertEllipsis(f *ast.Ellipsis) (*Ellipsis){

	return &Ellipsis{
		Ellipsis : t.Converttoken_Pos( f.Ellipsis),
		Elt : t.ConvertTypeExpr( f.Elt),
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertFuncType(f *ast.FuncType) (*FuncType){

	return &FuncType{
// other type &{}
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertUnaryExpr(f *ast.UnaryExpr) (*UnaryExpr){

	return &UnaryExpr{
		OpPos : t.Converttoken_Pos( f.OpPos),
		Op : t.Converttoken_Token( f.Op),
		X : t.ConvertTypeExpr( f.X),
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertNode(f *ast.Node) (*Node){

	return &Node{
		//Incomplete: &f2,
	}
}


// func (t* AddressTable) StrmapIdent(id string, f * ast.Ident) (*Ident){
//	f2 := t.ConvertPointerIdent(f)
//	t.Idents[id] = f2;
//	f.Report();
//	return f2
// }

// //func (t* AddressTable) StrmapBasicLit(id string, f * BasicLit) (*BasicLit){ t.BasicLits[id] =f; f.Report(); return f}
// func (t* AddressTable) StrmapBasicLit(id string, f * ast.BasicLit) (*BasicLit){
//	f2 := t.ConvertBasicLit(f)
//	t.BasicLits[id] = f2;
//	f.Report();
//	return f2
// }

// func (t* AddressTable) StrmapImportSpec(id string, f * ast.ImportSpec) (*ImportSpec){
//	f2 := t.ConvertImportSpec(f)
//	t.ImportSpecs[id] =f2;
//	f.Report();
//	return f2

// }
// func (t* AddressTable) StrmapGenDecl(id string, f * ast.GenDecl) (*GenDecl){
//	f2 := t.ConvertGenDecl(f)
//	t.GenDecls[id] =f2;
//	f.Report();
//	return f2
// }

func (t* AddressTable) ConvertBadDecl(f * ast.BadDecl) (*BadDecl){
	panic("todo"); return nil;
}

func (t* AddressTable) ConvertScope(f * ast.Scope) (*Scope){
	//panic("todo"); return nil;
	return &Scope{
		Objects : t.ConvertMapstringToPointerObject( f.Objects ),
	}
}

func (t* AddressTable) StrmapScope(id string, f * ast.Scope) (*ast.Scope){
	f2 := t.ConvertScope(f);
	t.Check(id,f2,f)
	t.Scopes[id] =f;
	//f.Report();
	return f}


func (t* AddressTable) StrmapArrayType(id string, f * ast.ArrayType) (*ast.ArrayType){
	f2 := t.ConvertArrayType(f);
	t.Check(id,f2,f)
	t.ArrayTypes[id] =f;
	return f
}
func (t* AddressTable) StrmapAssignStmt(id string, f * ast.AssignStmt) (*ast.AssignStmt){ f2 := t.ConvertAssignStmt(f); t.AssignStmts[id] =f; t.Check(id,f2,f); return f}
func (t* AddressTable) StrmapBadDecl(id string, f * ast.BadDecl) (*ast.BadDecl){ f2 := t.ConvertBadDecl(f); t.BadDecls[id] =f; t.Check(id,f2,f); return f}
func (t* AddressTable) StrmapBadExpr(id string, f * ast.BadExpr) (*ast.BadExpr){ f2 := t.ConvertBadExpr(f); t.BadExprs[id] =f; t.Check(id,f2,f); return f}
     func (t* AddressTable) StrmapBadStmt(id string, f * ast.BadStmt) (*ast.BadStmt){ f2 := t.ConvertBadStmt(f); t.BadStmts[id] =f; t.Check(id,f2,f); return f}
     func (t* AddressTable) StrmapBasicLit(id string, f * ast.BasicLit) (*ast.BasicLit){ f2 := t.ConvertBasicLit(f); t.BasicLits[id] =f; t.Check(id,f2,f); return f}
     func (t* AddressTable) StrmapBinaryExpr(id string, f * ast.BinaryExpr) (*ast.BinaryExpr){ f2 := t.ConvertBinaryExpr(f); t.BinaryExprs[id] =f; t.Check(id,f2,f); return f}
     func (t* AddressTable) StrmapBlockStmt(id string, f * ast.BlockStmt) (*ast.BlockStmt){ f2 := t.ConvertBlockStmt(f); t.BlockStmts[id] =f; t.Check(id,f2,f); return f}
    func (t* AddressTable) StrmapBranchStmt(id string, f * ast.BranchStmt) (*ast.BranchStmt){ f2 := t.ConvertBranchStmt(f); t.BranchStmts[id] =f; t.Check(id,f2,f); return f}
    func (t* AddressTable) StrmapCallExpr(id string, f * ast.CallExpr) (*ast.CallExpr){ f2 := t.ConvertCallExpr(f); t.CallExprs[id] =f; t.Check(id,f2,f); return f}
    func (t* AddressTable) StrmapCaseClause(id string, f * ast.CaseClause) (*ast.CaseClause){ f2 := t.ConvertCaseClause(f); t.CaseClauses[id] =f; t.Check(id,f2,f); return f}
    func (t* AddressTable) StrmapChanDir(id string, f * ast.ChanDir) (*ast.ChanDir){ f2 := t.ConvertChanDir(f); t.ChanDirs[id] =f; t.Check(id,f2,f); return f}
    func (t* AddressTable) StrmapChanType(id string, f * ast.ChanType) (*ast.ChanType){ f2 := t.ConvertChanType(f); t.ChanTypes[id] =f; t.Check(id,f2,f); return f}
    func (t* AddressTable) StrmapCommClause(id string, f * ast.CommClause) (*ast.CommClause){ f2 := t.ConvertCommClause(f); t.CommClauses[id] =f; t.Check(id,f2,f); return f}
    func (t* AddressTable) StrmapComment(id string, f * ast.Comment) (*ast.Comment){ f2 := t.ConvertComment(f); t.Comments[id] =f; t.Check(id,f2,f); return f}
    func (t* AddressTable) StrmapCompositeLit(id string, f * ast.CompositeLit) (*ast.CompositeLit){ f2 := t.ConvertCompositeLit(f); t.CompositeLits[id] =f; t.Check(id,f2,f); return f}
    func (t* AddressTable) StrmapDeclStmt(id string, f * ast.DeclStmt) (*ast.DeclStmt){ f2 := t.ConvertDeclStmt(f); t.DeclStmts[id] =f; t.Check(id,f2,f); return f}
    func (t* AddressTable) StrmapDeferStmt(id string, f * ast.DeferStmt) (*ast.DeferStmt){ f2 := t.ConvertDeferStmt(f); t.DeferStmts[id] =f; t.Check(id,f2,f); return f}
    func (t* AddressTable) StrmapEllipsis(id string, f * ast.Ellipsis) (*ast.Ellipsis){ f2 := t.ConvertEllipsis(f); t.Ellipsiss[id] =f; t.Check(id,f2,f); return f}
    func (t* AddressTable) StrmapEmptyStmt(id string, f * ast.EmptyStmt) (*ast.EmptyStmt){ f2 := t.ConvertEmptyStmt(f); t.EmptyStmts[id] =f; t.Check(id,f2,f); return f}
    func (t* AddressTable) StrmapExprStmt(id string, f * ast.ExprStmt) (*ast.ExprStmt){ f2 := t.ConvertExprStmt(f); t.ExprStmts[id] =f; t.Check(id,f2,f); return f}
    func (t* AddressTable) StrmapField(id string, f * ast.Field) (*ast.Field){ f2 := t.ConvertField(f); t.Fields[id] =f; t.Check(id,f2,f); return f}
    func (t* AddressTable) StrmapFieldList(id string, f * ast.FieldList) (*ast.FieldList){ f2 := t.ConvertFieldList(f); t.FieldLists[id] =f; t.Check(id,f2,f); return f}
    func (t* AddressTable) StrmapFile(id string, f * ast.File) (*ast.File){ f2 := t.ConvertFile(f); t.Files[id] =f; t.Check(id,f2,f); return f}
    func (t* AddressTable) StrmapForStmt(id string, f * ast.ForStmt) (*ast.ForStmt){ f2 := t.ConvertForStmt(f); t.ForStmts[id] =f; t.Check(id,f2,f); return f}
    func (t* AddressTable) StrmapFuncDecl(id string, f * ast.FuncDecl) (*ast.FuncDecl){ f2 := t.ConvertFuncDecl(f); t.FuncDecls[id] =f; t.Check(id,f2,f); return f}
    func (t* AddressTable) StrmapFuncLit(id string, f * ast.FuncLit) (*ast.FuncLit){ f2 := t.ConvertFuncLit(f); t.FuncLits[id] =f; t.Check(id,f2,f); return f}
    func (t* AddressTable) StrmapFuncType(id string, f * ast.FuncType) (*ast.FuncType){ f2 := t.ConvertFuncType(f); t.FuncTypes[id] =f; t.Check(id,f2,f); return f}
func (t* AddressTable) StrmapGenDecl(id string, f * ast.GenDecl) (*ast.GenDecl){ f2 := t.ConvertGenDecl(f); t.GenDecls[id] =f; t.Check(id,f2,f); return f}
func (t* AddressTable) StrmapGoStmt(id string, f * ast.GoStmt) (*ast.GoStmt){ f2 := t.ConvertGoStmt(f); t.GoStmts[id] =f; t.Check(id,f2,f); return f}

func (t* AddressTable) StrmapIfStmt(id string, f * ast.IfStmt) (*ast.IfStmt){ f2 := t.ConvertIfStmt(f); t.IfStmts[id] =f; t.Check(id,f2,f); return f}
func (t* AddressTable) StrmapImportSpec(id string, f * ast.ImportSpec) (*ast.ImportSpec){ f2 := t.ConvertImportSpec(f); t.ImportSpecs[id] =f; t.Check(id,f2,f); return f}
func (t* AddressTable) StrmapIncDecStmt(id string, f * ast.IncDecStmt) (*ast.IncDecStmt){ f2 := t.ConvertIncDecStmt(f); t.IncDecStmts[id] =f; t.Check(id,f2,f); return f}
func (t* AddressTable) StrmapIndexExpr(id string, f * ast.IndexExpr) (*ast.IndexExpr){ f2 := t.ConvertIndexExpr(f); t.IndexExprs[id] =f; t.Check(id,f2,f); return f}
func (t* AddressTable) StrmapInterfaceType(id string, f * ast.InterfaceType) (*ast.InterfaceType){ f2 := t.ConvertInterfaceType(f); t.InterfaceTypes[id] =f; t.Check(id,f2,f); return f}
func (t* AddressTable) StrmapIsExported(id string, f * ast.IsExported) (*ast.IsExported){ f2 := t.ConvertIsExported(f); t.IsExporteds[id] =f; t.Check(id,f2,f); return f}
func (t* AddressTable) StrmapKeyValueExpr(id string, f * ast.KeyValueExpr) (*ast.KeyValueExpr){ f2 := t.ConvertKeyValueExpr(f); t.KeyValueExprs[id] =f; t.Check(id,f2,f); return f}
func (t* AddressTable) StrmapLabeledStmt(id string, f * ast.LabeledStmt) (*ast.LabeledStmt){ f2 := t.ConvertLabeledStmt(f); t.LabeledStmts[id] =f; t.Check(id,f2,f); return f}
func (t* AddressTable) StrmapMapType(id string, f * ast.MapType) (*ast.MapType){ f2 := t.ConvertMapType(f); t.MapTypes[id] =f; t.Check(id,f2,f); return f}
func (t* AddressTable) StrmapNewIdent(id string, f * ast.NewIdent) (*ast.NewIdent){ f2 := t.ConvertNewIdent(f); t.NewIdents[id] =f; t.Check(id,f2,f); return f}
func (t* AddressTable) StrmapPackage(id string, f * ast.Package) (*ast.Package){ f2 := t.ConvertPackage(f); t.Packages[id] =f; t.Check(id,f2,f); return f}
func (t* AddressTable) StrmapParenExpr(id string, f * ast.ParenExpr) (*ast.ParenExpr){ f2 := t.ConvertParenExpr(f); t.ParenExprs[id] =f; t.Check(id,f2,f); return f}
func (t* AddressTable) StrmapRangeStmt(id string, f * ast.RangeStmt) (*ast.RangeStmt){ f2 := t.ConvertRangeStmt(f); t.RangeStmts[id] =f; t.Check(id,f2,f); return f}
func (t* AddressTable) StrmapReturnStmt(id string, f * ast.ReturnStmt) (*ast.ReturnStmt){ f2 := t.ConvertReturnStmt(f); t.ReturnStmts[id] =f; t.Check(id,f2,f); return f}
func (t* AddressTable) StrmapSelectStmt(id string, f * ast.SelectStmt) (*ast.SelectStmt){ f2 := t.ConvertSelectStmt(f); t.SelectStmts[id] =f; t.Check(id,f2,f); return f}
func (t* AddressTable) StrmapSelectorExpr(id string, f * ast.SelectorExpr) (*ast.SelectorExpr){ f2 := t.ConvertSelectorExpr(f); t.SelectorExprs[id] =f; t.Check(id,f2,f); return f}
func (t* AddressTable) StrmapSendStmt(id string, f * ast.SendStmt) (*ast.SendStmt){ f2 := t.ConvertSendStmt(f); t.SendStmts[id] =f; t.Check(id,f2,f); return f}
func (t* AddressTable) StrmapSliceExpr(id string, f * ast.SliceExpr) (*ast.SliceExpr){ f2 := t.ConvertSliceExpr(f); t.SliceExprs[id] =f; t.Check(id,f2,f); return f}
//func (t* AddressTable) StrmapSpec(id string, f * ast.Spec) (*ast.Spec){ f2 := t.ConvertSpec(f); t.Specs[id] =f; t.Check(id,f2,f); return f}
func (t* AddressTable) StrmapStarExpr(id string, f * ast.StarExpr) (*ast.StarExpr){
	f2 := t.ConvertStarExpr(f);
	t.StarExprs[id] =f; t.Check(id,f2,f);
	return f
}

func (t* AddressTable) StrmapStructType(id string, f * ast.StructType) (*ast.StructType){ f2 := t.ConvertStructType(f); t.StructTypes[id] =f; t.Check(id,f2,f); return f}
func (t* AddressTable) StrmapSwitchStmt(id string, f * ast.SwitchStmt) (*ast.SwitchStmt){ f2 := t.ConvertSwitchStmt(f); t.SwitchStmts[id] =f; t.Check(id,f2,f); return f}
func (t* AddressTable) StrmapTypeAssertExpr(id string, f * ast.TypeAssertExpr) (*ast.TypeAssertExpr){ f2 := t.ConvertTypeAssertExpr(f); t.TypeAssertExprs[id] =f; t.Check(id,f2,f); return f}
func (t* AddressTable) StrmapTypeSpec(id string, f * ast.TypeSpec) (*ast.TypeSpec){ f2 := t.ConvertTypeSpec(f); t.TypeSpecs[id] =f; t.Check(id,f2,f); return f}
func (t* AddressTable) StrmapTypeSwitchStmt(id string, f * ast.TypeSwitchStmt) (*ast.TypeSwitchStmt){ f2 := t.ConvertTypeSwitchStmt(f); t.TypeSwitchStmts[id] =f; t.Check(id,f2,f); return f}
func (t* AddressTable) StrmapUnaryExpr(id string, f * ast.UnaryExpr) (*ast.UnaryExpr){ f2 := t.ConvertUnaryExpr(f); t.UnaryExprs[id] =f; t.Check(id,f2,f); return f}
func (t* AddressTable) StrmapValueSpec(id string, f * ast.ValueSpec) (*ast.ValueSpec){ f2 := t.ConvertValueSpec(f); t.ValueSpecs[id] =f; t.Check(id,f2,f); return f}
func (t* AddressTable) StrmapCommentGroup(id string, f * ast.CommentGroup) (*ast.CommentGroup){ f2 := t.ConvertCommentGroup(f); t.CommentGroups[id] =f; t.Check(id,f2,f); return f}
func (t* AddressTable) StrmapSpec(id string, f * ast.Spec) (*ast.Spec){ f2 := t.ConvertSpec(f); t.Specs[id] =f; t.Check(id,f2,f); return f}
func (t* AddressTable) StrmapObject(id string, f * ast.Object) (*ast.Object){ f2 := t.ConvertObject(f); t.Objects[id] =f; t.Check(id,f2,f); return f}

func (t* AddressTable) ConvertDeferred(f * ast.Deferred) (*Deferred){

	if f == nil { return nil }
	//fmt.Printf("convert deferred %#v",f)


	switch v:= f.Set.(type) {

		
	case *map[string]*ast.TypeSpec:
		nt := NodeType_TYPESPEC
		return &Deferred {
			Id: &f.Id,
			Type: &nt,
		}
		
	case *map[string]*ast.Object:
		nt := NodeType_OBJECT
		return &Deferred {
			Id: &f.Id,
			Type: &nt,
		}

	default:
		fmt.Printf("type debug %T",v)
	}
	panic("todo"); return nil;
}

func (t* AddressTable) ConvertFieldDeferred(f * ast.FieldDeferred) (*FieldDeferred){

	if f == nil { return nil }
	//fmt.Printf("debug field %#v",f)

	//nt := NodeType_FIELDDEFERRED
	return &FieldDeferred {
		Id: &f.Id,
	}


	panic("todo"); return nil;
}

// func (t* AddressTable) ConvertDeferred2(f  ast.Deferred2) (*Deferred){
//	panic("todo"); return nil;
// }

func (t* AddressTable) ConvertString(f  ast.Foo3) (*string){
	//return &f
	if f == nil {
		return nil;
	}
	switch v:= f.(type) {
	case string:
//x		fmt.Printf("Native string: %s\n", v)
		return &v
	}
	panic("todo"); return nil;
}

func (t* AddressTable) ConvertObject(f * ast.Object) (*Object){

	if f == nil {
		panic("todo"); return nil;
	}

	return & Object{
		Deferred : t.ConvertDeferred(f.Deferred),
		//*Deferred `protobuf:"bytes,1,opt,name=Deferred,json=deferred" json:"Deferred,omitempty"`
		Kind:       t.ConvertString(f.Kind),
		Name:       t.ConvertString(f.Name),
		Decl:       t.ConvertFoo2(f.Decl),
		Data:       t.ConvertFoo3(f.Data),
	}
	//panic("todo"); return nil;
}


func (t* AddressTable) ConvertArrayStarIdent(f ast.Foo3) ([]*Ident){
	//fmt.Printf("ConvertArrayStarIdent %#v\n", f)

	switch v:= f.(type) {
	case []*ast.Ident:
		v2 := make([]*Ident, len(v))
		for i,j := range(v) {
			x:=t.ConvertIdent(j)
			//fmt.Printf("created %d %#v\n", i,*x.Name)
			v2[i]=x
		}
		//fmt.Printf("created %#v\n", f)
		return v2
	}
	panic("todo"); return nil;
}

func (t* AddressTable) StrmapIdent(id string, f * ast.Ident) (*ast.Ident){
	f2 := t.ConvertIdent(f);
	t.Idents[id] =f;
	t.Check(id,f2,f);
	return f
}

func (t* AddressTable) ConvertIdent(f *ast.Ident) (*Ident){
	if f == nil { return nil }
	//fmt.Printf("ConvertIdent %#v\n", f)
	return &Ident{
		NamePos:t.ConvertFoo3(f.NamePos),
		Name:&f.Name,
		Obj:t.ConvertFoo2(f.Obj),
	}
}

func (t* AddressTable) ConvertPointerIdent3(f ast.Foo3) (*Foo3){
	switch v:= f.(type) {
		case * ast.Ident:
		fmt.Printf("ConvertPointerIdent3 %#v\n", v)
	default:
		panic("no gittity")
	}
	panic("todo"); return nil;
}

func (t* AddressTable) ConvertPointerIdent(f ast.Foo2) (*Ident){
	//fmt.Printf("ConvertPointerIdent %#v\n", f)
	switch v:= f.(type) {
	case *ast.Ident:
		return t.ConvertIdent(v)
	default:
		fmt.Printf("type debug %T",v)
	}
	panic("todo"); return nil;
}


func (t* AddressTable) ConvertPointerIdent2(f ast.Foo2) (*Foo2){
	fmt.Printf("ConvertPointerIdent2 %#v\n", f)
	panic("todo"); return nil;
}
