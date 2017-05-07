# t4: t
# 	~/gocode/bin/peg  -inline -switch  t4.peg

# t3: t
# 	~/gocode/bin/peg  -inline -switch  t3.peg

vecdelete2.txt: parser ~/experiments/gogccintro/parser/tests/vecdelete.tu
	./parser -file ~/experiments/gogccintro/parser/tests/vecdelete.tu -name vecdel -consumer grapht

vecdelete.txt: parser ~/experiments/gogccintro/parser/tests/vecdelete.tu
	./parser -debuglevel 1000 -file ~/experiments/gogccintro/parser/tests/vecdelete.tu -consumer grapht

spec.txt: parser ~/experiments/gogccintro/parser/tests/naked_operator.tu
	./parser -debuglevel 1000 -file ~/experiments/gogccintro/parser/tests/spec.tu -consumer graph

op.txt: parser ~/experiments/gogccintro/parser/tests/naked_operator.tu
	./parser -debuglevel 1000 -file ~/experiments/gogccintro/parser/tests/naked_operator.tu -consumer graph


big.txt : parser
	./parser -file ~/experiments/gcc-1/build/gcc/tree-dump.c.001t.tu -consumer grapht

gcctest.txt : parser
	./parser -directory ~/experiments/gcc-1/build/gcc/tests/ -name gcctests -consumer grapht
#mem2

gcc.txt : parser
	./parser -directory ~/experiments/gcc-1/build/gcc/ -name gcc -consumer graph

tests.txt: parser
	./parser -debuglevel 2000 -directory ~/experiments/gogccintro/parser/tests/  -consumer grapht

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

parser :  main.go Parser.go Fields.go Trie.go Consumer.go CoOccurance.go Groups.go Ybc.go Gob.go readcache.go Ast_ffjson.go ContainmentHierarchy.go GraphConsumer.go

	GOPATH=/home/mdupont/gocode/ go build 

readcache :  parser
	./parser -operation readcache  -name gccpretty
