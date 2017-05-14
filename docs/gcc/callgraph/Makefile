
query: cayley callgraph.q
	cat callgraph.q |GOPATH=/home/mdupont/gocode ~/gocode/src/github.com/cayleygraph/cayley/cmd/cayley/cayley query  -i ~/experiments/gcc-1/build/gcc/tree-dump.c.001t.tu.pq --limit 999999999999999999 > test.json


test.json  : query
	echo ok

test.dot : test.json
	#jq -f dot.jq -r test.json
	#	jq -f dot.jq -s test.json  -r
	echo "digraph {" > test.dot
	cat test.json | jq -f dot.jq -s -r >>  test.dot
	echo "}" >> test.dot

test.png : test.dot
	dot -Tpng -o test.png test.dot

test.svg:
	dot -Tsvg -o test.svg test.dot


ws:
	./node_modules/wintersmith/bin/wintersmith  build
