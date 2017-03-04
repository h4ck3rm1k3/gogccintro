package filter
// load recurse command
import (
	"fmt"
	"encoding/json"
	//"log"
	//"os"
	"database/sql"
	"github.com/h4ck3rm1k3/gogccintro/models"
	"github.com/gocraft/dbr"
	"github.com/gocraft/dbr/dialect"
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
	fmt.Printf("Event: %s\n",eventName)
}

// EventKv receives a notification when various events occur along with
// optional key/value data
func (n *MyEventReceiver) EventKv(eventName string, kvs map[string]string) {
	fmt.Printf("EventKV: %s\n",eventName)
	fmt.Printf("KV: %s\n",kvs)
}

// EventErr receives a notification of an error if one occurs
func (n *MyEventReceiver) EventErr(eventName string, err error) error {
	fmt.Printf("errevnt: %s\n",eventName)
	fmt.Printf("errde: %s\n",err)

	return err }

// EventErrKv receives a notification of an error if one occurs along with
// optional key/value data
func (n *MyEventReceiver) EventErrKv(eventName string, err error, kvs map[string]string) error {
	fmt.Printf("errevt: %s\n",eventName)
	fmt.Printf("errde2: %s\n",err)
	fmt.Printf("errkv: %s\n",kvs)
	
	return err
}

// Timing receives the time an event took to happen
func (n *MyEventReceiver) Timing(eventName string, nanoseconds int64) {
	fmt.Printf("timing: %s\n",eventName)
	fmt.Printf("secs: %s\n",nanoseconds)
}

// TimingKv receives the time an event took to happen along with optional key/value data
func (n *MyEventReceiver) TimingKv(eventName string, nanoseconds int64, kvs map[string]string) {
	fmt.Printf("timing: %s\n",eventName)
	fmt.Printf("secs: %s\n",nanoseconds)
	fmt.Printf("timing kvs: %s\n",kvs)
}

type foo struct {
	NodeType      string        `json:"node_type"`       // node_type
	NodeID        string        `json:"node_id"`         // node_id

}

type Traversal struct {
	in * sql.DB
}

func (t Traversal) recurse_field(fieldname string, id sql.NullInt64){
	
}
func (t Traversal) recurse(id int){

	fmt.Printf("lookup id %d", id)
	fmt.Printf("in t %d", t)
	var err error
	n, err := models.GccTuParserNodeByID(t.in, id);

	if (err != nil){
	 	fmt.Printf("failed %s", err)
	 	return
	}
	
	//fmt.Printf("ok %s %s", n, err)
	b, err := json.Marshal(n)
	fmt.Printf("json %s %s\n", b, err)

	if (err != nil){
		fmt.Printf("failed %s", err)
		return
	}

	t.recurse_field("RefsArgs", n.RefsArgs)
	t.recurse_field("RefsArgt", n.RefsArgt)
	t.recurse_field("RefsBody", n.RefsBody)
	t.recurse_field("RefsBpos", n.RefsBpos)
	t.recurse_field("RefsChan", n.RefsChan)
	t.recurse_field("RefsCnst", n.RefsCnst)
	t.recurse_field("RefsCond", n.RefsCond)
	t.recurse_field("RefsCsts", n.RefsCsts)
	t.recurse_field("RefsDecl", n.RefsDecl)
	t.recurse_field("RefsDomn", n.RefsDomn)
	t.recurse_field("RefsE", n.RefsE)
	t.recurse_field("RefsElts", n.RefsElts)
	t.recurse_field("RefsExpr", n.RefsExpr)
	t.recurse_field("RefsFlds", n.RefsFlds)
	t.recurse_field("RefsFn", n.RefsFn)
	t.recurse_field("RefsIdx", n.RefsIdx)
	t.recurse_field("RefsInit", n.RefsInit)
	t.recurse_field("RefsLabl", n.RefsLabl)
	t.recurse_field("RefsLow", n.RefsLow)
	t.recurse_field("RefsMax", n.RefsMax)
	t.recurse_field("RefsMin", n.RefsMin)
	t.recurse_field("RefsMngl", n.RefsMngl)
	t.recurse_field("RefsName", n.RefsName)
	t.recurse_field("RefsOp0", n.RefsOp0)
	t.recurse_field("RefsOp1", n.RefsOp1)
	t.recurse_field("RefsOp2", n.RefsOp2)
	t.recurse_field("RefsPrms", n.RefsPrms)
	t.recurse_field("RefsPtd", n.RefsPtd)
	t.recurse_field("RefsPurp", n.RefsPurp)
	t.recurse_field("RefsRefd", n.RefsRefd)
	t.recurse_field("RefsRetn", n.RefsRetn)
	t.recurse_field("RefsScpe", n.RefsScpe)
	t.recurse_field("RefsSize", n.RefsSize)
	t.recurse_field("RefsType", n.RefsType)
	t.recurse_field("RefsUnql", n.RefsUnql)
	t.recurse_field("RefsVal", n.RefsVal)
	t.recurse_field("RefsValu", n.RefsValu)
	t.recurse_field("RefsVars", n.RefsVars)
	
	
}

func (t2 load_recurse_t) execute(in *sql.DB , out *sql.DB, t Transform) {
	//do_load_recurse(t)
	//fmt.Printf("load recurse test\n")
	// inside this function we select the 
	fmt.Printf("load_recurse: %#v\n", t.LoadRecurse)
	fmt.Printf("fileid: %d\n", )


	//log.SetOutput(os.Stderr)
	//logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	
	f,err := models.GccTuParserSourcefileByID(in,t.LoadRecurse.FileId)
	fmt.Printf("file: %s\n", f)
	if err != nil {
		fmt.Printf("could not find err: %s\n", err)
	}
	
	//i_sess := in.NewSession(nil)

	i_con := dbr.Connection{DB: in, EventReceiver: myReceiver, Dialect: dialect.SQLite3}
	i_sess := i_con.NewSession(nil)
	
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

		t := Traversal{
			in:in,
		}
		t.recurse(row);
		//
	}
}

// func do_load_recurse(t Transform) {
// 	//load_config()
// 	//execute_filter()
// 	//report_results()
// 	fmt.Printf("load recurse %#v\n", t)
// }

