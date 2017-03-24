package astproto
import (
	"fmt"
	"github.com/h4ck3rm1k3/gogccintro/fakego/ast"
)
// type Table struct {
// 	Scopes map[string] *Scope
// 	Objects map[string] *Object
// 	ArrayTypes map[string]*ArrayType
// 	AssignStmts map[string]*AssignStmt
// 	BadDecls map[string]*BadDecl
// 	BadExprs map[string]*BadExpr
// 	BadStmts map[string]*BadStmt
// 	BasicLits map[string]*BasicLit
// 	BinaryExprs map[string]*BinaryExpr
// 	BlockStmts map[string]*BlockStmt
// 	BranchStmts map[string]*BranchStmt
// 	CallExprs map[string]*CallExpr
// 	CaseClauses map[string]*CaseClause
// 	ChanDirs map[string]*ChanDir
// 	ChanTypes map[string]*ChanType
// 	CommClauses map[string]*CommClause
// 	CommentGroups map[string]*CommentGroup
// 	Comments map[string]*Comment
// 	CompositeLits map[string]*CompositeLit
// 	DeclStmts map[string]*DeclStmt

// 	DeferStmts map[string]*DeferStmt
// 	Ellipsiss map[string]*Ellipsis
// 	EmptyStmts map[string]*EmptyStmt
// 	ExprStmts map[string]*ExprStmt

// 	FieldLists map[string]*FieldList
// 	Fields map[string]*Field
// 	Files map[string]*File
// 	ForStmts map[string]*ForStmt
// 	FuncDecls map[string]*FuncDecl
// 	FuncLits map[string]*FuncLit
// 	FuncTypes map[string]*FuncType
// 	GenDecls map[string]*GenDecl
// 	GoStmts map[string]*GoStmt
// 	Idents map[string]*Ident
// 	IfStmts map[string]*IfStmt
// 	ImportSpecs map[string]*ImportSpec
// 	IncDecStmts map[string]*IncDecStmt
// 	IndexExprs map[string]*IndexExpr
// 	InterfaceTypes map[string]*InterfaceType
// 	IsExporteds map[string]*IsExported
// 	KeyValueExprs map[string]*KeyValueExpr
// 	LabeledStmts map[string]*LabeledStmt
// 	MapTypes map[string]*MapType
// 	NewIdents map[string]*NewIdent
// 	Nodes map[string]*Node
// 	Packages map[string]*Package
// 	ParenExprs map[string]*ParenExpr
// 	RangeStmts map[string]*RangeStmt
// 	ReturnStmts map[string]*ReturnStmt
// 	SelectStmts map[string]*SelectStmt
// 	SelectorExprs map[string]*SelectorExpr
// 	SendStmts map[string]*SendStmt
// 	SliceExprs map[string]*SliceExpr

// 	StarExprs map[string]*StarExpr

// 	StructTypes map[string]*StructType
// 	SwitchStmts map[string]*SwitchStmt
// 	TypeAssertExprs map[string]*TypeAssertExpr
// 	TypeSpecs map[string]*TypeSpec
// 	TypeSwitchStmts map[string]*TypeSwitchStmt
// 	UnaryExprs map[string]*UnaryExpr
// 	ValueSpecs map[string]*ValueSpec

// 	// interfaces
// 	Specs map[string]Spec
// 	Decls map[string]Decl
// 	Exprs map[string]Expr
// 	Stmts map[string]Stmt
// }

//func (t* AddressTable) StrmapArrayType(id string, f * ArrayType) (*ArrayType){ t.ArrayTypes[id] =f; f.Report(); return f}
func (t* AddressTable) StrmapAssignStmt(id string, f * AssignStmt) (*AssignStmt){ t.AssignStmts[id] =f; f.Report(); return f}
//func (t* AddressTable) StrmapBadDecl(id string, f * BadDecl) (*BadDecl){ t.BadDecls[id] =f; f.Report(); return f}
func (t* AddressTable) StrmapBadExpr(id string, f * BadExpr) (*BadExpr){ t.BadExprs[id] =f; f.Report(); return f}
func (t* AddressTable) StrmapBadStmt(id string, f * BadStmt) (*BadStmt){ t.BadStmts[id] =f; f.Report(); return f}
func (t* AddressTable) StrmapBasicLit(id string, f * BasicLit) (*BasicLit){ t.BasicLits[id] =f; f.Report(); return f}
func (t* AddressTable) StrmapBinaryExpr(id string, f * BinaryExpr) (*BinaryExpr){ t.BinaryExprs[id] =f; f.Report(); return f}
func (t* AddressTable) StrmapBlockStmt(id string, f * BlockStmt) (*BlockStmt){ t.BlockStmts[id] =f; f.Report(); return f}
func (t* AddressTable) StrmapBranchStmt(id string, f * BranchStmt) (*BranchStmt){ t.BranchStmts[id] =f; f.Report(); return f}
func (t* AddressTable) StrmapCallExpr(id string, f * CallExpr) (*CallExpr){ t.CallExprs[id] =f; f.Report(); return f}
func (t* AddressTable) StrmapCaseClause(id string, f * CaseClause) (*CaseClause){ t.CaseClauses[id] =f; f.Report(); return f}
//func (t* AddressTable) StrmapChanDir(id string, f * ChanDir) (*ChanDir){ t.ChanDirs[id] =f; f.Report(); return f}
//func (t* AddressTable) StrmapChanType(id string, f * ChanType) (*ChanType){ t.ChanTypes[id] =f; f.Report(); return f}
func (t* AddressTable) StrmapCommClause(id string, f * CommClause) (*CommClause){ t.CommClauses[id] =f; f.Report(); return f}
//func (t* AddressTable) StrmapComment(id string, f * Comment) (*Comment){ t.Comments[id] =f; f.Report(); return f}
func (t* AddressTable) StrmapCommentGroup(id string, f * CommentGroup) (*CommentGroup){ t.CommentGroups[id] =f; f.Report(); return f}
func (t* AddressTable) StrmapCompositeLit(id string, f * CompositeLit) (*CompositeLit){ t.CompositeLits[id] =f; f.Report(); return f}

//func (t* AddressTable) StrmapDeclStmt(id string, f * DeclStmt) (*DeclStmt){ t.DeclStmts[id] =f; f.Report(); return f}
func (t* AddressTable) StrmapDeferStmt(id string, f * DeferStmt) (*DeferStmt){ t.DeferStmts[id] =f; f.Report(); return f}
func (t* AddressTable) StrmapEllipsis(id string, f * Ellipsis) (*Ellipsis){
	/// 
	t.Ellipsiss[id] =f;
	f.Report();
	return f}
func (t* AddressTable) StrmapEmptyStmt(id string, f * EmptyStmt) (*EmptyStmt){ t.EmptyStmts[id] =f; f.Report(); return f}

//func (t* AddressTable) StrmapExprStmt(id string, f * ExprStmt) (*ExprStmt){ t.ExprStmts[id] =f; f.Report(); return f}
func (t* AddressTable) StrmapField(id string, f * Field) (*Field){ t.Fields[id] =f; f.Report(); return f}
func (t* AddressTable) StrmapFieldList(id string, f * FieldList) (*FieldList){ t.FieldLists[id] =f; f.Report(); return f}
func (t* AddressTable) StrmapFile(id string, f * File) (*File){ t.Files[id] =f; f.Report(); return f}
func (t* AddressTable) StrmapForStmt(id string, f * ForStmt) (*ForStmt){ t.ForStmts[id] =f; f.Report(); return f}

func (t* AddressTable) StrmapFuncLit(id string, f * FuncLit) (*FuncLit){ t.FuncLits[id] =f; f.Report(); return f}
func (t* AddressTable) StrmapFuncType(id string, f * FuncType) (*FuncType){ t.FuncTypes[id] =f; f.Report(); return f}
func (t* AddressTable) StrmapGenDecl(id string, f * GenDecl) (*GenDecl){ t.GenDecls[id] =f; f.Report(); return f}
func (t* AddressTable) StrmapGoStmt(id string, f * GoStmt) (*GoStmt){ t.GoStmts[id] =f; f.Report(); return f}
func (t* AddressTable) StrmapIdent(id string, f * Ident) (*Ident){ t.Idents[id] =f; f.Report(); return f}
func (t* AddressTable) StrmapIfStmt(id string, f * IfStmt) (*IfStmt){ t.IfStmts[id] =f; f.Report(); return f}
func (t* AddressTable) StrmapImportSpec(id string, f * ImportSpec) (*ImportSpec){ t.ImportSpecs[id] =f; f.Report(); return f}
func (t* AddressTable) StrmapIncDecStmt(id string, f * IncDecStmt) (*IncDecStmt){ t.IncDecStmts[id] =f; f.Report(); return f}
func (t* AddressTable) StrmapIndexExpr(id string, f * IndexExpr) (*IndexExpr){ t.IndexExprs[id] =f; f.Report(); return f}
func (t* AddressTable) StrmapInterfaceType(id string, f * InterfaceType) (*InterfaceType){ t.InterfaceTypes[id] =f; f.Report(); return f}
//func (t* AddressTable) StrmapIsExported(id string, f * IsExported) (*IsExported){ t.IsExporteds[id] =f; f.Report(); return f}
func (t* AddressTable) StrmapKeyValueExpr(id string, f * KeyValueExpr) (*KeyValueExpr){ t.KeyValueExprs[id] =f; f.Report(); return f}
func (t* AddressTable) StrmapLabeledStmt(id string, f * LabeledStmt) (*LabeledStmt){ t.LabeledStmts[id] =f; f.Report(); return f}
func (t* AddressTable) StrmapMapType(id string, f * MapType) (*MapType){ t.MapTypes[id] =f; f.Report(); return f}
//func (t* AddressTable) StrmapNewIdent(id string, f * NewIdent) (*NewIdent){ t.NewIdents[id] =f; f.Report(); return f}
//func (t* AddressTable) StrmapNode(id string, f * Node) (*Node){ t.Nodes[id] =f; f.Report(); return f}

func (t* AddressTable) StrmapPackage(id string, f * Package) (*Package){ t.Packages[id] =f; f.Report(); return f}
func (t* AddressTable) StrmapParenExpr(id string, f * ParenExpr) (*ParenExpr){ t.ParenExprs[id] =f; f.Report(); return f}
func (t* AddressTable) StrmapRangeStmt(id string, f * RangeStmt) (*RangeStmt){ t.RangeStmts[id] =f; f.Report(); return f}
//func (t* AddressTable) StrmapReturnStmt(id string, f * ReturnStmt) (*ReturnStmt){ t.ReturnStmts[id] =f; f.Report(); return f}
func (t* AddressTable) StrmapScope(id string, f * Scope) (*Scope){ t.Scopes[id] =f; f.Report(); return f}
func (t* AddressTable) StrmapSelectStmt(id string, f * SelectStmt) (*SelectStmt){ t.SelectStmts[id] =f; f.Report(); return f}
func (t* AddressTable) StrmapSelectorExpr(id string, f * SelectorExpr) (*SelectorExpr){ t.SelectorExprs[id] =f; f.Report(); return f}
func (t* AddressTable) StrmapSendStmt(id string, f * SendStmt) (*SendStmt){ t.SendStmts[id] =f; f.Report(); return f}
func (t* AddressTable) StrmapSliceExpr(id string, f * SliceExpr) (*SliceExpr){ t.SliceExprs[id] =f; f.Report(); return f}

//func (t* AddressTable) StrmapStarExpr(id string, f * StarExpr) (*StarExpr){ t.StarExprs[id] =f; f.Report(); return f}


func (t* AddressTable) StrmapSwitchStmt(id string, f * SwitchStmt) (*SwitchStmt){ t.SwitchStmts[id] =f; f.Report(); return f}
func (t* AddressTable) StrmapTypeAssertExpr(id string, f * TypeAssertExpr) (*TypeAssertExpr){ t.TypeAssertExprs[id] =f; f.Report(); return f}

func (t* AddressTable) StrmapTypeSwitchStmt(id string, f * TypeSwitchStmt) (*TypeSwitchStmt){ t.TypeSwitchStmts[id] =f; f.Report(); return f}
func (t* AddressTable) StrmapUnaryExpr(id string, f * UnaryExpr) (*UnaryExpr){ t.UnaryExprs[id] =f; f.Report(); return f}



//func (t* AddressTable) PtrmapArrayType(id string) (* ArrayType){ return t.ArrayTypes[id] }
func (t* AddressTable) PtrmapAssignStmt(id string) (* AssignStmt){ return t.AssignStmts[id] }
//func (t* AddressTable) PtrmapBadDecl(id string) (* BadDecl){ return t.BadDecls[id] }
func (t* AddressTable) PtrmapBadExpr(id string) (* BadExpr){ return t.BadExprs[id] }
func (t* AddressTable) PtrmapBadStmt(id string) (* BadStmt){ return t.BadStmts[id] }
func (t* AddressTable) PtrmapBasicLit(id string) (* BasicLit){ return t.BasicLits[id] }
func (t* AddressTable) PtrmapBinaryExpr(id string) (* BinaryExpr){ return t.BinaryExprs[id] }
func (t* AddressTable) PtrmapBlockStmt(id string) (* BlockStmt){ return t.BlockStmts[id] }
func (t* AddressTable) PtrmapBranchStmt(id string) (* BranchStmt){ return t.BranchStmts[id] }
func (t* AddressTable) PtrmapCallExpr(id string) (* CallExpr){ return t.CallExprs[id] }
func (t* AddressTable) PtrmapCaseClause(id string) (* CaseClause){ return t.CaseClauses[id] }
//func (t* AddressTable) PtrmapChanDir(id string) (* ChanDir){ return t.ChanDirs[id] }
//func (t* AddressTable) PtrmapChanType(id string) (* ChanType){ return t.ChanTypes[id] }
func (t* AddressTable) PtrmapCommClause(id string) (* CommClause){ return t.CommClauses[id] }
//func (t* AddressTable) PtrmapComment(id string) (* Comment){ return t.Comments[id] }
func (t* AddressTable) PtrmapCommentGroup(id string) (* CommentGroup){ return t.CommentGroups[id] }
func (t* AddressTable) PtrmapCompositeLit(id string) (* CompositeLit){ return t.CompositeLits[id] }
//func (t* AddressTable) PtrmapDeclStmt(id string) (* DeclStmt){ return t.DeclStmts[id] }
func (t* AddressTable) PtrmapDeferStmt(id string) (* DeferStmt){ return t.DeferStmts[id] }
func (t* AddressTable) PtrmapEllipsis(id string) (* Ellipsis){ return t.Ellipsiss[id] }
func (t* AddressTable) PtrmapEmptyStmt(id string) (* EmptyStmt){ return t.EmptyStmts[id] }
//func (t* AddressTable) PtrmapExprStmt(id string) (* ExprStmt){ return t.ExprStmts[id] }
func (t* AddressTable) PtrmapFile(id string) (* File){ return t.Files[id] }
func (t* AddressTable) PtrmapForStmt(id string) (* ForStmt){ return t.ForStmts[id] }
func (t* AddressTable) PtrmapFuncDecl(id string) (* FuncDecl){ return t.FuncDecls[id] }
func (t* AddressTable) PtrmapFuncLit(id string) (* FuncLit){ return t.FuncLits[id] }
func (t* AddressTable) PtrmapFuncType(id string) (* FuncType){ return t.FuncTypes[id] }
func (t* AddressTable) PtrmapGenDecl(id string) (* GenDecl){ return t.GenDecls[id] }
func (t* AddressTable) PtrmapGoStmt(id string) (* GoStmt){ return t.GoStmts[id] }
func (t* AddressTable) PtrmapIdent(id string) (* Ident){ return t.Idents[id] }
func (t* AddressTable) PtrmapIfStmt(id string) (* IfStmt){ return t.IfStmts[id] }
func (t* AddressTable) PtrmapImportSpec(id string) (* ImportSpec){ return t.ImportSpecs[id] }
func (t* AddressTable) PtrmapIncDecStmt(id string) (* IncDecStmt){ return t.IncDecStmts[id] }
func (t* AddressTable) PtrmapIndexExpr(id string) (* IndexExpr){ return t.IndexExprs[id] }
func (t* AddressTable) PtrmapInterfaceType(id string) (* InterfaceType){ return t.InterfaceTypes[id] }
//func (t* AddressTable) PtrmapIsExported(id string) (* IsExported){ return t.IsExporteds[id] }
func (t* AddressTable) PtrmapKeyValueExpr(id string) (* KeyValueExpr){ return t.KeyValueExprs[id] }
func (t* AddressTable) PtrmapLabeledStmt(id string) (* LabeledStmt){ return t.LabeledStmts[id] }
func (t* AddressTable) PtrmapMapType(id string) (* MapType){ return t.MapTypes[id] }
//func (t* AddressTable) PtrmapNewIdent(id string) (* NewIdent){ return t.NewIdents[id] }
//func (t* AddressTable) PtrmapNode(id string) (* Node){ return t.Nodes[id] }
func (t* AddressTable) PtrmapPackage(id string) (* Package){ return t.Packages[id] }
func (t* AddressTable) PtrmapParenExpr(id string) (* ParenExpr){ return t.ParenExprs[id] }
func (t* AddressTable) PtrmapRangeStmt(id string) (* RangeStmt){ return t.RangeStmts[id] }
//func (t* AddressTable) PtrmapReturnStmt(id string) (* ReturnStmt){ return t.ReturnStmts[id] }
func (t* AddressTable) PtrmapSelectStmt(id string) (* SelectStmt){ return t.SelectStmts[id] }
func (t* AddressTable) PtrmapSelectorExpr(id string) (* SelectorExpr){ return t.SelectorExprs[id] }
func (t* AddressTable) PtrmapSendStmt(id string) (* SendStmt){ return t.SendStmts[id] }
func (t* AddressTable) PtrmapSliceExpr(id string) (* SliceExpr){ return t.SliceExprs[id] }
//func (t* AddressTable) PtrmapStarExpr(id string) (* StarExpr){ return t.StarExprs[id] }
func (t* AddressTable) PtrmapStructType(id string) (* StructType){ return t.StructTypes[id] }
func (t* AddressTable) PtrmapSwitchStmt(id string) (* SwitchStmt){ return t.SwitchStmts[id] }
func (t* AddressTable) PtrmapTypeAssertExpr(id string) (* TypeAssertExpr){ return t.TypeAssertExprs[id] }


func (t* AddressTable) PtrmapTypeSwitchStmt(id string) (* TypeSwitchStmt){ return t.TypeSwitchStmts[id] }
func (t* AddressTable) PtrmapUnaryExpr(id string) (* UnaryExpr){ return t.UnaryExprs[id] }
func (t* AddressTable) PtrmapValueSpec(id string) (* ValueSpec){ return t.ValueSpecs[id] }


func CreateTable() (*AddressTable) {
	return &AddressTable{
		Scopes:make( map[string] *Scope),
		Objects:make( map[string] *Object),
//		ArrayTypes : make(map[string]*ArrayType),
		AssignStmts : make(map[string]*AssignStmt),
//		BadDecls : make(map[string]*BadDecl),
		BadExprs : make(map[string]*BadExpr),
		BadStmts : make(map[string]*BadStmt),
		BasicLits : make(map[string]*BasicLit),
BinaryExprs : make(map[string]*BinaryExpr),
BlockStmts : make(map[string]*BlockStmt),
BranchStmts : make(map[string]*BranchStmt),
CallExprs : make(map[string]*CallExpr),
     CaseClauses : make(map[string]*CaseClause),
//     ChanDirs : make(map[string]*ChanDir),
//     ChanTypes : make(map[string]*ChanType),
     CommClauses : make(map[string]*CommClause),
     CommentGroups : make(map[string]*CommentGroup),
//     Comments : make(map[string]*Comment),
     CompositeLits : make(map[string]*CompositeLit),
//     DeclStmts : make(map[string]*DeclStmt),

     DeferStmts : make(map[string]*DeferStmt),
     Ellipsiss : make(map[string]*Ellipsis),
     EmptyStmts : make(map[string]*EmptyStmt),
//     ExprStmts : make(map[string]*ExprStmt),

     FieldLists : make(map[string]*FieldList),
     Fields : make(map[string]*Field),
     Files : make(map[string]*File),
     ForStmts : make(map[string]*ForStmt),
     FuncDecls : make(map[string]*FuncDecl),
     FuncLits : make(map[string]*FuncLit),
     FuncTypes : make(map[string]*FuncType),
     GenDecls : make(map[string]*GenDecl),
     GoStmts : make(map[string]*GoStmt),
     Idents : make(map[string]*Ident),
     IfStmts : make(map[string]*IfStmt),
     ImportSpecs : make(map[string]*ImportSpec),
     IncDecStmts : make(map[string]*IncDecStmt),
     IndexExprs : make(map[string]*IndexExpr),
     InterfaceTypes : make(map[string]*InterfaceType),
//     IsExporteds : make(map[string]*IsExported),
     KeyValueExprs : make(map[string]*KeyValueExpr),
     LabeledStmts : make(map[string]*LabeledStmt),
     MapTypes : make(map[string]*MapType),
///     NewIdents : make(map[string]*NewIdent),
//     Nodes : make(map[string]*Node),
     Packages : make(map[string]*Package),
     ParenExprs : make(map[string]*ParenExpr),
     RangeStmts : make(map[string]*RangeStmt),
//     ReturnStmts : make(map[string]*ReturnStmt),
    SelectStmts : make(map[string]*SelectStmt),
    SelectorExprs : make(map[string]*SelectorExpr),
    SendStmts : make(map[string]*SendStmt),
    SliceExprs : make(map[string]*SliceExpr),

//    StarExprs : make(map[string]*StarExpr),

    StructTypes : make(map[string]*StructType),
    SwitchStmts : make(map[string]*SwitchStmt),
    TypeAssertExprs : make(map[string]*TypeAssertExpr),
    TypeSpecs : make(map[string]*TypeSpec),
    TypeSwitchStmts : make(map[string]*TypeSwitchStmt),
    UnaryExprs : make(map[string]*UnaryExpr),
    ValueSpecs : make(map[string]*ValueSpec),

		// interfaces
		Stmts : make(map[string]*Stmt),
		Exprs : make(map[string]*Expr),
		Specs : make(map[string]*Spec),
		Decls : make(map[string]*Decl),
	}
}

func (t* AddressTable) StrmapTypeSpec(id string, f * TypeSpec) (*TypeSpec){
	t.TypeSpecs[id] =f;
	f.Report();
	return f
}


func (t* Deferred) Report() (string){
	// needs access to table to resolve set
	return "TODO"
	//if t.Data != nil {
	//	return t.Data.Report()
	//} else {
	// i := t.Set
	// switch v:= i.(type) {
	// 	case *map[string]*Field :
	// 		if val, ok := (*v)[t.Id]; ok {
	// 			if val == nil {
	// 				fmt.Printf("did not find %s %s",v,t)
	// 			} else {
	// 				d := (*v)[t.Id]
	// 				return d.Report()					
	// 			}
	// 		}
	// 	case *map[string]*FieldList :
	// 		if val, ok := (*v)[t.Id]; ok {
	// 			if val == nil {
	// 				fmt.Printf("did not find %s %s",v,t)
	// 			} else {
	// 				d := (*v)[t.Id]
	// 				return d.Report()					
	// 			}
	// 		}
	// 	case *map[string]*TypeSpec :

	// 		if val, ok := (*v)[t.Id]; ok {
	// 			if val == nil {
	// 				fmt.Printf("did not find %s %s",v,t)
	// 			} else {
	// 				d := (*v)[t.Id]
	// 				return d.Report()					
	// 			}
	// 		}
				
						
	// 	default:
	// 		//t.Data = v[t.Id]
	// 		fmt.Printf("unknown type %s %s",v,t)
	// 		panic("unknown type")
	// 		return "unkown"
	// 	}
	// //}
	// return "huh?"
}

func (t* AddressTable) PtrmapTypeSpec(id string) (Foo2){
	if val, ok := t.TypeSpecs[id]; ok {
		if val == nil {
			t := NodeType_TYPESPEC
			t2 := NodeType_DEFERRED
			return Foo2{
				Type: &t2,
				Deferred:&Deferred{
				Id:&id,
				Type: &t,
				},
			}
		} else {
			t := NodeType_TYPESPEC
			return Foo2{
				Type: &t,
				Typespec : val,
			}
		}
		
	} else {
		//return &Deferred{ Id:id, Set: &t.TypeSpecs  }
		t := NodeType_TYPESPEC
		t2 := NodeType_DEFERRED
		return Foo2{
			Type: &t2,
			Deferred:&Deferred{
				Id:&id,
				Type: &t,
			},
		}
	}
}

func (t* AddressTable) PtrmapField(id string) (Foo2){
	t3 := NodeType_FIELD
	//return t.Fields[id]
	if val, ok := t.Fields[id]; ok {
		if val == nil {
			//return &Deferred{ Id:id, Set: &t.Fields  }

			t2 := NodeType_DEFERRED
			return Foo2{
				Type: &t2,
				Deferred:&Deferred{
				Id:&id,
				Type: &t3,
				},
			}
			
		} else {
			//return val
			return Foo2{
				Type: &t3,
				Field : val,
			}
		}
		
	} else {
		//return &Deferred{ Id:id, Set: &t.Fields  }
		t := NodeType_FIELD
		t2 := NodeType_DEFERRED
		return Foo2{
			Type: &t2,
			Deferred:&Deferred{
				Id:&id,
				Type: &t,
			},
		}

	}
}
func (t* AddressTable) PtrmapFieldList(id string) (Foo2){
	// return t.FieldLists[id]
	t3 := NodeType_FIELDLIST
	if val, ok := t.FieldLists[id]; ok {
		if val != nil {
			return Foo2{
				Type: &t3,
				Fieldlist : val,
			}			
		}		
	}

	t2 := NodeType_DEFERRED
	return Foo2{
		Type: &t2,
		Deferred:&Deferred{
			Id:&id,
			Type: &t3,
		},
	}
	
}

func (t* AddressTable) PtrmapObject(id string) (*Object){

	t2 := NodeType_OBJECT
	if val, ok := t.Objects[id]; ok {
		if val != nil {
			return val			
		}		
	} 

	return &Object{
		Deferred:&Deferred{
			Id:&id,
			Type: &t2,
		},
	}

}

func (t* AddressTable) StrmapObject(id string, f * Object) (*Object){ t.Objects[id] =f; f.Report(); return f}
func (t* AddressTable) StrmapValueSpec(id string, f * ValueSpec) (Foo2){
	//t.ValueSpecs[id] =f; f.Report(); return f
	s := t.ValueSpecs
	t2 := NodeType_VALUESPEC
	if val, ok := s[id]; ok {
		if val != nil {
			return Foo2{
				Type: &t2,
				Valuespec : val,
			}				
		}		
	}


	t3 := NodeType_DEFERRED
	return Foo2{
		Type: &t2,
		Deferred:&Deferred{
			Id:&id,
			Type: &t3,
		},
	}

}

func (t* AddressTable) StrmapFuncDecl(id string, f * FuncDecl) (Foo2){
	fmt.Printf("test")
	t2 := NodeType_FUNCDECL
	//t.FuncDecls[id] =f; f.Report(); return f
	s := t.FuncDecls
	if val, ok := s[id]; ok {
		if val != nil {
			return Foo2{
				Type: &t2,
				Funcdecl : val,
			}				
		}		
	}
	t3 := NodeType_DEFERRED
	return Foo2{
		Type: &t2,
		Deferred:&Deferred{
			Id:&id,
			Type: &t3,
		},
	}

}


func (t* AddressTable) StrmapDecl(id string, f Decl) (Decl){ t.Decls[id] = &f; f.Report(); return f}
func (t* AddressTable) StrmapExpr(id string, f Expr) (Expr){ t.Exprs[id] = &f; f.Report(); return f}
func (t* AddressTable) StrmapSpec(id string, f Spec) (Spec){ t.Specs[id] = &f; f.Report(); return f}
func (t* AddressTable) StrmapStmt(id string, f Stmt) (Stmt){ t.Stmts[id] = &f; f.Report(); return f}

func (t* AddressTable) PtrmapDecl(id string) (Decl){ return *t.Decls[id] }
func (t* AddressTable) PtrmapExpr(id string) (Expr){ return *t.Exprs[id] }
func (t* AddressTable) PtrmapStmt(id string) (Stmt){ return *t.Stmts[id] }
func (t* AddressTable) PtrmapSpec(id string) (Spec){ return *t.Specs[id] }



///func (t* AddressTable) StrmapStructType(id string, f * StructType) (*StructType){ t.StructTypes[id] =f; f.Report(); return f}
func (t* AddressTable) ConvertFoo3(f ast.Foo3) (*Foo3){
	return nil
}

func (t* AddressTable) ConvertFieldList(f *ast.FieldList) (*FieldList){
	return nil
}

func (t* AddressTable) ConvertStructType(f *ast.StructType) (*StructType){
	f2 := false
	return &StructType{
		Struct : t.ConvertFoo3(f.Struct),
		Fields : t.ConvertFieldList(f.Fields),
		Incomplete : &f2,
	}
}

func (t* AddressTable) StrmapStructType(id string, f * ast.StructType) (*ast.StructType){ t.StructTypes[id] = t.ConvertStructType(f); f.Report(); return f}
