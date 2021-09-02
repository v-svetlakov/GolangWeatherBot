package main

import (
	"bot/b_types"
	"encoding/json"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)


func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func weather(cityName string) *b_types.Weathers {
	res, err := http.Get(
				fmt.Sprintf(
				"https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric",
				cityName, os.Getenv("WEATHER_KEY")),
		)
	if err != nil{
		fmt.Println(err)
	}
	defer res.Body.Close()
	body, _    := ioutil.ReadAll(res.Body)
	data       := new(b_types.Weathers)
	jsonErr    := json.Unmarshal(body, &data)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}
	return data
}

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("API_KEY"))
	if err   != nil {
		//log.Panic(err)
		fmt.Println("hi")
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u            := tgbotapi.NewUpdate(0)
	u.Timeout     = 60
	updates, err := bot.GetUpdatesChan(u)
	for update   := range updates {
		if update.Message == nil {
			continue
		}
		msg      := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		UserName := update.Message.From.UserName

		switch update.Message.Command() {
		case "start":
			msg.Text = fmt.Sprintf("hi %s", UserName)
		case "weather":
			result := weather("Moscow")
			msg.Text = fmt.Sprintf(
				"Temp: %.2f \n" +
						"Temp_min: %.2f \n" +
						"Temp_max: %.2f \n" +
						"%s \n" +
						"Lat: %.2f \n" +
						"Lon: %.2f \n" +
						"visibility: %.2f \n",
				result.Main.Temp, result.Main.Temp_min, result.Main.Temp_max,
				result.Base,
				result.Coord.Lat, result.Coord.Lon,
				result.Visibility)

		default:
			msg.Text = fmt.Sprint("The bot responds to the following commands:\n" +
				"/start - greeting\n" +
				"/weather - will show the weather in Moscow")
		}
		//log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		bot.Send(msg)
	}
}

