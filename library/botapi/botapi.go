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

type GetMeT struct {
	OK      bool `json:"ok"`
	ResultT struct {
		Id                           int64  `json:"id"`
		IsBot                        bool   `json:"is_bot"`
		FirstName                    string `json:"first_name"`
		UserName                     string `json:"user_name"`
		Can_join_group               bool   `json:"can_join_group"`
		Can_read_all_groups_messages bool   `json:"can_read_all_groups_messages"`
		Support_inline_queries       bool   `json:"support_inline_queries"`
	} `json:"result"`
}

type SendMessageT struct {
	OK bool `json:"ok"`

	Result struct {
		MessageID                    int `json:"message_id"`
		GetUpdatesResultMessageFromT struct {
			ID           int    `json:"id"`
			IsBot        bool   `json:"is_bot"`
			FirstName    string `json:"first_name"`
			UserName     string `json:"user_name"`
			LanguageCode string `json:"language_code"`
		} `json:"from"`
		GetUpdatesResultMessageChatT struct {
			ID        int    `json:"id"`
			FirstName string `json:"first_name"`
			UserName  string `json:"user_name"`
			Type      string `json:"type"`
		} `json:"chat"`
		Date int    `json:"date"`
		Text string `json:"text"`
	} `json:"result"`
}

type GetUpdatesT struct {
	Ok     bool `json:"ok"`
	Result []struct {
		UpdateID int `json:"update_id"`
		Message  struct {
			MessageID int `json:"message_id"`
			From      struct {
				ID           int    `json:"id"`
				IsBot        bool   `json:"is_bot"`
				FirstName    string `json:"first_name"`
				LastName     string `json:"last"`
				UserName     string `json:"user_name"`
				LanguageCode string `json:"language_code"`
			} `json:"from"`
			Chat struct {
				ID        int    `json:"id"`
				FirstName string `json:"first_name"`
				LastName  string `json:"last"`
				UserName  string `json:"user_name"`
				Type      string `json:"type"`
			} `json:"chat"`
			Date int    `json:"date"`
			Text string `json:"text"`
		} `json:"message,omitempty"`
	} `json:"result"`
}

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

func GetUpdate() ([]string, int) {
	var lastMessage []string
	var id int
	body := getBodyByUrl(getUrlByMethod("getUpdates"))
	getUpdates := GetUpdatesT{}
	err := json.Unmarshal(body, &getUpdates)
	if err != nil {
		fmt.Printf("Error unmarshalling: %s", err.Error())
	}
	for _, update := range getUpdates.Result {
		lastMessage = append(lastMessage, fmt.Sprintf("%v", update.Message.Text))
		id = update.Message.Chat.ID
	}

	return lastMessage, id
}

func SendMessage(chatId int, message string) {
	url := getUrlByMethod("sendMessage")
	url = url + "?chat_id=" + strconv.Itoa(chatId) + "&text=" + message
	getBodyByUrl(url)

}
