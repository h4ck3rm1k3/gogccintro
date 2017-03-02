package filter
// load recurse command
import (
	"fmt"
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

func (t2 load_recurse_t) execute(t Transform) {
	//do_load_recurse(t)
	//fmt.Printf("load recurse test\n")
	// inside this function we select the 
	fmt.Printf("load_recurse: %#v\n", t.LoadRecurse)
	fmt.Printf("fileid: %d\n", t.LoadRecurse.FileId)
	fmt.Printf("node_type: %s\n", t.LoadRecurse.NodeType)
	for i,x := range t.LoadRecurse.Attrs {
		fmt.Printf("index: %d\n", i)
		fmt.Printf("Field: %s\n", x.Field)
		fmt.Printf("SubField: %s\n", x.SubField)
		fmt.Printf("SubFieldValue: %s\n", x.SubValue)
	}
}

// func do_load_recurse(t Transform) {
// 	//load_config()
// 	//execute_filter()
// 	//report_results()
// 	fmt.Printf("load recurse %#v\n", t)
// }

