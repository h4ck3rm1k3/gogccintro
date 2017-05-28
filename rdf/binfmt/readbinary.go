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
	unknown = iota  // c0 == 0
	intval = iota  // c0 == 1
	strval = iota  // c1 == 2
	intstrval = iota  // c1 == 3 // both string and int values.... mess
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
	IntOffset int // offset of this field type in the original array from disk
	StrOffset int // offset of this field type in the original array from disk

	FieldOffset int // offset of this field type in the int field list
	
	Range [] FieldValueInterface // an array of the range of this field, each unique subject is indexed.

	RangeTypes [] int // array of types of objects in the field
	DomainTypes [] int // array of types of subjects in the field

	RangeTypesCount [] int // array of counts per type of objects in the field
	DomainTypesCount [] int // array of count per type of subjects in the field
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
	
}

func CreateModel(otc int) (*Model){
	return &Model{
		Types: make([]ObjectType,otc),
		FieldNames : make(map[string] int), 
	}
}

func main() {	
	//string_predicates_pairs := ReadStrings("string_predicate2")
	int_predicates := ReadStrings("int_predicate")
	string_predicates := ReadStrings("string_predicate")

	object_types := ReadStrings("object_type")
	otc := len(object_types)
	m := CreateModel(otc)
	
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

	// assign each int predicate a position
	for _,s := range(int_predicates) {
		if _, ok := m.FieldNames[s]; ok {
			//
		} else {
			// create new field
			m.FieldNames[s] = len(m.FieldNames)
		}
	}
	
	for _,s := range(string_predicates) {
		if _, ok := m.FieldNames[s]; ok {
			//
		} else {
			// create new field
			m.FieldNames[s] = len(m.FieldNames)
		}
	}	
	m.Predicates = make( []GlobalFieldType, len(m.FieldNames))
	
	for i,s := range(int_predicates) {
		//pairs := instances(i) do we need instances?
		sn := fmt.Sprintf("%s_pairs",s)
		pairs := ReadPairs(sn)
		fmt.Printf("Int predicate : offset:%d name:%s count:%d\n", i,s, len(pairs))

		// now lookup the field type by name
		if valo, ok := m.FieldNames[s]; ok {

			ft := &m.Predicates[valo]

			if ft.FieldType == unknown {
				ft.FieldType = intval
				ft.Name = s
				ft.IntOffset = i
				ft.FieldOffset = valo
			} else {
				panic(s)
			}
			// now for each pair, look up the domain of the subject and object
					
		} else {panic(s)}
		
	}


	fmt.Printf("predicates %s\n", string_predicates)
	for i,s := range(string_predicates) {
		sn := fmt.Sprintf("s_%s_pairs",s)
		//fmt.Printf("read string pairs %d %s\n", i,sn)
		pairs := ReadPairs(sn)
		values := ReadStrings(s) // read the unique string values in
		fmt.Printf("string %d %s pairs:%d values%d\n", i,sn, len(pairs), len(values))

		// now lookup the field type by name
		if valo, ok := m.FieldNames[s]; ok {
			
			ft := &m.Predicates[valo]

			if ft.FieldType == unknown {
				ft.FieldType = intval
				ft.Name = s
				ft.StrOffset = i
				ft.FieldOffset = valo
			} else if ft.FieldType == intval {
				ft.FieldOffset = valo
				ft.StrOffset = i
				ft.FieldType = intstrval
			} else {
				panic(s)
				
			}
			// now for each pair, look up the domain of the subject and object
					
		} else {panic(s)}
	}

	//fmt.Printf("predicates pairs %s", string_predicates_pairs)
	
	// each instance types

}

