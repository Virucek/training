package main

import (
	"fmt"
	"time"
)

func main() {
	var n, a int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Printf("hello, world\n")
	}
	fmt.Println("Введите количество секунд для Sleep")
	fmt.Scan(&a)
	duration := time.Duration(a)*time.Second
	time.Sleep(duration)
}