package main

// Импортирум модули
import (
	"fmt"
	kitty "mykittybot/pkg/botapi"
	"mykittybot/pkg/filereader"
	"mykittybot/pkg/weatherapi"
    "strconv"
    "strings"
)

func main() {
	progSettings := filereader.SettingsRead("../../settings.ini")
	kitty.InitBot(progSettings["botTocken"])            // Инициализация телеграм-бота
	weatherapi.InitWehather(progSettings["weatherKey"]) // Инициализация API - погоды

	var lastMessages []string // Пул последних сообщений бота
	var ids []int

	messageQuerry := kitty.GetUpdatesT{} // Структура для хранения структуры сообщений чата телеграм
	weather := weatherapi.WeatherQueryT{} // Структура для хранения данных о погоде

    for {
		messageQuerry.GetUpdate()
		for _, messmessage := range messageQuerry.Result {
			lastMessages = append(lastMessages, fmt.Sprintf("%v", messmessage.Message.Text))
			ids = append(ids, messmessage.Message.Chat.ID)
			kitty.ClearQueue()
		}
		for n, message := range lastMessages {
            search := strings.ToLower(message)
            if search == "картинка" {
				kitty.SendPhotoByIrl(&ids[n], "https://www.ixbt.com/img/n1/news/2019/5/3/chrome-73-mode-sombre-android_large.jpg")
			}
            if search == "погода"{
                weather.GetWeather()
                kitty.SendMessage(&ids[n], "Погода в г.Железногорск сегодня")
                kitty.SendMessage(&ids[n], "")
                send := "Температура воздуха составляет: "+ strconv.Itoa(int(weather.Fact.Temp)) +" °C"
                kitty.SendMessage(&ids[n], send)

            }
            if search == "stop" || search == "стоп"{
				return
			}
			lastMessages = nil
		}
            	}
}
