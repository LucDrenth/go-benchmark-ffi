compile-c: 
	cc -shared -o ./simple_addition/c/libadd.dylib -fPIC ./simple_addition/c/add.c

test:
	go test ./...

bench:
	go test ./... -bench=.
