package main

import (
	"log"
	"talpidae-backend/server"
)

func main() {

	port := 8080
	s := server.NewServer(port)
	err := s.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
