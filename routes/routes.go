package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dustin-ward/minecraft-time-logging/data"
	"github.com/dustin-ward/minecraft-time-logging/parser"
	"github.com/gorilla/mux"
)

type PageVariables struct {
	Title string
	Users map[string]*data.User
	User  *data.User
}

// Starting page with all users
//
func TestFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: '/'")

	w.Header().Set("content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, "Minecraft Stat-tracker API is operational!")
}

// Return list of all users
//
func GetPlayers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: '/players'")

	w.Header().Set("content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(parser.Users)
}

// Return an individual user
//
func GetPlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	if username, ok := vars["username"]; !ok {
		log.Fatal("GetPlayer: Missing username in params")
	} else {
		if user, exists := parser.Users[username]; !exists {
			w.Write([]byte("User doesnt exist"))
		} else {
			fmt.Println("Endpoint Hit: '/player/" + user.Username + "'")
			json.NewEncoder(w).Encode(user)
		}
	}
}
