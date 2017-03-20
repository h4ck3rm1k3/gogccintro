package ast
import (
	//"fmt"
	//"strings"
)

type Foo2 interface {
	Report() string
}

type Foo3 interface {

}

type File struct {
	Package Foo3
	Name Foo3
	Decls []Decl
	Scope Foo2
	Imports []*ImportSpec
	Unresolved []*Ident
	Comments   []*CommentGroup
}

type FuncType struct {
	Func Foo3
	Params Foo2
	Results *FieldList
}

type SelectorExpr struct {
	X Foo2
	Sel Foo2
}



type Ident struct {
	NamePos Foo3
	Name string
	Obj Foo2
}

type Object struct {
	Kind Foo3
	Name Foo3
	Decl Foo2
	Data Foo3
}

type InterfaceType struct {
	Interface Foo3
	Methods *FieldList
	Incomplete bool
}

type StructType struct {
	Struct Foo3
	Fields *FieldList
	Incomplete bool
}

type Field struct {
	Names []*Ident
	Type Foo2
	Tag Foo2
}


type StarExpr struct {
	Star Foo3
	X Foo2
}


type FieldList struct {
	Opening Foo3
	List [] *Field
	Closing Foo3

}


type Scope struct {

	Objects map[string]*Object
}


type TypeSpec struct {
	Name Foo2
	Type Foo2
}

type BasicLit struct {
	ValuePos Foo3
	Kind Foo3
	Value Foo3
}



type ImportSpec struct {
	Path Foo2
	EndPos Foo3
}


type GenDecl struct {
	TokPos Foo3
	Tok Foo3
	Lparen Foo3
	Specs []Spec
	Rparen Foo3
}

type FuncDecl struct {
	Recv Foo2
	Name Foo3
	Type Foo2
	Body Foo2
}
type BlockStmt struct {
	Lbrace Foo3
	List []Stmt
	Rbrace Foo3
}

type(
	BadExpr struct {
		From, To Foo2 // position range of bad expression
	}

	// An Ident node represents an identifier.
	

	// An Ellipsis node stands for the "..." type in a
	// parameter list or the "..." length in an array type.
	//
	Ellipsis struct {
		Ellipsis Foo2 // position of "..."
		Elt      Foo2      // ellipsis element type (parameter lists only); or nil
	}

	// A BasicLit node represents a literal of basic type.
	

	// A FuncLit node represents a function literal.
	FuncLit struct {
		Type *FuncType  // function type
		Body *BlockStmt // function body
	}

	// A CompositeLit node represents a composite literal.
	CompositeLit struct {
		Type   Foo2      // literal type; or nil
		Lbrace Foo3
		Elts   []Expr    // []list of composite elements; or nil
		Rbrace Foo3
	}

	// A ParenFoo2 node represents a parenthesized expression.
	ParenExpr struct {
		Lparen Foo3
		X      Foo2      // parenthesized expression
		Rparen Foo3
	}

	// A SelectorExpr node represents an expression followed by a selector.
	
	// An IndexExpr node represents an expression followed by an index.
	IndexExpr struct {
		X      Foo2      // expression
		Lbrack Foo3
		Index  Foo2      // index expression
		Rbrack Foo3
	}

	// An SliceExpr node represents an expression followed by slice indices.
	SliceExpr struct {
		X      Foo2      // expression
		Lbrack Foo3
		Low    Foo2      // begin of slice range; or nil
		High   Foo2      // end of slice range; or nil
		Max    Foo2      // maximum capacity of slice; or nil
		Slice3 bool      // true if 3-index slice (2 colons present)
		Rbrack Foo3
	}

	// A TypeAssertExpr node represents an expression followed by a
	// type assertion.
	//
	TypeAssertExpr struct {
		X      Foo2      // expression
		Lparen Foo3
		Type   Foo2      // asserted type; nil means type switch X.(type)
		Rparen Foo3
	}

	// A CallExpr node represents an expression followed by an argument list.
	CallExpr struct {
		Fun      Foo2      // function expression
		Lparen   Foo3
		Args     []Expr
		Ellipsis Foo3
		Rparen   Foo3
	}

	// A StarExpr node represents an expression of the form "*" Expression.
	// Semantically it could be a unary "*" expression, or a pointer type.
	//
	

	// A UnaryExpr node represents a unary expression.
	// Unary "*" expressions are represented via StarExpr nodes.
	//
	UnaryExpr struct {
		OpPos Foo3
		Op    Foo3 // operator
		X     Foo2        // operand
	}

	// A BinaryExpr node represents a binary expression.
	BinaryExpr struct {
		X     Foo2        // left operand
		OpPos Foo3   // position of Op
		Op    Foo3  // operator
		Y     Foo2        // right operand
	}

	// A KeyValueExpr node represents (key : value) pairs
	// in composite literals.
	//
	KeyValueExpr struct {
		Key   Foo2
		Colon Foo2 // position of ":"
		Value Foo2
	}

)


type (
	// A BadStmt node is a placeholder for statements containing
	// syntax errors for which no correct statement nodes can be
	// created.
	//
	BadStmt struct {
		From, To Foo2 // position range of bad statement
	}

	// A DeclStmt node represents a declaration in a statement list.
	DeclStmt struct {
		Decl Decl // *GenDecl with CONST, TYPE, or VAR token
	}

	// An EmptyStmt node represents an empty statement.
	// The "position" of the empty statement is the position
	// of the immediately following (explicit or implicit) semicolon.
	//
	EmptyStmt struct {
		Semicolon Foo2 // position of following ";"
		Implicit  bool      // if set, ";" was omitted in the source
	}

	// A LabeledStmt node represents a labeled statement.
	LabeledStmt struct {
		Label *Ident
		Colon Foo2 // position of ":"
		Stmt  Foo2
	}

	// An ExprStmt node represents a (stand-alone) expression
	// in a statement list.
	//
	ExprStmt struct {
		X Foo2 // expression
	}

	// A SendStmt node represents a send statement.
	SendStmt struct {
		Chan  Foo2
		Arrow Foo2 // position of "<-"
		Value Foo2
	}

	// An IncDecStmt node represents an increment or decrement statement.
	IncDecStmt struct {
		X      Foo2
		TokPos Foo3   // position of Tok
		Tok     Foo3
 // INC or DEC
	}

	// An AssignStmt node represents an assignment or
	// a short variable declaration.
	//
	AssignStmt struct {
		Lhs    []Expr
		TokPos Foo3   // position of Tok
		Tok     Foo3
 // assignment token, DEFINE
		Rhs    []Expr
	}

	// A GoStmt node represents a go statement.
	GoStmt struct {
		Go   Foo2 // position of "go" keyword
		Call *CallExpr
	}

	// A DeferStmt node represents a defer statement.
	DeferStmt struct {
		Defer Foo3
		Call  *CallExpr
	}

	// A ReturnStmt node represents a return statement.
	ReturnStmt struct {
		Return  Foo3
		Results []Expr    // result expressions; or nil
	}

	// A BranchStmt node represents a break, continue, goto,
	// or fallthrough statement.
	//
	BranchStmt struct {
		TokPos Foo3
		Tok     Foo3
 // keyword token (BREAK, CONTINUE, GOTO, FALLTHROUGH)
		Label  *Ident      // label name; or nil
	}

	// A BlockStmt node represents a braced statement list.
	
	// An IfStmt node represents an if statement.
	IfStmt struct {
		If   Foo3
		Init Foo2      // initialization statement; or nil
		Cond Foo2      // condition
		Body Foo2
		Else Foo2 // else branch; or nil
	}

	// A CaseClause represents a case of an expression or type switch statement.
	CaseClause struct {
		Case  Foo3
		List  []Expr   // list of expressions or types; nil means default case
		Colon Foo3
		Body  []Stmt    // statement list; or nil
	}

	// A SwitchStmt node represents an expression switch statement.
	SwitchStmt struct {
		Switch Foo3  // position of "switch" keyword
		Init   Foo2       // initialization statement; or nil
		Tag    Foo2       // tag expression; or nil
		Body   Foo2 // CaseClauses only
	}

	// An TypeSwitchStmt node represents a type switch statement.
	TypeSwitchStmt struct {
		Switch Foo2  // position of "switch" keyword
		Init   Foo2       // initialization statement; or nil
		Assign Foo2       // x := y.(type) or y.(type)
		Body   Foo2 // CaseClauses only
	}

	// A CommClause node represents a case of a select statement.
	CommClause struct {
		Case  Foo2 // position of "case" or "default" keyword
		Comm  Foo2      // send or receive statement; nil means default case
		Colon Foo2 // position of ":"
		Body  []Stmt    // statement list; or nil
	}

	// An SelectStmt node represents a select statement.
	SelectStmt struct {
		Select Foo2  // position of "select" keyword
		Body   Foo2 // CommClauses only
	}

	// A ForStmt represents a for statement.
	ForStmt struct {
		For  Foo3 // position of "for" keyword
		Init Foo2      // initialization statement; or nil
		Cond Foo2      // condition; or nil
		Post Foo2      // post iteration statement; or nil
		Body Foo2
	}

	// A RangeStmt represents a for statement with a range clause.
	RangeStmt struct {
		For        Foo3
		Key, Value Foo2       // Key, Value may be nil
		TokPos     Foo3
		Tok         Foo3 // ILLEGAL if Key == nil, ASSIGN, DEFINE
		X          Foo2//Expr        // value to range over
		Body       Foo2//BlockStmt
	}


)


type ArrayType struct{
	Lbrack Foo3
	Elt Foo2
}

type StarArrayType struct{}
type SwitchCaseClause struct{}
type ValueSpec struct{
	Names []*Ident
	Type Foo2
	Values []Expr
}

type MapType struct{
	Map Foo3
	Key Foo2
	Value Foo2
}

type BadDecl struct {}
type ChanDir struct {}
type ChanType struct {}
type Comment struct {}
type CommentGroup struct {
	List []*Comment // len(List) > 0
}
type IsExported struct {}
type NewIdent struct {}
type Node struct {}
type Package struct {
	Name    string             // package name
	Scope   Foo2             // package scope across all files
	Imports map[string]*Object // map of package id -> package object
	Files   map[string]*File   // Go source files by filename
}




// type Foo interface {
// 	End() token.Pos
// 	Pos() token.Pos
// 	declNode()
// //	ast.Node
// 	specNode()
// }

type Decl interface {
	Report() string
}

type Stmt interface {
	Report() string
}

type Expr interface {
	Report() string
}

type Spec interface {
	Report() string
}

// deferred
type Deferred struct {
	Id    string 
	Data  Foo2
	Set   interface{} 
}
