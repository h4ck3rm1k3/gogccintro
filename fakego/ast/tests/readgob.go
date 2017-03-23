package tests

import (
	//"strings"
	"fmt"
	"os"
	lager "github.com/lowentropy/go-lager"
	"github.com/h4ck3rm1k3/gogccintro/fakego/ast"
)

func DoUnmarshal(filename string ) {
	lager.Register(ast.Deferred{})
	lager.Register(ast.Deferred2{})
	lager.Register(ast.Scope{})

	lager.Register(ast.ArrayType{})
	lager.Register(ast.AssignStmt{})
	lager.Register(ast.BadDecl{})
	lager.Register(ast.BadExpr{})
	lager.Register(ast.BadStmt{})
	lager.Register(ast.BasicLit{})
	lager.Register(ast.BinaryExpr{})
	lager.Register(ast.BlockStmt{})
	lager.Register(ast.BranchStmt{})
	lager.Register(ast.CallExpr{})
	lager.Register(ast.CaseClause{})
	lager.Register(ast.ChanDir{})
	lager.Register(ast.ChanType{})
	lager.Register(ast.CommClause{})
	lager.Register(ast.Comment{})
	lager.Register(ast.CompositeLit{})
	lager.Register(ast.DeclStmt{})
	lager.Register(ast.DeferStmt{})
	lager.Register(ast.Ellipsis{})
	lager.Register(ast.EmptyStmt{})
	lager.Register(ast.ExprStmt{})
	lager.Register(ast.Field{})
	lager.Register(ast.FieldList{})
	lager.Register(ast.File{})
	lager.Register(ast.ForStmt{})
	lager.Register(ast.FuncDecl{})
	lager.Register(ast.FuncLit{})
	lager.Register(ast.FuncType{})
	lager.Register(ast.GenDecl{})
	lager.Register(ast.GoStmt{})
	lager.Register(ast.Ident{})
	lager.Register(ast.IfStmt{})
	lager.Register(ast.ImportSpec{})
	lager.Register(ast.IncDecStmt{})
	lager.Register(ast.IndexExpr{})
	lager.Register(ast.InterfaceType{})
	lager.Register(ast.IsExported{})
	lager.Register(ast.KeyValueExpr{})
	lager.Register(ast.LabeledStmt{})
	lager.Register(ast.MapType{})
	lager.Register(ast.NewIdent{})
	lager.Register(ast.Package{})
	lager.Register(ast.ParenExpr{})
	lager.Register(ast.RangeStmt{})
	lager.Register(ast.ReturnStmt{})
	lager.Register(ast.SelectStmt{})
	lager.Register(ast.SelectorExpr{})
	lager.Register(ast.SendStmt{})
	lager.Register(ast.SliceExpr{})
	//lager.Register(ast.Spec{})
	lager.Register(ast.StarExpr{})
	lager.Register(ast.StructType{})
	lager.Register(ast.SwitchStmt{})
	lager.Register(ast.TypeAssertExpr{})
	lager.Register(ast.TypeSpec{})
	lager.Register(ast.TypeSwitchStmt{})
	lager.Register(ast.UnaryExpr{})
	lager.Register(ast.ValueSpec{})
	lager.Register(ast.CommentGroup{})

	lager.Register(ast.Object{})
	
//lager.Register(Foo{})                     // register the Foo type
	fo, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	decoder,err := lager.NewDecoder(fo)
	if err != nil {
		panic(err)
		//fmt.Printf("Error %s\n",err)
	}

	obj,err := decoder.Read(); 
	if err != nil {
		panic(err)
	}
	fmt.Printf("ok %s", obj)

}



