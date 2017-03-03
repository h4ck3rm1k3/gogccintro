package filter

import (
	//"github.com/knq/dburl"
	"database/sql"
)
type Transform struct {
  Filter     string
  LoadRecurse StartNode
}

type ITransform interface{
  execute(in *sql.DB , out *sql.DB, t Transform)
}

func DoTransform(in *sql.DB , out *sql.DB, t Transform) {
	ld := load_recurse_t{}
	
	funcs := map[string] ITransform {
		"load_recurse":ld}
	funcs[t.Filter].execute(in,out,t)
}

