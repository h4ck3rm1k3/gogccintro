protobuf:
	#/usr/bin/protoc generated.proto --python_out=testpython
	PATH=$(PATH):/home/mdupont/gocode/bin	protoc generated.proto --go_out=plugins=grpc:proto

test:
	GOPATH=/home/mdupont/gocode/ go test -v .

gentest:
	GOPATH=/home/mdupont/gocode/ go test -v . -test.run TestGen
