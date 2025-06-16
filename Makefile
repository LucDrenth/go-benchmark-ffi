compile-c: 
	cc -shared -o ./c/libadd.dylib -fPIC ./c/add.c

test:
	go test ./...

bench:
	go test ./... -bench=.
