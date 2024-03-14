package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type SuperHero struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Power   string   `json:"power"`
	Creator *Creator `json:"creator"`
}

type Creator struct {
	OrgName  string `json:"orgName"`
	Category string `json:"category"`
}

var superHeros []SuperHero

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/heroes", getHeroes).Methods("GET")
	r.HandleFunc("/heroes/{id}", getHero).Methods("GET")
	r.HandleFunc("/heroes", createMovie).Methods("POST")
	r.HandleFunc("/heroes/{id}", updateHero).Methods("PUT")
	r.HandleFunc("/heroes/{id}", deleteHero).Methods("DELETE")

	fmt.Printf("Server started at port 8000\n")
	
	log.Fatal(http.ListenAndServe(":8000", r))
	fmt.Println("SuperHeroes CRUD Api")
}
