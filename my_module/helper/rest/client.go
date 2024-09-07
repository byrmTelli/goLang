package rest

import "fmt"

func Rest1() {
	fmt.Println("Rest from helper/helper_2/rest")
}

func rest2() {
	fmt.Println("Rest from helper/helper_2/rest and this function is not accesible from other packages")
}