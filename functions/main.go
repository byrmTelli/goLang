package main

import "fmt"

func main() {


	// numbers:= []int{10,20,2,3,5}

	total:= sum(1,2,5,1,1,1)

	fmt.Println(total)


	// total,multiple:= add(5,10)


	// fmt.Printf("total: %d \n multip: %d",total,multiple)
}


func sum(numbers ... int) int {
	sum:=0
	
	for _,value := range numbers {
		sum+=value
	}

	return sum
}

// func sum (x int,y int,z int) int {
// 	return x+y+z
// }



// func sum(numbers []int) int {
// 	sum:=0

// 	for _,value:=range numbers {
// 		sum+=value
// 	}

// 	return sum
// }


// func add (x int, y int) {
// 	fmt.Println(x+y)
// }

// func add(x int, y int) (int,int) {
// 	return x + y,x*y
// }