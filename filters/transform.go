package filter

type Transform struct {
  Filter     string
  LoadRecurse StartNode
}

type ITransform interface{
  execute(t Transform)
}

func DoTransform(t Transform) {
	ld := load_recurse_t{}
	
	funcs := map[string] ITransform {
		"load_recurse":ld}
	funcs[t.Filter].execute(t)
}

