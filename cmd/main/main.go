package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JohnnyChangTW/go-learning-3-book-management-application/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("starting server")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
