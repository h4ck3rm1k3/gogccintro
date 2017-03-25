package astproto
import (
	"fmt"
	//"github.com/h4ck3rm1k3/gogccintro/fakego/ast"
	//"github.com/h4ck3rm1k3/gogccintro/fakego/ast/proto"
	//"github.com/h4ck3rm1k3/gogccintro/fakego/token"
)

func (t *AddressTable) Check(id string , f2 interface{},f interface{}) {
//	fmt.Printf("Check %s\n", f)
//	fmt.Printf("Check %T\n", f)
	fmt.Printf("Check %s %#v => %#v\n", id, f,f2)

	t.Idsold[id]=f
	t.Idsnew[id]=f
	t.OldNew[f]=f2
	
	// switch v:= f.(type) {
	// default:
	// 	fmt.Printf("other %s\n", v)
	// }
}
