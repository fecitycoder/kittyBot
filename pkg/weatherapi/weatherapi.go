package weatherapi

import "net/http"

const weatherBaseUrl = "https://api.weather.yandex.ru/v2/informers"

var weatherKey string

func InitWehather(key string) {
	weatherKey = key
}

func getBodyByUrl(){
    client := &http
	.Client{} 
    req, err := http.NewRequest( 
         "GET", "https://google.com", nil, 
    ) 
    // добавляем заголовки
    req.Header.Add("Accept", "text/html")   // добавляем заголовок Accept
    req.Header.Add("User-Agent", "MSIE/15.0")   // добавляем заголовок User-Agent
  
    resp, err := client.Do(req) 
    if err != nil { 
        fmt.Println(err) 
        return
    } 
    defer resp.Body.Close() 
    io.Copy(os.Stdout, resp.Body)
}