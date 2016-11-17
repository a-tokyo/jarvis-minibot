package main

import (
	"fmt"
	"log"
	"os"
	"github.com/ramin0/chatbot"

	"encoding/json"
	"strings"
	"net/http"
)

// Autoload environment variables in .env
import _ "github.com/joho/godotenv/autoload"

func chatbotProcess(session chatbot.Session, message string) (string, error) {
	resp := ExtractValues(message)
	newsSource := resp.Entities.Source[0].Value
	data := GetArticles(newsSource)

	return `<div class="jarvis-response"`+data[0].Title+`</div>`, nil
}

func main() {
	chatbot.WelcomeMessage = "What would you like to know about?"
	chatbot.ProcessFunc(chatbotProcess)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Start the server
	fmt.Printf("Listening on port %s...\n", port)
	log.Fatalln(chatbot.Engage(":" + port))
}

func ExtractValues(message string) (*wITResponse){
    resp := new(wITResponse)
    headers := map[string]string{"Authorization": "Bearer DJVBMEA4W6PG23OLEWV745OKU5XRY4MR"}
    GetJSON("https://api.wit.ai/message?v=20161114&q=" + message, resp, headers)
    return resp
}
type wITResponse struct{
    Entities entity `json:"entities"`
    MsgID string `json:"msg_id"`
    Text string `json:"_text"`
}
type entity struct{
    Source []source `json:"source"`
}

type source struct{
    Confidence float64 `json:"confidence"`
    Type string `json:"type"`
    Value string `json:"value"`
    Suggested bool `json:"suggested"`
}

func GetJSON(url string, target interface{},headers map[string]string) error {
    client := &http.Client{}

    url = strings.Replace(url, " ", "+", -1)
    req, err := http.NewRequest("GET", url, nil)

    for key, value := range headers {
        req.Header.Add(key,value)
    }
    resp, err := client.Do(req)

    if err != nil {
        return err
    }
    defer resp.Body.Close()

    return json.NewDecoder(resp.Body).Decode(target)
}

type newsResponse struct{
	Status string `json:"status"`
	Source string `json:"source"`
	SortBy string `json:"sortBy"`
	Articles []article `json:"articles"`
}

type article struct{
	Author string `json:"author"`
	Title string `json:"title"`
	Description string `json:"description"`
	URL string `json:"url"`
	URLToImage string `json:"urlToImage"`
	PublishedAt string `json:"publishedAt"`
}
/*
I will call this function in the main file this should return me the articles I have requested as we agreed,
This function should return list of articles and error
You should also handle if you recieve an error so you should give me back the error so I can display it to the user
*/
func GetArticles(source string) ([]article) {
	resp := new(newsResponse) // or &Foo{}
	GetJSON("https://newsapi.org/v1/articles?source="+source+"&apiKey=93c2574743664eb19825dba1e1729988", resp, nil)
	return resp.Articles
}

/*
This one should get the sources available on the api in other words the sources that we support,
please note that all these variables are optional so you might get all of these as nulls
and you can carry on with the request with no problem
This function should return list of sources and error
You should also handle if you recieve error so you should give me back the error so I can display it to the user
*/
func getSources(category string, lang string,country string){

}

//Weather API
// TODO Modularize
//Weather Datastructures
//Weather Response, this structure is the struct form of the parsed weather API JSON response
type weatherResponse struct{
	weather []weatherDescription `json:"weather"`
	main weatherData `json:"main"`
	wind weatherWind `json:"wind"`
	visibility float64 `json:"visibility"`
}
//Weather Data, this structure contains the main weather data
type weatherData struct{
	temp float64 `json:"temp"`
	pressure float64 `json:"pressure"`
	humidity float64 `json:"humidity"`
	temp_min float64 `json:"temp_min"`
	temp_max float64 `json:"temp_max"`
}

//Weather Description, this structure contains the weather description in english. i.e: main:"Clouds" description: "Few Clouds"
type weatherDescription struct{
	id int32 `json:"temp"`
	main string `json:"main"`
	description string `json:"description"`
	icon string `json:"icon"`
}

//Weather Description, this structure contains the weather description in english. i.e: main:"Clouds" description: "Few Clouds"
type weatherWind struct{
	speed float32 `json:"speed"`
	deg float32 `json:"deg"`
}

/*
The query string is as follows: "cityName" or "cityName,countryInitials" i.e: "cairo" or "cairo,eg"
This one should get the weather full state from openweathermap.org API.
The GetJSON Func mutates the response defined in the first line and fills it with the data
To know how to correctly get your neded values please check the 4 weather Data structures.
The object has 2 levels, i.e myWeather.main.temp gives the temperature.
If an error was thrown by the getJSON func, this error is returned as is preceded by a nil
*/
func GetWeather(query string) (*weatherResponse, error) {
	response := new(weatherResponse)
	err := GetJSON("http://api.openweathermap.org/data/2.5/weather?q="+ query +"&appid=84c9e58ea74f56dfbffb9c5594fa45f5&units=metric", response, nil)
	if err != nil {
		return nil, err
	}
	return response, nil
}
