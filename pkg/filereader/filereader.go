package filereader

import (
	"log"
	"os"
	"strings"
)

func SettingsRead(path string) map[string]string {
	set := make(map[string]string)
	data, err := os.ReadFile(path)
	if err != nil {
		log.Print(err)
	}
	lines := strings.Split(string(data), "\n") // Разделяем строки на массив
	for _, line := range lines {
		chunc := strings.Split(strings.TrimRight(line, "\r\n"), "=")
		if len(chunc) == 2 {
			key, value := strings.Trim(chunc[0], "\t "), strings.Trim(chunc[1], "\t ")

			last := len(value) - 1

			if last > 0 && value[0] == '"' && value[last] == '"' {
				value = string(value[1:last])
			}
			set[key] = string(value)
		}
	}
	return set
}
