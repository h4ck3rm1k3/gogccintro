package ast
import(
	"fmt"
	"strings"
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

		"DeclStmt",
		"DeferStmt",
		"Ellipsis",
		"EmptyStmt",

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

	Interfaces := []string{
		// interfaces
		"Expr",
		"Stmt",
		"Decl",
		"Node",
	}

	for j,n := range Names {
		n2 := strings.ToUpper(n)
		n3 := strings.ToLower(n)
		fmt.Printf("%s = %d, // ENUM BASE\n",n2,j)
		fmt.Printf("optional %s %s = %d, // InterOAface:BASE\n",n,n3,j)			
		
		for _,i := range Interfaces {
			if strings.Contains(n,i) {				
				fmt.Printf("%s = %d, // ENUM %s\n",n2,j,i)
				fmt.Printf("optional %s %s = %d, // Interface:%s\n",n,n3,j,i)			
			}
		}
		
		fmt.Printf("func (t* %s) Report() (string){ r := fmt.Sprintf(\"%s:%s\",t);fmt.Println(r);return r }\n",n,n,"%+v")
		fmt.Printf("func (t* Table) Ptrmap%s(id string) (* %s){ if val,ok := t.%ss[id]; ok { return val } else {  return Future%s(id) } }\n",n,n,n,n)
		fmt.Printf("func (t* Table) Ptrmap%s(id string) (* %s){ if val,ok := t.%ss[id]; ok { return val } else {  return Future%s(id) } }\n",n,n,n,n)
		
		//fmt.Printf("func (t* Table) Strmap%s(id string, f * %s) (*%s){ t.%ss[id] =f; f.Report(); return f}\n",n,n,n,n)
		fmt.Printf("func (t* Table) Strmap%s(id string, f * ast.%s) (*%s){ f2 := t.Convert%s(f); t.%ss[id] =f2; return f2}\n",n,n,n,n,n)
		fmt.Printf("type %s struct {}\n", n)
		fmt.Printf("%ss map[string]*%s\n", n,n)
		fmt.Printf("%ss : make(map[string]*%s),\n", n,n)
		fmt.Printf("lager.Register(ast.%s{})\n", n)

	}
}
