package main

import (
	"broker/cmd/api/routes"
	"fmt"
	"log"
	"net/http"
)

const webPort = "80"

func main() {
	app := routes.Config{}

	log.Printf("starting broker on port %s\n", webPort)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.Routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
