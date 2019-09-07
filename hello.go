package main

import (
	"fmt"
	"time"
)

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Printf("hello, world\n")
	}
	duration := time.Duration(10)*time.Second
	time.Sleep(duration)
}