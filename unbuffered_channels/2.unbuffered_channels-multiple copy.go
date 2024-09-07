package main

import (
	"fmt"
	"time"
)

func main() {

	myChannel := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			myChannel <- i
			fmt.Println("Sent data: ",i)
			time.Sleep(1*time.Second)
		}
		close(myChannel)
	}()


	//infinite loop to see the data sended by the function that defined upper.
	for{
		data,isOpen := <- myChannel

		if isOpen == false {
			break
		}

		fmt.Println("Recieved data: ",data)
	}

}
