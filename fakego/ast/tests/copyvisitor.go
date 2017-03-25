package tests
import (
	"fmt"
	//"reflect"
	//"reflect"
	"github.com/h4ck3rm1k3/gogccintro/fakego/ast"
)
// typed copy visitor

type CreateCopyVisitor struct {

}

func (v*CreateCopyVisitor) VisitScope(t *ast.Table, s * ast.Scope) {
	for name,strt:= range s.Objects { //

		fmt.Printf("func (t* AddressTable) Convert%s(f *ast.%s) (*%s){\n",name,name,name)
		fmt.Printf("\tf2 := false\n")
		fmt.Printf("\treturn &%s{\n",name)
		//fmt.Printf("Struct type %T\n",strt.Decl)
		//fmt.Printf("Structval %#v\n",strt.Decl)
		
		switch  v3 := strt.Decl.(type){
		case *ast.Deferred:
			i := v3.Set
			j := v3.Id
			switch v2:= i.(type) {
			case *map[string]*ast.Field :
				if val, ok := (*v2)[j]; ok {
					if val == nil {
						fmt.Printf("did not find %s %s",v,t)
					} else {
						fmt.Printf("field %#v\n",val)
						//return d.Report()					
					}
				}
			case *map[string]*ast.FieldList :
				if val, ok := (*v2)[j]; ok {
					if val == nil {
						fmt.Printf("did not find %s %s",v,t)
					} else {
						fmt.Printf("fieldlist %#v\n",val)						
					}
				}
			case *map[string]*ast.FuncDecl :
				if val, ok := (*v2)[j]; ok {
					if val == nil {
						fmt.Printf("did not find %s %s",v,t)
					} else {
						fmt.Printf("funcdecl %#v\n",val)
					}
				}
			//case interface{}:
			//	fmt.Printf("Random Interface %+v\n",v)
			case *ast.ValueSpec:
				fmt.Printf("ValueSpec %+v\n",v)
				
			case *map[string]*ast.TypeSpec :
				
				if val, ok := (*v2)[j]; ok {
					if val == nil {
						fmt.Printf("did not find %s %s",v,t)
					} else {
						d := (*v2)[j]
						//fmt.Printf("typespec %#v\n",d)
						//fmt.Printf("typespec name : %s\n",d.Name)
						v.VisitTypeSpec(t,s,d)
						//return d.Report()					
					}
				}
			default:
				fmt.Printf("other type %s\n",v)			
				
			}
			
		default:
			fmt.Printf("other type %s\n",v)			
		}
		// case ast.Deferred : 
		// 	fmt.Printf("Def %+v\n",v)
			
		// 	// case ast.StructType:
		// // 	for _,fld := range strt.Decl.Fields.List {
		// // 		for _,fn := range fld.Names {
		// // 			fmt.Printf("Field %s\n",fn)
		// // 		}			
		// // 	}
		// }
		fmt.Printf("\t\tIncomplete: &f2,\n")
		fmt.Printf("\t}\n}\n\n")
	}
}

func (v*CreateCopyVisitor) 	VisitTypeSpec(t *ast.Table, s * ast.Scope, st * ast.TypeSpec) {

	switch  nt := st.Name.(type){
	case * ast.Ident:
		//v.VisitTypeSpecIdent(t,s,st,nt.Name)
		//fmt.Printf("\tName %s\n",nt.Name)
		
		switch  tt := st.Type.(type){

		case *ast.StructType:
			v.VisitStructType(t,s,nt, tt)
		case *ast.InterfaceType:
			v.VisitInterfaceType(t,s,nt, tt)
		case *ast.Ident:
			//v.VisitInterfaceType(t,s,nt, tt)
			fmt.Printf("identifier %+v \n",nt)
		default:
			fmt.Printf("type other type2 %T \n",tt)
			fmt.Printf("  - %#v\n",tt)
		}
		
	default:
		fmt.Printf("name other type2 %T \n",nt)
		fmt.Printf("  - %#v\n",nt)
	}


}

func (v*CreateCopyVisitor) 	VisitStructType(t *ast.Table, s * ast.Scope, name * ast.Ident, st * ast.StructType) {
	//fmt.Printf("st %#v\n", st)
	// Fields.fieldList
	for i,k := range st.Fields.List {
		//fmt.Printf("st %d %#v\n", i, k)
		for ii,kk := range k.Names {
			v.VisitField(t,s, name,st,i,k,ii,kk )
		}
	}
}

func TypeName(f ast.Foo2) string{
	switch  nt := f.(type) {
	case *ast.Ident :
		return fmt.Sprintf("%s\n",nt.Name)
	case *ast.SelectorExpr :
		return fmt.Sprintf("%s.%s",
			IdentName(nt.X),
			IdentName(nt.Sel))
	case *ast.ArrayType :
		return fmt.Sprintf("[]%s\n",IdentName(nt.Elt))
	case *ast.MapType :
		return fmt.Sprintf("map[%s]%s\n",IdentName(nt.Key),TypeName(nt.Value))		
	case *ast.StarExpr:
		return fmt.Sprintf("* %s",IdentName(nt.X))
	default:
		fmt.Printf("unknown %T",nt)
	}
	return "unknown"

}

func IdentName(f ast.Foo2) string{
	switch  nt := f.(type) {
	case *ast.Ident :
		return nt.Name

	case *ast.StarExpr :
		return fmt.Sprintf("* %s",IdentName(nt.X))
	default:
		fmt.Printf("unknown %T",nt)
	}
	return "unknown"
}


func TypeName2(f ast.Foo2) string{
	switch  nt := f.(type) {
	case *ast.Ident :
		return fmt.Sprintf("Type%s",nt.Name)
	case *ast.SelectorExpr :
		return fmt.Sprintf("%s_%s",
			IdentName2(nt.X),
			IdentName2(nt.Sel))
	case *ast.ArrayType :
		return fmt.Sprintf("Array%s",IdentName2(nt.Elt))
	case *ast.MapType :
		return fmt.Sprintf("Map%sTo%s",IdentName2(nt.Key),TypeName2(nt.Value))		
	case *ast.StarExpr:
		return fmt.Sprintf("Pointer%s",IdentName2(nt.X))
	default:
		fmt.Printf("unknown %T",nt)
	}
	return "unknown"

}

func IdentName2(f ast.Foo2) string{
	switch  nt := f.(type) {
	case *ast.Ident :
		return nt.Name
	case *ast.StarExpr :
		return fmt.Sprintf("Star%s",IdentName(nt.X))
	default:
		fmt.Printf("unknown %T",nt)
	}
	return "unknown"
}

func (v*CreateCopyVisitor) VisitField(t *ast.Table, s * ast.Scope, sn * ast.Ident, st * ast.StructType, pos int,
	f * ast.Field,
	pos2 int,
	fn * ast.Ident){
	fmt.Printf("\t\t%s : t.Convert%s( f.%s),\n",fn.Name, TypeName2(f.Type), fn.Name)
}

func (v*CreateCopyVisitor) 	VisitInterfaceType(t *ast.Table, s * ast.Scope, name * ast.Ident, st * ast.InterfaceType) {

}


func (v*CreateCopyVisitor) 	VisitTypeSpecIdent(t *ast.Table, s * ast.Scope, st * ast.TypeSpec, i * ast.Ident){
	
}

func (v*CreateCopyVisitor) 	VisitStructIdent(t *ast.Table, s * ast.Scope, st * ast.StructType, i * ast.Ident){

}

func (v*CreateCopyVisitor) 	VisitFieldTypeStar(
t *ast.Table, s * ast.Scope, sn * ast.Ident, st * ast.StructType, pos int,
	f * ast.Field,
	pos2 int,
	fn * ast.Ident,	
	ft * ast.Type){

}
