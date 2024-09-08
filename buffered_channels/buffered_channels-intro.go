package main

import (
	"fmt"
	"time"
)

func main() {
	// // the difference between unbuffered and buffered channels is a size definition in the make function
	// // you can define a size to create buffered channel.
	// myChannel := make(chan int, 2)

	// myChannel <- 1
	// myChannel <- 2

	// // by pushing 3rd item to channel will cause to get deadlock error.
	// myChannel <-3
	// fmt.Println("End of the main func")

	messages:= make(chan string,2)

	go func(){

		data1 := <-messages
		fmt.Println("First: ", data1)
		
		data2 := <-messages
		fmt.Println("Second: ", data2)

		data3 := <-messages
		fmt.Println("Third: ", data3)

	}()

	messages <- "Hello"
	messages <- "World"
	messages <- "World2"

	time.Sleep(1*time.Second)
	fmt.Println("End of the main func")
}