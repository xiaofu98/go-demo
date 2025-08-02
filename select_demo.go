package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func(){
		time.Sleep(100*time.Microsecond)
		ch1 <- 100
	}()

	go func(){
		time.Sleep(100*time.Microsecond)
		ch2 <- 200
	}()
	select{
	case val := <- ch1:
		fmt.Println("ch1:",val)
	case val := <- ch2:
		fmt.Println("ch2:",val)
	}

}
