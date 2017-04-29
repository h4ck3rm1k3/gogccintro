package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Stat struct {
	TheCount int
}

//t.Parent.Node(t)
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
	//Vals map [string] SValue
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

func (t *SAttributeNames) Vals(n *TestNode, t2 *TestConsumer) {
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
		t2.CoOccurance.Fields(n.NodeType,key)
		
		for key2, _ := range n.Vals {
			if key2 != key {
				if strings.Compare(key, key2) > 0 {
					t2.CoOccurance.Fields(key, key2)
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

func (t *SAttributeVector) Vals(n *TestNode, t2 *TestConsumer) {
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
		val.Vals(n, t2)
	} else {
		o := &SAttributeNames{}
		t.TheVals[key] = o
		o.Vals(n, t2)
	}
}

type SNodeType struct {
	Count          Stat
	AttributeCount map[int]*SAttributeVector // Attribute County
}

func (t *SNodeType) Report(nodetype string) {
	t.Count.Report(2)
	fmt.Printf("\t\tlen attrcount %d:\n", len(t.AttributeCount))
	for k, v := range t.AttributeCount {
		fmt.Printf("\t\tattrcount %d:\n", k)
		v.Report(nodetype, k)
	}
}

func (t *SNodeType) Node(n *TestNode, t2 *TestConsumer) {
	t.Count.Count()
	key := len(n.Vals)
	if t.AttributeCount == nil {
		t.AttributeCount = make(map[int]*SAttributeVector)
	}
	if val, ok := t.AttributeCount[key]; ok {
		val.Vals(n, t2)
	} else {
		o := &SAttributeVector{}
		t.AttributeCount[key] = o
		o.Vals(n, t2)
	}
	//t.AttributeCount[len(n.Vals)].Vals(n)
}

type TestConsumer struct {
	Count       Stat
	NodeTypes   map[string]*SNodeType

	// parser stats
	Transitions   map[int]map[int]int
	States        map[int]int

	
	CoOccurance *CoOccurance
}

func (t *TestConsumer) StateUsage(from int) {
	old := t.States[from]
	t.States[from]=old + 1	
}

func (t *TestConsumer) StateTransition(from int, to int) {

	if val, ok := t.Transitions[from]; ok {
		old := val[to]
		val[to]=old + 1
	} else {
		v := make(map[int]int)
	        v[to]=1
		t.Transitions[from] = v
	}
}

func (t *TestConsumer) Report() {

	//
	for k, v := range t.States {
		fmt.Printf("\tS %20s %10d\n", StateLookup[Token(k)], v)
	}
	// report on transitions

	to_count := make(map[int]int)
	
	for k, v := range t.Transitions {
		c:=0
		for k2, v2 := range v {
			fmt.Printf("\tT %20s %20s %10d\n", StateLookup[Token(k)], StateLookup[Token(k2)], v2)
			c = c+ v2
			to_count[k2] = to_count[k2] + v2
		}
		//
		fmt.Printf("\tAS %20s %10d\n", StateLookup[Token(k)], t.States[k]/c)
		
	}
	for k, v := range to_count {
		fmt.Printf("\tTAS %20s %10d\n", StateLookup[Token(k)], t.States[k]/v)
	}
		
	t.Count.Report(1)
	
	//fmt.Printf("Report %#v\n", t.NodeTypes)
	for k, v := range t.NodeTypes {
		fmt.Printf("\tNT %s %d\n", k, v.Count.TheCount)
		//v.Report(k)
	}
	
	t.CoOccurance.Report()
}

func (t *TestConsumer) Node(n *TestNode) {
	t.Count.Count()

	if t.NodeTypes == nil {
		t.NodeTypes = make(map[string]*SNodeType)
	}
	if val, ok := t.NodeTypes[n.NodeType]; ok {
		val.Node(n, t)
	} else {
		o := &SNodeType{}
		t.NodeTypes[n.NodeType] = o
		o.Node(n, t)
	}
}

type TestNode struct {
	Parent   *TestConsumer
	NodeType string
	NodeId   string
	Vals     map[string][]string
}

func (t *TestNode) Finish() {
	t.Parent.Node(t)
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

func (t *TestConsumer) NodeType(nodetype string, nodeid string) TreeNode {
	return &TestNode{
		Parent:   t,
		NodeType: nodetype,
		NodeId:   nodeid,
		Vals:     make(map[string][]string),
	}
}

func NewConsumer() *TestConsumer {
	return &TestConsumer{
		Transitions   :make(map[int]map[int]int),
		States        :make(map[int]int),	
		CoOccurance: &CoOccurance{},
	}
}
