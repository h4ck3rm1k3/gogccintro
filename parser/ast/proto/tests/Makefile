dtests : tests
	./tests ~/experiments/gcc_py_introspector/tests/lsof.tu.proto create

ptests : tests
	./tests ~/experiments/gcc_py_introspector/tests/lsof.tu.proto print

tests : read_protobuf.go create_graph.go
	GOPATH=/home/mdupont/gocode/ go build

