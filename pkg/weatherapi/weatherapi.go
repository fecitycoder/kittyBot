package weatherapi

const weatherBaseUrl = "https://api.weather.yandex.ru/v2/informers"

var weatherKey string

func InitWehather(key string) {
	weatherKey = key
}

func getBodyByUrl() {

}
