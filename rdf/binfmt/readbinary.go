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

func main() {
	//string_predicates_pairs := ReadStrings("string_predicate2")
	int_predicates := ReadStrings("int_predicate")
	for i,s := range(int_predicates) {
		//pairs := instances(i) do we need instances?
		sn := fmt.Sprintf("%s_pairs",s)
		pairs := ReadPairs(sn)
		fmt.Printf("%d %s %#v\n", i,s, pairs)
	}

	string_predicates := ReadStrings("string_predicate")	
	fmt.Printf("predicates %s\n", string_predicates)
	for i,s := range(string_predicates) {
		sn := fmt.Sprintf("s_%s_pairs",s)
		//fmt.Printf("read string pairs %d %s\n", i,sn)
		pairs := ReadPairs(sn)
		fmt.Printf("string %d %s %#v\n", i,sn, pairs)
	}

	//fmt.Printf("predicates pairs %s", string_predicates_pairs)
	
	// each instance types
	object_types := ReadStrings("object_type")
	for _,s := range(object_types) {
		//fmt.Printf("reading type %s\n", s)
		objects := ReadInstances(s)
		fmt.Printf("object type %s: %#v\n", s,objects)
	}

}

