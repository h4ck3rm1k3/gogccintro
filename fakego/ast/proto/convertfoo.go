package astproto
import (
	"fmt"
	"github.com/h4ck3rm1k3/gogccintro/fakego/ast"
)

func (t* AddressTable) ConvertFoo2(f ast.Foo2) (*Foo2){
	if f!= nil  {
		//fmt.Printf("ConvertFoo2 %s %#v %T\n", f,f,f)

		switch v:= f.(type) {
			
			
		case *ast.TypeSpec:
			t2 := NodeType_TYPESPEC
			o  := t.ConvertTypeSpec(v)
			r  := &Foo2{
	 			Type: &t2,
	 			Typespec:o,
			}
			return r
			
		case *ast.AssignStmt:
			t2 := NodeType_ASSIGNSTMT
			o  := t.ConvertAssignStmt(v)
			r  := &Foo2{
	 			Type: &t2,
	 			Assignstmt:o,
			}
			return r

		case *ast.Deferred:
			t2 := NodeType_DEFERRED
			o  := t.ConvertDeferred(v)
			r  := &Foo2{
	 			Type: &t2,
	 			Deferred:o,
			}
//			fmt.Printf("chk object:%s\n", o.String())
//			fmt.Printf("foo2 object:%s\n", r.String())
			return r
			
		case *ast.FieldDeferred:
			t2 := NodeType_FIELDDEFERRED
			o  := t.ConvertFieldDeferred(v)
			r  := &Foo2{
	 			Type: &t2,
	 			Fielddeferred:o,
			}
			//fmt.Printf("chk object:%s\n", o.String())
			//fmt.Printf("foo2 object:%s\n", r.String())
			return r
			
		case *ast.FuncDecl:
			t2 := NodeType_FUNCDECL
			o  := t.ConvertFuncDecl(v)
			r  := &Foo2{
	 			Type: &t2,
	 			Funcdecl:o,
			}
			return r
			
		case *ast.ValueSpec:
			t2 := NodeType_VALUESPEC
			o  := t.ConvertValueSpec(v)
			r  := &Foo2{
	 			Type: &t2,
	 			Valuespec:o,
			}
			return r
			
		case *ast.Object:
			t2 := NodeType_OBJECT
			o  := t.ConvertObject(v)
			r  := &Foo2{
	 			Type: &t2,
	 			Object:o,
			}
//			fmt.Printf("chk object:%s\n", o.String())
//			fmt.Printf("foo2 object:%s\n", r.String())
			return r

			
		default:
			fmt.Printf("ConvertFoo2 Unknown : %T\n val :%#v\n str:%s\n", f,f,f)
		}
	}	
	return nil
}
