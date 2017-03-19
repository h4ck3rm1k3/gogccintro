package node_types

import (

	//"reflect"
	"fmt"

	//"go/ast"
	//"go/parser"
	//"go/token"

	"github.com/h4ck3rm1k3/gogccintro/go/ast"
	"github.com/h4ck3rm1k3/gogccintro/go/parser"
	"github.com/h4ck3rm1k3/gogccintro/go/token"
)

func MapOne(from interface {}, to interface {}) {
	fmt.Println(from)
	fmt.Println(to)
}

func AstParse(file string) {
	fset := token.NewFileSet()
	f, e := parser.ParseFile(fset, file, nil, 0)
	if e != nil {
		//log.Fatal(e)
		fmt.Printf("fatal %s",e)
	}
	ast.Print(fset, f)
}

func DoMap() {

	AstParse("/usr/share/go-1.7/src/go/ast/ast.go")
/*
	AstParse("types.go")
	
	MapOne(
		reflect.TypeOf((*ast.FuncDecl)(nil)),
		reflect.TypeOf((*NodeTypeFunctionDecl)(nil)),
	)
*/
//	VistAst()
}
