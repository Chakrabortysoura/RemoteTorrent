package main

import (
	"RemoteTorrent/api"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Start of the Server")
	server := http.Server{
		Addr:    ":8080",
		Handler: api.NewRequestMatcher(),
	}
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Error starting the server", err)
	}
}
