package product

type Product struct {
	ProductID    int    `json: "productId"`
	Manufacturer string `json: "manufacturer"`
	Sku          string `json: "sku"`
	Upc          string `json: "upc"`
	PricePerUnit int    `json: "pricePerUnit"`
	Stock        int    `json: "stock"`
	ProductName  string `json: "productName"`
}
