package main

import (
	"fmt"
	"net/http"

	"std_lib_server/middleware"
)

func handlePing(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func handleBaseID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.Write([]byte("Id:" + id))
}

func handlePostBaseID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.Write([]byte("Id:" + id))
}

func main() {
	v1 := http.NewServeMux()
	router := http.NewServeMux()
	v1.Handle("/v1/", http.StripPrefix("/v1", router))
	router.HandleFunc("/ping", handlePing)
	router.HandleFunc("/base/{id}", handleBaseID)
	router.HandleFunc("POST /base/{id}", handlePostBaseID)

	server := http.Server{
		Addr:    ":8080",
		Handler: middleware.Log(router),
	}

	fmt.Println(`
 _____     _    _                        _                 _                  _      _   ___ ___ ___ ___
|  _  |___|_|  |_|___    ___ _ _ ___ ___|_|___ ___    ____| |_    ___ ___ ___| |_   |_| | . |   | . |   |
|     | . | |  | |_ -|  |  _| | |   |   | |   | . |  | .' |  _|  | . | . |  _|  _|   _  | . | | | . | | |
|__|__|  _|_|  |_|___|  |_| |___|_|_|_|_|_|_|_|_  |  |__,_|_|    |  _|___|_| |_|    |_| |___|___|___|___|
      |_|                                     |___|              |_|
	`)

	server.ListenAndServe()
}
