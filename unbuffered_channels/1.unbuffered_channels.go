package main

import "fmt"

func main() {

	myChannel := make(chan string)
	doneChannel := make(chan bool)
	go func() {
		//this routine does not throw deadlock by itself
		message := <-myChannel
		fmt.Println(message)
		doneChannel <- true
	}()

	go func() {
		myChannel <- "Hello WORLD!"
	}()

	<-doneChannel
	fmt.Println("End of the main function!")

	// go func() {
	// 	myChannel <- "Hello Bayram!"
	// }()

	// //Blocking

	// message,isOpen := <-myChannel

	// fmt.Println(message,isOpen)
}
