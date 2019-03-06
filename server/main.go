package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	log.Println("starting server...")

	s := &server{
		port: 8080,
		db:   testData(),
	}

	go func() {
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", s.port), s))
	}()

	log.Printf("server started on localhost:%d\n", s.port)
	<-stop
}
