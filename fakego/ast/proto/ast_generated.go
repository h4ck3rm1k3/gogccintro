package astproto

import (
	"strings"
	"fmt"
	"os"
	"bufio"
	//"bytes"
	//"encoding/json"
	//"gopkg.in/yaml.v2"
	lager "github.com/lowentropy/go-lager"
)

func DoMarshal(i interface{}, filename string ) {

	fo, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()
	w := bufio.NewWriter(fo)

	//var buffer bytes.Buffer
	encoder := lager.NewEncoder(w)
	encoder.Write(i)
	encoder.Finish()    
	//body, _ := yaml.Marshal(i)
	//buffer.WriteString("\n")
	//fmt.Printf("%s",buffer.String())
}

func (f* Ident ) ToString( ) (string) {	return *f.Name}

func (f* Field ) ToString( ) (string) {
	keys := make([]string, 0, len(f.Names))
	for _,k := range f.Names {
		keys = append(keys, k.ToString())
	}
	return "[" + strings.Join(keys, ", ") + "]"
}

//func (t* FieldList) Report() (string){ r := fmt.Sprintf("FieldList:%+v",t);/*fmt.Println(r)*/;return r }
func (f *FieldList) Report( ) (string) {
	keys := make([]string, 0, len(f.List))
	for _,k := range f.List {
		keys = append(keys, k.Report())
	}
	return "FieldList[\n\t" + strings.Join(keys, ", ") + "]"
}

func (t* Scope) Report() (string){

	DoMarshal(t,"scope.gob")
	
	for i,j:= range t.Objects {
		fmt.Printf("Scope:%s :\n\t%v\n",i,j.Report())

	}
 	return "Scope..." }


//func (t* ArrayType) Report() (string){	r := fmt.Sprintf("ArrayType:%s",t.Elt.Report());/*fmt.Println(r)*/;return r}
     func (t* AssignStmt) Report() (string){ r := fmt.Sprintf("AssignStmt:%+v",t);/*fmt.Println(r)*/;return r }
//     func (t* BadDecl) Report() (string){ r := fmt.Sprintf("BadDecl:%+v",t);/*fmt.Println(r)*/;return r }
     func (t* BadExpr) Report() (string){ r := fmt.Sprintf("BadExpr:%+v",t);/*fmt.Println(r)*/;return r }
     func (t* BadStmt) Report() (string){ r := fmt.Sprintf("BadStmt:%+v",t);/*fmt.Println(r)*/;return r }
     func (t* BasicLit) Report() (string){ r := fmt.Sprintf("BasicLit:%+v",t);/*fmt.Println(r)*/;return r }
     func (t* BinaryExpr) Report() (string){ r := fmt.Sprintf("BinaryExpr:%+v",t);/*fmt.Println(r)*/;return r }
     func (t* BlockStmt) Report() (string){ r := fmt.Sprintf("BlockStmt:%+v",t);/*fmt.Println(r)*/;return r }
     func (t* BranchStmt) Report() (string){ r := fmt.Sprintf("BranchStmt:%+v",t);/*fmt.Println(r)*/;return r }
     func (t* CallExpr) Report() (string){ r := fmt.Sprintf("CallExpr:%+v",t);/*fmt.Println(r)*/;return r }
     func (t* CaseClause) Report() (string){ r := fmt.Sprintf("CaseClause:%+v",t);/*fmt.Println(r)*/;return r }
//     func (t* ChanDir) Report() (string){ r := fmt.Sprintf("ChanDir:%+v",t);/*fmt.Println(r)*/;return r }
//     func (t* ChanType) Report() (string){ r := fmt.Sprintf("ChanType:%+v",t);/*fmt.Println(r)*/;return r }

    func (t* CommClause) Report() (string){ r := fmt.Sprintf("CommClause:%+v",t);/*fmt.Println(r)*/;return r }
    //func (t* Comment) Report() (string){ r := fmt.Sprintf("Comment:%+v",t);/*fmt.Println(r)*/;return r }
    func (t* CommentGroup) Report() (string){ r := fmt.Sprintf("CommentGroup:%+v",t);/*fmt.Println(r)*/;return r }
    func (t* CompositeLit) Report() (string){ r := fmt.Sprintf("CompositeLit:%+v",t);/*fmt.Println(r)*/;return r }
    //func (t* Decl) Report() (string){ r := fmt.Sprintf("Decl:%+v",t);/*fmt.Println(r)*/;return r }
//    func (t* DeclStmt) Report() (string){ r := fmt.Sprintf("DeclStmt:%+v",t);/*fmt.Println(r)*/;return r }
    func (t* DeferStmt) Report() (string){ r := fmt.Sprintf("DeferStmt:%+v",t);/*fmt.Println(r)*/;return r }
    func (t* Ellipsis) Report() (string){ r := fmt.Sprintf("Ellipsis:%+v",t);/*fmt.Println(r)*/;return r }
    func (t* EmptyStmt) Report() (string){ r := fmt.Sprintf("EmptyStmt:%+v",t);/*fmt.Println(r)*/;return r }
    //func (t* Expr) Report() (string){ r := fmt.Sprintf("Expr:%+v",t);/*fmt.Println(r)*/;return r }
//    func (t* ExprStmt) Report() (string){ r := fmt.Sprintf("ExprStmt:%+v",t);/*fmt.Println(r)*/;return r }

func (t* Field) Report() (string){
	if t == nil {
		return "Field:NIL"
	}
	keys := make([]string, 0, len(t.Names))
	for _,k := range t.Names {
		keys = append(keys, k.ToString())
	}
	r := fmt.Sprintf("\tField(Names:%s Type:<%s>)\n","[" + strings.Join(keys, ", ") + "]", t.Type.Report());	
	return r
}

func (t* File) Report() (string){
	r := fmt.Sprintf("File:%+v",t);
	for _,t := range t.Decls {
		t.Report()
	}
	/*fmt.Println(r)*/;
	return r
}
    func (t* ForStmt) Report() (string){ r := fmt.Sprintf("ForStmt:%+v",t);/*fmt.Println(r)*/;return r }
    func (t* FuncDecl) Report() (string){ r := fmt.Sprintf("FuncDecl:%+v",t);/*fmt.Println(r)*/;return r }
    func (t* FuncLit) Report() (string){ r := fmt.Sprintf("FuncLit:%+v",t);/*fmt.Println(r)*/;return r }

func (t* FuncType) Report() (string){
	r := fmt.Sprintf("FuncType:%+v",t);/*fmt.Println(r)*/;
	return r
}
    func (t* GenDecl) Report() (string){ r := fmt.Sprintf("GenDecl:%+v",t);/*fmt.Println(r)*/;return r }
    func (t* GoStmt) Report() (string){ r := fmt.Sprintf("GoStmt:%+v",t);/*fmt.Println(r)*/;return r }
func (t* Ident) Report() (string){
	r := fmt.Sprintf("Ident(%s)",t.Name);
	/*fmt.Println(r)*/;return r
}
    func (t* IfStmt) Report() (string){ r := fmt.Sprintf("IfStmt:%+v",t);/*fmt.Println(r)*/;return r }
    func (t* ImportSpec) Report() (string){ r := fmt.Sprintf("ImportSpec:%+v",t);/*fmt.Println(r)*/;return r }
    func (t* IncDecStmt) Report() (string){ r := fmt.Sprintf("IncDecStmt:%+v",t);/*fmt.Println(r)*/;return r }
    func (t* IndexExpr) Report() (string){ r := fmt.Sprintf("IndexExpr:%+v",t);/*fmt.Println(r)*/;return r }
    func (t* InterfaceType) Report() (string){ r := fmt.Sprintf("InterfaceType:%+v",t);/*fmt.Println(r)*/;return r }
//    func (t* IsExported) Report() (string){ r := fmt.Sprintf("IsExported:%+v",t);/*fmt.Println(r)*/;return r }
    func (t* KeyValueExpr) Report() (string){ r := fmt.Sprintf("KeyValueExpr:%+v",t);/*fmt.Println(r)*/;return r }
    func (t* LabeledStmt) Report() (string){ r := fmt.Sprintf("LabeledStmt:%+v",t);/*fmt.Println(r)*/;return r }
func (t* MapType) Report() (string){
	r := fmt.Sprintf("MapType:%s->%s",t.Key.Report(),t.Value.Report());
	/*fmt.Println(r)*/;return r
}
//    func (t* NewIdent) Report() (string){ r := fmt.Sprintf("NewIdent:%+v",t);/*fmt.Println(r)*/;return r }
//    func (t* Node) Report() (string){ r := fmt.Sprintf("Node:%+v",t);/*fmt.Println(r)*/;return r }

    func (t* Package) Report() (string){ r := fmt.Sprintf("Package:%+v",t);/*fmt.Println(r)*/;return r }
    func (t* ParenExpr) Report() (string){ r := fmt.Sprintf("ParenExpr:%+v",t);/*fmt.Println(r)*/;return r }
    func (t* RangeStmt) Report() (string){ r := fmt.Sprintf("RangeStmt:%+v",t);/*fmt.Println(r)*/;return r }
//    func (t* ReturnStmt) Report() (string){ r := fmt.Sprintf("ReturnStmt:%+v",t);/*fmt.Println(r)*/;return r }
    func (t* SelectStmt) Report() (string){ r := fmt.Sprintf("SelectStmt:%+v",t);/*fmt.Println(r)*/;return r }
func (t* SelectorExpr) Report() (string){
	r := fmt.Sprintf("SelectorExpr:X:%s Sel:%s",
		t.X.Report(),
		t.Sel.Report(),
	);
	/*fmt.Println(r)*/;
	return r
}
    func (t* SendStmt) Report() (string){ r := fmt.Sprintf("SendStmt:%+v",t);/*fmt.Println(r)*/;return r }
    func (t* SliceExpr) Report() (string){ r := fmt.Sprintf("SliceExpr:%+v",t);/*fmt.Println(r)*/;return r }
//func (t* Spec) Report() (string){ r := fmt.Sprintf("Spec:%+v",t);/*fmt.Println(r)*/;return r }

//func (t* StarExpr) Report() (string){	r := fmt.Sprintf("StarExpr:%s",t.X.Report());/*fmt.Println(r)*/;	return r }
    //func (t* Stmt) Report() (string){ r := fmt.Sprintf("Stmt:%+v",t);/*fmt.Println(r)*/;return r }
func (t* StructType) Report() (string){
	
	r := fmt.Sprintf("StructType:\n\tFields:%s\n",t.Fields.Report());
	/*fmt.Println(r)*/;
	return r
}
    func (t* SwitchStmt) Report() (string){ r := fmt.Sprintf("SwitchStmt:%+v",t);/*fmt.Println(r)*/;return r }
    func (t* TypeAssertExpr) Report() (string){ r := fmt.Sprintf("TypeAssertExpr:%+v",t);/*fmt.Println(r)*/;return r }

func (t* TypeSpec) Report() (string){
	if t == nil {
		return "TypeSpec:NULL TODO"
	}
	
	r := fmt.Sprintf("TypeSpec: Name:%s\nType:%s", t.Name.Report(), t.Type.Report());
	///*fmt.Println(r)*/;
	return r
}
    func (t* TypeSwitchStmt) Report() (string){ r := fmt.Sprintf("TypeSwitchStmt:%+v",t);/*fmt.Println(r)*/;return r }
    func (t* UnaryExpr) Report() (string){ r := fmt.Sprintf("UnaryExpr:%+v",t);/*fmt.Println(r)*/;return r }
    func (t* ValueSpec) Report() (string){ r := fmt.Sprintf("ValueSpec:%+v",t);/*fmt.Println(r)*/;return r }


func (t* Object) Report() (string){
	r := fmt.Sprintf("Object: Kind:%s Name:%s Decl:%+v",t.Kind, t.Name, t.Decl.Report());
	/*fmt.Println(r)*/;
	
	return r
}


func (t* Foo2) Report() (string){
	return "todo"
}

func (t* Decl) Report() (string){
	return "todo"
}

func (t* Expr) Report() (string){
	return "todo"
}

func (t* Spec) Report() (string){
	return "todo"
}

func (t* Stmt) Report() (string){
	return "todo"
}
