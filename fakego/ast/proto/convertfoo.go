package astproto
import (
	"fmt"
	"github.com/h4ck3rm1k3/gogccintro/fakego/ast"
)

func (t* AddressTable) ConvertFoo2(f ast.Foo2) (*Foo2){
	if f!= nil  {
		//fmt.Printf("ConvertFoo2 %s %#v %T\n", f,f,f)

		switch v:= f.(type) {
			
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
