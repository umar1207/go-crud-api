package main

import (
	"fmt"
	"net/http"

	"github.com/umar1207/go-crud-api/services"
)

func main() {
	api := http.NewServeMux()
	api.HandleFunc("GET /items", services.GetItems)
	api.HandleFunc("GET /items/{id}", services.GetItemById)
	api.HandleFunc("POST /items/create", services.CreateItem)
	api.HandleFunc("POST /items/update/{id}", services.UpdateItem)
	api.HandleFunc("DELETE /items/delete/{id}", services.DeleteItem)

	fmt.Println("Listening on PORT 8080")
	http.ListenAndServe(":8080", api)
}
