genhugo:
	~/gocode/bin/hugo --theme=Introspector


query: cayley callgraph.q
	cat callgraph.q |GOPATH=/home/mdupont/gocode ~/gocode/src/github.com/cayleygraph/cayley/cmd/cayley/cayley query  -i ~/experiments/gcc-1/build/gcc/tree-dump.c.001t.tu.pq --limit -1 > test.json


test.json  : query
	echo ok

test.dot : test.json
	bash ./gendot.sh

test.png : test.dot
	dot -Tpng -o test.png test.dot

test.svg:
	dot -Tsvg -o test.svg test.dot

hugo:
	#	strace -f  -eopen,stat
	~/gocode/bin/hugo serve --bind 0.0.0.0 --log -v --verboseLog  --logFile log.txt --enableGitInfo --theme=Introspector --baseURL http://new-host:1313/gogccintro/gcc/CallGraph/public
