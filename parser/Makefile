# t4: t
# 	~/gocode/bin/peg  -inline -switch  t4.peg

# t3: t
# 	~/gocode/bin/peg  -inline -switch  t3.peg

op.txt: parser ~/experiments/gogccintro/parser/tests/naked_operator.tu
	./parser -debuglevel 1000 -file ~/experiments/gogccintro/parser/tests/naked_operator.tu

gcc.txt : parser
	./parser -directory ~/experiments/gcc-1/build/gcc/

tests.txt: parser
	./parser -directory ~/experiments/gogccintro/parser/tests/ 

test : parser tests.txt

#	./parser ~/experiments/gcc-1/build/gcc/tests > gcctest.txt
#	./parser ~/experiments/gcc-1/build/gcc/types > gcc_types.txt


#./gcc/tree-dump.ii
#mdupont@debian-build-speed:~/experiments/gcc-1/build$ grep "strg: lngt" ./gcc/tree-dump.c.001t.tu

#	./parser ~/experiments/gcc_py_introspector/tests
#	./parser ~/experiments/gcc_py_introspector/tests2/
#	./parser ~/experiments/lsof/

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

#test2.peg.go: test2.peg
#	~/gocode/bin/peg  -inline -switch  test2.peg

parser :  main.go Parser.go Fields.go Trie.go Consumer.go CoOccurance.go Groups.go

	GOPATH=/home/mdupont/gocode/ go build 

