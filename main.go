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

	return data[0].Title, nil
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
