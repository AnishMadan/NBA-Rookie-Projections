package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Player is a struct that represents a single player
type Player struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var players []Player = []Player{}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/players", addPlayer).Methods("POST")

	router.HandleFunc("/players", getAllPlayers).Methods("GET")

	router.HandleFunc("/players/{id}", getPlayer).Methods("GET")

	router.HandleFunc("/players/{id}", updatePlayer).Methods("PUT")

	router.HandleFunc("/players/{id}", patchPlayer).Methods("PATCH")

	router.HandleFunc("/players/{id}", deletePlayer).Methods("DELETE")

	error := http.ListenAndServe(":5000", router)

	if error != nil {
		log.Fatal(error)
	}
}

func getPlayer(w http.ResponseWriter, r *http.Request) {
	// get the ID of the post from the route parameter
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		// there was an error
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}

	// error checking
	if id >= len(players) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified ID"))
		return
	}

	post := players[id]

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func getAllPlayers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(players)
}

func addPlayer(w http.ResponseWriter, r *http.Request) {
	// get Item value from the JSON body
	var newPlayer Player
	json.NewDecoder(r.Body).Decode(&newPlayer)

	players = append(players, newPlayer)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(players)
}

func updatePlayer(w http.ResponseWriter, r *http.Request) {
	// get the ID of the post from the route parameters
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}

	// error checking
	if id >= len(players) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified ID"))
		return
	}

	// get the value from JSON body
	var updatedPlayer Player
	json.NewDecoder(r.Body).Decode(&updatedPlayer)

	players[id] = updatedPlayer

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedPlayer)
}

func patchPlayer(w http.ResponseWriter, r *http.Request) {
	// get the ID of the post from the route parameters
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}

	// error checking
	if id >= len(players) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified ID"))
		return
	}

	// get the current value
	post := &players[id]
	json.NewDecoder(r.Body).Decode(post)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func deletePlayer(w http.ResponseWriter, r *http.Request) {
	// get the ID of the post from the route parameters
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}

	// error checking
	if id >= len(players) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified ID"))
		return
	}

	// Delete the post from the slice
	// https://github.com/golang/go/wiki/SliceTricks#delete
	players = append(players[:id], players[id+1:]...)

	w.WriteHeader(200)
}
