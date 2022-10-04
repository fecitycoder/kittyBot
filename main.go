package main

// Импортирум модули
import (
	"mykittybot/library/botapi"
	"mykittybot/library/filereader"
	"strings"
)

func main() {
	progSettings := filereader.SettingsRead("settings.ini")
	tocken := progSettings["botTocken"]
	botapi.InitBot(&tocken)
	var id int
	last2 := [5]string{"Мне 15 лет", "Привет. Меня зовут Вася", "Я мальчик", "Работаю в ПАО МГОК", "Живу в твоем воображении"}

	for {
		messages := botapi.GetUpdate(&id)
		if len(messages) != 0 {
			lastMsg := strings.ToLower(messages[0])
			if strings.Contains(lastMsg, "привет") {
				botapi.SendMessage(&id, last2[1])
			}

			if strings.Contains(lastMsg, "лет") {
				botapi.SendMessage(&id, last2[0])
			}

			if strings.Contains(lastMsg, "работаеш") {
				botapi.SendMessage(&id, last2[3])
			}

			if strings.Contains(lastMsg, "живеш") {
				botapi.SendMessage(&id, last2[4])
			}

		}
	}

}

// "Ну здравствуй! Дорогой мой незнакомец..."
