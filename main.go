package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
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

var productList []Product

func init() {
	productsJSON := `
[{
    "productId": 1,
    "manufacturer": "Babblestorm",
    "sku": "4092",
    "upc": "852-B10-167",
    "productName": "Pepper - Yellow Bell",
    "stock": 34
  }, {
    "productId": 2,
    "manufacturer": "Realcube",
    "sku": "9324",
    "upc": "942-V42-187",
    "productName": "Pastry - Choclate Baked",
    "stock": 1
  }, {
    "productId": 3,
    "manufacturer": "Skinte",
    "sku": "5468",
    "upc": "946-I67-209",
    "productName": "Cheese - Pont Couvert",
    "stock": 17
  }, {
    "productId": 4,
    "manufacturer": "Bluejam",
    "sku": "2271",
    "upc": "767-K80-766",
    "productName": "Wine - Magnotta - Cab Franc",
    "stock": 20
  }]`

	err := json.Unmarshal([]byte(productsJSON), &productList)
	if err != nil {
		log.Fatal(err)
	}
}

func getNextID() int {
	hightestID := 1
	for _, product := range productList {
		if hightestID < product.ProductID {
			hightestID = product.ProductID
		}
	}

	return hightestID + 1

}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		productsJson, err := json.Marshal(productList)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(productsJson)

	case http.MethodPost:
		var newProduct Product
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bodyBytes, &newProduct)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if newProduct.ProductID != 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		newProduct.ProductID = getNextID()

		productList = append(productList, newProduct)
		w.WriteHeader(http.StatusCreated)
	}

}

func main() {
	http.HandleFunc("/products", productsHandler)
	http.ListenAndServe(":5000", nil)
}
