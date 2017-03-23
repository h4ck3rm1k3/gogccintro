
gen: gogccintro
	./gogccintro

test: gogccintro 
	GOPATH=/home/mdupont/gocode/ GOTRACEBACK=crash go test ./... -v

t:gogccintro 
	GOPATH=/home/mdupont/gocode/ go test --run TestFactory -v

gogccintro : gogccintro.go  filters/*.go models/*.go node_types/*.go tree/*.go node_types/examples/*.go
	GOPATH=/home/mdupont/gocode/ go build
