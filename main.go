package main

import "log"

func main() {
	store, err := NewDB()
	if err != nil {
		log.Fatal(err)
	}
	server := NewAPIServer(":4000", store)
	server.Run()
}
