package main

import (
	"fmt"
	"package/compute"
)

func main() {
	params := &compute.IntParams{
		P1: 20,
		P2: 30,
	}
	fmt.Println(params.Add())
}
