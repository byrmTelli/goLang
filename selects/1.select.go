package main

import (
	"fmt"
	"time"
)

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	var data1 string
	var data2 string

	go func(){
		time.Sleep(5*time.Second)
		channel1 <- "Hello"
	}()

	go func(){
		time.Sleep(3*time.Second)
		channel2 <- "World!"
	}()

	//select is like switch case if the data comes one of these cases, the reciever case will run first and program will finish by that.

	//But if you use for loop like implemented below you can solve the issue.

		for len(data1) == 0 || len(data2) == 0{
			select{
				case data1 = <-channel1:
					fmt.Println("Recieved data from channel1",data1)
				case data2 = <-channel2:
					fmt.Println("Recieved data from channel2",data2)
				default:
					fmt.Println("No data recieved yet.")
			}
			time.Sleep(1*time.Second)
		}

	
}