package main

import (
	"fmt"
	"log"
	"os"	
	"github.com/ramin0/chatbot"
	"langprocessor"
	"newsapi"
)

// Autoload environment variables in .env
import _ "github.com/joho/godotenv/autoload"

func chatbotProcess(session chatbot.Session, message string) (string, error) {	
	resp := langprocessor.ExtractValues(message)
	newsSource := resp.Entities.Source[0].Value 
	data := newsapi.GetArticles(newsSource)

	return data[0].Title, nil
}

func main() {
	chatbot.WelcomeMessage = "What's your name?"
	chatbot.ProcessFunc(chatbotProcess)
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	
	// Start the server
	fmt.Printf("Listening on port %s...\n", port)
	log.Fatalln(chatbot.Engage(":" + port))
}
