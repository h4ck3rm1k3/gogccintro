package main

import (
	"sync"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"io/ioutil"
	//"bytes"

)

type Stat struct {
	TheCount int
}

func (t *Stat) AppendTo(t2 *Stat) {
	t2.TheCount = t2.TheCount + t.TheCount 
}
	
func (t *Stat) Count() {
	t.TheCount++
}

func (t *Stat) Report(indent int) {
	for i := 0; i < indent; i++ {
		fmt.Printf("\t")
	}
	fmt.Printf("Count %d\n", t.TheCount)
}

type SValue struct {
	Count Stat

}

func (t *SValue) Report() {
	t.Count.Report(7)
}

func (t *SValue) Val(v string) {
	t.Count.Count()
}

// the values of the attribute, maybe just the chars or some abstraction
type SAttributeValString struct {
	Count   Stat
	TheVals map[string]*SValue
}

func (t *SAttributeValString) Report() {
	t.Count.Report(6)
	fmt.Printf("\t\t\t\t\t\t\tlen vals %d:\n", len(t.TheVals))
	for k, v := range t.TheVals {
		fmt.Printf("\t\t\t\t\t\tValue %s \n", k)
		if t.Count.TheCount > 1 {
			v.Report()
		}
	}
}

func (t *SAttributeValString) Vals(vals *[]string) {
	t.Count.Count()
	key := strings.Join(*vals, "|")
	//t.TheVals[val].Val(val)
	if t.TheVals == nil {
		t.TheVals = make(map[string]*SValue)
	}
	if val, ok := t.TheVals[key]; ok {
		val.Val(key)
	} else {
		o := &SValue{}
		t.TheVals[key] = o
		o.Val(key)
	}
}

type SAttributeVals struct {
	Count   Stat
	TheVals map[int]*SAttributeValString
}

func (t *SAttributeVals) Report() {
	t.Count.Report(5)
	fmt.Printf("\t\t\t\t\t\tlen vals %d:\n", len(t.TheVals))
	for k, v := range t.TheVals {
		fmt.Printf("\t\t\t\t\tLen %d\n", k)
		v.Report()
	}
}

func (t *SAttributeVals) Vals(vals *[]string) {
	t.Count.Count()
	//t.TheVals[len(*vals)].Vals(vals)
	key := len(*vals)
	if t.TheVals == nil {
		t.TheVals = make(map[int]*SAttributeValString)
	}
	if val, ok := t.TheVals[key]; ok {
		val.Vals(vals)
	} else {
		o := &SAttributeValString{}
		t.TheVals[key] = o
		o.Vals(vals)
	}
}

// simple node type 
type SAttributeNames struct {
	Count   Stat
	TheVals map[string]*SAttributeVals
}

func (t *SAttributeNames) Report() {
	t.Count.Report(4)
	fmt.Printf("\t\t\t\tlen attrs %d:\n", len(t.TheVals))
	for k, v := range t.TheVals {
		fmt.Printf("\t\t\t\tAttrName: %s \n", k)
		v.Report()
	}
}

func (t *SAttributeNames) Vals(n *TestNode, t2 *TestConsumer, t3 *SNodeType) {
	for key, v := range n.Vals {
		//t.TheVals[k].Vals(&v)
		if t.TheVals == nil {
			t.TheVals = make(map[string]*SAttributeVals)
		}
		if val, ok := t.TheVals[key]; ok {
			val.Vals(&v)
		} else {
			o := &SAttributeVals{}
			t.TheVals[key] = o
			o.Vals(&v)
		}
	}
	// now lets make the pairs
	for key, _ := range n.Vals {

		// peg each node type
		//t2.CoOccurance.Fields(n.NodeType,key)
		
		for key2, _ := range n.Vals {
			if key2 != key {
				if strings.Compare(key, key2) > 0 {
					t3.CoOccurance.Fields(key, key2)
				}
			}
		}
	}
}

// different vectors of attributes
type SAttributeVector struct {
	Count   Stat
	TheVals map[string]*SAttributeNames
}

func (t *SAttributeVector) Report(nodetype string, attrcount int) {
	t.Count.Report(3)
	fmt.Printf("\t\t\tlen attrlist %d:\n", len(t.TheVals))
	for k, _ := range t.TheVals {
		fmt.Printf("\t%s\t%d\tattrlist:%s \n", nodetype, attrcount, k)
		//v.Report()
	}
}

func (t *SAttributeVector) Vals(n *TestNode, t2 *TestConsumer, t3 *SNodeType) {
	t.Count.Count()
	keys := make([]string, 0, len(n.Vals))
	for k := range n.Vals {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	key := strings.Join(keys, "|")
	//t.TheVals[key].Vals(n)
	if t.TheVals == nil {
		t.TheVals = make(map[string]*SAttributeNames)
	}
	if val, ok := t.TheVals[key]; ok {
		val.Vals(n, t2, t3)
	} else {
		o := &SAttributeNames{}
		t.TheVals[key] = o
		o.Vals(n, t2, t3)
	}
}

type SNodeType struct {
	Count          Stat
	AttributeCount map[int]*SAttributeVector // Attribute County

	AttrNames map[string]int // simple count
	CoOccurance *CoOccurance
}

// copy all the data to the other 
func (t *SNodeType) AppendTo(t2 *SNodeType) {

	t.Count.AppendTo(&t2.Count)

	t.CoOccurance.AppendTo(t2.CoOccurance)
	
	// copy the attrnames
	for k, val1 := range t.AttrNames {
		if (k[0]>= '0' && k[0] <= '9'){
		}else {
			//fmt.Printf("\t\tattr\t%s\t%d\n", k,v)
			if val, ok := t2.AttrNames[k]; ok {
				t2.AttrNames[k] =val + val1
			} else {
				t2.AttrNames[k] = val1
			}
		}
	}
}
	
func (t *SNodeType) Report(nodetype string) {
	fmt.Printf("Report on node type %s:\n", nodetype)
	
	t.Count.Report(2)
	//fmt.Printf("\t\tlen attrcount %d:\n", len(t.AttributeCount))
	fmt.Printf("\t\tlen attrnames %d:\n", len(t.AttrNames))


	// sort attribute names desc
	n := map[int][]string{}
	
	for k, v := range t.AttrNames {

		if (k[0]>= '0' && k[0] <= '9'){
		} else {
			//fmt.Printf("\t\tattr\t%s\t%d\n", k,v)
			n[v] = append(n[v], k)
		}
	}

	var a []int
	for k := range n {
		a = append(a, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	for _, k := range a {
		for _, s := range n[k] {
			fmt.Printf("ATTR : %s, %d\n", s, k)
		}
	}
	
	t.CoOccurance.Report()


	// for k, _ := range t.AttributeCount {
	// 	fmt.Printf("\t\tattrcount %d:\n", k)
	// 	//v.Report(nodetype, k)	
	// }
}

func (t *SNodeType) Node(n *TestNode, t2 *TestConsumer) {
	t.Count.Count()
	key := len(n.Vals)
	if t.AttributeCount == nil {
		t.AttributeCount = make(map[int]*SAttributeVector)
	}
	if val, ok := t.AttributeCount[key]; ok {
		val.Vals(n, t2, t)
	} else {
		o := &SAttributeVector{

		}
		t.AttributeCount[key] = o
		o.Vals(n, t2, t)
	}
	//t.AttributeCount[len(n.Vals)].Vals(n)

	for k, _ := range n.Vals {
		if val, ok := t.AttrNames[k]; ok {
			t.AttrNames[k] =val + 1		
		} else {
			t.AttrNames[k] =1
		}
	}
	
}

type ProgramData struct {
	Count         Stat
	NodeTypes     map[string]*SNodeType

	// a summary node type of all the types rolled togehter
	Summary       * SNodeType
	
	// parser stats
	Transitions   map[int]map[int]int
	States        map[int]int	
}

type TestConsumer struct {
	Lock          sync.RWMutex//= sync.RWMutex{}
	Filename      string
	Data ProgramData
	//Cache *       Cache
}

func (t *TestConsumer) StateUsage(from int) {
	t.Lock.Lock()
	defer t.Lock.Unlock()
	old := t.Data.States[from]
	t.Data.States[from]=old + 1	
}

func (t *TestConsumer) StateTransition(from int, to int) {
	t.Lock.Lock()
	defer t.Lock.Unlock()

	if val, ok := t.Data.Transitions[from]; ok {
		old := val[to]
		val[to]=old + 1
	} else {
		v := make(map[int]int)
	        v[to]=1
		t.Data.Transitions[from] = v
	}
}

func (t *TestConsumer) Write() {
	b, err := t.Data.MarshalJSON()
	if err != nil {
		panic(err)
	}	
	//DEBUG : fmt.Printf("NodeTypes:%s\n",b)

	err = ioutil.WriteFile(fmt.Sprintf("%s.json",t.Filename), b, 0644)
	if (err != nil){panic(err)}
}

func (t *TestConsumer) Read() {

	b,err := ioutil.ReadFile(fmt.Sprintf("%s.json",t.Filename))
	if (err != nil){panic(err)}
	
	err = t.Data.UnmarshalJSON(b)
	if err != nil {
		panic(err)
	}	
	//DEBUG :
	//fmt.Printf("Read:%s\n",t.Data)
}

func (t *TestConsumer) Rollup() {
	// roll up the nodetypes into the summary
	// Summary=sum(NodeTypes)

	//  append each type to the summary
	for _, v := range t.Data.NodeTypes {
		v.AppendTo(t.Data.Summary)
	}
	fmt.Printf("total report:\n")
	t.Data.Summary.Report("total")
	t.Data.Summary.ReportDisjoint("total")
}
	
func (t *TestConsumer) Report() {

	t.Rollup()
	
	for k, v := range t.Data.States {
		fmt.Printf("\tS %20s %10d\n", StateLookup[Token(k)], v)
	}
	// report on transitions

	to_count := make(map[int]int)
	
	for k, v := range t.Data.Transitions {
		c:=0
		for k2, v2 := range v {
			fmt.Printf("\tT %20s %20s %10d\n", StateLookup[Token(k)], StateLookup[Token(k2)], v2)
			c = c+ v2
			to_count[k2] = to_count[k2] + v2
		}
		//
		fmt.Printf("\tAS %20s %10d\n", StateLookup[Token(k)], t.Data.States[k]/c)
		
	}
	for k, v := range to_count {
		fmt.Printf("\tTAS %20s %10d\n", StateLookup[Token(k)], t.Data.States[k]/v)
	}
		
	t.Data.Count.Report(1)
	
	//fmt.Printf("Report %#v\n", t.NodeTypes)
	for k, v := range t.Data.NodeTypes {
		fmt.Printf("\tNT %s %d\n", k, v.Count.TheCount)
		v.Report(k)
		v.ReportDisjoint(k)
	}
	
	//t.Cache.Add(x,b)
	//t.Cache.Close()
}

func (t *TestConsumer) NodeImp(n *TestNode) {

	t.Lock.Lock()
	defer t.Lock.Unlock()

	t.Data.Count.Count()

	if t.Data.NodeTypes == nil {
		t.Data.NodeTypes = make(map[string]*SNodeType)
	}
	if val, ok := t.Data.NodeTypes[n.NodeType]; ok {
		val.Node(n, t)
	} else {
		o := &SNodeType{
			CoOccurance: &CoOccurance{},
			AttrNames: make(map[string]int),
		}
		t.Data.NodeTypes[n.NodeType] = o
		o.Node(n, t)
	}
}

type TestNode struct {
	Filename string
	NodeType string
	NodeId   string
	Vals     map[string][]string
	//StrVals     map[string]string
	//IntVals     map[string]int
}

func (t *TestNode) Finish(Parent TreeConsumer) {
	Parent.Node(t)
}

func (t *TestNode) Report() {

	l := len(t.Vals)
	fmt.Printf("type %s size:%d\n", t.NodeType, l)
	for k, v := range t.Vals {
		fmt.Printf("\tkv '%s' -> '%s'\n", k, v)
		//v.Report()
	}
}

func (t *TestNode) SetAttr(name string, vals []string) {
	if name == "lngt:" {
		i, err := strconv.Atoi(vals[0])
		if err != nil {
			panic(vals[0])
		}
		// truncate the strg to the right length
		d := t.Vals["strg:"]
		if len(d) > 0 {
			//fmt.Printf("check %s\n",d)
			d2 := d[0]
			//fmt.Printf("check d2 %s\n",d2)
			if len(d2) > i {
				//fmt.Printf("Trunc %s to %d\n",d,i)
				t.Vals["strg:"][0] = d2[0:i]
			}
		}
	}
	if len(name) == 0 {
		panic("null name")
	}
	if name[0] == '\t' || name[0] == ' ' {
		panic("space in name")
	}
	//fmt.Printf("adding type %s name:%s vals:%s\n", t.NodeType, name, vals)
	t.Vals[name] = vals
}

func (t *TestConsumer) NodeType(nodetype string, nodeid string, filename string) TreeNode {
	return &TestNode{
		Filename : filename, 
		NodeType: nodetype,
		NodeId:   nodeid,
		Vals:     make(map[string][]string),
	}
}

func NewConsumer(Filename string ) *TestConsumer {
	return &TestConsumer{
		Filename : Filename,
		Lock :sync.RWMutex{},
		Data : ProgramData {
			Summary : &SNodeType{
				CoOccurance: &CoOccurance{},
				AttrNames: make(map[string]int),
			},
			Transitions   :make(map[int]map[int]int),
			States        :make(map[int]int),
		},
		//Cache :  LoadCache(),
	}
}

func (t *TestConsumer) Node(n TreeNode)  {

	switch v:= n.(type) {
	case * TestNode:
		t.NodeImp(v)
	default:
		panic("unknown")
	}
}
