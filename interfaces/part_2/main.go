package main

import "fmt"

// interfaces part-2

type IEncoder interface {
	Encode(value string)
	Decode(value string)
}



type XEncoder struct {
}


func (xEncoder *XEncoder) Encode(value string){
	fmt.Println("XEncoder ile encode edildi.")
}

func (xEncoder *XEncoder) Decode(value string){
	fmt.Println("XEncoder ile decode edildi.")
}

type YEncoder struct {

}

func (yEncoder *YEncoder) Encode(value string){
	fmt.Println("YEncoder ile encode edildi.")
}

func (yEncoder *YEncoder) Decode(value string){
	fmt.Println("YEncoder ile decode edildi.")
}


func main() {

	var encoder IEncoder
	encoder = &YEncoder{}

	encoder.Encode("12345")

}