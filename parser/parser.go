package parser

import (
	"bufio"
	"os"
	"regexp"
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

	var startTime time.Time

	lineNo := 1
	for scanner.Scan() {
		line := scanner.Text()

		// Log starting time
		if startTime.IsZero() {
			startTime = util.GetTimestamp(line)
		}

		// Fatal Error Catch
		r, _ := regexp.Compile(`\[.+\/FATAL\].+crash`)
		if idx := r.FindStringIndex(line); idx != nil {
			// for user := range Users {
			// 	if Users[user].InSession {
			// 		EndSession(user, util.GetTimestamp(line))
			// 	}
			// }
		}

		// Joining the game
		r, _ = regexp.Compile(`\w+ joined the game`)
		if idx := r.FindStringIndex(line); idx != nil {
			username := line[idx[0] : idx[1]-16]
			data.CreateUser(username)
			// if !IsUser(username) {
			// 	CreateUser(username)
			// }

			// if Users[username].InSession {
			// 	util.ParsingError(path, lineNo, "User "+username+" already in session")
			// }

			// session := data.Session{Start: util.GetTimestamp(line)}
			// Users[username].InSession = true
			// Users[username].Sessions = append(Users[username].Sessions, session)
		}

		// Leaving the game
		r, _ = regexp.Compile(`\w+ left the game`)
		if idx := r.FindStringIndex(line); idx != nil {
			// username := line[idx[0] : idx[1]-14]
			// if !IsUser(username) {
			// 	util.ParsingError(path, lineNo, "User "+username+" does not exist")
			// }

			// if !Users[username].InSession {
			// 	util.ParsingError(path, lineNo, "User "+username+" not in session")
			// }

			// EndSession(username, util.GetTimestamp(line))
		}

		// Message sent
		r, _ = regexp.Compile(`: <[a-zA-Z0-9_]{2,16}>`)
		if idx := r.FindStringIndex(line); idx != nil {
			// username := line[idx[0]+3 : idx[1]-1]
			// message := data.Message{Timestamp: util.GetTimestamp(line), Content: line[idx[1]+1:]}
			// // fmt.Println("Found msg", message, "from", username)
			// Users[username].Messages = append(Users[username].Messages, message)
			// Users[username].MessageCount += 1
		}

		lineNo++
	}

	// End all sessions
	// SHOULDNT HAPPEN... BUT ADD ERROR CATCH HERE
	// It happened... but im still not fixing it
	// for _, user := range Users {
	// 	if user.InSession {
	// 		EndSession(user.Username, timestamp)
	// 	}
	// }

	util.ErrorCheck(scanner.Err())
}

func IsUser(username string) bool {
	if _, exists := Users[username]; exists {
		return true
	} else {
		return false
	}
}

// func EndSession(username string, timestamp time.Time) {
// 	s := Users[username].Sessions[len(Users[username].Sessions)-1]
// 	s.End = timestamp
// 	s.Duration = timestamp.Sub(s.Start)
// 	Users[username].Sessions[len(Users[username].Sessions)-1] = s
// 	Users[username].InSession = false

// 	Users[username].TotalTime += s.Duration
// }
