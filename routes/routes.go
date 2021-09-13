package routes

import (
	"html/template"
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
func HomePage(w http.ResponseWriter, r *http.Request) {
	PageVariables := PageVariables{
		Title: "Users:",
		Users: parser.Users,
	}

	t, err := template.ParseFiles("public/html/main.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}

	err = t.Execute(w, PageVariables)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}

// Page for an individual player
//
func UserPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if username, ok := vars["username"]; !ok {
		log.Fatal("Missing username in params")
	} else {
		if user, exists := parser.Users[username]; !exists {
			w.Write([]byte("User doesnt exist"))
		} else {
			PageVariables := PageVariables{
				Title: username,
				User:  user,
			}

			t, err := template.ParseFiles("public/html/user.html")
			if err != nil {
				log.Print("template parsing error: ", err)
			}

			err = t.Execute(w, PageVariables)
			if err != nil {
				log.Print("template executing error: ", err)
			}
		}
	}

}
