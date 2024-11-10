package main

import "fmt"
import "time"

func main() {
	ch := make(chan int)
	go myfunc(ch)
	for {
		select {
		case v := <-ch:
			fmt.Println("Here we go ", v)
		default:
			fmt.Println("Nothing.")
		}
		time.Sleep(1 * time.Second)
	}
}

func myfunc(ch chan int) {
	for {
		ch <- 1
		time.Sleep(1*time.Second)
	}
}
