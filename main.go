package main
import (
	"fmt"
	"log"
	"os"
	"net/http"
	"encoding/json"
	"github.com/ramin0/chatbot"
)

// Autoload environment variables in .env
import _ "github.com/joho/godotenv/autoload"

type Response struct{
	status string
	source string
	sortBy string
	articles []string
}

func getJSON(url string, target interface{}) error {
    r, err := http.Get(url)
    if err != nil {
        return err
    }
    defer r.Body.Close()

    return json.NewDecoder(r.Body).Decode(target)
}

func chatbotProcess(session chatbot.Session, message string) (string, error) {
	// if strings.EqualFold(message, "chatbot") {
	// 	return "", fmt.Errorf("This can't be, I'm the one and only %s!", message)
	// }
	
	resp := new(Response) // or &Foo{}
    getJSON("https://newsapi.org/v1/articles?source=techcrunch&apiKey=93c2574743664eb19825dba1e1729988", resp)
    println(resp.status)

	return fmt.Sprintf("Hello %s, my name is chatbot. What was yours again?", resp.articles), nil
}

func main() {
	// Uncomment the following lines to customize the chatbot
	chatbot.WelcomeMessage = "What's your name?"
	chatbot.ProcessFunc(chatbotProcess)

	// Use the PORT environment variable
	port := os.Getenv("PORT")
	// Default to 3000 if no PORT environment variable was defined
	if port == "" {
		port = "3000"
	}

	// Start the server
	fmt.Printf("Listening on port %s...\n", port)
	log.Fatalln(chatbot.Engage(":" + port))
}
