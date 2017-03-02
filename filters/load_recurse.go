package filter
// load recurse command
import (
	"fmt"
)

type StartNode struct {
	FileID int
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
	fmt.Printf("transform: %#v\n", t)
}

// func do_load_recurse(t Transform) {
// 	//load_config()
// 	//execute_filter()
// 	//report_results()
// 	fmt.Printf("load recurse %#v\n", t)
// }

