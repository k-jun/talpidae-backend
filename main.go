package main

import (
	"log"
	"os"
	"strconv"
	"talpidae-backend/server"
)

func main() {

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8080
	}
	s := server.NewServer(port)
	err = s.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
