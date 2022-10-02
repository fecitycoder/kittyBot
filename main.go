package main

// Импортирум модули
import (
	"fmt"
	"mykittybot/library/filereader"
)

func main() {

	// Если чтение данных прошло успено
	// выводим их в консоль
	line := filereader.SettingsRead("settings.ini")
	fmt.Println(line["botTocken"])
}
