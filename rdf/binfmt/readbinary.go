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
	intval = iota  // c0 == 0
	strval = iota  // c1 == 1
	sizeofint = 8
)
// read in strings
// read in objects using strings
// read in predicate list
func ReadRaw(filename string) (*bytes.Buffer) {
	f:= fmt.Sprintf("../data/%s",filename)
	fmt.Printf("going to open %s ->%s\n", filename, f)
	//data/int_predicate_str_sizes.dat
	buf, err := ioutil.ReadFile(f)
	if err!=nil { panic(err)}
	return bytes.NewBuffer(buf)
}

func ReadInt(filename string) ([]int) {
	b := ReadRaw(filename)
	count := b.Len() / sizeofint
	one := make([]byte,1)
	var res = make([]int,count)
	
	for i:= 0; i< count ; i++ {
		b.Read(one)
		res[i]=int(one[0])
		fmt.Printf("read int %d,%d,%d\n", i, one[0],res[i])
	}
	return res
}

func ReadStrings(filename string) []string{
	// first read the string lengths as ints
	sizes := ReadInt(filename + "_str_sizes.dat")
	count := len(sizes)
	// next read the string value
	data := ReadRaw(filename + "_str.dat")
	strings := make([]string,count) // array of strings the size of the ints
	//pos := 0
	for i,l := range(sizes) {
		one := make([]byte,l)
		s,e := data.Read(one)
		if e != nil {
			panic(e)
		}
		st := string(one)
		fmt.Printf("read string %d, %v, %d, %v %s\n", i, one, l, s, st)
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

func ReadPairs(filename string) {
	// []NodePair
}

func main() {
	object_types := ReadStrings("object_type")
	string_predicates := ReadStrings("string_predicate")
	//string_predicates_pairs := ReadStrings("string_predicate2")
	int_predicates := ReadStrings("int_predicate")

	fmt.Printf("predicates %s", string_predicates)
	//fmt.Printf("predicates pairs %s", string_predicates_pairs)
	
	// each instance types
	for _,s := range(object_types) {
		objects := ReadInstances(s)
		fmt.Printf("predicates %s", objects)
	}
	
	for i,s := range(int_predicates) {
		//pairs := instances(i) do we need instances?
		//pairs := ReadPairs(s)
		fmt.Printf("%d %s", i,s)
	}

	fmt.Printf("f")
}

