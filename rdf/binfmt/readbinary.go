package main
// read rdf file
import (
	"bytes"
	//"strconv"
	//"encoding/binary"
	"fmt"
	//"bufio"
	//"os"
	//"strings"
	"io/ioutil"
	//"container/list"
)
const (
//	unknown = iota  // c0 == 0
	intval = iota  // c0 == 0
	strval = iota  // c1 == 1
	max_field_count = 2
	//intstrval = iota  // c1 == 3 // both string and int values.... mess
	sizeofint = 8
)
// read in strings
// read in objects using strings
// read in predicate list
func ReadRaw(filename string) (*bytes.Buffer) {
	f:= fmt.Sprintf("../data/%s.dat",filename)
	//fmt.Printf("going to open %s ->%s\n", filename, f)
	//data/int_predicate_str_sizes.dat
	buf, err := ioutil.ReadFile(f)
	if err!=nil { panic(err)}
	return bytes.NewBuffer(buf)
}

func ReadBytes(filename string) ([]int) {
	b := ReadRaw(filename)
	count := b.Len()
	one := make([]byte,1)
	var res = make([]int,count)
	
	for i:= 0; i< count ; i++ {
		b.Read(one)
		res[i]=int(one[0])
		//fmt.Printf("read int %d,%d,%d\n", i, one[0],res[i])
	}
	return res
}

func bytes_to_int(one []byte) (int){
	return int(
		int(one[0])<<(8*0) |
			int(one[1])<<(8*1) |
			int(one[2])<<(8*2) |
			int(one[3])<<(8*3) |
			int(one[4])<<(8*4) |
			int(one[5])<<(8*5) |
			int(one[6])<<(8*6) |
			int(one[7])<<(8*7))
}

func ReadInt(filename string) ([]int) {
	b := ReadRaw(filename)
	count := b.Len() / sizeofint
	one := make([]byte,sizeofint)
	var res = make([]int,count)
	
	for i:= 0; i< count ; i++ {
		b.Read(one)
		res[i]= bytes_to_int(one)
	}
	return res
}

func ReadStrings(filename string) []string{
	// first read the string lengths as ints
	sizes := ReadBytes(filename + "_str_sizes")
	count := len(sizes)
	// next read the string value
	data := ReadRaw(filename + "_str")
	strings := make([]string,count) // array of strings the size of the ints
	//pos := 0
	for i,l := range(sizes) {
		one := make([]byte,l)
		_,e := data.Read(one)
		if e != nil {
			panic(e)
		}
		st := string(one)
		//fmt.Printf("read string %d, %v, %d, %v %s\n", i, one, l, s, st)
		//fmt.Printf("read string %s\n",  st)
		strings[i] = st
	}
	return strings
}

func ReadInstances(filename string) []int{
	return ReadInt(filename)
	//return v
}

type NodePair struct {
	From int
	To int
}

func ReadPairs(filename string) ([]NodePair){
	// []NodePair
	b := ReadRaw(filename)
	count := b.Len() / (sizeofint * 2)
	//one := NodePair{}
	var res = make([]NodePair,count)

	one := make([]byte,sizeofint)
	two := make([]byte,sizeofint)	
	
	for i:= 0; i< count ; i++ {
		b.Read(one)
		b.Read(two)
		res[i]=NodePair{
			From: bytes_to_int(one),
			To:  bytes_to_int(two),
		}		
	}
	return res
}

type FieldValueInterface interface {
	// this will resolve to an integer or string field collection.. maybe something else
}

/*
global field type shared for all fields
*/
type GlobalFieldType struct {
	FieldType int// 
	Name string
	Offset [max_field_count]int // offset of this field type in the original array from disk for each array type

	FieldOffset int // offset of this field type in the int field list
	
	Range [] FieldValueInterface // an array of the range of this field, each unique subject is indexed.

	RangeTypes [] int // array of types of objects in the field
	DomainTypes [] int // array of types of subjects in the field

	RangeTypesCount [] int // array of counts per type of objects in the field
	DomainTypesCount [] int // array of count per type of subjects in the field

	// this next field is only for strings
	StringValues []string // for the string type what are the values
	// this is for integers
	Pairs  []NodePair 
}

/*
the field type is inside an object type
*/
type ContainedFieldType struct {
	Global * GlobalFieldType
	Domain [] int // array of the object occuring in the domain, in the object type, used to reduce the cardinality of each predicate
	Range [] int // array of the object occuring in the range

	RangeTypes [] int // array of types of objects in the field
	RangeTypesCount [] int // array of counts per type of objects in the field
	
	Instances [] int // indexed by domain index what is the value index. values of each object.

/*
How to find all the domain indexes for each object?
look up the type of the object, go over each field, lookup the offset of the domain.
we append a contains field object to each node type that contains the offsets for those objects
*/
	ObjectIndex map[int]int // lookup global object id and find the offset in the index, used to find instance data
}

type ObjectType struct {
	Name string
	Offset int // offset of this object type in array
	// each object type represents a type of node, we will start with one type and then refine it into subset
	// a list of int fields (store the fieldid), used to
	FieldNames map[string] int
	Fields [] ContainedFieldType
	// a list of instances (store the object), index, we read this from the file
	Instances [] int
}


func (m *Model)Refine(){
	// given a model we try and create types that are more contained
}

type FieldValue struct {
	FieldType * ContainedFieldType // pointer to the field object description containing name etc.
	Domainoffset int // the offset of this subject in the domain array
	Rangeoffset int // the offset of the object in the range array 
}

/*
TODO : be able to create optimal representation of row of data for a given table
without generating code. Use arrays of fields to describe the format.

maybe use unsafe object , create array of bytes and use fields to get the offsets to the values
*/

/*
description of the object instance, still will require lookups.
*/
type ObjectInstance struct {
	// object type, the type of object we are instance of
	object_type int // contains 
	Fields [] FieldValue // ordered 
	// 
	//incoming_fields  how to find all the incoming pointers?
}

type Model struct {
	// a list of int fields (store the fieldid), used to

	Types [] ObjectType // all the object types
	
	Instances [] ObjectInstance // all object instances, to find the type of that objects, indexed by object id

	FieldNames map[string] int // name lookup
	Predicates [] GlobalFieldType // load predicate lists into here, offset contains the position on the original array
	PredicateTypes [max_field_count]PredicateType

	
}

func CreateModel(otc int) (*Model){
	return &Model{
		Types: make([]ObjectType,otc),
		FieldNames : make(map[string] int), 
	}
}

type PredicateType struct {
	Values []string //:= {,string_predicates}
	Name  string
	Type   int
}


func main() {	
	object_types := ReadStrings("object_type")
	otc := len(object_types)
	m := CreateModel(otc)

	m.PredicateTypes[intval]=PredicateType{
		Name : "int",
		Values:ReadStrings("int_predicate"),
		Type: intval,
	}
	m.PredicateTypes[strval]=PredicateType{
		Name : "string",
		Values :ReadStrings("string_predicate"),
		Type: strval,
	}

	for mi,s := range(object_types) {
		//fmt.Printf("reading type %s\n", s)
		ot := m.Types[mi]
		ot.Name = s
		ot.Offset = mi
		objects := ReadInstances(s)
		ot.Instances = objects//make(Instances [],int)	
		//fmt.Printf("object type %s: len %d\n", s, ic)
		//fmt.Printf("object type %#v\n", ot)		
		// construct an array with all possible predicates for each object
	}

	// load up the global name map, FieldNames
	for _,pt := range(m.PredicateTypes) {
		for _,s := range(pt.Values) {
			if _, ok := m.FieldNames[s]; ok {
				//
			} else {
				// create new field
				m.FieldNames[s] = len(m.FieldNames)
			}
		}
	}

	// create a global predicates
	m.Predicates = make( []GlobalFieldType, len(m.FieldNames))
	
	for pti,pt := range(m.PredicateTypes) {
		for i,s := range(pt.Values) {
			sn := fmt.Sprintf("%s_pairs",s)

	
			// now lookup the field type by name
			if valo, ok := m.FieldNames[s]; ok {
				ft := &m.Predicates[valo]
				
				// ft.FieldType == unknown {
				ft.FieldType = pt.Type
				ft.Name = s
				
				if (pt.Type == strval) {							
					ft.StringValues = ReadStrings(s)

					fmt.Printf("%s predicate, offset:%d name:%s count:%d\n",
						pt.Name,
						i,s, len(ft.StringValues))

				} else if (pt.Type == intval) {							
					pairs := ReadPairs(sn)
					ft.Pairs=pairs
					fmt.Printf("%s predicate, offset:%d name:%s count:%d\n",
						pt.Name,
						i,s, len(ft.Pairs))

				}
				
				
				ft.Offset[pti] = i
				ft.FieldOffset = valo

				// now for each pair, look up the domain of the subject and object
			} else {panic(s)}

		}
	}

	//fmt.Printf("predicates pairs %s", string_predicates_pairs)
	
	// each instance types

}

