package main

import (
	"context"
	"fmt"
	"time"
)

type Product struct {
	Id int64
	Name string
}

var productChannel = make(chan Product)


func getProductDetailFromExternalService(productId int64) {

	time.Sleep(2*time.Second)

	productChannel <- Product{ Id :productId , Name:"Çamarşır Makinesi"}
}


func main() {

	cntx,canc:= context.WithTimeout(context.Background(),time.Second *3 )
	defer canc()


	go getProductDetailFromExternalService(10)

	select {
		case productDetail:= <- productChannel:
			fmt.Println("Ürün detayları getirildi.",productDetail)
		case <- cntx.Done():
			fmt.Println("Timeout occured, context cancelled.")
	}

	fmt.Println("End of the main function.")
}
