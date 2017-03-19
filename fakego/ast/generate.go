package ast
import(
	"fmt"
)
func Generate() {
	Names := []string{
		
		"ArrayType",
		"AssignStmt",
		"BadDecl",
		"BadExpr",
		"BadStmt",
		"BasicLit",
		"BinaryExpr",
		"BlockStmt",
		"BranchStmt",
		"CallExpr",
		"CaseClause",
		"ChanDir",
		"ChanType",
		"CommClause",
		"Comment",
		"CompositeLit",
		"Decl",
		"DeclStmt",
		"DeferStmt",
		"Ellipsis",
		"EmptyStmt",
		"Expr",
		"ExprStmt",
		"Field",
		"FieldList",
		"File",
		"ForStmt",
		"FuncDecl",
		"FuncLit",
		"FuncType",
		"GenDecl",
		"GoStmt",
		"Ident",
		"IfStmt",
		"ImportSpec",
		"IncDecStmt",
		"IndexExpr",
		"InterfaceType",
		"IsExported",
		"KeyValueExpr",
		"LabeledStmt",
		"MapType",
		"NewIdent",
		"Node",
		"Package",
		"ParenExpr",
//		"RECV",
		"RangeStmt",
		"ReturnStmt",
//		"SEND",
		"SelectStmt",
		"SelectorExpr",
		"SendStmt",
		"SliceExpr",
		"Spec",
		"StarExpr",
		"Stmt",
		"StructType",
		"SwitchStmt",
		"TypeAssertExpr",
		"TypeSpec",
		"TypeSwitchStmt",
		"UnaryExpr",
		"ValueSpec",
		"CommentGroup",
		"Spec",
		"Object",
	}

	for _,n := range Names {
		fmt.Printf("func (t* %s) Report() (string){ r := fmt.Sprintf(\"%s:%s\",t);fmt.Println(r);return r }\n",n,n,"%+v")
		fmt.Printf("func (t* Table) Ptrmap%s(id string) (* %s){ return t.%ss[id] }\n",n,n,n)
		fmt.Printf("func (t* Table) Strmap%s(id string, f * %s) (*%s){ t.%ss[id] =f; f.Report(); return f}\n",n,n,n,n)
		fmt.Printf("type %s struct {}\n", n)
		fmt.Printf("%ss map[string]*%s\n", n,n)
		fmt.Printf("%ss : make(map[string]*%s),\n", n,n)

	}
}
