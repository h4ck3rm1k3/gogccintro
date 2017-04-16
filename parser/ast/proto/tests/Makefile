test :
	GOPATH=/home/mdupont/gocode/ go test 

dtests : tests
	./tests ~/experiments/gcc_py_introspector/tests/lsof.tu.proto create > r



ptests : tests
	./tests ~/experiments/gcc_py_introspector/tests/lsof.tu.proto print

tests : read_protobuf.go create_graph.go Report.go
	GOPATH=/home/mdupont/gocode/ go build 
