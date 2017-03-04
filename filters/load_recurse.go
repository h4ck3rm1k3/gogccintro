package filter
// load recurse command
import (
	"fmt"
	"encoding/json"
	"bytes"
	//"log"
	"os"
	"database/sql"
	"github.com/h4ck3rm1k3/gogccintro/models"
	"github.com/h4ck3rm1k3/gogccintro/tree"
	"github.com/gocraft/dbr"
	"github.com/gocraft/dbr/dialect"
	"time"
	"runtime"
	"sync"
)

type StartNode struct {
	FileId int
	NodeType string
	Attrs[] struct {
		Field string
		SubField string
		SubValue string
	}
}


type load_recurse_t struct {
}

var myReceiver = &MyEventReceiver{}

// MyEventReceiver is a sentinel EventReceiver; use it if the caller doesn't supply one
type MyEventReceiver struct{}

// Event receives a simple notification when various events occur
func (n *MyEventReceiver) Event(eventName string) {
	// fmt.Printf("Event: %s\n",eventName)
}

// EventKv receives a notification when various events occur along with
// optional key/value data
func (n *MyEventReceiver) EventKv(eventName string, kvs map[string]string) {
	// fmt.Printf("EventKV: %s\n",eventName)
	// fmt.Printf("KV: %s\n",kvs)
}

// EventErr receives a notification of an error if one occurs
func (n *MyEventReceiver) EventErr(eventName string, err error) error {
	// fmt.Printf("errevnt: %s\n",eventName)
	// fmt.Printf("errde: %s\n",err)

	return err }

// EventErrKv receives a notification of an error if one occurs along with
// optional key/value data
func (n *MyEventReceiver) EventErrKv(eventName string, err error, kvs map[string]string) error {
	// fmt.Printf("errevt: %s\n",eventName)
	// fmt.Printf("errde2: %s\n",err)
	// fmt.Printf("errkv: %s\n",kvs)
	
	return err
}

// Timing receives the time an event took to happen
func (n *MyEventReceiver) Timing(eventName string, nanoseconds int64) {
	// fmt.Printf("timing: %s\n",eventName)
	// fmt.Printf("secs: %s\n",nanoseconds)
}

// TimingKv receives the time an event took to happen along with optional key/value data
func (n *MyEventReceiver) TimingKv(eventName string, nanoseconds int64, kvs map[string]string) {
	// fmt.Printf("timing: %s\n",eventName)
	// fmt.Printf("secs: %s\n",nanoseconds)
	// fmt.Printf("timing kvs: %s\n",kvs)
}


type Traversal struct {
	in * sql.DB
	out * sql.DB
	outf * os.File
	transform * Transform
	tree * tree.TreeMap
	wg * sync.WaitGroup
}

func (t * Traversal) recurse_field(fieldname string, id sql.NullInt64, fromid int , from_type string){
	if id.Valid {
		var v uint64 = uint64(id.Int64)
		if t.tree.SetBitFirst(v){
			//fmt.Printf("waitgroup before %v\n", t.wg)
			//fmt.Printf("going to recurse %d\n", v)
			t.wg.Add(1) // add this to the waitgroup
			//fmt.Printf("waitgroup %v\n", t.wg)
			go func() {
				//fmt.Printf("in recurse %d\n", v)
				//fmt.Printf("waitgroup %v\n", t.wg)
				defer t.wg.Done() // 
				t.recurse_node_id(int(v),fieldname, fromid, from_type)
				//fmt.Printf("after recurse %d\n", v)
			}()
		} else {
			//fmt.Printf("missed %d\n", v)
		}
	} else {
		//fmt.Printf("null %s\n", fieldname)
	}
}

func (t * Traversal) recurse(id int,fieldname string, fromid int){
	
	fmt.Printf("recurse start %d %s %d\n", id, fieldname, fromid)
	var err error
	n, err := models.GccTuParserNodeByID(t.in, id);
	if (err != nil){
	 	fmt.Printf("failed %s", err)
	 	return
	}
	t.recurse_node(n,fieldname, fromid, "entry")
}

func (t * Traversal) recurse_node_id(id int,fieldname string, fromid int, fromtype string){
	
	fmt.Printf("recurse field_node_id %d %s %d %s\n", id, fieldname, fromid, fromtype)
	var err error
	n, err := models.GccTuParserNodeBySourceFileIDNodeID(t.in, t.transform.LoadRecurse.FileId, id)
	if (err != nil){
	 	fmt.Printf("failed %s", err)
	 	return
	}
	t.recurse_node(n,fieldname, fromid, fromtype)
}

func (t * Traversal) recurse_node(n * models.GccTuParserNode,fieldname string, fromid int, from_type string){

	id := n.NodeID

	// save the nodes
	t.tree.Nodes[id]=n
	
	fmt.Printf("recurse nodeid:%d type:%s from (%s,%d,%s)\n", id, n.NodeType, fieldname, fromid, from_type)

	t.recurse_field("RefsArgs", n.RefsArgs, id, n.NodeType)
	t.recurse_field("RefsArgt", n.RefsArgt, id, n.NodeType)
	t.recurse_field("RefsBody", n.RefsBody, id, n.NodeType)
	t.recurse_field("RefsBpos", n.RefsBpos, id, n.NodeType)
	t.recurse_field("RefsChan", n.RefsChan, id, n.NodeType)
	// if the type is field_decl
	if n.NodeType == "field_decl" {
		t.recurse_field("RefsChain", n.RefsChain, id, n.NodeType)
	}
	t.recurse_field("RefsCnst", n.RefsCnst, id, n.NodeType)
	t.recurse_field("RefsCond", n.RefsCond, id, n.NodeType)
	t.recurse_field("RefsCsts", n.RefsCsts, id, n.NodeType)
	t.recurse_field("RefsDecl", n.RefsDecl, id, n.NodeType)
	t.recurse_field("RefsDomn", n.RefsDomn, id, n.NodeType)
	t.recurse_field("RefsE", n.RefsE, id, n.NodeType)
	t.recurse_field("RefsElts", n.RefsElts, id, n.NodeType)
	t.recurse_field("RefsExpr", n.RefsExpr, id, n.NodeType)
	t.recurse_field("RefsFlds", n.RefsFlds, id, n.NodeType)
	t.recurse_field("RefsFn", n.RefsFn, id, n.NodeType)
	t.recurse_field("RefsIdx", n.RefsIdx, id, n.NodeType)
	t.recurse_field("RefsInit", n.RefsInit, id, n.NodeType)
	t.recurse_field("RefsLabl", n.RefsLabl, id, n.NodeType)
	t.recurse_field("RefsLow", n.RefsLow, id, n.NodeType)
	t.recurse_field("RefsMax", n.RefsMax, id, n.NodeType)
	t.recurse_field("RefsMin", n.RefsMin, id, n.NodeType)
	t.recurse_field("RefsMngl", n.RefsMngl, id, n.NodeType)
	t.recurse_field("RefsName", n.RefsName, id, n.NodeType)
	t.recurse_field("RefsOp0", n.RefsOp0, id, n.NodeType)
	t.recurse_field("RefsOp1", n.RefsOp1, id, n.NodeType)
	t.recurse_field("RefsOp2", n.RefsOp2, id, n.NodeType)
	t.recurse_field("RefsPrms", n.RefsPrms, id, n.NodeType)
	t.recurse_field("RefsPtd", n.RefsPtd, id, n.NodeType)
	t.recurse_field("RefsPurp", n.RefsPurp, id, n.NodeType)
	t.recurse_field("RefsRefd", n.RefsRefd, id, n.NodeType)
	t.recurse_field("RefsRetn", n.RefsRetn, id, n.NodeType)
	//t.recurse_field("RefsScpe", n.RefsScpe, id, n.NodeType)
	t.recurse_field("RefsSize", n.RefsSize, id, n.NodeType)
	t.recurse_field("RefsType", n.RefsType, id, n.NodeType)
	t.recurse_field("RefsUnql", n.RefsUnql, id, n.NodeType)
	t.recurse_field("RefsVal", n.RefsVal, id, n.NodeType)
	t.recurse_field("RefsValu", n.RefsValu, id, n.NodeType)
	t.recurse_field("RefsVars", n.RefsVars, id, n.NodeType)
	
	//fmt.Printf("recurse done %d\n", id, n.NodeType)

}

func (t2 load_recurse_t) execute(in *sql.DB , out *sql.DB, outf * os.File, t * Transform) {
	runtime.GOMAXPROCS(1000)

	//fmt.Printf("load recurse test\n")
	// inside this function we select the 
	fmt.Printf("load_recurse: %#v\n", t.LoadRecurse)


	f,err := models.GccTuParserSourcefileByID(in,t.LoadRecurse.FileId)
	fmt.Printf("file: %s\n", f)
	if err != nil {
		fmt.Printf("could not find err: %s\n", err)
	}
	
	//i_sess := in.NewSession(nil)

	i_con := dbr.Connection{DB: in, EventReceiver: myReceiver, Dialect: dialect.SQLite3}
	i_sess := i_con.NewSession(nil)

	max_id := 0
	
	// get the max id from the file for a bitmap
	i_sess.Select("max(gcc_tu_parser_node.node_id)").
		From("gcc_tu_parser_node").
		Where(
		dbr.Eq("gcc_tu_parser_node.source_file_id",t.LoadRecurse.FileId),
	).LoadValue(&max_id)
	fmt.Printf("max id: %d\n", max_id)
	treemap := tree.NewTreeMap(max_id)
	fmt.Printf("treemap: %v\n", treemap)
	//treemap.Nodes[0]=nil
	fmt.Printf("node_type: %s\n", t.LoadRecurse.NodeType)

	for i,x := range t.LoadRecurse.Attrs {
		fmt.Printf("index: %d\n", i)
		fmt.Printf("Field: %s\n", x.Field)

		sql := i_sess.
			Select("gcc_tu_parser_node.id").
			From("gcc_tu_parser_node").

			Join(
			dbr.I("gcc_tu_parser_node").As("subobject"),
			dbr.And(
				dbr.Eq("gcc_tu_parser_node.source_file_id",t.LoadRecurse.FileId),
				dbr.Eq("subobject.source_file_id",t.LoadRecurse.FileId),
				dbr.Eq("gcc_tu_parser_node.node_type",t.LoadRecurse.NodeType),				
				dbr.Eq(
					fmt.Sprintf("gcc_tu_parser_node.%s",x.Field),
					dbr.I("subobject.node_id"),
				),
				dbr.Eq(
					fmt.Sprintf("subobject.%s",x.SubField),
					x.SubValue,
				),
			),
		)
		buf := dbr.NewBuffer()
		err := sql.Build(i_con.Dialect, buf)
		if err != nil {
			fmt.Printf("build err: %s\n", err)
		}
		fmt.Printf("build buf: %s\n", buf)


		fmt.Printf("SubField: %s\n", x.SubField)
		fmt.Printf("SubFieldValue: %s\n", x.SubValue)
		fmt.Printf("sql: %s\n", sql)

		var row int
		row = 0 
		sql.LoadValue(&row)
		fmt.Printf("node_id: %d\n", row)
		var wg sync.WaitGroup
		t2 := Traversal{
			in:in,
			out:out,
			transform:t,
			tree:treemap,
			wg: &wg,
		}

		start := time.Now()
		t2.recurse(row,"start",0);
		duration := time.Now().Sub(start)
		fmt.Printf("duration: %s\n", duration)
		//
		fmt.Println("Waiting To Finish")
		t2.wg.Wait()
		duration = time.Now().Sub(start)
		fmt.Printf("duration2: %s\n", duration)
		fmt.Printf("results: %v\n", t2.tree.Nodes)

		// write to  test file
		var buffer bytes.Buffer

		body, _ := json.Marshal(treemap.Nodes)
		buffer.Write(body)
		buffer.WriteString("\n")
		
		//fmt.Println(buffer.String())
		outf.WriteString(buffer.String())
		// now lets do a type analysis of this
		
	}
	
}


