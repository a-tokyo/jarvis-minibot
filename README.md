# Jarvis Mini-Bot

### Go Installation -> Google search
### Setup GOPATH
run these two commands on the terminal 
```
echo "export PATH=$GOPATH/bin:$PATH" >> ~/.bashrc
source ~/.bashrc 
```
*disclaimer: this is only tested on Ubuntu Linux*
then run this 
go get github.com/kardianos/govendor
go get github.com/pilu/fresh
### Setup project 
- Clone project
- `cd` to the project
- `govendor sync` or `govendor install` not sure which one works :D
- `fresh` will start the server and listen to changes
- If you want to run individual file just run `go nameoffile.go` in the terminal 