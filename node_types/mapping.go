package node_types

import (
	"reflect"
	"regexp"
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
	Field reflect.StructField
	TypeName string
}
	
type StructMap struct {
	Parent * TypesMap
	Name string
	Names map[string] * FieldMap
}

type TypesMap struct {
	Types map[string] *StructMap
}

func CreateTypesMap() (*TypesMap){
	return &TypesMap{
		Types: make(map[string] *StructMap),
	}
}

func (s*StructMap) Add(name string, f reflect.StructField, typename string){
	s.Names[name]=&FieldMap{
		TypeName:typename,
		Field:f,
	}
}

func (s*StructMap) map_details(indent int, name string, v interface{}){
	//indents := Indent(indent)
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

		
		if ok,_ := regexp.MatchString("^[_a-z]",sf.Name); ok {
			fmt.Printf("skipping lower case %s\n", sf.Name)
			
		} else {
			//fmt.Printf("not lower case %s\n", sf.Name, ok, err)

			sfi := f.Interface()
			
			//"v=%s\n\tve:%s\n\t"
			//"vt:%s\n"
			//fmt.Printf("%s FIELD i:%d\n",indents,i)
			//fmt.Printf("%s \tsf:%s/%s\n",indents,sf,extract_name(sf))
			//fmt.Printf("%s \tsfi:%s/%s\n",indents,sfi,extract_name(sfi))
			//fmt.Printf("%s \tf:%s/%s\n",indents,f,extract_name(f))
			//fmt.Printf("%s \tName:%s\n",indents,sf.Name)
			s.Add(sf.Name,sf,extract_name(sfi))
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
				s.Parent.MapOne(indent +1,sf.Name, sfi)
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
				s.Parent.MapOne(indent +1,sf.Name,sfi)
				break
			case reflect.Ptr:
				//fmt.Printf("some interface\n")
				//MapOne(indent +1,sf.Name,sfi)
				break
			case reflect.Map:
				//fmt.Printf("some interface\n")
				//MapOne(indent +1,sf.Name,sfi)
				break
			case reflect.Int64:
				//fmt.Printf("some interface\n")
				//MapOne(indent +1,sf.Name,sfi)
				break
			default:
				fmt.Printf("some other thing %s\n" , sf.Type.Kind())
			}
			//fmt.Printf("-----------------\n")

		}
		
	}
}

func (t*TypesMap)CreateStructMap(name string) (*StructMap){
	return &StructMap{
		Parent: t,
		Name: name,
		Names: make(map[string] * FieldMap),
	}
}

func extract_name(v interface{}) (string) {

	t := reflect.TypeOf(v)
	ts := fmt.Sprintf("%s",t)
	//fmt.Printf("val %s\n",v)
	//fmt.Printf("\ttype %s\n",ts)
	return ts

}

func (t*TypesMap)MapOne(indent int, role string, v interface{}){

	// extract the name of the interface
	name := extract_name(v)
	
	if _, ok := t.Types[name]; ok {
		//fmt.Printf("Name exists already: %s",name)
	}else{
		sm := t.CreateStructMap(name)
		sm.map_details(indent, name, v)
		//sm.Report(indent)
		t.Types[name]=sm
	}
}

func (t*TypesMap)MapType(v * models.GccTuParserNode, out NodeInterface){
	//MapOne(v)
	t.MapOne(1,fmt.Sprintf("new %s",v.NodeType),out)
	t.MapOne(1,v.NodeType,v)
}

func (t*TypesMap)Report(){
	fmt.Printf("TypeMaps Report\n")
	for name,field := range t.Types {
		fmt.Printf("report on Name:%s\n",name)
		//fmt.Printf("Field:%s",field)
		field.Report(1)
	}
}

func (s*StructMap) flatten(indents string){
	for name,field := range s.Names {
		if ok,_ := regexp.MatchString("^\\*",field.TypeName); ok {
			fmt.Printf("%s PTR %s: %s\n", indents, name, field.TypeName)
		}else {
			fmt.Printf("%s F %s: %s\n", indents, name, field.TypeName)

			if field.TypeName != "sql.NullInt64" {
				if s2, ok := s.Parent.Types[field.TypeName]; ok {
					s2.flatten(indents)
				}
			}
		}
	}
}

func (s*StructMap) Report(indent int){
	indents := Indent(indent)
	fmt.Printf("%s report %s\n", indents, s.Name)
	s.flatten(indents)
	// for name,field := range s.Names {
	// 	fmt.Printf("%s F %s: %s\n", indents, name, field.TypeName)
	// 	if s2, ok := s.Parent.Types[field.TypeName]; ok {
	// 		for name2,field2 := range s2.Names {
	// 			fmt.Printf("\t%s F2 %s: %s\n", indents, name2, field2.TypeName)
	// 			if s3, ok := s.Parent.Types[field2.TypeName]; ok {
	// 				for name3,field3 := range s3.Names {
	// 					fmt.Printf("\t\t%s F3 %s: %s\n", indents, name3, field3.TypeName)
	// 				}
	// 			}
	// 		}
	// 	}
	// }
}
