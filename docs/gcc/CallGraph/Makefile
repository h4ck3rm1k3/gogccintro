genhugo:
	~/gocode/bin/hugo --theme=Introspector


query: cayley callgraph.q
	cat callgraph.q |GOPATH=/home/mdupont/gocode ~/gocode/src/github.com/cayleygraph/cayley/cmd/cayley/cayley query  -i ~/experiments/gcc-1/build/gcc/tree-dump.c.001t.tu.pq --limit -1 > test.json


test.json  : query
	echo ok

static/callgraph/treedump/test.dot : static/callgraph/treedump/test.json ./gendot.sh
	bash ./gendot.sh

static/callgraph/treedump/test.png : static/callgraph/treedump/test.dot
	dot -Tpng -o static/callgraph/treedump/test.png static/callgraph/treedump/test.dot

static/callgraph/treedump/test.svg:
	dot -Tsvg -o static/callgraph/treedump/test.svg static/callgraph/treedump/test.dot

out : static/callgraph/treedump/test.svg static/callgraph/treedump/test.png

hugo: out browserify
	#	strace -f  -eopen,stat
	~/gocode/bin/hugo serve --bind 0.0.0.0 --log -v --verboseLog  --logFile log.txt --enableGitInfo --theme=Introspector --baseURL http://new-host:1313/gogccintro/gcc/CallGraph/public


browserify :
	browserify --debug static/cayley/js/cayley_static.js -o static/cayley/js/bundle.js
