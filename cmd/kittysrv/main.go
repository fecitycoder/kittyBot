package main

// Импортирум модули
import (
	"fmt"
	"mykittybot/pkg/botapi"
	"mykittybot/pkg/filereader"
	"mykittybot/pkg/weatherapi"
)

func main() {
	progSettings := filereader.SettingsRead("../../settings.ini")
	botapi.InitBot(progSettings["botTocken"])
	weatherapi.InitWehather(progSettings["weatherKey"])
	messageQuerry := new(botapi.GetUpdatesT)
	messageQuerry.GetUpdate()

	var lastMessages []string //
	weather := weatherapi.WeatherQueryT{}

	for _, messmessage := range messageQuerry.Result {
		lastMessages = append(lastMessages, fmt.Sprintf("%v", messmessage.Message.Text))
		fmt.Println(messmessage.Message.Chat.ID, lastMessages)
		botapi.ClearQueue()
	}
	weather.GetWeather()
	fmt.Println(weather)
}
