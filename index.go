package main

import "fmt"

func main() {
	ch := make(chan int, 4)
	ch <- 1
	ch <- 2
	ch <- 100000000
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

}
