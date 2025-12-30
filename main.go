package main

import (
	"fmt"
	"net/http"
	"proj3/router"
)

// FONCTION PRINCIPAL
func main() {
	r := router.New()
	fmt.Println("http://localhost:8080/home")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println("Erreur serveur:", err)
	}
}
