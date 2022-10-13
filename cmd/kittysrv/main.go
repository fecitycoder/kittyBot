package main

// Импортирум модули
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mykittybot/pkg/botapi"
	"mykittybot/pkg/filereader"
	"mykittybot/pkg/weatherapi"
	"net/http"
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

	client := &http.Client{}
	req, err := http.NewRequest(
		"GET", "https://api.weather.yandex.ru/v2/informers?lat=52.339204&lon=35.350873", nil,
	)

	req.Header.Add("X-Yandex-API-Key", progSettings["weatherKey"])

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &weather)

	fmt.Println(weather.Fact.Temp)
}
