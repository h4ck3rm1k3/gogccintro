package astproto
import (
	"fmt"
	"strings"
//	"encoding/json"	
	"github.com/h4ck3rm1k3/gogccintro/fakego/ast"
//	"github.com/h4ck3rm1k3/gogccintro/fakego/ast/proto"
	//"github.com/golang/protobuf/proto"
	//"github.com/h4ck3rm1k3/gogccintro/fakego/token"
)

func ProcessBuf(x string){
	fmt.Printf("\t\tCheck %s\n", x)
}

func (t *AddressTable) Check(id string , f2 interface{},f interface{}) {
//	fmt.Printf("Check %s\n", f)
//	fmt.Printf("Check %T\n", f)
	//fmt.Printf("Check %s %#v => %#v\n", id, f,f2)

	t.Idsold[id]=f
	t.Idsnew[id]=f
	t.OldNew[f]=f2

	
	switch v:= f.(type) {
	case *ast.Ident: switch v2:= f2.(type) {case *Ident:ProcessBuf(fmt.Sprintf("*Ident:%#v\n",v2.String()))}
	case *ast.BasicLit: switch v2:= f2.(type) {case *BasicLit: ProcessBuf(fmt.Sprintf("*BasicLit:%s\n",v2.String))}
	case *ast.ImportSpec: switch v2:= f2.(type) {case *ImportSpec: ProcessBuf(fmt.Sprintf("*ImportSpec:%s\n",v2.String()))}
	case *ast.GenDecl: switch v2:= f2.(type) {case *GenDecl: ProcessBuf(fmt.Sprintf("*astproto.GenDecl:%s\n",v2.String()))}
	case *ast.Object: switch v2:= f2.(type) {case *Object:
		fmt.Printf("*astproto.Object:%s\n",v2.String())}
	case *ast.FieldList: switch v2:= f2.(type) {case *FieldList: ProcessBuf(fmt.Sprintf("*astproto.FieldList:%s\n",v2.String()))}
	case *ast.ArrayType: switch v2:= f2.(type) {case *ArrayType: ProcessBuf(fmt.Sprintf("*astproto.ArrayType:%s\n",v2.String()))}
	case *ast.AssignStmt: switch v2:= f2.(type) {case *AssignStmt: ProcessBuf(fmt.Sprintf("*astproto.AssignStmt:%s\n",v2.String()))}
	case *ast.BinaryExpr: switch v2:= f2.(type) {case *BinaryExpr: ProcessBuf(fmt.Sprintf("*astproto.BinaryExpr:%s\n",v2.String()))}
	case *ast.BlockStmt: switch v2:= f2.(type) {case *BlockStmt: ProcessBuf(fmt.Sprintf("*astproto.BlockStmt:%s\n",v2.String()))}
	case *ast.CallExpr: switch v2:= f2.(type) {case *CallExpr: ProcessBuf(fmt.Sprintf("*astproto.CallExpr:%s\n",v2.String()))}
	case *ast.CaseClause: switch v2:= f2.(type) {case *CaseClause: ProcessBuf(fmt.Sprintf("*astproto.CaseClause:%s\n",v2.String()))}
	case *ast.CompositeLit: switch v2:= f2.(type) {case *CompositeLit: ProcessBuf(fmt.Sprintf("*astproto.CompositeLit:%s\n",v2.String()))}
	case *ast.Field: switch v2:= f2.(type) {case *Field: ProcessBuf(fmt.Sprintf("*astproto.Field:%s\n",v2.String()))}
	case *ast.File: switch v2:= f2.(type) {case *File: ProcessBuf(fmt.Sprintf("*astproto.File:%s\n",v2.String()))}
	case *ast.ForStmt: switch v2:= f2.(type) {case *ForStmt: ProcessBuf(fmt.Sprintf("*astproto.ForStmt:%s\n",v2.String()))}
	case *ast.FuncDecl: switch v2:= f2.(type) {case *FuncDecl: ProcessBuf(fmt.Sprintf("*astproto.FuncDecl:%s\n",v2.String()))}
	case *ast.FuncType: switch v2:= f2.(type) {case *FuncType: ProcessBuf(fmt.Sprintf("*astproto.FuncType:%s\n",v2.String()))}
	case *ast.IfStmt: switch v2:= f2.(type) {case *IfStmt: ProcessBuf(fmt.Sprintf("*astproto.IfStmt:%s\n",v2.String()))}
	case *ast.IncDecStmt: switch v2:= f2.(type) {case *IncDecStmt: ProcessBuf(fmt.Sprintf("*astproto.IncDecStmt:%s\n",v2.String()))}
	case *ast.IndexExpr: switch v2:= f2.(type) {case *IndexExpr: ProcessBuf(fmt.Sprintf("*astproto.IndexExpr:%s\n",v2.String()))}
	case *ast.InterfaceType: switch v2:= f2.(type) {case *InterfaceType: ProcessBuf(fmt.Sprintf("*astproto.InterfaceType:%s\n",v2.String()))}
	case *ast.MapType: switch v2:= f2.(type) {case *MapType: ProcessBuf(fmt.Sprintf("*astproto.MapType:%s\n",v2.String()))}
	case *ast.RangeStmt: switch v2:= f2.(type) {case *RangeStmt: ProcessBuf(fmt.Sprintf("*astproto.RangeStmt:%s\n",v2.String()))}
	case *ast.ReturnStmt: switch v2:= f2.(type) {case *ReturnStmt: ProcessBuf(fmt.Sprintf("*astproto.ReturnStmt:%s\n",v2.String()))}
	case *ast.Scope: switch v2:= f2.(type) {case *Scope: ProcessBuf(fmt.Sprintf("*astproto.Scope:%s\n",v2.String()))}
	case *ast.SelectorExpr: switch v2:= f2.(type) {case *SelectorExpr: ProcessBuf(fmt.Sprintf("*astproto.SelectorExpr:%s\n",v2.String()))}
	case *ast.SliceExpr: switch v2:= f2.(type) {case *SliceExpr: ProcessBuf(fmt.Sprintf("*astproto.SliceExpr:%s\n",v2.String()))}
	case *ast.StarExpr: switch v2:= f2.(type) {case *StarExpr: ProcessBuf(fmt.Sprintf("*astproto.StarExpr:%s\n",v2.String()))}
	case *ast.StructType: switch v2:= f2.(type) {case *StructType: ProcessBuf(fmt.Sprintf("*astproto.StructType:%s\n",v2.String()))}
	case *ast.SwitchStmt: switch v2:= f2.(type) {case *SwitchStmt: ProcessBuf(fmt.Sprintf("*astproto.SwitchStmt:%s\n",v2.String()))}
	case *ast.TypeSpec: switch v2:= f2.(type) {case *TypeSpec: ProcessBuf(fmt.Sprintf("*astproto.TypeSpec:%s\n",v2.String()))}
	case *ast.UnaryExpr: switch v2:= f2.(type) {case *UnaryExpr: ProcessBuf(fmt.Sprintf("*astproto.UnaryExpr:%s\n",v2.String()))}
	case *ast.ValueSpec: switch v2:= f2.(type) {case *ValueSpec: ProcessBuf(fmt.Sprintf("*astproto.ValueSpec:%s\n",v2.String()))}
		
	default:
		//fmt.Printf("case %T to %T\n", v, v2)
		switch v2:= f2.(type) {
		default:

			fmt.Printf("case %T: switch v2:= f2.(type) {case %s: ProcessBuf(fmt.Sprintf(\"%T:%%s\\n\",v2.String()))}\n",
				v,
				strings.Replace(fmt.Sprintf("%T",v2),
					"astproto.","",-1,
				),
				v2)
		}
		
	}

//	body, _ := json.Marshal(f)
//	fmt.Printf("json : %s\n",body)

}
