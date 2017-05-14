package main

import (
	//"sync"
	"fmt"
	"os"
	//"sort"
	"strconv"
	"strings"
	//"io/ioutil"
	//"bytes"
	//"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/graph"
	_ "github.com/cayleygraph/cayley/graph/bolt"
	//"github.com/cayleygraph/cayley/quad"
)


type TestGraphConsumer struct {
	//Lock          sync.RWMutex//= sync.RWMutex{}
	Filename      string
	Outfile       *os.File 
	//Store         graph.QuadWriter
	//Transaction   * graph.Transaction
	//Data ProgramData
	//Cache *       Cache
}

func (t *TestGraphConsumer) StateUsage(from int) {}
func (t *TestGraphConsumer) StateTransition(from int, to int) {}
func (t *TestGraphConsumer) Write() {}
func (t *TestGraphConsumer) Read() {}
func (t *TestGraphConsumer) Report() {
	t.Outfile.Close()
	// if t.Transaction != nil {
	// 	err := t.Store.ApplyTransaction(t.Transaction)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	t.Transaction = cayley.NewTransaction()
	// }
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

func AbstractValue(nt string , ft string) (string){
	//
	if ft =="accs"{  		return "Scalar"
	} else if ft =="algn"{		return "Scalar"
	} else if ft =="args"{          return "Contained"     // "function_decl"->"parm_decl"		
	} else if ft =="argt"{		return "TypeReference" // param_decl -> some type
	} else if ft =="base"{		return "TypeReference" // record type -> record type
	} else if ft =="bases"{         return "Scalar"		
	} else if ft =="bfld"{		return "TypeReference" // record type -> record type
	} else if ft =="binf"{		return "TypeReference" // record type/union
	} else if ft =="body"{		return "Contained"
	} else if ft =="bpos"{		return "ContainedScalar"
	} else if ft =="chain"{		return "Skip"          // skip for now
	} else if ft =="chan"{		return "Skip"          // todo
	} else if ft =="clas"{		return "TypeReference"   // method -> record type
	} else if ft =="cnst"{		return "ContainedScalar" // const decl -> integer constant
	} else if ft =="cond"{          return "Contained"      // stmts and exprs
	} else if ft =="crnt"{		return "TypeReference"  // templates overload/ function decl/template_decl
	} else if ft =="csts"{		return "ContainedList"  // list of constants
	} else if ft =="ctor"{		return "Scalar"
	} else if ft =="dcls"{		return "ContainedList" // list of decls in a namespace
	} else if ft =="decl"{		return "Contained"     // target_expr/var_decl/aggr_init_expr
	} else if ft =="domn"{		return "TypeReference"
	} else if ft =="else"{		return "Contained"     // if/bind
	} else if ft =="elts"{		return "TypeReference" // array element types
	} else if ft =="expr"{		return "Contained"
	} else if ft =="externbody"{    return "Error"	       // errr
	} else if ft =="flds"{		return "Contained"     // types contain fields and others
	} else if ft =="fn"{		return "Contained"     // types contain fields and others
	} else if ft =="fncs"{		return "Contained"     // types contain fields and others
	} else if ft =="idx"{		return "Contained"     // types contain fields and others
	} else if ft =="init"{		return "Contained"
	} else if ft =="inst"{		return "ContainedList" // template / list
	} else if ft =="int"{		return "Scalar"
	} else if ft =="labl"{		return "Contained"     // goto-label
	} else if ft =="lang"{		return "Scalar"
	} else if ft =="line"{		return "Scalar"
	} else if ft =="link"{		return "Scalar"
	} else if ft =="lngt"{		return "Scalar"
	} else if ft =="low"{		return "Contained"     // integer const
	} else if ft =="max"{		return "Contained"
	} else if ft =="min"{ 		return "Contained"
	} else if ft =="mngl"{		return "Contained"
	} else if ft =="name"{		return "Contained"
	} else if ft =="nodetype"{	return "Type"
	} else if ft =="note"{          return "Scalar"
	} else if ft =="op_0"{		return "Contained"
	} else if ft =="op_1"{		return "Contained"
	} else if ft =="op_2"{		return "Contained"
	} else if ft =="orig"{		return "Reference"
	} else if ft =="prec"{		return "Scalar"
	} else if ft =="prms"{		return "ContainedList" // -> tree_list
	} else if ft =="ptd"{		return "Reference"
	} else if ft =="purp"{		return "Contained"
	} else if ft =="qual"{		return "Reference"
	} else if ft =="refd"{		return "Reference"
	} else if ft =="retn"{		return "Contained"
	} else if ft =="rslt"{		return "Contained"
	} else if ft =="scpe"{		return "Container"
	} else if ft =="sign"{		return "Scalar"
	} else if ft =="size"{		return "Reference"
	} else if ft =="spcs"{		return "ContainedList"
	} else if ft =="spec"{		return "Scalar"
	} else if ft =="srcp"{		return "Contained"
	} else if ft =="strg"{		return "Contained"
	} else if ft =="tag"{		return "Scalar"
	} else if ft =="then"{		return "Contained"
	} else if ft =="type"{		return "Reference"
	} else if ft =="unql"{		return "Reference"
	} else if ft =="used"{		return "Scalar"
	} else if ft =="val"{		return "Contained"
	} else if ft =="valu"{		return "Contained"
	} else if ft =="vars"{		return "Contained"
	} else if ft =="vfld"{		return "Contained"
	} else {	                return "Contained" }
	return ""	
}

func (t *TestGraphConsumer) NodeImp(n *TestGraphNode) {
	//t.Lock.Lock()
	//defer t.Lock.Unlock()
	//if t.Transaction != nil{
	nid := fmt.Sprintf("@%s",n.NodeId)
	//t.Transaction.AddQuad(quad.Make(nid, "nodetype", n.NodeType,""))//, n.Filename
	t.Outfile.WriteString(fmt.Sprintf("\"%s\" \"nodetype\" \"%s\" \"\".\n",nid,Escape(n.NodeType)))
	for key, v := range n.Vals {
		k := strings.Replace(key," ","",-1)
		k = strings.Replace(k,":","",-1)
		vals := strings.Join(v, "|")
		//t.Transaction.AddQuad(quad.Make(nid, k, vals,""))//, n.Filename
		//"71014" "type:" "43207" "" .
		v2 := Escape(vals)
		t.Outfile.WriteString(fmt.Sprintf("\"%s\" \"%s\" \"%s\" \"\".\n",nid,k,v2))

		// now we abstract the fieldname based on the nodetype and fieldname
		k2 := AbstractValue(n.NodeType,k)
		if k2 != "" {
			t.Outfile.WriteString(fmt.Sprintf("\"%s\" \"%s\" \"%s\" \"\".\n",nid,k2,v2))
		}
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

// from calay
var escaper = strings.NewReplacer(
	"\\", "\\\\",
	"\"", "\\\"",
	"\n", "\\n",
	"\r", "\\r",
	"\t", "\\t",
)

// from cayley
func Escape(s string) string {
	//TODO(barakmich): Proper escaping.
	return escaper.Replace(string(s))
}

func NewGraphTransactionConsumer(Filename string ) *TestGraphConsumer {

	filename := fmt.Sprintf("%s.nq",Filename)
	fmt.Printf("creating %s\n", filename)
	fi, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	return &TestGraphConsumer{
		Filename : Filename,
		Outfile: fi,
	}
}

func NewGraphConsumer(Filename string ) *TestGraphConsumer {

	boltfile := fmt.Sprintf("%s.bolt",Filename)
	
	// Initialize the database
	graph.InitQuadStore("bolt", boltfile, nil)
	//store, err := cayley.NewGraph("bolt", boltfile, nil)
	fmt.Printf("writing bolt %s\n", boltfile)
	// if err != nil {
	// 	panic(err)
	// }
	return &TestGraphConsumer{
		//Transaction : nil,
		Filename : Filename,
		//Store: store,
		//Lock :sync.RWMutex{},
	}
}

func NewMemoryGraphConsumer(Filename string ) *TestGraphConsumer {

	//store, err := cayley.NewMemoryGraph()

	//if err != nil {
	//panic(err)
	//}
	return &TestGraphConsumer{
		//Transaction : nil,
		Filename : Filename,
		//Store: store,
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
