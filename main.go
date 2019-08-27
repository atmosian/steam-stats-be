package main

import (
	"log"
	"net/http"
	// "os"

	"github.com/atmosian/steam-stats-be/api"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/players/{player-id}/achievments", api.GetAchievementsByPlayerID)
	router.HandleFunc("/players/{player-id}/games", api.GetOwnedGamesByPlayerID)
	
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", router))
}
