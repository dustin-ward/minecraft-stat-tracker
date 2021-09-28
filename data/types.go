package data

import "time"

// type User struct {
// 	Id           int
// 	Username     string
// 	TotalTime    time.Duration
// 	MessageCount int
// 	InSession    bool
// 	Sessions     []Session
// 	Messages     []Message
// }

type User struct {
	id           uint32 `json:"id"`
	username     string `json:"username"`
	totalTime    uint32 `json:"totalTime"`
	messageCount uint32 `json:"messageCount"`
	deathCount   uint32 `json:"messageCount"`
}

// type Session struct {
// 	Start    time.Time
// 	End      time.Time
// 	Duration time.Duration
// }

type Session struct {
	id    uint32    `json:"id"`
	user  uint32    `json:"user"`
	start time.Time `json:"start"`
	end   time.Time `json:"end"`
}

// type Message struct {
// 	Timestamp time.Time
// 	Content   string
// }

type Message struct {
	id        uint32    `json:"id"`
	user      uint32    `json:"user"`
	content   string    `json:"content"`
	timestamp time.Time `json:"timestamp"`
}

type Death struct {
	id        uint32    `json:"id"`
	user      uint32    `json:"user"`
	content   string    `json:"content"`
	timestamp time.Time `json:"timestamp"`
}
