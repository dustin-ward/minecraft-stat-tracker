package data

import "time"

type User struct {
	Id           int
	Username     string
	TotalTime    float64
	MessageCount int
	InSession    bool
	Sessions     []Session
	Messages     []Message
}

type Session struct {
	Start time.Time
	End   time.Time
}

type Message struct {
	Timestamp time.Time
	Content   string
}
