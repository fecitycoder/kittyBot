package main

// Импортирум модули
import (
	"fmt"
	"mykittybot/library/botapi"
	"mykittybot/library/filereader"
)

func main() {
	progSettings := filereader.SettingsRead("settings.ini")
	tocken := progSettings["botTocken"]
	botapi.InitBot(&tocken)
	var id int
	messages := botapi.GetUpdate(&id)
	//last := strings.ToLower(messages[len(messages)-1])
	for _, mes := range messages {
		fmt.Println(mes, id)
	}
}
