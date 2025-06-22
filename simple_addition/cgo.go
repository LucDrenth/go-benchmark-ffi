package main

/*
#include <./c/add.c>
*/
import "C"

func addCgo(a, b int64) int64 {
	return int64(C.do_add(
		(C.int64_t)(a),
		(C.int64_t)(b),
	))
}
