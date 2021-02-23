package main

import (
	"log"
	"myProject/test/config"
	"net/http"
	"strconv"
)

func main() {
	port := config.Config.Port
	path := config.Config.Path

	http.Handle("/notes/", http.StripPrefix("/notes/", http.FileServer(http.Dir(path))))
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}
