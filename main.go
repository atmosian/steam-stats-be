package main

import (
	"log"
	"net/http"

	"github.com/atmosian/steam-stats-be/api"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/players/{player-id}/achievments", api.GetAchievementsByPlayerID)
	router.HandleFunc("/players/{player-id}/games", api.GetOwnedGamesByPlayerID)

	// Bind to a port and pass our router in
	port := ":8000"
	log.Printf("[INFO] Web server has been started on %s port", port)
	log.Fatal(http.ListenAndServe(port, router))
}
