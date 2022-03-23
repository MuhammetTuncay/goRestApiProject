package main

import (
	"fmt"
	"gogo/app"
	controllers "gogo/controller"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Use(app.JwtAuthentication) // Middleware'e JWT kimlik doğrulaması eklenir

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")

	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")

	port := os.Getenv("PORT") // Environment dosyasından port bilgisi getirilir
	if port == "" {
		port = "5858" //localhost:8000
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router) // Uygulamamız localhost:8000/api altında istekleri dinlemeye başlar
	if err != nil {
		fmt.Print(err)
	}
}
