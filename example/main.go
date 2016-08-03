package main

import (
	"fmt"
	"github.com/TranDuyThanh/gomath"
)

func main() {
	var a, i int64 = 1, 2
	for ; i <= 1000; i++ {
		a = gomath.LCM(a, i)
	}
	fmt.Println(a)
}
