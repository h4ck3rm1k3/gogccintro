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
	Role string
	Name string
	Path string
}
	
type StructMap struct {
	Parent * TypesMap
	Name string
	Names map[string] * FieldMap
}

type TypesMap struct {
	Types map[string] *StructMap
	Unique map[string] bool
}

func CreateTypesMap() (*TypesMap){
	return &TypesMap{
		Types: make(map[string] *StructMap),
		Unique : make(map[string] bool),
	}
}

func(s*StructMap) ShortName() string{
	re := regexp.MustCompile("^\\*?[_A-Za-z]+\\.")
	val := re.ReplaceAllString(s.Name,"")
	return val
}

func(s*FieldMap) ShortName() string{
	re := regexp.MustCompile("^\\*?[_A-Za-z]+\\.")
	val := re.ReplaceAllString(s.Name,"")
	return val
}

func(s*FieldMap) TopPath() string{
	// re := regexp.MustCompile("^([_A-Za-z]+)\\.([_A-Za-z]+)\\.")
	// val := re.ReplaceAllString(s.Path,"$1.$2")
	re := regexp.MustCompile("^[_A-Za-z]+\\.([_A-Za-z]+)")
	val := re.FindStringSubmatch(s.Path)
	for _, v := range val { 
		//fmt.Printf("CHECK %s\n",v)
		return v
	}
	//fmt.Printf("SKIP %s\n",s.Path)
	return s.Path
}

func(s*FieldMap) TopPathField() string{
	re := regexp.MustCompile("^[_A-Za-z]+\\.([_A-Za-z]+)")
	groups := re.FindStringSubmatch(s.Path)
	v := groups[1]
	return v

}

func(s*FieldMap) ShortTypeName() string{
	re := regexp.MustCompile("([_A-Za-z]+).([_A-Za-z]+)")
	groups := re.FindStringSubmatch(s.TypeName)
	v := groups[2]
	return v
}

func(s*FieldMap) ShortPath() string{
	return s.Path
}

func(s*FieldMap) PathLength() int{
	var c int
	c = 0
	for _, char := range s.Path { 
		if char == '.' {
			c++
		}
	}
	return c
}

func(s*StructMap) LocalName() string{
	re := regexp.MustCompile("^(\\*)node_types\\.")
	val := re.ReplaceAllString(s.Name,"$1")
	return val
}
	
func (s*StructMap) Add(name string, f reflect.StructField, typename string, role string, name2 string, path string){
	if _, ok := s.Names[name]; ok {
		//fmt.Printf("duplicate %s : %s\n", name,s.Names[name])
	} else{
		//fmt.Printf("adding %s : %s\n", name,name2)
		s.Names[name]=&FieldMap{
			TypeName:typename,
			Field:f,
			Role: role,
			Name: name2,
			Path: path,
		}
	}
}

func (s*StructMap) map_details(indent int, role string, name string, v interface{}){
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
			s.Add(sf.Name,sf,extract_name(sfi), name, role, "to")
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

func (t*TypesMap)MapOne(indent int, role string, v interface{}) (*StructMap){

	// extract the name of the interface
	name := extract_name(v)
	
	if sm, ok := t.Types[name]; ok {
		//fmt.Printf("Name exists already: %s",name)
		return sm
	}else{
		sm := t.CreateStructMap(name)
		sm.map_details(indent, role, name, v)
		//sm.Report(indent)
		t.Types[name]=sm
		return sm
	}
}

func (t*TypesMap) unify(from * StructMap, to * StructMap) {
	unique := make(map[string] bool)	
	fmt.Printf("func (to %s) Mapping%s(from %s){\n", to.LocalName(),from.ShortName(), from.Name)
	for name,_ := range from.Names {
		if toField, ok := to.Names[name]; ok {
			var new string		
			if toField.PathLength() > 0 {
				new = fmt.Sprintf("%s.Mapping%s(from)\n",
					toField.TopPath(),
					from.ShortName(),
				)
			} else {
				new = fmt.Sprintf("Mapping%s(%s.%s,from.%s)\n",
					name,
					toField.TopPath(), //to.. .
					name, // Refs ...
					name,
				)
			}
			if _, ok := unique[new]; ok {	
			} else {
				fmt.Printf("\t%s", new)
				unique[new]=true
				}
		} else {
			//fmt.Printf("%s.%s<=?\n", fromField.Path,name)
		}
	}	
	// end of func
	fmt.Printf("}\n")
}

func (t*TypesMap) GenerateConstructor(from * StructMap, to * StructMap) {
	unique := make(map[string] bool)
	fmt.Printf("func (t * TUFile) Create%s(from %s ) %s {\n\treturn &%s{",
		to.ShortName(), // create%2
		from.Name, // from 
		to.LocalName(), // return ttype
		to.ShortName(), // create type
	)
	
	for name,_ := range from.Names {
		if toField, ok := to.Names[name]; ok {
			var new string
			// subobject
			if toField.PathLength() > 0 {

				new = fmt.Sprintf("%s : t.Create%s(from),\n",
					toField.TopPathField(),
					toField.TopPathField(),
					//from.ShortName(),
				)
			} else {
				//                1field : 2FieldCreate(from.3Field)
				//				fmt.Printf("DEBUG:%v\n",toField)
				//fmt.Printf("DEBUG:%s\n",toField.ShortTypeName())
				new = fmt.Sprintf("%s: t.CreateRef%s(from.%s),\n",
					name,
					toField.ShortTypeName(), //to.. .
					name, // Refs ...
				)
			}
			if _, ok := unique[new]; ok {	
			} else {
				fmt.Printf("\t\t%s", new)
				unique[new]=true
				}
		}else{
			//fmt.Printf("%s.%s<=?\n", fromField.Path,name)
		}
	}	
	// end of func
	fmt.Printf("\t}\n}\n")
}

func (t*TypesMap) GenerateRefConstructor(from * StructMap, to * StructMap) {
	//unique := make(map[string] bool)
	fmt.Printf("%sMap map [int64] * %s\n",to.ShortName(),to.ShortName())
	fmt.Printf("%sMap : make(map[int64] *%s),\n",to.ShortName(),to.ShortName())
	//fmt.Printf("func (t * TUFile) Lookup%s(id int64 ) (*%s,bool) {\n\treturn t.%sMap[id]\n}\n",to.ShortName(),to.ShortName(),to.ShortName())
		
	fmt.Printf("func (t * TUFile) CreateRef%s(id sql.NullInt64 ) %s {\n",
		to.ShortName(), // create%2
//		from.Name, // from 
		to.LocalName(), // return type
	)

	fmt.Printf("\tif id.Valid {\n\t\tif node,ok := t.%sMap[id.Int64]; ok {\n\t\t\treturn node\n\t\t} else {\n\t\t\treturn Create%s(t.LookupGccNode(id.Int64))\n\t\t}\t\n\t}\n\treturn nil\n}\n\n",to.ShortName(),to.ShortName())
}

func (t*TypesMap)GenerateFields(from * StructMap, to * StructMap){

	for name,fromField := range from.Names {
		if toField, ok := to.Names[name]; ok {

			var new string
			
			if toField.PathLength() > 0 {
			} else {
				//fmt.Printf("Mapping%s(from %s){\n", to.LocalName(),from.ShortName(), from.Name)
				
				new = fmt.Sprintf("func Mapping%s(to NodeInterface, from %s) { to.Load(from) }\n",
					//toField.TypeName,
					name,
					fromField.TypeName,
					//name,

				)

			}
			
			if _, ok := t.Unique[new]; ok {
				
			} else {
				fmt.Printf("%s", new)
				t.Unique[new]=true
				}
		}else{
			//fmt.Printf("%s.%s<=?\n", fromField.Path,name)
		}
	}

}

func (t*TypesMap)MapType(v * models.GccTuParserNode, out NodeInterface){
	//MapOne(v)
	to := t.MapOne(1,fmt.Sprintf("new %s",v.NodeType),out)
	from := t.MapOne(1,v.NodeType,v)

	from.flatten("\t","from",from)
	to.flatten("\t","to",to)

	//t.GenerateFields(from,to)
	//t.unify(from,to)
	//t.GenerateConstructor(from, to)
	t.GenerateRefConstructor(from, to)
	
}

func (t*TypesMap)Report(){
	
	// fmt.Printf("TypeMaps Report\n")
	// for _,field := range t.Types {
	// 	//fmt.Printf("report on Name:%s\n",name)
	// 	//fmt.Printf("Field:%s",field)
	// 	field.Report(1)
	// }
}

func (s*StructMap) flatten(indents string, path string, top*StructMap){
	for name,field := range s.Names {
		if ok,_ := regexp.MatchString("^\\*",field.TypeName); ok {
			//fmt.Printf("%s %s PTR %s: %s %s %s\n", indents, path, name, field.TypeName, field.Role, field.Name)
			top.Add(name,field.Field,field.TypeName, field.Role, field.Name, path)
		}else {
			//fmt.Printf("%s %s F %s: %s %s %s\n", indents, path, name, field.TypeName, field.Role, field.Name)
			top.Add(name,field.Field,field.TypeName, field.Role, field.Name, path)
			
			if field.TypeName != "sql.NullInt64" {
				if s2, ok := s.Parent.Types[field.TypeName]; ok {
					s2.flatten(indents, fmt.Sprintf("%s.%s",path, name), top)
				}
			}
		}
	}
}

func (s*StructMap) Report(indent int){
	indents := Indent(indent)
	//fmt.Printf("%s report %s\n", indents, s.Name)
	s.flatten(indents,"to",s)

	// for name,field := range s.Names {
	// 	fmt.Printf("%s F %s: TYPE:%s ROLE:%s NAME:%s PATH:%s\n", indents, name, field.TypeName, field.Role, field.Name, field.Path)
	// }
	
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
