package weatherapi

var (
	Condition = map[string]string{
		"clear":                            "ясно",
		"partly-cloudy":                    "малооблачно",
		"cloudy":                           "облачно с прояснениями",
		"overcast":                         "пасмурно",
		"partly-cloudy-and-light-rain":     "малооблачно, небольшой дождь",
		"partly-cloudy-and-rain":           "малооблачно, дождь",
		"overcast-and-rain":                "значительная облачность, сильный дождь",
		"overcast-thunderstorms-with-rain": "сильный дождь с грозой",
		"cloudy-and-light-rain":            "облачно, небольшой дождь",
		"overcast-and-light-rain ":         "значительная облачность, небольшой дождь",
		"cloudy-and-rain":                  "облачно, дождь",
		"overcast-and-wet-snow":            "дождь со снегом",
		"partly-cloudy-and-light-snow":     "небольшой снег",
		"partly-cloudy-and-snow":           "малооблачно, снег",
		"overcast-and-snow":                "снегопад",
		"cloudy-and-light-snow":            "облачно, небольшой снег",
		"overcast-and-light-snow":          "значительная облачность, небольшой снег",
		"cloudy-and-snow":                  "облачно, снег",
	}

	Daytime = map[string]string{"d": "светлое время суток", "n": "темное время суток"}

	Season = map[string]string{"summer": "лето", "autumn": "осень", "winter": "зима", "spring": "весна"}

	WindDir = map[string]string{
		"nw": "северо-западное",
		"n":  "северное",
		"ne": "северо-восточное",
		"e":  "восточное",
		"se": "юго-восточное",
		"s":  "южное",
		"sw": "юго-западное",
		"w":  "западное",
		"с":  "штиль",
	}
)

type WeatherQueryT struct {
	ServerTimeUnix uint32 `json:"now"`    // Время сервера в Unix формате
	ServerTimeDT   string `json:"now_dt"` // Время сервера в формате Data time
	Info           struct {
		Url       string `json:"url"`
		Lattitude string `json:"lat"`
		Longitude string `json:"lon"`
	} `json:"info"` //
	Fact struct {
		TimeUnix     uint32  `json:"obs_time"`    // Время замера погодных данных в Unix формате
		Temp         int8    `json:"temp"`        // Температура в ((С)
		TempFeelLike int8    `json:"feel_like"`   // Температура по ощущеметру
		Icon         string  `json:"icon"`        // Название иконки https://yastatic.net/weather/i/icons/funky/dark/< icon >.svg
		Conditions   string  `json:"conditions"`  // Код погодного описания
		Wind_speed   float32 `json:"wind_speed"`  // Скорость ветра (в м/с)
		Wind_gust    float32 `json:"wind_gust"`   // Скорость порывов ветра (в м/с)
		Wind_dir     string  `json:"wind_dir"`    // Направление ветра
		Pressure_mm  int8    `json:"pressure_mm"` // Давление (в мм рт.ст.).
		Pressure_pa  int8    `json:"pressure_pa"` // Давление (в гексопаскалях).
		Humidity_mm  int8    `json:"humidity_mm"` // Влажность воздуха
		Daytime      string  `json:"daytime"`     // Время суток
		Polar        bool    `json:"polar"`       // Признак того, что время суток, указанное в поле daytime является полярным.
		Season       string  `json:"season"`      // Время года в данном населенном пункте
	} `json:"fact"` //	Объект содержит информацию о погоде на данный момент

}
