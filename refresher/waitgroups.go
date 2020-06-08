package main

import (
	"fmt"
	"time"
)

func myFunc() {
	time.Sleep(1 * time.Second)
	fmt.Println("Finished Exc Goroutine")
}

func main() {
	fmt.Println("Working With Goroutine")
	go myFunc()
	fmt.Println("Finished running my ")

}
