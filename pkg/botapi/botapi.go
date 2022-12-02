package botapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

const telegramBaseUrl = "https://api.telegram.org/bot"

var botTocken string

func InitBot(tocken string) {
	botTocken = tocken
}

func getUrlByMethod(methodName string) string {
	return telegramBaseUrl + botTocken + "/" + methodName
}

func getBodyByUrl(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	return body
}

func GetMe() {
	body := getBodyByUrl(getUrlByMethod("getMe"))
	getMe := GetMeT{}
	err := json.Unmarshal(body, &getMe)
	if err != nil {
		fmt.Printf("Error unmarshalling: %s", err.Error())
	}
	fmt.Printf("%v", getMe)
}

func (getUpdates *GetUpdatesT) GetUpdate() {
	// Формируем url запроса и получаем ответ, который складываем в body
	body := getBodyByUrl(getUrlByMethod("getUpdates"))
	err := json.Unmarshal(body, getUpdates)
	if err != nil {
		fmt.Printf("Error unmarshalling: %s", err.Error())
	}
}

func ClearQueue() {
	url := getUrlByMethod("deleteWebhook")
	url = url + "?drop_pending_updates=true"
	getBodyByUrl(url)
}

func SendMessage(chatId *int, message string) {
	url := getUrlByMethod("sendMessage")
	url = url + "?chat_id=" + strconv.Itoa(*chatId) + "&text=" + message
	getBodyByUrl(url)
}

func SendPhotoByIrl(chatId *int, photo string) {
	url := getUrlByMethod("sendPhoto")
	url = url + "?chat_id=" + strconv.Itoa(*chatId) + "&photo=" + photo
	getBodyByUrl(url)
}
