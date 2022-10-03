package main

// Импортирум модули
import (
	"fmt"
	"mykittybot/library/botapi"
	"mykittybot/library/filereader"
)

func main() {
	progSettings := filereader.SettingsRead("settings.ini")
	botapi.InitBot(progSettings["botTocken"])

	messages, _ := botapi.GetUpdate()
	//last := strings.ToLower(messages[len(messages)-1])
	for _, mes := range messages {
		fmt.Println(mes)
	}
}
