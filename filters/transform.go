package filters

import (
	//"github.com/knq/dburl"
	"os"
	"database/sql"
)
type Transform struct {
  Filter     string
  LoadRecurse StartNode
}

type ITransform interface{
  execute(in *sql.DB , out *sql.DB, outf * os.File, t * Transform)
}

func DoTransform(in *sql.DB , out *sql.DB, outf *os.File, t * Transform) {
	ld := load_recurse_t{}
	
	funcs := map[string] ITransform {
		"load_recurse":ld}
	funcs[t.Filter].execute(in,out,outf,t)
}

