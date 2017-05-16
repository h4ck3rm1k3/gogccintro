echo "digraph {" > static/callgraph/treedump/test.dot
cat static/callgraph/treedump/test.json | jq -f dot.jq -s -r >>  static/callgraph/treedump/test.dot
echo "}" >> static/callgraph/treedump/test.dot
