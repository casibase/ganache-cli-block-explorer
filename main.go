package main

import (
	"ganache-cli-block-explorer/conf"
	"ganache-cli-block-explorer/router"
	"log"
	"net/http"
	"time"
)

func main() {

	log.Println("!!!!INITIALIZING SERVER!!!!")
	config := conf.LoadConfig("conf/app.yaml")
	gorilla := router.InitRouter(config)

	// http server
	// Note: Here gorilla is like passing our own server handler into net/http, by default its false
	srv := &http.Server{
		Handler: gorilla,
		Addr:    config.ServerAddr,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Printf("SERVER STARTED at ADDRESS : %s\n", config.ServerAddr)
	log.Fatal(srv.ListenAndServe())
}
