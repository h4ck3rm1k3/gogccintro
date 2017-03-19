package ast

import (
	"strings"
	"fmt"
)

func (f* Ident ) ToString( ) (string) {
	return f.Name
}

func (f* Field ) ToString( ) (string) {
	keys := make([]string, 0, len(f.Names))
	for _,k := range f.Names {
		keys = append(keys, k.ToString())
	}
	return "[" + strings.Join(keys, ", ") + "]"
}
	
func (f *FieldList) ToString( ) (string) {
	keys := make([]string, 0, len(f.List))
	for _,k := range f.List {
		keys = append(keys, k.ToString())
	}
	return "[" + strings.Join(keys, ", ") + "]"
}


func (t* Scope) Report() (string){
	for i,j:= range t.Objects {
		fmt.Printf("Scope:%d -: %v\n",i,j)
	}
 	return "TODO" }

// func (t* ArrayType) Report() (string){ return "TODO" }
// func (t* AssignStmt) Report() (string){ return "TODO" }
// func (t* BadDecl) Report() (string){ return "TODO" }
// func (t* BadExpr) Report() (string){ return "TODO" }
// func (t* BadStmt) Report() (string){ return "TODO" }
// func (t* BasicLit) Report() (string){ return "TODO" }
// func (t* BinaryExpr) Report() (string){ return "TODO" }
// func (t* BlockStmt) Report() (string){ return "TODO" }
// func (t* BranchStmt) Report() (string){ return "TODO" }
// func (t* CallExpr) Report() (string){ return "TODO" }
// func (t* CaseClause) Report() (string){ return "TODO" }
// func (t* ChanDir) Report() (string){ return "TODO" }
// func (t* ChanType) Report() (string){ return "TODO" }
// func (t* CommClause) Report() (string){ return "TODO" }
// func (t* Comment) Report() (string){ return "TODO" }
// func (t* CommentGroup) Report() (string){ return "TODO" }
// func (t* CompositeLit) Report() (string){ return "TODO" }
// //func (t* Decl) Report() (string){ return "TODO" }
// func (t* DeclStmt) Report() (string){ return "TODO" }
// func (t* DeferStmt) Report() (string){ return "TODO" }
// func (t* Ellipsis) Report() (string){ return "TODO" }
// func (t* EmptyStmt) Report() (string){ return "TODO" }
// //func (t* Expr) Report() (string){ return "TODO" }
// func (t* ExprStmt) Report() (string){ return "TODO" }
// func (t* Field) Report() (string){ return "TODO" }
// func (t* FieldList) Report() (string){ return "TODO" }
// func (t* File) Report() (string){ return "TODO" }
// func (t* ForStmt) Report() (string){ return "TODO" }
// func (t* FuncDecl) Report() (string){ return "TODO" }
// func (t* FuncLit) Report() (string){ return "TODO" }
// func (t* FuncType) Report() (string){ return "TODO" }
// func (t* GenDecl) Report() (string){
// 	fmt.Printf("GenDecl: Specs(%s)",t.Specs.Report())
// 	return "TODO"
// }
// func (t* GoStmt) Report() (string){ return "TODO" }
// func (t* Ident) Report() (string){ return "TODO" }
// func (t* IfStmt) Report() (string){ return "TODO" }
// func (t* ImportSpec) Report() (string){ return "TODO" }
// func (t* IncDecStmt) Report() (string){ return "TODO" }
// func (t* IndexExpr) Report() (string){ return "TODO" }
// func (t* InterfaceType) Report() (string){ return "TODO" }
// func (t* IsExported) Report() (string){ return "TODO" }
// func (t* KeyValueExpr) Report() (string){ return "TODO" }
// func (t* LabeledStmt) Report() (string){ return "TODO" }
// func (t* MapType) Report() (string){ return "TODO" }
// func (t* NewIdent) Report() (string){ return "TODO" }
// func (t* Node) Report() (string){ return "TODO" }
// func (t* Package) Report() (string){ return "TODO" }
// func (t* ParenExpr) Report() (string){ return "TODO" }
// func (t* RangeStmt) Report() (string){ return "TODO" }
// func (t* ReturnStmt) Report() (string){ return "TODO" }
// func (t* SelectStmt) Report() (string){ return "TODO" }
// func (t* SelectorExpr) Report() (string){ return "TODO" }
// func (t* SendStmt) Report() (string){ return "TODO" }
// func (t* SliceExpr) Report() (string){ return "TODO" }
// //func (t* Spec) Report() (string){ return "TODO" }
// func (t* StarExpr) Report() (string){ return "TODO" }
// //func (t* Stmt) Report() (string){ return "TODO" }
// func (t* StructType) Report() (string){ return "TODO" }
// func (t* SwitchStmt) Report() (string){ return "TODO" }
// func (t* Object) Report() (string){ return "TODO" }

// func (t* TypeAssertExpr) Report() (string){ return "TODO" }
// func (t* TypeSpec) Report() (string){ return "TODO" }
// func (t* TypeSwitchStmt) Report() (string){ return "TODO" }
// func (t* UnaryExpr) Report() (string){ return "TODO" }
// func (t* ValueSpec) Report() (string){ return "TODO" }

// type ArrayType struct {}
// type AssignStmt struct {}
// type BadDecl struct {}
// type BadExpr struct {}
// type BadStmt struct {}
// type BasicLit struct {}
// type BinaryExpr struct {}
// type BlockStmt struct {}
// type BranchStmt struct {}
// type CallExpr struct {}
// type CaseClause struct {}
// type ChanDir struct {}
// type ChanType struct {}
// type CommClause struct {}
// type Comment struct {}
// type CommentGroup struct {}
// type CompositeLit struct {}
// type Decl struct {}
// type DeclStmt struct {}
// type DeferStmt struct {}
// type Ellipsis struct {}
// type EmptyStmt struct {}
// type Expr struct {}
// type ExprStmt struct {}
// type Field struct {}
// type FieldList struct {}
// type File struct {}
// type ForStmt struct {}
// type FuncDecl struct {}
// type FuncLit struct {}
// type FuncType struct {}
// type GenDecl struct {}
// type GoStmt struct {}
// type Ident struct {}
// type IfStmt struct {}
// type ImportSpec struct {}
// type IncDecStmt struct {}
// type IndexExpr struct {}
// type InterfaceType struct {}
// type IsExported struct {}
// type KeyValueExpr struct {}
// type LabeledStmt struct {}
// type MapType struct {}
// type NewIdent struct {}
// type Node struct {}
// type Package struct {}
// type ParenExpr struct {}
// type RangeStmt struct {}
// type ReturnStmt struct {}
// type SelectStmt struct {}
// type SelectorExpr struct {}
// type SendStmt struct {}
// type SliceExpr struct {}
// type Spec struct {}
// type StarExpr struct {}
// type Stmt struct {}
// type StructType struct {}
// type SwitchStmt struct {}
// type TypeAssertExpr struct {}
// type TypeSpec struct {}
// type TypeSwitchStmt struct {}
// type UnaryExpr struct {}
// type ValueSpec struct {}


     func (t* ArrayType) Report() (string){ r := fmt.Sprintf("ArrayType:%+v",t);fmt.Println(r);return r }
     func (t* AssignStmt) Report() (string){ r := fmt.Sprintf("AssignStmt:%+v",t);fmt.Println(r);return r }
     func (t* BadDecl) Report() (string){ r := fmt.Sprintf("BadDecl:%+v",t);fmt.Println(r);return r }
     func (t* BadExpr) Report() (string){ r := fmt.Sprintf("BadExpr:%+v",t);fmt.Println(r);return r }
     func (t* BadStmt) Report() (string){ r := fmt.Sprintf("BadStmt:%+v",t);fmt.Println(r);return r }
     func (t* BasicLit) Report() (string){ r := fmt.Sprintf("BasicLit:%+v",t);fmt.Println(r);return r }
     func (t* BinaryExpr) Report() (string){ r := fmt.Sprintf("BinaryExpr:%+v",t);fmt.Println(r);return r }
     func (t* BlockStmt) Report() (string){ r := fmt.Sprintf("BlockStmt:%+v",t);fmt.Println(r);return r }
     func (t* BranchStmt) Report() (string){ r := fmt.Sprintf("BranchStmt:%+v",t);fmt.Println(r);return r }
     func (t* CallExpr) Report() (string){ r := fmt.Sprintf("CallExpr:%+v",t);fmt.Println(r);return r }
     func (t* CaseClause) Report() (string){ r := fmt.Sprintf("CaseClause:%+v",t);fmt.Println(r);return r }
     func (t* ChanDir) Report() (string){ r := fmt.Sprintf("ChanDir:%+v",t);fmt.Println(r);return r }
     func (t* ChanType) Report() (string){ r := fmt.Sprintf("ChanType:%+v",t);fmt.Println(r);return r }

    func (t* CommClause) Report() (string){ r := fmt.Sprintf("CommClause:%+v",t);fmt.Println(r);return r }
    func (t* Comment) Report() (string){ r := fmt.Sprintf("Comment:%+v",t);fmt.Println(r);return r }
    func (t* CommentGroup) Report() (string){ r := fmt.Sprintf("CommentGroup:%+v",t);fmt.Println(r);return r }
    func (t* CompositeLit) Report() (string){ r := fmt.Sprintf("CompositeLit:%+v",t);fmt.Println(r);return r }
    //func (t* Decl) Report() (string){ r := fmt.Sprintf("Decl:%+v",t);fmt.Println(r);return r }
    func (t* DeclStmt) Report() (string){ r := fmt.Sprintf("DeclStmt:%+v",t);fmt.Println(r);return r }
    func (t* DeferStmt) Report() (string){ r := fmt.Sprintf("DeferStmt:%+v",t);fmt.Println(r);return r }
    func (t* Ellipsis) Report() (string){ r := fmt.Sprintf("Ellipsis:%+v",t);fmt.Println(r);return r }
    func (t* EmptyStmt) Report() (string){ r := fmt.Sprintf("EmptyStmt:%+v",t);fmt.Println(r);return r }
    //func (t* Expr) Report() (string){ r := fmt.Sprintf("Expr:%+v",t);fmt.Println(r);return r }
    func (t* ExprStmt) Report() (string){ r := fmt.Sprintf("ExprStmt:%+v",t);fmt.Println(r);return r }
    func (t* Field) Report() (string){ r := fmt.Sprintf("Field:%+v",t);fmt.Println(r);return r }
    func (t* FieldList) Report() (string){ r := fmt.Sprintf("FieldList:%+v",t);fmt.Println(r);return r }
    func (t* File) Report() (string){ r := fmt.Sprintf("File:%+v",t);fmt.Println(r);return r }
    func (t* ForStmt) Report() (string){ r := fmt.Sprintf("ForStmt:%+v",t);fmt.Println(r);return r }
    func (t* FuncDecl) Report() (string){ r := fmt.Sprintf("FuncDecl:%+v",t);fmt.Println(r);return r }
    func (t* FuncLit) Report() (string){ r := fmt.Sprintf("FuncLit:%+v",t);fmt.Println(r);return r }
    func (t* FuncType) Report() (string){ r := fmt.Sprintf("FuncType:%+v",t);fmt.Println(r);return r }
    func (t* GenDecl) Report() (string){ r := fmt.Sprintf("GenDecl:%+v",t);fmt.Println(r);return r }
    func (t* GoStmt) Report() (string){ r := fmt.Sprintf("GoStmt:%+v",t);fmt.Println(r);return r }
    func (t* Ident) Report() (string){ r := fmt.Sprintf("Ident:%+v",t);fmt.Println(r);return r }
    func (t* IfStmt) Report() (string){ r := fmt.Sprintf("IfStmt:%+v",t);fmt.Println(r);return r }
    func (t* ImportSpec) Report() (string){ r := fmt.Sprintf("ImportSpec:%+v",t);fmt.Println(r);return r }
    func (t* IncDecStmt) Report() (string){ r := fmt.Sprintf("IncDecStmt:%+v",t);fmt.Println(r);return r }
    func (t* IndexExpr) Report() (string){ r := fmt.Sprintf("IndexExpr:%+v",t);fmt.Println(r);return r }
    func (t* InterfaceType) Report() (string){ r := fmt.Sprintf("InterfaceType:%+v",t);fmt.Println(r);return r }
    func (t* IsExported) Report() (string){ r := fmt.Sprintf("IsExported:%+v",t);fmt.Println(r);return r }
    func (t* KeyValueExpr) Report() (string){ r := fmt.Sprintf("KeyValueExpr:%+v",t);fmt.Println(r);return r }
    func (t* LabeledStmt) Report() (string){ r := fmt.Sprintf("LabeledStmt:%+v",t);fmt.Println(r);return r }
    func (t* MapType) Report() (string){ r := fmt.Sprintf("MapType:%+v",t);fmt.Println(r);return r }
    func (t* NewIdent) Report() (string){ r := fmt.Sprintf("NewIdent:%+v",t);fmt.Println(r);return r }
    func (t* Node) Report() (string){ r := fmt.Sprintf("Node:%+v",t);fmt.Println(r);return r }
    func (t* Object) Report() (string){ r := fmt.Sprintf("Object:%+v",t);fmt.Println(r);return r }
    func (t* Package) Report() (string){ r := fmt.Sprintf("Package:%+v",t);fmt.Println(r);return r }
    func (t* ParenExpr) Report() (string){ r := fmt.Sprintf("ParenExpr:%+v",t);fmt.Println(r);return r }
    func (t* RangeStmt) Report() (string){ r := fmt.Sprintf("RangeStmt:%+v",t);fmt.Println(r);return r }
    func (t* ReturnStmt) Report() (string){ r := fmt.Sprintf("ReturnStmt:%+v",t);fmt.Println(r);return r }
    func (t* SelectStmt) Report() (string){ r := fmt.Sprintf("SelectStmt:%+v",t);fmt.Println(r);return r }
    func (t* SelectorExpr) Report() (string){ r := fmt.Sprintf("SelectorExpr:%+v",t);fmt.Println(r);return r }
    func (t* SendStmt) Report() (string){ r := fmt.Sprintf("SendStmt:%+v",t);fmt.Println(r);return r }
    func (t* SliceExpr) Report() (string){ r := fmt.Sprintf("SliceExpr:%+v",t);fmt.Println(r);return r }
//func (t* Spec) Report() (string){ r := fmt.Sprintf("Spec:%+v",t);fmt.Println(r);return r }

    func (t* StarExpr) Report() (string){ r := fmt.Sprintf("StarExpr:%+v",t);fmt.Println(r);return r }
    //func (t* Stmt) Report() (string){ r := fmt.Sprintf("Stmt:%+v",t);fmt.Println(r);return r }
    func (t* StructType) Report() (string){ r := fmt.Sprintf("StructType:%+v",t);fmt.Println(r);return r }
    func (t* SwitchStmt) Report() (string){ r := fmt.Sprintf("SwitchStmt:%+v",t);fmt.Println(r);return r }
    func (t* TypeAssertExpr) Report() (string){ r := fmt.Sprintf("TypeAssertExpr:%+v",t);fmt.Println(r);return r }
    func (t* TypeSpec) Report() (string){ r := fmt.Sprintf("TypeSpec:%+v",t);fmt.Println(r);return r }
    func (t* TypeSwitchStmt) Report() (string){ r := fmt.Sprintf("TypeSwitchStmt:%+v",t);fmt.Println(r);return r }
    func (t* UnaryExpr) Report() (string){ r := fmt.Sprintf("UnaryExpr:%+v",t);fmt.Println(r);return r }
    func (t* ValueSpec) Report() (string){ r := fmt.Sprintf("ValueSpec:%+v",t);fmt.Println(r);return r }
