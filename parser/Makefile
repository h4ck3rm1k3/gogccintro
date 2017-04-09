test : example2
	./parser ~/experiments/gcc_py_introspector/tests
	./parser ~/experiments/gcc_py_introspector/tests2/
	./parser ~/experiments/lsof/

debug: test.peg
	~/gocode/bin/pigeon -debug test.peg 

example.go1: test.peg
	~/gocode/bin/pigeon test.peg  | tee example.go1

example.go : example.go1
	cat example.go1 | goimports>  example.go

example : example.go
	go build example.go

rtest : example
	go test
#-debug

test2.peg.go: test2.peg
	~/gocode/bin/peg  -inline -switch  test2.peg

example2 : test2.peg.go main.go
	GOPATH=/home/mdupont/gocode/ go build 

peg: 
	for x in test_*.peg; do echo $(x); ~/gocode/bin/peg  -inline -switch  $(x); done

