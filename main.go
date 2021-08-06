package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/dustin-ward/minecraft-time-logging/data"
	"github.com/dustin-ward/minecraft-time-logging/parser"

	"github.com/gorilla/mux"
)

type PageVariables struct {
	Title string
	Users map[string]*data.User
	User  *data.User
}

func main() {
	dir, _ := os.Getwd()
	parser.Parse(dir + "/logs_1.txt")

	// fmt.Println(parser.WorkingDate)

	// for _, user := range parser.Users {
	// 	fmt.Println("=USER=====================")
	// 	fmt.Println("Username:", user.Username)
	// 	fmt.Println("TotalMessages:", user.MessageCount)
	// 	fmt.Println("Messages:")
	// 	for _, m := range user.Messages {
	// 		fmt.Println("   ", m.Timestamp, m.Content)
	// 	}
	// 	fmt.Println("TotalTime:", user.TotalTime)
	// 	fmt.Println("Sessions:")
	// 	for _, s := range user.Sessions {
	// 		fmt.Println("    Start:", s.Start, " End:", s.End, "Duration:", s.Duration)
	// 	}
	// }

	r := mux.NewRouter()
	r.HandleFunc("/", HomePage)
	r.HandleFunc("/user/{username}", UserPage)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))
	log.Fatal(http.ListenAndServe(":8080", r))
}

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
