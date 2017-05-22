// read rdf file

package main

import (	
	"strconv"
	"encoding/binary"
	"unsafe"
	"fmt"
	"os"
	"strings"
	"io/ioutil"
	//"bytes"
)
//import "encoding/binary"
//import "github.com/wallix/triplestore"
import "github.com/iand/ntriples"

const (  
	intval = iota  // c0 == 0
	strval = iota  // c1 == 1
)
// details on the node itself
type SAttributeValBlob struct {
	BytePosition uintptr // pos in file
	CardinalPos int // pos in hash
	Size uintptr // size of object
	Bytes []byte // bytes of object
	Value int
	Incount int //  incoming count- how many times has this be referenced
	Datatype int // the type of the field
}

func (t * SAttributeValBlob) Val(val int) {
	t.Incount++
}

func (t * SAttributeValBlob) Report() {
	fmt.Printf("%d %d %d %x %d %d %d\n",t.BytePosition,
		t.CardinalPos,
		t.Size,
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
	fmt.Printf("Name %s, Count %d, Max %d, Next %d\n",name, t.Count, t.Max, t.NextPos)
	d := make([]byte,t.NextPos)
	
	//Total.Report()
	for _, v := range t.TheVals {
		// 	//v.Report()
		for i:=0; i <8; i++{
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
			Size : unsafe.Sizeof(val),
			Value : val,
			Incount : 0, // will be incremented
		}
		valo.Bytes = make([]byte,valo.Size)
		binary.LittleEndian.PutUint32(valo.Bytes,uint32(val))
		t.Count++
		if val > t.Max{
			t.Max = val
		}
		valo.Val(val)
		t.NextPos=t.NextPos+uintptr(valo.Size)
		t.TheVals[val] = valo
	}	
}

type StringAttributeVals struct {
	TheVals map[string]*SAttributeValBlob
}

// for each attribute name, if int or string instanciate one of these
type SAttributeNames struct {
	IntVals map[string]*IntAttributeVals
	//IntValTotal IntAttributeVals
	StrVals map[string]*StringAttributeVals
}

func (t * SAttributeNames) Report() {
	for k, v := range t.IntVals {
		//fmt.Printf("Name:%s\t",k)
		v.Report(k)
	}
}

func (t * SAttributeNames) IntVal(key string, val int) {
	//key = "all"
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
				//fmt.Printf("%d TYPE:%s\n", si,o)
				datamap.IntVal(o,si)// peg node type
			} else {
				//fmt.Printf("%d %s OTHER:%s\n", si,p,o)
			}
		} else {
			datamap.IntVal("id",si)
			datamap.IntVal(p,oi)
			//fmt.Printf("%d %s %d\n", si,p,oi)
		}		
	}
	datamap.Report()
}
