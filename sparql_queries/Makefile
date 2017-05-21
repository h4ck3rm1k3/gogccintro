
sparql_queries : recurse.go
	GOPATH=~/gocode go build

t: sparql_queries
	#strace -f -s9999
	./sparql_queries
