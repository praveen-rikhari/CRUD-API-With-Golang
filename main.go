package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type SuperHero struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Power   string   `json:"power"`
	Creator *Creator `json:"creator"`
}

type Creator struct {
	OrgName  string `json:"orgName"`
	Category string `json:"category"`
}

var superHeros []SuperHero

func getHeroes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(superHeros)
}

func deleteHero(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range superHeros {
		if item.ID == params["id"] {
			superHeros = append(superHeros[:index], superHeros[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(superHeros)
}

func getHero(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range superHeros {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createHero(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var superHero SuperHero
	_ = json.NewDecoder(r.Body).Decode(&superHero)
	superHero.ID = strconv.Itoa(rand.Intn(100000000))
	superHeros = append(superHeros, superHero)
	json.NewEncoder(w).Encode(superHero)
}

func updateHero(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range superHeros {
		if item.ID == params["id"] {
			superHeros = append(superHeros[:index], superHeros[index+1:]...)
			var superHero SuperHero
			_ = json.NewDecoder(r.Body).Decode(&superHero)
			superHero.ID = params["id"]
			superHeros = append(superHeros, superHero)
			json.NewEncoder(w).Encode(superHero)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()

	superHeros = append(superHeros, SuperHero{
		ID:    "1",
		Name:  "Batman",
		Power: "Money, Will Power",
		Creator: &Creator{
			OrgName:  "DC",
			Category: "Dark, Crime, Thriller",
		},
	})
	superHeros = append(superHeros, SuperHero{
		ID:    "2",
		Name:  "Ironman",
		Power: "Tech, Suit",
		Creator: &Creator{
			OrgName:  "Marvel",
			Category: "Action, Sci-Fi, Thriller",
		},
	})

	r.HandleFunc("/heroes", getHeroes).Methods("GET")
	r.HandleFunc("/heroes/{id}", getHero).Methods("GET")
	r.HandleFunc("/heroes", createHero).Methods("POST")
	r.HandleFunc("/heroes/{id}", updateHero).Methods("PUT")
	r.HandleFunc("/heroes/{id}", deleteHero).Methods("DELETE")

	fmt.Printf("Server started at port 8000\n")

	log.Fatal(http.ListenAndServe(":8000", r))
	fmt.Println("SuperHeroes CRUD Api")
}
