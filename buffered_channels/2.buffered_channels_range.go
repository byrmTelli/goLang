package main

import (
	"fmt"
	"time"
)

func main() {

	// This code will throw deadlock. The reason is  first go rutine function will send first 4 data one by one.
	// After that for loop will read the data and it will continue like this. But in the end bufferChannel is still not closed this is the reason of deadlock.
	buffereChannel := make(chan int, 4)

	go func() {
		for i := 1; i <= 10; i++ {
			buffereChannel <- i
			fmt.Println("Send data: ",i)
		}

		//if you add this close method the process is not gonna throw deadlock.
		close(buffereChannel)
	}()

	time.Sleep(3*time.Second)

	for data:= range buffereChannel {
		fmt.Println("Recieved data: ",data)
		time.Sleep(1*time.Second)
	} 
}
