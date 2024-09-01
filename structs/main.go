package main

import "fmt"

func main() {

	//Creating new variable from struct

	var customer1 = Customer{id: 1, name: "bayram", age: 23 ,adress:Adress {city:"Kayseri",district:"Yahalı"}}

	customer1.ToString()
	customer1.changeName("Bayram Telli")

	customer1.ToString()
	// fmt.Println(customer1)


	// var customer2 = Customer{id:2, name:"emir",age:13}
	// fmt.Println(customer2)

	// customer2.age = 31
	// customer2.adress = Adress {city:"Ankara",district: "Kızılcahamam"}
	// fmt.Println(customer2)




}


//Function which we send the pointer of struct and change the value within adress of sctruct.


// func changeName(customer *Customer) {
// 	customer.name = "Eren Bülbül"
// }


//Structs 

type Customer struct {
	id   int64
	name string
	age  int
	adress Adress
}


type Adress struct {
	city string
	district string
}


//Struct Functions

func (customer *Customer) changeName(newName string){
	customer.name = newName
}


func(customer *Customer) ToString(){
	fmt.Printf("Name: %s, Age: %d\n", customer.name,customer.age)
}