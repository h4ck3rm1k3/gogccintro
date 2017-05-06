package main

import (
	//"sync"
	"fmt"
	//"sort"
	"strconv"
	"strings"
	//"io/ioutil"
	//"bytes"
	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/graph"
	_ "github.com/cayleygraph/cayley/graph/bolt"
	"github.com/cayleygraph/cayley/quad"
)


type TestGraphConsumer struct {
	//Lock          sync.RWMutex//= sync.RWMutex{}
	Filename      string
	Store         graph.QuadWriter
	//Data ProgramData
	//Cache *       Cache
}

func (t *TestGraphConsumer) StateUsage(from int) {}
func (t *TestGraphConsumer) StateTransition(from int, to int) {}
func (t *TestGraphConsumer) Write() {}
func (t *TestGraphConsumer) Read() {}
func (t *TestGraphConsumer) Report() {

	// buf := bytes.NewBuffer(nil)
	// for _, c := range t.Store {
	// 	buf.Reset()
	// 	w := gml.NewWriter(buf)
	// 	n, err := quad.Copy(w, quad.NewReader(c.quads))
	// 	if err != nil {
	// 		t.Fatalf("write failed after %d quads: %v", n, err)
	// 	}
	// 	if err = w.Close(); err != nil {
	// 		t.Fatal("error on close:", err)
	// 	}
	// 	if c.data != buf.String() {
	// 		t.Fatalf("wrong output:\n%s\n\nvs\n\n%s", buf.String(), c.data)
	// 	}
	// }
	// err = ioutil.WriteFile(fmt.Sprintf("%s.graphml",t.Filename), buf, 0644)
	// if (err != nil){panic(err)}
}

func (t *TestGraphConsumer) NodeImp(n *TestGraphNode) {
	//t.Lock.Lock()
	//defer t.Lock.Unlock()
	t.Store.AddQuad(quad.Make(n.NodeId, "nodetype", n.NodeType,""))//, n.Filename

	for key, v := range n.Vals {
		vals := strings.Join(v, "|")
		t.Store.AddQuad(quad.Make(n.NodeId, key, vals,""))//, n.Filename

	}
	
}

type TestGraphNode struct {
	Filename string
	NodeType string
	NodeId   string
	Vals     map[string][]string
	//StrVals     map[string]string
	//IntVals     map[string]int
}

func (t *TestGraphNode) Finish(Parent TreeConsumer) {
	Parent.Node(t)
}

func (t *TestGraphNode) Report() {

}

func (t *TestGraphNode) SetAttr(name string, vals []string) {
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

func (t *TestGraphConsumer) NodeType(nodetype string, nodeid string, filename string) TreeNode {
	return &TestGraphNode{
		Filename: filename, 
		NodeType: nodetype,
		NodeId:   nodeid,
		Vals:     make(map[string][]string),
	}
}

func NewGraphConsumer(Filename string ) *TestGraphConsumer {

	boltfile := fmt.Sprintf("%s.bolt",Filename)
	
	// Initialize the database
	graph.InitQuadStore("bolt", boltfile, nil)
	store, err := cayley.NewGraph("bolt", boltfile, nil)
	fmt.Printf("writing bolt %s\n", boltfile)
	if err != nil {
		panic(err)
	}
	return &TestGraphConsumer{
		Filename : Filename,
		Store: store,
		//Lock :sync.RWMutex{},
	}
}

func NewMemoryGraphConsumer(Filename string ) *TestGraphConsumer {

	store, err := cayley.NewMemoryGraph()
	if err != nil {
		panic(err)
	}
	return &TestGraphConsumer{
		Filename : Filename,
		Store: store,
		//Lock :sync.RWMutex{},
	}
}



func (t *TestGraphConsumer) Node(n TreeNode)  {

	switch v:= n.(type) {
	case * TestGraphNode:
		t.NodeImp(v)
	default:
		panic("unknown")
	}
}
