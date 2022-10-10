package main

// Импортирум модули
import (
	"fmt"
	"mykittybot/pkg/botapi"
	"mykittybot/pkg/filereader"
)

func main() {
	progSettings := filereader.SettingsRead("../../settings.ini")
	botapi.InitBot(progSettings["botTocken"])
	messageQuerry := new(botapi.GetUpdatesT)
	messageQuerry.GetUpdate()

	var lastMessages []string //

	for _, messmessage := range messageQuerry.Result {
		lastMessages = append(lastMessages, fmt.Sprintf("%v", messmessage.Message.Text))
		fmt.Println(messmessage.Message.Chat.ID, lastMessages)
		botapi.ClearQueue()
	}

}
