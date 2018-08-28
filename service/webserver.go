package service

import (
	"log"
	"net/http"
)

func StartWebServer(port string) {
	log.Println("Starting HTTP service at " + port)
	r := NewRouter()
	http.Handle("/", r)

	err := http.ListenAndServe(":"+port, r)

	if err != nil {
		log.Println("An error ocurred starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}
}
