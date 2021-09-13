package util

import (
	"log"
	"time"
)

func ErrorCheck(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func ParsingError(path string, lineNo int, message string) {
	log.Fatal("PARSING ERROR: ", path, ":", lineNo, "\n", message)
}

func GetTimestamp(line string) time.Time {
	t, err := time.Parse("02Jan2006 15:04:05", line[1:23])
	if err != nil {
		log.Fatal(err.Error())
	}
	return t
}
