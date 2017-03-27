package astproto
import (
	"fmt"
	"github.com/h4ck3rm1k3/gogccintro/fakego/ast"
)

func (t* AddressTable) ConvertFoo2(f ast.Foo2) (*Foo2){
	if f!= nil  {
		//fmt.Printf("ConvertFoo2 %s %#v %T\n", f,f,f)

		switch v:= f.(type) {
		case *ast.Object:
			t2 := NodeType_OBJECT
			o  := t.ConvertObject(v)
			r  := &Foo2{
	 			Type: &t2,
	 			Object:o,
			}
			fmt.Printf("chk object:%s\n", o.String())
			fmt.Printf("foo2 object:%s\n", r.String())
			return r
		default:
			fmt.Printf("ConvertFoo2 %T\n val :%#v\n str:%s\n", f,f,f)
			fmt.Printf("unknown %T\n", v)
		}
	}	
	return nil
}
