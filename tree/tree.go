package tree

import (
	"sync"
	"fmt"
	"encoding/json"
	"strconv"
	"io/ioutil"
	"reflect"
	"github.com/polaris1119/bitmap"
	"github.com/h4ck3rm1k3/gogccintro/models"
)

type TreeMap struct {
	Bitmap * bitmap.Bitmap
	Nodes map[int] * models.GccTuParserNode
	Mutex sync.RWMutex
}

func (t * TreeMap) FindName(n * models.GccTuParserNode) string{
	if n.NodeType == "identifier_node" {
		return n.AttrsString
	} else if n.RefsName.Valid {
		return t.FindName(t.Nodes[int(n.RefsName.Int64)])
	} else {
		return t.FindConstInt(n) // just return  the int value
	}
}

func (t * TreeMap) FindConstInt(n * models.GccTuParserNode) string{
	if n.NodeType == "integer_cst" {
		return n.AttrsTypeSize // misnamed
	} else {
		return "TODO2"
	}
}

func NewTreeMap(size int) *TreeMap {
	return &TreeMap{
		Mutex: sync.RWMutex{},
		Bitmap: bitmap.NewBitmapSize(size),
		Nodes : make(map[int] * models.GccTuParserNode),
	}
}

// return true if we are the first to set this bit, false otherwise
func (this* TreeMap) SetBitFirst( pos uint64 ) bool {
	if this.Bitmap.GetBit(pos)==0  {
		this.Mutex.RLock()
		this.Bitmap.SetBit(pos,1)
		this.Mutex.RUnlock()
		return true
	}
	return false
}

type Receiver interface {
	StartGraph()
	StartNode(n * models.GccTuParserNode)
	ReferenceNode(n * models.GccTuParserNode, name string, o * models.GccTuParserNode)
	ReferenceAttribute(n * models.GccTuParserNode, name string, value string)
	EndNode(n * models.GccTuParserNode)
	EndGraph()
}


func (this * TreeMap) ResolveReferences(r Receiver ){
	//fmt.Printf("start resolve %v\n", r)
	r.StartGraph()
	//fmt.Printf("start resolve %v\n", r)
	for _,v := range this.Nodes {
		r.StartNode(v)
		objValue := reflect.ValueOf(v).Elem()		
		for _,fn := range StrFields {		
			field := objValue.FieldByName(fn).String()
			if field != "" {
				r.ReferenceAttribute(v, fn, field)				
			}
		}	
		
		for _,fn := range RefFields {
			field := objValue.FieldByName(fn)
			valid := field.FieldByName("Valid").Bool()
			rid := field.FieldByName("Int64").Int()			
			if valid {
				o:=this.Nodes[int(rid)]
				if o != nil {
					//s2:=this.FindName(o)
					//fmt.Printf("\treflect %d %s %v %d %v %s\n", k,fn,rid,o.NodeID, o.NodeType, s2)
					r.ReferenceNode(v, fn, o)				
				} else{
					//fmt.Printf("\treflect %d %s %v NULL\n", k,fn,rid)
				}
			}			
		}
		
		for _,fn := range StrRefFields {
			field := objValue.FieldByName(fn)
			rid := field.String()
			if rid != "" {
				i, err := strconv.Atoi(rid)			
				if err == nil {
					o:=this.Nodes[i]
					if o != nil {
						//s2:=this.FindName(o)
						//fmt.Printf("\treflect2 %d %s %v %d %v %s\n", k,fn,rid,o.NodeID, o.NodeType, s2)
						r.ReferenceNode(v, fn, o)
					} 
				}
			}
		}
		r.EndNode(v)
	}
	r.EndGraph()
}

func NewTreeMapFromFile(filename string) *TreeMap {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("err %v\n", err)
		panic(err)
	}
	//fmt.Printf("content %v\n", content)

	treemap := NewTreeMap(0)
	//Do something
	err = json.Unmarshal(content, &treemap.Nodes)
	if err != nil {
		fmt.Printf("err %v\n", err)
		panic(err)
	}
	return treemap
}

var StrFields = []string{
		"AttrsString",
		"AttrsTypeName",
		"AttrsNote",
		"AttrsTypeSize",
		"AttrsAddr",
	}

var StrRefFields = []string{
		"AttrsType",
	}

var RefFields = []string{
		"RefsArgs",
		//"RefsScpe",
		"RefsArgt",
		"RefsBody",
		"RefsBpos",
		"RefsChan",
		"RefsChain",
		"RefsCnst",
		"RefsCond",
		"RefsCsts",
		"RefsDecl",
		"RefsDomn",
		"RefsE",
		"RefsElts",
		"RefsExpr",
		"RefsFlds",
		"RefsFn",
		"RefsIdx",
		"RefsInit",
		"RefsLabl",
		"RefsLow",
		"RefsMax",
		"RefsMin",
		"RefsMngl",
		"RefsName",
		"RefsOp0",
		"RefsOp1",
		"RefsOp2",
		"RefsPrms",
		"RefsPtd",
		"RefsPurp",
		"RefsRefd",
		"RefsRetn",
		"RefsSize",
		"RefsType",
		"RefsUnql",
		"RefsVal",
		"RefsValu",
		"RefsVars",
	}

