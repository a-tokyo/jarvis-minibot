# Jarvis Minibot

## Description
Jarvis minibot is a simple bot that uses natural language processing to analyze the users' question and respond with data about the subject from 3rd Party APIs.
The currently supported subjects are News and Weather.
> Deployment url: https://jarvis-minibot.herokuapp.com

## Demo
Go to: http://ramin0-chatbot-ui.herokuapp.com
Paste the url https://jarvis-minibot.herokuapp.com in the url input field
The chatbot is now ready to use.

## Setup

### Go Installation
Visit <a href="https://golang.org">https://golang.org</a>

### Setup GOPATH
run these two commands on the terminal
```
echo "export PATH=$GOPATH/bin:$PATH" >> ~/.bashrc
source ~/.bashrc
```
then run this
go get github.com/kardianos/govendor
go get github.com/pilu/fresh
### Setup project
- Clone project
- run `cd jarvis-minibot`
- run `govendor sync`
- `fresh` will start the server and listen to changes
- instead of fresh you can run `go run main.go`
- If you want to run individual file just run `go nameoffile.go` in the terminal

## Contribution

- Currently we need to add more 3rd party APIs to the bot.
  > Once you implement the API communication code, adjust the bot to reply about specific entities with this API's response.

- Also we need to modularize the current APIs into seperate files and packages.
