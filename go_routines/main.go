package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// Async
	wg := sync.WaitGroup{}

	wg.Add(3)
	startTime := time.Now()

	go func(){
		defer wg.Done()
		fmt.Println("Function 1")
		time.Sleep(3*time.Second)
	}()

	go	func(){
		defer wg.Done()
		fmt.Println("Function 2")
		time.Sleep(4*time.Second)
	}()

	go	func(){
		defer wg.Done()
		fmt.Println("Function 3")
		time.Sleep(5*time.Second)
	}()

	wg.Wait()

	fmt.Println("End of the main function.")
	fmt.Println("Passed time: ",time.Now().Sub(startTime))

	// // Sync

	// startTime := time.Now()

	// func(){
	// 	fmt.Println("Function 1")
	// 	time.Sleep(3*time.Second)
	// }()

	// 	func(){
	// 	fmt.Println("Function 2")
	// 	time.Sleep(4*time.Second)
	// }()

	// 	func(){
	// 	fmt.Println("Function 3")
	// 	time.Sleep(5*time.Second)
	// }()



	// fmt.Println("End of the main function.")
	// fmt.Println("Passed time: ",time.Now().Sub(startTime))

	// wg := sync.WaitGroup{};

	// wg.Add(2)

	// go func () {
	// 	defer wg.Done()
	// 	fmt.Println("First routine worked.")
	// }()

	// go func () {
	// 	defer wg.Done()
	// 	fmt.Println("Second routine worked.")
	// }()

	// wg.Wait()

	// fmt.Println("Main function end.")
}