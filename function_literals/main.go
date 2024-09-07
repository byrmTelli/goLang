package main

import (
	"fmt"
	"reflect"
)

func main() {

	fmt.Println("MErhaba")
	func (x int , y int) {
		fmt.Println(x+y)
	}(2,2)

	//func assingment to variable

	add:= func(x int ,y int) int {
		return x+y
	}

	// will return =>  func (int,int) int
	fmt.Println(reflect.TypeOf(add))

	//will return => 8
	fmt.Println("result: ",add(3,5))

	multiply:= func(x int ,y int) int {
		return x*y
	}

	a,b := calculator(3,5,add,multiply)

	fmt.Println(a,b)

}

func calculator (x int, y int, f1 func(int, int) int,f2 func(int ,int ) int) (int,int) {
	return f1(x,y),f2(x,y)
}