package main

// Импортирум модули
import (
	"mykittybot/library/botapi"
	"mykittybot/library/filereader"
)

func main() {
	progSettings := filereader.SettingsRead("settings.ini")
	botapi.InitBot(&progSettings["botTocken"])

}
