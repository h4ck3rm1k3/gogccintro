echo "digraph {" > test.dot
cat test.json | jq -f dot.jq -s -r >>  test.dot
echo "}" >> test.dot
