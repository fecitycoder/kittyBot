package botapi

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const telegramBaseUrl = "https://api.telegram.org/bot"
const methodGetMe = "getMe"
const methodGetUpdates = "getUpdates"
const methodSendMessage = "sendMessage"

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
		Date int `json:"date"`
	} `json:result"`
}

type getUpdatesT struct {
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
