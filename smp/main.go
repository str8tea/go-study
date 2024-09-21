package main

import (
	"fmt"

	"github.com/str8tea/go-study/prime"
)

func main() {
	for i := 0; i < 100; i++ {
		if prime.IsPrime(i) {
			fmt.Printf("%d-素数\n", i)
		}
	}
}
