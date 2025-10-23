package main

import (
	"flag"
	"fmt"
	"load_balancer/balancer"
	server "load_balancer/cmd"
	"load_balancer/proxy"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// healthCheck runs a routine for check status of the backends every 2 mins
func healthCheck(serverPool *server.ServerPool) {
	t := time.NewTicker(time.Minute * 2)
	for range t.C {
		log.Println("Starting health check...")
		serverPool.HealthCheck()
		log.Println("Health check completed")
	}
}

func main() {
	var serverPool server.ServerPool
	var serverList string
	var port int
	flag.StringVar(&serverList, "backends", "", "Load balanced backends, use commas to separate")
	flag.IntVar(&port, "port", 3030, "Port to serve")
	flag.Parse()

	if len(serverList) == 0 {
		log.Fatal("Please provide one or more backends to load balance")
	}

	// parse servers
	tokens := strings.Split(serverList, ",")
	for _, tok := range tokens {
		serverUrl, err := url.Parse(tok)
		if err != nil {
			log.Fatal(err)
		}

		proxy := proxy.SetUpProxy(serverUrl, serverPool)
		serverPool.AddBackend(&server.Backend{
			URL:          serverUrl,
			Alive:        true,
			ReverseProxy: proxy,
		})
		log.Printf("Configured server: %s\n", serverUrl)
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		balancer.LB(w, r, &serverPool)
	}

	// create http server
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: http.HandlerFunc(handler),
	}

	// start health checking
	go healthCheck(&serverPool)

	log.Printf("Load Balancer started at :%d\n", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
