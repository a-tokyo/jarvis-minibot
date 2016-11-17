# Jarvis Mini-Bot

## Description
Jarvis minibot is a simple bot that uses natural language processing to analyze the users' question and respond with data about the subject from 3rd Party APIs.
The currently supported subjects are News and Weather.


### Go Installation
Visit <a href="https://golang.org/">https://golang.org/</a>

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
