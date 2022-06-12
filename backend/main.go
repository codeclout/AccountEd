package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	// v := rand.Int()
	v2 := rand.Intn(10000000000000)

	// fmt.Println(v)
	fmt.Println(v2)
}
