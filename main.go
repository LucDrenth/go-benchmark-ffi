package main

import "fmt"

func main() {
	loadPurego()

	var a int64 = 10
	var b int64 = 25

	goAddition := addGo(a, b)
	cgoAddition := addCgo(a, b)
	puregoAddition := addPurego(a, b)

	fmt.Println(goAddition, cgoAddition, puregoAddition)
}
