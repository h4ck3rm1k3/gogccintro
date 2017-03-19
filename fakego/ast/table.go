package ast

type Table struct {
	Scopes map[string] *Scope
	Objects map[string] *Object
	ArrayTypes map[string]*ArrayType
	AssignStmts map[string]*AssignStmt
	BadDecls map[string]*BadDecl
	BadExprs map[string]*BadExpr
	BadStmts map[string]*BadStmt
	BasicLits map[string]*BasicLit
	BinaryExprs map[string]*BinaryExpr
	BlockStmts map[string]*BlockStmt
	BranchStmts map[string]*BranchStmt
	CallExprs map[string]*CallExpr
	CaseClauses map[string]*CaseClause
	ChanDirs map[string]*ChanDir
	ChanTypes map[string]*ChanType
	CommClauses map[string]*CommClause
	CommentGroups map[string]*CommentGroup
	Comments map[string]*Comment
	CompositeLits map[string]*CompositeLit
	DeclStmts map[string]*DeclStmt

	DeferStmts map[string]*DeferStmt
	Ellipsiss map[string]*Ellipsis
	EmptyStmts map[string]*EmptyStmt
	ExprStmts map[string]*ExprStmt

	FieldLists map[string]*FieldList
	Fields map[string]*Field
	Files map[string]*File
	ForStmts map[string]*ForStmt
	FuncDecls map[string]*FuncDecl
	FuncLits map[string]*FuncLit
	FuncTypes map[string]*FuncType
	GenDecls map[string]*GenDecl
	GoStmts map[string]*GoStmt
	Idents map[string]*Ident
	IfStmts map[string]*IfStmt
	ImportSpecs map[string]*ImportSpec
	IncDecStmts map[string]*IncDecStmt
	IndexExprs map[string]*IndexExpr
	InterfaceTypes map[string]*InterfaceType
	IsExporteds map[string]*IsExported
	KeyValueExprs map[string]*KeyValueExpr
	LabeledStmts map[string]*LabeledStmt
	MapTypes map[string]*MapType
	NewIdents map[string]*NewIdent
	Nodes map[string]*Node
	Packages map[string]*Package
	ParenExprs map[string]*ParenExpr
	RangeStmts map[string]*RangeStmt
	ReturnStmts map[string]*ReturnStmt
	SelectStmts map[string]*SelectStmt
	SelectorExprs map[string]*SelectorExpr
	SendStmts map[string]*SendStmt
	SliceExprs map[string]*SliceExpr

	StarExprs map[string]*StarExpr

	StructTypes map[string]*StructType
	SwitchStmts map[string]*SwitchStmt
	TypeAssertExprs map[string]*TypeAssertExpr
	TypeSpecs map[string]*TypeSpec
	TypeSwitchStmts map[string]*TypeSwitchStmt
	UnaryExprs map[string]*UnaryExpr
	ValueSpecs map[string]*ValueSpec

	// interfaces
	Specs map[string]Spec
	Decls map[string]Decl
	Exprs map[string]Expr
	Stmts map[string]Stmt
}


func (t* Table) PtrmapArrayType(id string) (* ArrayType){ return t.ArrayTypes[id] }
func (t* Table) PtrmapAssignStmt(id string) (* AssignStmt){ return t.AssignStmts[id] }
func (t* Table) PtrmapBadDecl(id string) (* BadDecl){ return t.BadDecls[id] }
func (t* Table) PtrmapBadExpr(id string) (* BadExpr){ return t.BadExprs[id] }
func (t* Table) PtrmapBadStmt(id string) (* BadStmt){ return t.BadStmts[id] }
func (t* Table) PtrmapBasicLit(id string) (* BasicLit){ return t.BasicLits[id] }
func (t* Table) PtrmapBinaryExpr(id string) (* BinaryExpr){ return t.BinaryExprs[id] }
func (t* Table) PtrmapBlockStmt(id string) (* BlockStmt){ return t.BlockStmts[id] }
func (t* Table) PtrmapBranchStmt(id string) (* BranchStmt){ return t.BranchStmts[id] }
func (t* Table) PtrmapCallExpr(id string) (* CallExpr){ return t.CallExprs[id] }
func (t* Table) PtrmapCaseClause(id string) (* CaseClause){ return t.CaseClauses[id] }
func (t* Table) PtrmapChanDir(id string) (* ChanDir){ return t.ChanDirs[id] }
func (t* Table) PtrmapChanType(id string) (* ChanType){ return t.ChanTypes[id] }
func (t* Table) PtrmapCommClause(id string) (* CommClause){ return t.CommClauses[id] }
func (t* Table) PtrmapComment(id string) (* Comment){ return t.Comments[id] }
func (t* Table) PtrmapCommentGroup(id string) (* CommentGroup){ return t.CommentGroups[id] }
func (t* Table) PtrmapCompositeLit(id string) (* CompositeLit){ return t.CompositeLits[id] }

func (t* Table) PtrmapDeclStmt(id string) (* DeclStmt){ return t.DeclStmts[id] }
func (t* Table) PtrmapDeferStmt(id string) (* DeferStmt){ return t.DeferStmts[id] }
func (t* Table) PtrmapEllipsis(id string) (* Ellipsis){ return t.Ellipsiss[id] }
func (t* Table) PtrmapEmptyStmt(id string) (* EmptyStmt){ return t.EmptyStmts[id] }

func (t* Table) PtrmapExprStmt(id string) (* ExprStmt){ return t.ExprStmts[id] }
func (t* Table) PtrmapField(id string) (* Field){ return t.Fields[id] }
func (t* Table) PtrmapFieldList(id string) (* FieldList){ return t.FieldLists[id] }
func (t* Table) PtrmapFile(id string) (* File){ return t.Files[id] }
func (t* Table) PtrmapForStmt(id string) (* ForStmt){ return t.ForStmts[id] }
func (t* Table) PtrmapFuncDecl(id string) (* FuncDecl){ return t.FuncDecls[id] }
func (t* Table) PtrmapFuncLit(id string) (* FuncLit){ return t.FuncLits[id] }
func (t* Table) PtrmapFuncType(id string) (* FuncType){ return t.FuncTypes[id] }
func (t* Table) PtrmapGenDecl(id string) (* GenDecl){ return t.GenDecls[id] }
func (t* Table) PtrmapGoStmt(id string) (* GoStmt){ return t.GoStmts[id] }
func (t* Table) PtrmapIdent(id string) (* Ident){ return t.Idents[id] }
func (t* Table) PtrmapIfStmt(id string) (* IfStmt){ return t.IfStmts[id] }
func (t* Table) PtrmapImportSpec(id string) (* ImportSpec){ return t.ImportSpecs[id] }
func (t* Table) PtrmapIncDecStmt(id string) (* IncDecStmt){ return t.IncDecStmts[id] }
func (t* Table) PtrmapIndexExpr(id string) (* IndexExpr){ return t.IndexExprs[id] }
func (t* Table) PtrmapInterfaceType(id string) (* InterfaceType){ return t.InterfaceTypes[id] }
func (t* Table) PtrmapIsExported(id string) (* IsExported){ return t.IsExporteds[id] }
func (t* Table) PtrmapKeyValueExpr(id string) (* KeyValueExpr){ return t.KeyValueExprs[id] }
func (t* Table) PtrmapLabeledStmt(id string) (* LabeledStmt){ return t.LabeledStmts[id] }
func (t* Table) PtrmapMapType(id string) (* MapType){ return t.MapTypes[id] }
func (t* Table) PtrmapNewIdent(id string) (* NewIdent){ return t.NewIdents[id] }
func (t* Table) PtrmapNode(id string) (* Node){ return t.Nodes[id] }
func (t* Table) PtrmapPackage(id string) (* Package){ return t.Packages[id] }
func (t* Table) PtrmapParenExpr(id string) (* ParenExpr){ return t.ParenExprs[id] }
func (t* Table) PtrmapRangeStmt(id string) (* RangeStmt){ return t.RangeStmts[id] }
func (t* Table) PtrmapReturnStmt(id string) (* ReturnStmt){ return t.ReturnStmts[id] }
func (t* Table) PtrmapSelectStmt(id string) (* SelectStmt){ return t.SelectStmts[id] }
func (t* Table) PtrmapSelectorExpr(id string) (* SelectorExpr){ return t.SelectorExprs[id] }
func (t* Table) PtrmapSendStmt(id string) (* SendStmt){ return t.SendStmts[id] }
func (t* Table) PtrmapSliceExpr(id string) (* SliceExpr){ return t.SliceExprs[id] }

func (t* Table) PtrmapStarExpr(id string) (* StarExpr){ return t.StarExprs[id] }

func (t* Table) PtrmapStructType(id string) (* StructType){ return t.StructTypes[id] }
func (t* Table) PtrmapSwitchStmt(id string) (* SwitchStmt){ return t.SwitchStmts[id] }
func (t* Table) PtrmapTypeAssertExpr(id string) (* TypeAssertExpr){ return t.TypeAssertExprs[id] }
func (t* Table) PtrmapTypeSpec(id string) (* TypeSpec){ return t.TypeSpecs[id] }
func (t* Table) PtrmapTypeSwitchStmt(id string) (* TypeSwitchStmt){ return t.TypeSwitchStmts[id] }
func (t* Table) PtrmapUnaryExpr(id string) (* UnaryExpr){ return t.UnaryExprs[id] }
func (t* Table) PtrmapValueSpec(id string) (* ValueSpec){ return t.ValueSpecs[id] }
func (t* Table) StrmapArrayType(id string, f * ArrayType) (*ArrayType){ t.ArrayTypes[id] =f; f.Report(); return f}
func (t* Table) StrmapAssignStmt(id string, f * AssignStmt) (*AssignStmt){ t.AssignStmts[id] =f; f.Report(); return f}
func (t* Table) StrmapBadDecl(id string, f * BadDecl) (*BadDecl){ t.BadDecls[id] =f; f.Report(); return f}
func (t* Table) StrmapBadExpr(id string, f * BadExpr) (*BadExpr){ t.BadExprs[id] =f; f.Report(); return f}
func (t* Table) StrmapBadStmt(id string, f * BadStmt) (*BadStmt){ t.BadStmts[id] =f; f.Report(); return f}
func (t* Table) StrmapBasicLit(id string, f * BasicLit) (*BasicLit){ t.BasicLits[id] =f; f.Report(); return f}
func (t* Table) StrmapBinaryExpr(id string, f * BinaryExpr) (*BinaryExpr){ t.BinaryExprs[id] =f; f.Report(); return f}
func (t* Table) StrmapBlockStmt(id string, f * BlockStmt) (*BlockStmt){ t.BlockStmts[id] =f; f.Report(); return f}
func (t* Table) StrmapBranchStmt(id string, f * BranchStmt) (*BranchStmt){ t.BranchStmts[id] =f; f.Report(); return f}
func (t* Table) StrmapCallExpr(id string, f * CallExpr) (*CallExpr){ t.CallExprs[id] =f; f.Report(); return f}
func (t* Table) StrmapCaseClause(id string, f * CaseClause) (*CaseClause){ t.CaseClauses[id] =f; f.Report(); return f}
func (t* Table) StrmapChanDir(id string, f * ChanDir) (*ChanDir){ t.ChanDirs[id] =f; f.Report(); return f}
func (t* Table) StrmapChanType(id string, f * ChanType) (*ChanType){ t.ChanTypes[id] =f; f.Report(); return f}
func (t* Table) StrmapCommClause(id string, f * CommClause) (*CommClause){ t.CommClauses[id] =f; f.Report(); return f}
func (t* Table) StrmapComment(id string, f * Comment) (*Comment){ t.Comments[id] =f; f.Report(); return f}
func (t* Table) StrmapCommentGroup(id string, f * CommentGroup) (*CommentGroup){ t.CommentGroups[id] =f; f.Report(); return f}
func (t* Table) StrmapCompositeLit(id string, f * CompositeLit) (*CompositeLit){ t.CompositeLits[id] =f; f.Report(); return f}

func (t* Table) StrmapDeclStmt(id string, f * DeclStmt) (*DeclStmt){ t.DeclStmts[id] =f; f.Report(); return f}
func (t* Table) StrmapDeferStmt(id string, f * DeferStmt) (*DeferStmt){ t.DeferStmts[id] =f; f.Report(); return f}
func (t* Table) StrmapEllipsis(id string, f * Ellipsis) (*Ellipsis){ t.Ellipsiss[id] =f; f.Report(); return f}
func (t* Table) StrmapEmptyStmt(id string, f * EmptyStmt) (*EmptyStmt){ t.EmptyStmts[id] =f; f.Report(); return f}

func (t* Table) StrmapExprStmt(id string, f * ExprStmt) (*ExprStmt){ t.ExprStmts[id] =f; f.Report(); return f}
func (t* Table) StrmapField(id string, f * Field) (*Field){ t.Fields[id] =f; f.Report(); return f}
func (t* Table) StrmapFieldList(id string, f * FieldList) (*FieldList){ t.FieldLists[id] =f; f.Report(); return f}
func (t* Table) StrmapFile(id string, f * File) (*File){ t.Files[id] =f; f.Report(); return f}
func (t* Table) StrmapForStmt(id string, f * ForStmt) (*ForStmt){ t.ForStmts[id] =f; f.Report(); return f}
func (t* Table) StrmapFuncDecl(id string, f * FuncDecl) (*FuncDecl){ t.FuncDecls[id] =f; f.Report(); return f}
func (t* Table) StrmapFuncLit(id string, f * FuncLit) (*FuncLit){ t.FuncLits[id] =f; f.Report(); return f}
func (t* Table) StrmapFuncType(id string, f * FuncType) (*FuncType){ t.FuncTypes[id] =f; f.Report(); return f}
func (t* Table) StrmapGenDecl(id string, f * GenDecl) (*GenDecl){ t.GenDecls[id] =f; f.Report(); return f}
func (t* Table) StrmapGoStmt(id string, f * GoStmt) (*GoStmt){ t.GoStmts[id] =f; f.Report(); return f}
func (t* Table) StrmapIdent(id string, f * Ident) (*Ident){ t.Idents[id] =f; f.Report(); return f}
func (t* Table) StrmapIfStmt(id string, f * IfStmt) (*IfStmt){ t.IfStmts[id] =f; f.Report(); return f}
func (t* Table) StrmapImportSpec(id string, f * ImportSpec) (*ImportSpec){ t.ImportSpecs[id] =f; f.Report(); return f}
func (t* Table) StrmapIncDecStmt(id string, f * IncDecStmt) (*IncDecStmt){ t.IncDecStmts[id] =f; f.Report(); return f}
func (t* Table) StrmapIndexExpr(id string, f * IndexExpr) (*IndexExpr){ t.IndexExprs[id] =f; f.Report(); return f}
func (t* Table) StrmapInterfaceType(id string, f * InterfaceType) (*InterfaceType){ t.InterfaceTypes[id] =f; f.Report(); return f}
func (t* Table) StrmapIsExported(id string, f * IsExported) (*IsExported){ t.IsExporteds[id] =f; f.Report(); return f}
func (t* Table) StrmapKeyValueExpr(id string, f * KeyValueExpr) (*KeyValueExpr){ t.KeyValueExprs[id] =f; f.Report(); return f}
func (t* Table) StrmapLabeledStmt(id string, f * LabeledStmt) (*LabeledStmt){ t.LabeledStmts[id] =f; f.Report(); return f}
func (t* Table) StrmapMapType(id string, f * MapType) (*MapType){ t.MapTypes[id] =f; f.Report(); return f}
func (t* Table) StrmapNewIdent(id string, f * NewIdent) (*NewIdent){ t.NewIdents[id] =f; f.Report(); return f}
func (t* Table) StrmapNode(id string, f * Node) (*Node){ t.Nodes[id] =f; f.Report(); return f}
func (t* Table) StrmapPackage(id string, f * Package) (*Package){ t.Packages[id] =f; f.Report(); return f}
func (t* Table) StrmapParenExpr(id string, f * ParenExpr) (*ParenExpr){ t.ParenExprs[id] =f; f.Report(); return f}
func (t* Table) StrmapRangeStmt(id string, f * RangeStmt) (*RangeStmt){ t.RangeStmts[id] =f; f.Report(); return f}
func (t* Table) StrmapReturnStmt(id string, f * ReturnStmt) (*ReturnStmt){ t.ReturnStmts[id] =f; f.Report(); return f}
func (t* Table) StrmapSelectStmt(id string, f * SelectStmt) (*SelectStmt){ t.SelectStmts[id] =f; f.Report(); return f}
func (t* Table) StrmapSelectorExpr(id string, f * SelectorExpr) (*SelectorExpr){ t.SelectorExprs[id] =f; f.Report(); return f}
func (t* Table) StrmapSendStmt(id string, f * SendStmt) (*SendStmt){ t.SendStmts[id] =f; f.Report(); return f}
func (t* Table) StrmapSliceExpr(id string, f * SliceExpr) (*SliceExpr){ t.SliceExprs[id] =f; f.Report(); return f}

func (t* Table) StrmapStarExpr(id string, f * StarExpr) (*StarExpr){ t.StarExprs[id] =f; f.Report(); return f}

func (t* Table) StrmapStructType(id string, f * StructType) (*StructType){ t.StructTypes[id] =f; f.Report(); return f}
func (t* Table) StrmapSwitchStmt(id string, f * SwitchStmt) (*SwitchStmt){ t.SwitchStmts[id] =f; f.Report(); return f}
func (t* Table) StrmapTypeAssertExpr(id string, f * TypeAssertExpr) (*TypeAssertExpr){ t.TypeAssertExprs[id] =f; f.Report(); return f}
func (t* Table) StrmapTypeSpec(id string, f * TypeSpec) (*TypeSpec){ t.TypeSpecs[id] =f; f.Report(); return f}
func (t* Table) StrmapTypeSwitchStmt(id string, f * TypeSwitchStmt) (*TypeSwitchStmt){ t.TypeSwitchStmts[id] =f; f.Report(); return f}
func (t* Table) StrmapUnaryExpr(id string, f * UnaryExpr) (*UnaryExpr){ t.UnaryExprs[id] =f; f.Report(); return f}
func (t* Table) StrmapValueSpec(id string, f * ValueSpec) (*ValueSpec){ t.ValueSpecs[id] =f; f.Report(); return f}

func (t* Table) StrmapDecl(id string, f Decl) (Decl){ t.Decls[id] =f; f.Report(); return f}
func (t* Table) StrmapExpr(id string, f Expr) (Expr){ t.Exprs[id] =f; f.Report(); return f}
func (t* Table) StrmapSpec(id string, f Spec) (Spec){ t.Specs[id] =f; f.Report(); return f}
func (t* Table) StrmapStmt(id string, f Stmt) (Stmt){ t.Stmts[id] =f; f.Report(); return f}

func (t* Table) PtrmapDecl(id string) (Decl){ return t.Decls[id] }
func (t* Table) PtrmapExpr(id string) (Expr){ return t.Exprs[id] }
func (t* Table) PtrmapStmt(id string) (Stmt){ return t.Stmts[id] }
func (t* Table) PtrmapSpec(id string) (Spec){ return t.Specs[id] }
func (t* Table) PtrmapObject(id string) (*Object){ return t.Objects[id] }


func (t* Table) StrmapObject(id string, f * Object) (*Object){ t.Objects[id] =f; f.Report(); return f}
func (t* Table) StrmapScope(id string, f * Scope) (*Scope){ t.Scopes[id] =f; f.Report(); return f}

func CreateTable() (*Table) {
	return &Table{
		Scopes:make( map[string] *Scope),
		Objects:make( map[string] *Object),
		ArrayTypes : make(map[string]*ArrayType),
		AssignStmts : make(map[string]*AssignStmt),
		BadDecls : make(map[string]*BadDecl),
		BadExprs : make(map[string]*BadExpr),
		BadStmts : make(map[string]*BadStmt),
		BasicLits : make(map[string]*BasicLit),
BinaryExprs : make(map[string]*BinaryExpr),
BlockStmts : make(map[string]*BlockStmt),
BranchStmts : make(map[string]*BranchStmt),
CallExprs : make(map[string]*CallExpr),
     CaseClauses : make(map[string]*CaseClause),
     ChanDirs : make(map[string]*ChanDir),
     ChanTypes : make(map[string]*ChanType),
     CommClauses : make(map[string]*CommClause),
     CommentGroups : make(map[string]*CommentGroup),
     Comments : make(map[string]*Comment),
     CompositeLits : make(map[string]*CompositeLit),
     DeclStmts : make(map[string]*DeclStmt),

     DeferStmts : make(map[string]*DeferStmt),
     Ellipsiss : make(map[string]*Ellipsis),
     EmptyStmts : make(map[string]*EmptyStmt),
     ExprStmts : make(map[string]*ExprStmt),

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
     IsExporteds : make(map[string]*IsExported),
     KeyValueExprs : make(map[string]*KeyValueExpr),
     LabeledStmts : make(map[string]*LabeledStmt),
     MapTypes : make(map[string]*MapType),
     NewIdents : make(map[string]*NewIdent),
     Nodes : make(map[string]*Node),
     Packages : make(map[string]*Package),
     ParenExprs : make(map[string]*ParenExpr),
     RangeStmts : make(map[string]*RangeStmt),
     ReturnStmts : make(map[string]*ReturnStmt),
    SelectStmts : make(map[string]*SelectStmt),
    SelectorExprs : make(map[string]*SelectorExpr),
    SendStmts : make(map[string]*SendStmt),
    SliceExprs : make(map[string]*SliceExpr),

    StarExprs : make(map[string]*StarExpr),

    StructTypes : make(map[string]*StructType),
    SwitchStmts : make(map[string]*SwitchStmt),
    TypeAssertExprs : make(map[string]*TypeAssertExpr),
    TypeSpecs : make(map[string]*TypeSpec),
    TypeSwitchStmts : make(map[string]*TypeSwitchStmt),
    UnaryExprs : make(map[string]*UnaryExpr),
    ValueSpecs : make(map[string]*ValueSpec),

		// interfaces
		Stmts : make(map[string]Stmt),
		Exprs : make(map[string]Expr),
		Specs : make(map[string]Spec),
		Decls : make(map[string]Decl),
	}
}
