package main

import "fmt"

func main() {

	defer cleanUp()

	var condition = true

	if condition {
		panic("An error occurred.")
	}

	//Defer order check

	// defer fmt.Println("Defer")
	// defer fmt.Println("Defer 1")
	// defer fmt.Println("Defer 2")
	// defer fmt.Println("Defer 3")
	// fmt.Println("Go")
	// fmt.Println("Burası main fonksiyonu")
}

func cleanUp (){
	fmt.Println("@@@@@@@\n\n\nCleanup worked.\n\n\n")
}