package botapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

const telegramBaseUrl = "https://api.telegram.org/bot"

var botTocken string

func InitBot(tocken *string) {
	botTocken = *tocken
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

	body, err := ioutil.ReadAll(response.Body)
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

func GetUpdate(id *int) []string {
	var lastMessage []string
	body := getBodyByUrl(getUrlByMethod("getUpdates"))
	getUpdates := GetUpdatesT{}
	err := json.Unmarshal(body, &getUpdates)
	if err != nil {
		fmt.Printf("Error unmarshalling: %s", err.Error())
	}
	for _, update := range getUpdates.Result {
		lastMessage = append(lastMessage, fmt.Sprintf("%v", update.Message.Text))
		*id = update.Message.Chat.ID
	}
	getBodyByUrl(getUrlByMethod("setWebhook?drop_pending_updates=true"))
	return lastMessage
}

func SendMessage(chatId *int, message string) {
	url := getUrlByMethod("sendMessage")
	url = url + "?chat_id=" + strconv.Itoa(*chatId) + "&text=" + message
	getBodyByUrl(url)

}
