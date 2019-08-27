package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

// GetAchievementsByPlayerID getting player achievements
func GetAchievementsByPlayerID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Achievements...\n"))
}

// GetOwnedGamesByPlayerID getting player games info
func GetOwnedGamesByPlayerID(w http.ResponseWriter, r *http.Request) {
	apiKey := os.Getenv("STEAM_API_KEY")
	vars := mux.Vars(r)
	steamID := vars["player-id"]
	var response Response

	httpClient := &http.Client{
		Timeout: 3 * time.Second,
	}

	// Build a request
	req, _ := http.NewRequest("GET", "http://api.steampowered.com/IPlayerService/GetOwnedGames/v0001/", nil)
	q := req.URL.Query()
	q.Add("key", apiKey)
	q.Add("steamid", steamID)
	q.Add("include_appinfo", "1")
	q.Add("format", "json")
	req.URL.RawQuery = q.Encode()
	
	httpResp, httpErr := httpClient.Do(req)

	if httpErr != nil {
		log.Fatal(httpErr)
		http.Error(w, httpErr.Error(), 500)
		return
	}

	defer httpResp.Body.Close()

	// A HTTP response deserialization
	body, readErr := ioutil.ReadAll(httpResp.Body)
	if readErr != nil {
		log.Fatal(readErr)
		http.Error(w, readErr.Error(), 500)
		return
	}

	jsonErr := json.Unmarshal(body, &response)
	if jsonErr != nil {
		log.Fatal(jsonErr)
		http.Error(w, jsonErr.Error(), 500)
		return
	}

	// Build a response
	json.NewEncoder(w).Encode(&response.Body.Games)
}

// Game structure
type Game struct {
	Appid                    int32  `json:"appid"`
	Name                     string `json:"name"`
	PlaytimeForever          int32  `json:"playtime_forever"`
	ImgIconURL               string `json:"img_icon_url"`
	ImgLogoURL               string `json:"img_logo_url"`
	HasCommunityVisibleStats bool   `json:"has_community_visible_stats"`
}

type PlayerGames struct {
	GamesCount int32  `json:"game_count"`
	Games      []Game `json:"games"`
}

type Response struct {
	Body PlayerGames `json:"response"`
}
