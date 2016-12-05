package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/ramin0/chatbot"
)

// Autoload environment variables in .env
import _ "github.com/joho/godotenv/autoload"

const witAPIKey = "Bearer PAJE7ECQBN2BM26FL3BNOUVH7GDTOIJ5"
const newsAPIKey = "93c2574743664eb19825dba1e1729988"
const weatherAPIKey = "84c9e58ea74f56dfbffb9c5594fa45f5"
const port = "3000"
const errorMessage = "I'm embarrassed! \n Sorry, I can't answer this question :("

func main() {

	chatbot.WelcomeMessage = `What would you like to know about?
	 I can answer questions about:
		Weather: Give me the weather in Cairo
		News: What is the news in techcrunch?
	  `
	// chatbot.WelcomeMessage = `
	// <div style="
	// 	box-shadow: 0 3px 6px rgba(0,0,0,0.16), 0 3px 6px rgba(0,0,0,0.23);
	// 	background: rgba(255,255,255,0.8);
	//   	padding: 16px;
	// 	border-radius: 6px;
	// 	color: rgba(0,0,0,0.44);
	// ">
	// 	<h3>What would you like to know about?</h3>
	// 	<h5>I can asnwer questions about</h5>
	// 	<ul style="list-style: none;">
	// 		<li><h5>Weather: <blockquote>Give me the weather in Cairo?</blockquote></h5></li>
	// 		<li><h5>News: <blockquote>What is the news in Techcrunch?</blockquote></h5></li>
	// 	</ul>
	// </div>
	// `
	chatbot.ProcessFunc(chatbotProcess)

	port := os.Getenv("PORT")
	// Default to 3000 if no PORT environment variable was defined
	if port == "" {
		port = "3000"
	}
	// Start the server
	fmt.Printf("Listening on port %s...\n", port)
	log.Fatalln(chatbot.Engage(":" + port))
}

func chatbotProcess(session chatbot.Session, message string) (string, error) {
	isMobile := strings.HasPrefix(message, "mobileSession/")
	intent, value, err := extractValues(message)
	if err != nil {
		return "", err
	}
	switch intent {
	case "weather":
		data, error := getWeather(value)
		if error != nil {
			return "", error
		}
		if !isMobile {
			return weatherToHTMLString(data), nil
		} else {
			return weatherToJSONString(data), nil
		}
	case "news":
		data, error := getArticles(strings.ToLower(value))
		if error != nil {
			return "", error
		}
		if !isMobile {
			return articlesToHTMLString(data)
		} else {
			return articlesToJSONString(data)
		}
	}
	return errorMessage, nil
}

func getJSON(url string, target interface{}, headers map[string]string) error {
	client := &http.Client{}

	url = strings.Replace(url, " ", "+", -1)
	req, err := http.NewRequest("GET", url, nil)

	for key, value := range headers {
		req.Header.Add(key, value)
	}
	resp, err := client.Do(req)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

type wITResponse struct {
	Entities entity `json:"entities"`
	MsgID    string `json:"msg_id"`
	Text     string `json:"_text"`
}
type entity struct {
	Source []source `json:"source"`
	Intent []intent `json:"intent"`
}
type intent struct {
	Confidence float64 `json:"confidence"`
	Type       string  `json:"confidence"`
	Value      string  `json:"value"`
}

type source struct {
	Confidence float64 `json:"confidence"`
	Type       string  `json:"type"`
	Value      string  `json:"value"`
	Suggested  bool    `json:"suggested"`
}

func extractValues(message string) (string, string, error) {
	resp := new(wITResponse)
	headers := map[string]string{"Authorization": witAPIKey}
	err := getJSON("https://api.wit.ai/message?v=20161114&q="+message, resp, headers)
	if err != nil {
		return "", "", err
	}

	if len(resp.Entities.Intent) == 0 || len(resp.Entities.Source) == 0 {
		return "", "", errors.New(errorMessage)
	}
	return resp.Entities.Intent[0].Value, resp.Entities.Source[0].Value, nil
}

type newsResponse struct {
	Status   string    `json:"status"`
	Source   string    `json:"source"`
	SortBy   string    `json:"sortBy"`
	Articles []article `json:"articles"`
	Message  string    `json:"message"`
}

type article struct {
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	URLToImage  string `json:"urlToImage"`
	PublishedAt string `json:"publishedAt"`
}

func getArticles(source string) (*newsResponse, error) {
	resp := new(newsResponse) // or &Foo{}
	err := getJSON("https://newsapi.org/v1/articles?source="+source+"&apiKey="+newsAPIKey, resp, nil)
	return resp, err
}

func articlesToHTMLString(newsResponse *newsResponse) (string, error) {
	if strings.EqualFold("error", newsResponse.Status) {
		return "", errors.New(newsResponse.Message)
	}
	htmlString := `<ul style="list-style: none">`
	for _, article := range newsResponse.Articles {
		htmlString = htmlString +
			`<li
			style="
				background: rgba(0,0,0,0.05);
				padding: 5px 10px;
				border-radius: 6px;
				margin: 20px 0;
				margin-left: -40px;
				box-shadow: 0 3px 6px rgba(0,0,0,0.16), 0 3px 6px rgba(0,0,0,0.23);
				background: rgba(255,255,255,0.8);"
		> <a style="text-decoration: none;" href=" ` + article.URL + `">
		<h3 style="color: rgba(0,0,0,0.5);">` +
			article.Title + "</h3>" +
			`<h5 style="color: rgba(0,0,0,0.4);">` + article.Description + `</h5><li>`
	}
	htmlString += "</ul>"
	return htmlString, nil
}

/*
This one takes a newsResponse pointer and returns a JSON string of Articles in JSON format
i.e:
{
	article_1: {url: "url", title: "title", "description: "description"},
	article_1: {url: "url", title: "title", "description: "description"}
}
*/

func articlesToJSONString(newsResponse *newsResponse) (string, error) {
	if strings.EqualFold("error", newsResponse.Status) {
		return "", errors.New(newsResponse.Message)
	}
	jsonString := ``
	for i, article := range newsResponse.Articles {
		// jsonString += "Article Title: " + article.Title
		// jsonString += "\nArticle Description" + article.Description + "\n \n \n"

		jsonString = jsonString + `article_` + strconv.Itoa(i) + `: ` +
			`{url: "` + article.URL + `", ` +
			`title: "` + article.Title + `", ` +
			`description: "` + article.Description + `"}`
		if i != len(newsResponse.Articles)-1 {
			jsonString += `,`
		}
	}
	return jsonString, nil
}

// Weather API

// TODO Modularize
/*
How to deal with this piece of code ?
If wIT returns weather as an entinty, get the city* and the country abbr.(optional).
Call getWeather("city,countryInitials") or getWeather("city"); this returns a weatherResponse structure,
let's call this weatherResponse struct weatherState.
Call weatherToHTMLString(weatherState), this returns an HTML div in a string containing
the weather information in english language.
*/

//Weather Datastructures

//Weather Response, this structure is the struct form of the parsed weather API JSON response
type weatherResponse struct {
	Weather []WeatherDescription `json:"weather"`
	Main    WeatherData          `json:"main"`
	Wind    WeatherWind          `json:"wind"`
	Name    string               `json:"name"`
}

//Weather Data, this structure contains the main weather data
type WeatherData struct {
	Temp     float64 `json:"temp"`
	Pressure float64 `json:"pressure"`
	Humidity float64 `json:"humidity"`
	TempMin  float64 `json:"temp_min"`
	TempMax  float64 `json:"temp_max"`
}

//Weather Description, this structure contains the weather description in english. i.e: main:"Clouds" description: "Few Clouds"
type WeatherDescription struct {
	ID          int32  `json:"temp"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

//Weather Description, this structure contains the weather description in english. i.e: main:"Clouds" description: "Few Clouds"
type WeatherWind struct {
	Speed float64 `json:"speed"`
	Deg   float64 `json:"deg"`
}

/*
The query string is as follows: "cityName" or "cityName,countryInitials" i.e: "cairo" or "cairo,eg"
This one should get the weather full state from openweathermap.org API.
The getJSON Func mutates the response defined in the first line and fills it with the data
To know how to correctly get your neded values please check the 4 weather Data structures.
The object has 2 levels, i.e myWeather.main.temp gives the temperature.
If an error was thrown by the getJSON func, this error is returned as is preceded by a nil
*/
func getWeather(query string) (*weatherResponse, error) {
	response := new(weatherResponse)
	err := getJSON("http://api.openweathermap.org/data/2.5/weather?q="+query+"&appid="+weatherAPIKey+"&units=metric", response, nil)
	if err != nil {
		return nil, err
	}
	fmt.Println(response.Main.Temp)
	return response, nil
}

/*
This one takes a weatherResponse struct pointer and returns the weather data in an HTML string.
The returned string should be injected to the bot as the reply.
*/
func weatherToHTMLString(weatherState *weatherResponse) string {
	return `<div style="
    box-shadow: 0 3px 6px rgba(0,0,0,0.16), 0 3px 6px rgba(0,0,0,0.23);
	background: rgba(255,255,255,0.8);
    padding: 16px;
	max-width: 392px;
	border-radius: 6px;
	color: rgba(0,0,0,0.44);
">
    <h2>Weather in ` + weatherState.Name + `</h2>
    <img style="
    height: 100px;
    width: 100px;
" src="http://openweathermap.org/img/w/` + weatherState.Weather[0].Icon + `.png">
    <h4>Temprature is ` + floatToFixed(weatherState.Main.Temp) + `°C</h4>
    <h5>` + weatherState.Weather[0].Main + ` - ` + weatherState.Weather[0].Description + ` </h5>
    <ul style="list-style= none;">
        <li><h6>Min: ` + floatToFixed(weatherState.Main.TempMin) + `°C</h6></li>
        <li><h6>Max: ` + floatToFixed(weatherState.Main.TempMax) + `°C</h6></li>
    </ul>
		<h4>Wind status: </h4>
		<ul style="list-style= none;">
        <li><h6>Speed: ` + floatToFixed(weatherState.Wind.Speed) + `m/s</h6></li>
        <li><h6>Direction: ` + floatToFixed(weatherState.Wind.Deg) + `°</h6></li>
    </ul>
</div>`
}

/*
This one takes a weatherResponse pointer and returns a JSON string representing the weather
i.e:
{
	icon: "url.png",
	temprature: "3°C",
	weatherState: "clouds",
	weatherDescription: "few clouds",
	minTemprature: "3°C",
	maxTemprature: "3°C",
	windSpeed: "6 m/s",
	windDirection: "south west"
}
*/

func weatherToJSONString(weatherState *weatherResponse) string {
	return "The weather in " + weatherState.Name + " is " + floatToFixed(weatherState.Main.Temp) + "°C \n"
	// return `{icon: "http://openweathermap.org/img/w/` + weatherState.Weather[0].Icon + `.png",
	// temprature: "` + floatToFixed(weatherState.Main.Temp) + `°C",
	// weatherState: "` + weatherState.Weather[0].Main + `",
	// weatherDescription: "` + weatherState.Weather[0].Description + `",
	// minTemprature: "` + floatToFixed(weatherState.Main.TempMin) + `°C",
	// maxTemprature: "` + floatToFixed(weatherState.Main.TempMax) + `°C",
	// windSpeed ": ` + floatToFixed(weatherState.Wind.Speed) + `m/s",
	// windDirection: "` + floatToFixed(weatherState.Wind.Deg) + `"}`
}

// ./END Weather API

// Utilities Functions

func floatToFixed(num float64) string {
	intNum := int(num)
	return strconv.Itoa(intNum)
}
