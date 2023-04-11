package handlers

import (
	"log"
	"net/http"
	"os"
	"src/codetwitter/middlew"
	"src/codetwitter/routers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Manejadores() {

	router := mux.NewRouter()
	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"

	}
	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
