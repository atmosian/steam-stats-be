package api

import "net/http"

// GetAchievementsByPlayerID getting player achievements
func GetAchievementsByPlayerID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Achievements...\n"))
}
