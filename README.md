### Creating Handlers
- http.Handle
- http.HandleFunc
Example:
```go
package main

import (
	"net/http"
)

type fooHandler struct {
	Message string
}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(f.Message))
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bar is called"))
}

func main() {
	http.Handle("/foo", &fooHandler{Message: "Foo is called"})
	http.HandleFunc("/bar", barHandler)
	http.ListenAndServe(":5000", nil)
}

```
### Working with Json
- Marshal -- encodes
- Unmarshal -- decodes
```go
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
```

### Request Object
#### Request.Method string  
#### Request.Header Header (map[string][]string)  
#### Request.Body  
- io.ReadCloser: Mostly json Marshall/unmarshall is used
- Returns EOF when not present