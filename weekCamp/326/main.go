package main

import "fmt"

func main() {
	prim := make([]int, 0, 4)
	prim = append(prim, 1)
	fmt.Printf("%d\n", prim[0])
}
