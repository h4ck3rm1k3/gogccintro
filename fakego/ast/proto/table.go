package astproto
import (
	"fmt"
	"github.com/h4ck3rm1k3/gogccintro/fakego/ast"
	//"github.com/h4ck3rm1k3/gogccintro/fakego/token"
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

// func (t* AddressTable) ConvertFieldList(f *ast.FieldList) (*FieldList){
// 	return nil
// }

// func (t* AddressTable) ConvertStructType(f *ast.StructType) (*StructType){
// 	
// 	return &StructType{
// 		Struct : t.ConvertFoo3(f.Struct),
// 		Fields : t.ConvertFieldList(f.Fields),
// 		Incomplete : &f2,
// 	}
// }

func (t* AddressTable) StrmapStructType(id string, f * ast.StructType) (*ast.StructType){ t.StructTypes[id] = t.ConvertStructType(f); f.Report(); return f}


func (t* AddressTable) Converttoken_Pos(f ast.Foo3) (*Foo3){
	return nil
}

func (t* AddressTable) Converttoken_Pos2(f ast.Foo3) (*Foo2){
	return nil
}

func (t* AddressTable) Converttoken_Token(f ast.Foo3) (*Foo3){
	return nil
}

func (t* AddressTable) ConvertTypeExpr(f ast.Foo3) (*Foo2){
	return nil
}

func (t* AddressTable) ConvertPointerFieldList(f ast.Foo3) (*FieldList){
	return nil
}

func (t* AddressTable) ConvertArrayStarIdent(f ast.Foo3) ([]*Ident){
	return nil
}

func (t* AddressTable) ConvertTypebool(f ast.Foo3) (*bool){
	return nil
}

func (t* AddressTable) ConvertArrayExpr(f ast.Foo3) ([]*Expr){
	return nil
}

func (t* AddressTable) ConvertPointerBasicLit(f ast.Foo3) (*Foo2){
	return nil
}

func (t* AddressTable) ConvertTypeStmt(f ast.Foo3) (*Foo2){
	return nil
}

func (t* AddressTable) ConvertPointerBlockStmt(f ast.Foo3) (*Foo2){
	return nil
}

	
func (t* AddressTable) ConvertMapType(f *ast.MapType) (*MapType){
//o	
	return &MapType{
		Map : t.Converttoken_Pos( f.Map),
		Key : t.ConvertTypeExpr( f.Key),
		Value : t.ConvertTypeExpr( f.Value),
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertField(f *ast.Field) (*Field){
	
	return &Field{
		//Doc : t.ConvertPointerCommentGroup( f.Doc),
		Names : t.ConvertArrayStarIdent( f.Names),
		Type : t.ConvertTypeExpr( f.Type),
		Tag : t.ConvertPointerBasicLit( f.Tag),
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
		Slice3 : t.ConvertTypebool( f.Slice3),
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
		Methods : t.ConvertPointerFieldList( f.Methods),
		Incomplete : t.ConvertTypebool( f.Incomplete),
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertIfStmt(f *ast.IfStmt) (*IfStmt){
	
	return &IfStmt{
		If : t.Converttoken_Pos( f.If),
		Init : t.ConvertTypeStmt( f.Init),
		Cond : t.ConvertTypeExpr( f.Cond),
		Body : t.ConvertPointerBlockStmt( f.Body),
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
// 	return &BadDecl{
// 		From : t.Converttoken_Pos( f.From),
// 		To : t.Converttoken_Pos( f.To),
// 		Incomplete: &f2,
// 	}
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
	return nil
}

func (t* AddressTable) ConvertDeferStmt(f *ast.DeferStmt) (*DeferStmt){
	
	return &DeferStmt{
		Defer : t.Converttoken_Pos( f.Defer),
		Call : t.ConvertPointerCallExpr( f.Call),
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertKeyValueExpr(f *ast.KeyValueExpr) (*KeyValueExpr){
	
	return &KeyValueExpr{
		Key : t.ConvertTypeExpr( f.Key),
		Colon : t.Converttoken_Pos2( f.Colon),
		Value : t.ConvertTypeExpr( f.Value),
		//Incomplete: &f2,
	}
}

// func (t* AddressTable) ConvertSEND(f *ast.SEND) (*SEND){
// 	
// 	return &SEND{
// // other type &{}
// 		Incomplete: &f2,
// 	}
// }

// func (t* AddressTable) ConvertisWhitespace(f *ast.isWhitespace) (*isWhitespace){
// 	
// 	return &isWhitespace{
// // other type &{}
// 		Incomplete: &f2,
// 	}
// }

func (t* AddressTable) ConvertSelectorExpr(f *ast.SelectorExpr) (*SelectorExpr){
	
	return &SelectorExpr{
		X : t.ConvertTypeExpr( f.X),
		Sel : t.ConvertPointerIdent2( f.Sel),
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
		Body : t.ConvertPointerBlockStmt( f.Body),
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertStructType(f *ast.StructType) (*StructType){
	
	return &StructType{
		Struct : t.Converttoken_Pos( f.Struct),
		Fields : t.ConvertPointerFieldList( f.Fields),
		Incomplete : t.ConvertTypebool( f.Incomplete),
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

func (t* AddressTable) ConvertArrayStmt(args []ast.Stmt) ([]*Stmt){
	return nil
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

func (t* AddressTable) ConvertExpr(f *ast.Expr) (*Expr){
	
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
// 	return &RECV{
// // other type &{}
// 		//Incomplete: &f2,
// 	}
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
		Label : t.ConvertPointerIdent( f.Label),
		Colon : t.Converttoken_Pos2( f.Colon),
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
		Name : t.ConvertPointerIdent2( f.Name),
		Type : t.ConvertTypeExpr( f.Type),
		//Comment : t.ConvertPointerCommentGroup( f.Comment),
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertDecl(f *ast.Decl) (*Decl){
	
	return &Decl{
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertPointerFuncType(f *ast.FuncType) (*FuncType){
	return nil
}

func (t* AddressTable) ConvertPointerFuncType2(f ast.Foo2) (*Foo2){
	return nil
}

func (t* AddressTable) ConvertFuncDecl(f *ast.FuncDecl) (*FuncDecl){
	
	return &FuncDecl{
		//Doc : t.ConvertPointerCommentGroup( f.Doc),
		Recv : t.ConvertPointerFieldList2( f.Recv),
		Name : t.ConvertPointerIdent3( f.Name),
		Type : t.ConvertPointerFuncType2( f.Type),
		Body : t.ConvertPointerBlockStmt( f.Body),
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertFuncLit(f *ast.FuncLit) (*FuncLit){
	
	return &FuncLit{
		Type : t.ConvertPointerFuncType( f.Type),
		Body : t.ConvertPointerBlockStmt2( f.Body),
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertForStmt(f *ast.ForStmt) (*ForStmt){
	
	return &ForStmt{
		For : t.Converttoken_Pos( f.For),
		Init : t.ConvertTypeStmt( f.Init),
		Cond : t.ConvertTypeExpr( f.Cond),
		Post : t.ConvertTypeStmt( f.Post),
		Body : t.ConvertPointerBlockStmt( f.Body),
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertPointerIdent2(f ast.Foo2) (*Foo2){
	return nil
}

func (t* AddressTable) ConvertPointerIdent3(f ast.Foo3) (*Foo3){
	return nil
}

func (t* AddressTable) ConvertPointerIdent(f ast.Foo2) (*Ident){
	return nil
}
func (t* AddressTable) ConvertIdent(f *ast.Ident) (*Ident){
	
	return &Ident{
// other type &{}
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertSelectStmt(f *ast.SelectStmt) (*SelectStmt){
	
	return &SelectStmt{
		Select : t.Converttoken_Pos2( f.Select),
		Body : t.ConvertPointerBlockStmt( f.Body),
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
	return nil
}
	
func (t* AddressTable) ConvertFieldList(f *ast.FieldList) (*FieldList){
	
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
		//Name : t.ConvertPointerIdent( f.Name),
		//Path : t.ConvertPointerBasicLit( f.Path),
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
	return nil
}

func (t* AddressTable) ConvertMapstringToPointerFile(f map[string]*ast.File) (map[string]*File){
	return nil
}

func (t* AddressTable) ConvertArrayStarField(f []*ast.Field) ([]*Field){
	return nil
}


func (t* AddressTable) ConvertPointerBlockStmt2(f *ast.BlockStmt) (*BlockStmt){
	return nil
}

func (t* AddressTable) ConvertPointerScope(f *ast.Scope) (*Scope){
	return nil
}

func (t* AddressTable) ConvertPointerScope2(f ast.Foo2) (*Foo2){
	return nil
}

func (t* AddressTable) ConvertTypestring(f string) (*string){
	return nil
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
		Body : t.ConvertPointerBlockStmt( f.Body),
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
		Body : t.ConvertPointerBlockStmt( f.Body),
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
		Implicit : t.ConvertTypebool( f.Implicit),
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
func (t* AddressTable) ConvertArrayDecl(f []ast.Decl)([]*Decl){ return nil}

func (t* AddressTable) ConvertArrayStarImportSpec (f []*ast.ImportSpec)([]*ImportSpec){ return nil}
func (t* AddressTable) ConvertArrayStarCommentGroup(f []*ast.CommentGroup)([]*CommentGroup){ return nil}
	
func (t* AddressTable) ConvertFile(f *ast.File) (*File){
	
	return &File{
		//Doc : t.ConvertPointerCommentGroup( f.Doc),
		Package : t.Converttoken_Pos( f.Package),
		Name : t.ConvertPointerIdent3( f.Name),
		Decls : t.ConvertArrayDecl( f.Decls),
		Scope : t.ConvertPointerScope2( f.Scope),
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
		Label : t.ConvertPointerIdent( f.Label),
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertCommentGroup(f *ast.CommentGroup) (*CommentGroup){
	
	return &CommentGroup{
		//List : t.ConvertArrayStarComment( f.List),
		//Incomplete: &f2,
	}
}

func (t* AddressTable) ConvertArraySpec(f []ast.Spec) ([]*Spec){return nil}

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
		Call : t.ConvertPointerCallExpr( f.Call),
		//Incomplete: &f2,
	}
}

// func (t* AddressTable) ConvertstripTrailingWhitespace(f *ast.stripTrailingWhitespace) (*stripTrailingWhitespace){
// 	
// 	return &stripTrailingWhitespace{
// // other type &{}
// 		//Incomplete: &f2,
// 	}
// }

func (t* AddressTable) ConvertEllipsis(f *ast.Ellipsis) (*Ellipsis){
	
	return &Ellipsis{
		Ellipsis : t.Converttoken_Pos2( f.Ellipsis),
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
