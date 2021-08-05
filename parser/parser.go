// THERE ARE 8 DAYS
package parser

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/dustin-ward/minecraft-time-logging/data"
	"github.com/dustin-ward/minecraft-time-logging/util"
)

var Users = make(map[string]*data.User)
var WorkingDate time.Time

func Parse(path string) {
	file, err := os.Open(path)
	util.ErrorCheck(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var prevHrs int
	var timestamp time.Time
	for scanner.Scan() {
		line := scanner.Text()
		var hrs, mins, secs int

		// Initial Date
		r, _ := regexp.Compile(`^\d+-\d+-\d+`)
		if idx := r.FindStringIndex(line); idx != nil {
			if WorkingDate, err = time.Parse("2006-01-02", line[idx[0]:idx[1]]); err != nil {
				log.Fatal(err.Error())
			}
			prevHrs, _ = strconv.Atoi(line[idx[1]+1 : idx[1]+3])
		} else {
			hrs, _ = strconv.Atoi(line[1:3])
			mins, _ = strconv.Atoi(line[4:6])
			secs, _ = strconv.Atoi(line[7:9])

			if hrs < prevHrs {
				WorkingDate = WorkingDate.AddDate(0, 0, 1)
			}
			prevHrs = hrs
		}

		timestamp = WorkingDate.Add(time.Hour*time.Duration(hrs) + time.Minute*time.Duration(mins) + time.Second*time.Duration(secs))
		fmt.Println(timestamp)

		// Joining the game
		r, _ = regexp.Compile(`\w+ joined the game`)
		if idx := r.FindStringIndex(line); idx != nil {
			username := line[idx[0] : idx[1]-16]
			if _, exists := Users[username]; !exists {
				Users[username] = &data.User{
					Username: username,
					Messages: []data.Message{},
				}
			}

			if Users[username].InSession {
				log.Fatal("ERROR: User", username, "already in session")
			}

			session := data.Session{Start: timestamp}
			Users[username].InSession = true
			Users[username].Sessions = append(Users[username].Sessions, session)
		}

		// Leaving the game
		r, _ = regexp.Compile(`\w+ left the game`)
		if idx := r.FindStringIndex(line); idx != nil {
			username := line[idx[0] : idx[1]-14]
			if _, exists := Users[username]; !exists {
				log.Fatal("ERROR: User", username, "does not exist")
			}

			if !Users[username].InSession {
				log.Fatal("ERROR: User", username, "not in session")
			}

			EndSession(username, timestamp)
		}

		// Message sent
		r, _ = regexp.Compile(`<.+>`)
		if idx := r.FindStringIndex(line); idx != nil {
			username := line[idx[0]+1 : idx[1]-1]
			message := data.Message{Timestamp: timestamp, Content: line[idx[1]+1:]}
			Users[username].Messages = append(Users[username].Messages, message)
			Users[username].MessageCount += 1
		}
		fmt.Print("\n")
	}

	// End all sessions
	for _, user := range Users {
		if user.InSession {
			EndSession(user.Username, timestamp)
		}
	}

	util.ErrorCheck(scanner.Err())
}

func EndSession(username string, timestamp time.Time) {
	s := Users[username].Sessions[len(Users[username].Sessions)-1]
	s.End = timestamp
	s.Duration = timestamp.Sub(s.Start)
	Users[username].Sessions[len(Users[username].Sessions)-1] = s
	Users[username].InSession = false

	Users[username].TotalTime += s.Duration
}
