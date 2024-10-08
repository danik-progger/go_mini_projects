package main
import (
	 "net/http"
	 "log"
	 "github.com/gorilla/mux"
	 _ "github.com/jinzhu/gorm/dialects/mysql"
	 "mysql_bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookstoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}