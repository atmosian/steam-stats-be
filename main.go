package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/atmosian/steam-stats-be/api"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/players/{player-id}/achievments", api.GetAchievementsByPlayerID)
	
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", router))
}
