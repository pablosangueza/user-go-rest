package main

import (
	"log"
	"net/http"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
)

func main() {
	// create the Swagger UI handler
	swaggerHandler := httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	)

	// create the API handler
	apiHandler := middleware.Spec("",
		nil,
		loads.Embedded(listUsers),
	)

	// create the HTTP server
	http.Handle("/swagger/", swaggerHandler)
	http.Handle("/swagger/doc.json", apiHandler)
	//http.HandleFunc("/users", listUsers)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
