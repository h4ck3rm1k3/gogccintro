protobuf:
	#/usr/bin/protoc generated.proto --python_out=testpython
	PATH=$(PATH):/home/mdupont/gocode/bin	protoc generated.proto --go_out=plugins=grpc:proto

test: protobuf
	GOPATH=/home/mdupont/gocode/ go test -v ./proto/test/

generate:
	GOPATH=/home/mdupont/gocode/ go test -v . -test.run TestGen
