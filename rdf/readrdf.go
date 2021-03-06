// read rdf file
package main

import (
	"bufio"
	"bytes"
	"strconv"
	"encoding/binary"
	//"unsafe"
	"fmt"
	"os"
	"strings"
	"io/ioutil"
	"container/list"
	//"bytes"
)
//import "encoding/binary"
//import "github.com/wallix/triplestore"
import "github.com/iand/ntriples"

const (  
	intval = iota  // c0 == 0
	strval = iota  // c1 == 1
	sizeofint = 8
)
// details on the node itself
type SAttributeValBlob struct {
	BytePosition uintptr // pos in file
	CardinalPos int // pos in hash
//	Size uintptr // size of object always sizeofint
	Bytes []byte // bytes of object
	Value int
	Incount int //  incoming count- how many times has this be referenced
	Datatype int // the type of the field
}

type SAttributeValStrBlob struct {
	BytePosition uintptr // pos in file
	CardinalPos int // pos in hash
	Size int // size of object
	Bytes []byte // bytes of object
	Value string
	Incount int //  incoming count- how many times has this be referenced
	Datatype int // the type of the field
}

func (t * SAttributeValBlob) Val(val int) {
	t.Incount++
}

func (t * SAttributeValStrBlob) Val(val string) {
	t.Incount++
}

func (t * SAttributeValBlob) Report() {
	fmt.Printf("%d %d %d %x %d %d %d\n",t.BytePosition,
		t.CardinalPos,
		sizeofint,
		t.Bytes,
		t.Value,
		t.Incount,
		t.Datatype)
}

// depending on the values create a blob object for storage
type IntAttributeVals struct {
	TheVals map[int]*SAttributeValBlob
	Count int
	Max int
	NextPos uintptr // the next position int the block
}

func (t * IntAttributeVals) Report(name string) {
	fmt.Printf("Int Name %s, Count %d, Max %d, Next %d\n",name, t.Count, t.Max, t.NextPos)
	d := make([]byte,t.NextPos)
	
	for _, v := range t.TheVals {
		for i:=0; i < sizeofint; i++{
			d[v.BytePosition+uintptr(i)]=v.Bytes[i]
		}
	}
	fn := fmt.Sprintf("data/%s.dat",name)
	err := ioutil.WriteFile(fn, d, 0644)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%x\n",d)
}

func (t * IntAttributeVals) Val(val int) {
	if (t.TheVals == nil){
		t.TheVals = make(map[int]*SAttributeValBlob)
	}
	if valo, ok := t.TheVals[val]; ok {
		valo.Val(val)
	} else {
		valo := &SAttributeValBlob{
			BytePosition: t.NextPos,
			CardinalPos : t.Count,
			Datatype : intval,
//			Size : sizeofint,
			Value : val,
			Incount : 0, // will be incremented
		}
		valo.Bytes = make([]byte,sizeofint)
		binary.LittleEndian.PutUint32(valo.Bytes,uint32(val))
		t.Count++
		if val > t.Max{
			t.Max = val
		}
		valo.Val(val)
		t.NextPos=t.NextPos+uintptr(sizeofint)
		t.TheVals[val] = valo
	}	
}

type StrAttributeVals struct {
	TheVals map[string]*SAttributeValStrBlob
	Count int
	Max int
	NextPos uintptr // the next position int the block

}

func (t * StrAttributeVals) Report(name string) {

	fmt.Printf("Str Name %s, Count %d, Max %d, Next %d\n",name, t.Count, t.Max, t.NextPos)
	fn := fmt.Sprintf("data/%s_str.dat",name)
	f, err := os.Create(fn)
	if err != nil {	panic(err)}
	w := bufio.NewWriter(f)
	//////////
	fn2 := fmt.Sprintf("data/%s_str_sizes.dat",name)
	f2, err := os.Create(fn2)
	if err != nil {	panic(err)}
	w2 := bufio.NewWriter(f2)
	buf3 := make([]byte,1)	
	///////////////
	
	for _, v := range t.TheVals {
		//
		//fmt.Printf("DEBUGSTR : %#v, %d, %s len:%d buf:%x\n",v,v.Size,v.Value,len(v.Value))
		// write the string
		w.Write(v.Bytes)
		/// write the int
		l := len(v.Value)
		s2 :=byte(l)
		buf3[0]=s2
		w2.Write(buf3)
		////

	}
	w.Flush()
	w2.Flush()
}


func (t * StrAttributeVals) Val(val string) (int) {
	if (t.TheVals == nil){
		t.TheVals = make(map[string]*SAttributeValStrBlob)
	}
	if valo, ok := t.TheVals[val]; ok {
		valo.Val(val)
		return valo.CardinalPos
	} else {
		valo := &SAttributeValStrBlob{
			BytePosition: t.NextPos,
			CardinalPos : t.Count,
			Datatype : strval,
			Size : len(val),
			Value : val,
			Incount : 0, // will be incremented
		}
		valo.Bytes = make([]byte,valo.Size)
		valo.Bytes = []byte(val)
		t.Count++
		if len(val) > t.Max{ // max len
			t.Max = len(val)
		}
		valo.Val(val)
		t.NextPos=t.NextPos+uintptr(valo.Size)
		t.TheVals[val] = valo
		return valo.CardinalPos
	}	
}

type NodePair struct {
	From int
	To int
}

type NodePairArray struct {
	Count int
	Pairs * list.List	
}

func CreateNodePairArray() (* NodePairArray){
	return &NodePairArray{	Count :0, Pairs : list.New()	}
}

func( t * NodePairArray) Report(p string){
	fmt.Printf("Pairs Name %s, Count %d\n",p, t.Count)
	//bytes = make([]byte,t.Count * unsafe.Sizeof(NodePair))
	var b bytes.Buffer
	
	b.Grow(int(uintptr(t.Count) * sizeofint * 2))
	
	for e := t.Pairs.Front(); e != nil; e = e.Next() {
		switch t := e.Value.(type) {
		case NodePair :
			b2 := make([]byte,sizeofint)
			binary.LittleEndian.PutUint32(b2,uint32(t.From))
			b.Write(b2)
			binary.LittleEndian.PutUint32(b2,uint32(t.To))
			b.Write(b2)

		default:
			panic(t)
		}
	}
	fn := fmt.Sprintf("data/%s_pairs.dat",p)
	err := ioutil.WriteFile(fn, b.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
}

func( t * NodePairArray) Add(s int, o int){
	t.Count ++
	t.Pairs.PushBack(NodePair{From:s, To:o})
}

// for each attribute name, if int or string instanciate one of these
type SAttributeNames struct {
	IntVals map[string]*IntAttributeVals

	Pairs map[string]*NodePairArray  // one array per field for node ref fields
	
	//IntValTotal IntAttributeVals
	StrVals map[string]*StrAttributeVals

	//PredicateIds map[string]int
}

func (t * SAttributeNames) Report() {

	// k:= "object_type"
	// v := t.StrVals[k]
	// fmt.Printf("Writing %s\n",k)
	// v.Report(k)
	// //v.ReportSizes(k)
	// return
	
	for k, v := range t.StrVals {
		fmt.Printf("Writing %s\n",k)
		v.Report(k)
		//v.ReportSizes(k)
	}

	for k, v := range t.Pairs {
		v.Report(k)
	}
	for k, v := range t.IntVals {
		//fmt.Printf("Name:%s\t",k)
		v.Report(k)
	}

}

func (t * SAttributeNames) Pair(p string, s int, o int) {

	if (t.Pairs == nil){
		t.Pairs = make(map[string]*NodePairArray)
	}

	if valo, ok := t.Pairs[p]; ok {
		valo.Add(s,o)
	} else {
		valo := CreateNodePairArray()
		t.Pairs[p] = valo
		valo.Add(s,o)
	}

}
	
func (t * SAttributeNames) IntVal(key string, val int) {

	if (t.IntVals == nil){
		t.IntVals = make(map[string]*IntAttributeVals)
	}

	if valo, ok := t.IntVals[key]; ok {
		valo.Val(val)
	} else {
		valo := &IntAttributeVals{
			Count:0,
			Max : 0,
			NextPos : 0,			
		}
		t.IntVals[key] = valo
		valo.Val(val)
	}
	
}

// func (t * SAttributeNames) PredVal(key string) (int){
// /*
// assign each predicate an id
// */
// 	if (t.PredicateIds == nil){
// 		t.PredicateIds = make(map[string]int)
// 	}
// 	if valo, ok := t.PredicateIds[key]; ok {
// 		return valo
// 	} else {
// 		valo := len(t.PredicateIds)
// 		t.PredicateIds[key] = valo
// 		return valo
// 	}

// }

func (t * SAttributeNames) StrVal(key string, val string) (int){

	if (t.StrVals == nil){
		t.StrVals = make(map[string]*StrAttributeVals)
	}

	if valo, ok := t.StrVals[key]; ok {
		return valo.Val(val)
	} else {
		valo := &StrAttributeVals{
			Count:0,
			Max : 0,
			NextPos : 0,			
		}
		t.StrVals[key] = valo
		return valo.Val(val)
	}
	
}

func main(){
	//var buff bytes.Buffer
	ntfile, err := os.Open("test.ntriples")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		os.Exit(1)
	}
	defer ntfile.Close()

	datamap := SAttributeNames{}
	
	prefix := "https://h4ck3rm1k3.github.io/gogccintro/gcc/ontology/2017/05/20/gcc_compiler.owl#"
	stype := "http://www.w3.org/1999/02/22-rdf-syntax-ns#type"
	named := "http://www.w3.org/2002/07/owl#NamedIndividual"
	count := 0
	r := ntriples.NewReader(ntfile)

	
	for triple, err := r.Read(); err == nil;  triple, err = r.Read() {
		count++
		s := triple.Subject()
		s = strings.Replace(s,prefix,"",1)
		si, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		
		o := triple.Object()
		if o == named {
			continue
		}
		p := triple.Predicate()
		o = strings.Replace(o,prefix,"",1)
		p = strings.Replace(p,prefix,"",1)
		
		oi, err := strconv.Atoi(o)
		if err != nil {
			datamap.IntVal("id",si)
			//panic(err)
			if p == stype {
				fmt.Printf("%d TYPE:%s\n", si,o)
				datamap.IntVal(o,si)// peg node type as predicate...

				datamap.StrVal("object_type",o) // other string value
			} else {
				//fmt.Printf("%d %s OTHER:%s\n", si,p,o)
				oid := datamap.StrVal(p,o) // other string value
				// the  object id are the indexes to the file where these will be output
				sp2 := fmt.Sprintf("s_%s",p)
				datamap.Pair( sp2, si,oid)  // use the cardinality of the string for the index

				datamap.StrVal("string_predicate",p) // other string value
				datamap.StrVal("string_predicate2",sp2) // other string value
			}
		} else {
			datamap.StrVal("int_predicate",p) // other string value
			datamap.IntVal("id",si)
			
			//datamap.IntVal(p,oi)
			datamap.Pair(p, si,oi)
			//fmt.Printf("%d %s %d\n", si,p,oi)
		}		
	}
	datamap.Report()
}
