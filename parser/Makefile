example.go: test.peg
	~/gocode/bin/pigeon  test.peg | goimports>  example.go

example : example.go
	go build example.go

test : example
	go test
#-debug
