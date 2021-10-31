## ServMux and Handlers
- Create ServeMux and attach handlers, then service that ServeMux with `ListenAndServe`
- You can use default ServeMux
#### Using Handler
- define custom types (Struct), and implement ServeHTTP function
```go
type fooHandler struct {
	Message string
}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(f.Message))
}

```
- Register in serveMux
```go
func main() {
	fh := &fooHandler{Message: "foo is called"}
	http.Handle("/foo", fh)
	http.ListenAndServe(":5000", nil)
}
```
#### Using HandleFunc
- Any method which has signature `func(http.ResponseWriter, *http.Request)` can be converted to HandlerFunc type
Example: 
```go
func fooHandler(w http.ResponseWriter, r *http.Request) {
}
```
After you created your function, convert it to `HandlerFunc` and attach it to serveMux OR  
you can use `HandleFunc` directly
```go
package main

import (
	"log"
	"net/http"
	"time"
)

func timeHandler(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(time.RFC1123)
	w.Write([]byte("The time is: " + tm))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/time", timeHandler)

	log.Println("Listening..")
	http.ListenAndServe(":3001", mux)
}
```
### Using closure to pass variables to handler

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

### URL Path parameters
- Add new handler
```go
	http.HandleFunc("/products/", productHandler)
```

### Middleware
- [ client <---> server <---> HTTP Mux <----Middleware----> Handler ]
- Is client authenticated ? 
- Do they have enough permissions ?
- logging
- load session data

#### Wrap your handler with Adapter function
## References
1. https://www.alexedwards.net/blog/a-recap-of-request-handling
