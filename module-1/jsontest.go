package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Product struct {
	ProductID    int    `json: "productId"`
	Manufacturer string `json: "manufacturer"`
	Sku          string `json: "sku"`
	Upc          string `json: "upc"`
	PricePerUnit int    `json: "pricePerUnit"`
	Stock        int    `json: "stock"`
	ProductName  string `json: "productName"`
}

func main() {
	fmt.Println("Hello Playground")

	product := &Product{
		ProductID:    1,
		Manufacturer: "ACME",
		Sku:          "12345",
		Upc:          "12345-12345-12345",
		ProductName:  "Gizmo",
		Stock:        28,
	}
	product := Product{}
	productJson, err := json.Marshal(product)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(productJson))
}
