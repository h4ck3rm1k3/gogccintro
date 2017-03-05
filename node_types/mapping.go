package node_types

import (
	"reflect"
	"fmt"
	"github.com/h4ck3rm1k3/gogccintro/models"
	//"github.com/h4ck3rm1k3/gogccintro/tree"
)

func MapOne(v interface{}){
	ve := reflect.ValueOf(v).Elem()
	vt := ve.Type()
	for i := 0; i < vt.NumField(); i++ {
		sf := vt.Field(i)
		f := ve.FieldByIndex([]int{i})

		fmt.Printf(
			//"v=%s\n\tve:%s\n\t"
			//"vt:%s\n"
			"FIELD\ti:%d\n\tsf:%s\n\tf:%n\n\tName:%s\n\tType:%s\n\tPath%s\n",
			//v,
			//ve,
			//vt,
			i,
			sf,
			f,
			sf.Name,
			sf.Type,
			sf.PkgPath,
		)
		//fmt.Printf("typ:%s\n",reflect.TypeOf(sf.Type).String)
		switch sf.Type.Kind() {
		case reflect.Int:
			fmt.Printf("someint1\n")
			break
		case reflect.String:
			fmt.Printf("a string\n")
			break
			
		case reflect.Struct:
			fmt.Printf("some struct\n")
			break
		case reflect.Bool:
			fmt.Printf("some bool\n")
			break
		case reflect.Array:
			fmt.Printf("some array\n")
			break
		case reflect.Slice:
			fmt.Printf("some slice\n")
			break
		case reflect.Interface:
			fmt.Printf("some interface\n")
			break
		default:
			fmt.Printf("some other thing %s\n" , sf.Type.Kind())
		}
		fmt.Printf("-----------------\n")
	}


}
func MapType(v * models.GccTuParserNode, out NodeInterface){
	MapOne(v)
	MapOne(out)
}
