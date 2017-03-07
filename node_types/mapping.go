package node_types

import (
	"reflect"
	"fmt"
	"github.com/h4ck3rm1k3/gogccintro/models"
	//"github.com/h4ck3rm1k3/gogccintro/tree"
)

func Indent(indent int) (string) {
	b := make([]rune, indent)
	for i := 0; i < indent; i++ {
		b[i]='\t'
	}
	return string(b)
}

type FieldMap struct {
	
}

type StructMap struct {
	Parent * TypesMap
	Name string
	Names map[string] reflect.StructField
}

type TypesMap struct {
	Types map[string] *StructMap
}

func CreateTypesMap() (*TypesMap){
	return &TypesMap{
		Types: make(map[string] *StructMap),
	}
}


func (t*StructMap) Add(name string, f reflect.StructField){
	t.Names[name]=f
}


func (t*StructMap) Report(indent int){
	indents := Indent(indent)
	fmt.Printf("%s report %s\n", indents, t.Name)
	for name,field := range t.Names {
		fmt.Printf("%s F %s : %s\n", indents, name, field)
	}
}

func (t*StructMap) map_details(indent int, name string, v interface{}){
	ve := reflect.ValueOf(v)
	vt2 := reflect.TypeOf(v)
	if vt2 != nil {
		vt2k := vt2.Kind()
		switch vt2k {
		case reflect.Struct:
			break
		case reflect.Ptr:
			ve = ve.Elem()
			break
		default:
			fmt.Printf("kind %s\n",vt2k)		
		}
	} else {
		//fmt.Printf("null v %s", v)
		return
	}

	vt := ve.Type()
	
	for i := 0; i < vt.NumField(); i++ {
		sf := vt.Field(i)
		//sfi := sf.Interface()
		f := ve.FieldByIndex([]int{i})
		sfi := f.Interface()
		
			//"v=%s\n\tve:%s\n\t"
			//"vt:%s\n"
		// fmt.Printf("%s FIELD i:%d\n",indents,i)
		// fmt.Printf("%s \tsf:%s\n",indents,sf)
		// fmt.Printf("%s \tsfi:%s\n",indents,sfi)
		// fmt.Printf("%s \tf:%s\n",indents,f)
		// fmt.Printf("%s \tName:%s\n",indents,sf.Name)
		t.Add(sf.Name,sf)
		// fmt.Printf("%s \tType:%s\n",indents,sf.Type)
		// fmt.Printf("%s \tTypeKind:%s\n",indents,sf.Type.Kind())

		//fmt.Printf("typ:%s\n",reflect.TypeOf(sf.Type).String)
		switch sf.Type.Kind() {
		case reflect.Int:
			//fmt.Printf("%ssomeint1\n")
			break
		case reflect.String:
			//fmt.Printf("a string\n")
			break
			
		case reflect.Struct:
			//fmt.Printf("some struct\n")
			t.Parent.MapOne(indent +1,sf.Name, sfi)
			break
		case reflect.Bool:
			//fmt.Printf("some bool\n")
			break
		case reflect.Array:
			//fmt.Printf("some array\n")
			break
		case reflect.Slice:
			//fmt.Printf("some slice\n")
			break
		case reflect.Interface:
			//fmt.Printf("some interface\n")
			t.Parent.MapOne(indent +1,sf.Name,sfi)
			break
		case reflect.Ptr:
			//fmt.Printf("some interface\n")
			//MapOne(indent +1,sf.Name,sfi)
			break
		default:
			fmt.Printf("some other thing %s\n" , sf.Type.Kind())
		}
		//fmt.Printf("-----------------\n")
	}
}

func (t*TypesMap)CreateStructMap(name string) (*StructMap){
	return &StructMap{
		Parent: t,
		Name: name,
		Names: make(map[string] reflect.StructField),
	}
}

func (t*TypesMap)MapOne(indent int, name string, v interface{}){
	//indents := Indent(indent)
	sm := t.CreateStructMap(name)
	sm.map_details(indent, name, v)

	sm.Report(indent)

}

func (t*TypesMap)MapType(v * models.GccTuParserNode, out NodeInterface){
	//MapOne(v)
	t.MapOne(1,v.NodeType,out)
}
