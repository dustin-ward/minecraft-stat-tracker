package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dustin-ward/minecraft-time-logging/parser"
)

func main() {
	dir, _ := os.Getwd()
	parser.Parse(dir + "/logs_1.txt")

	fmt.Println(parser.WorkingDate)

	for _, user := range parser.Users {
		fmt.Println("=USER=====================")
		fmt.Println("Username:", user.Username)
		fmt.Println("Messages:")
		for _, m := range user.Messages {
			fmt.Println("   ", m.Timestamp, m.Content)
		}
		fmt.Println("Sessions:")
		for _, s := range user.Sessions {
			fmt.Println("    Start:", s.Start, " End:", s.End)
		}
	}
}

func Error(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
