package newsapi

import (
	"net/http"
	"encoding/json"
)

// https://newsapi.org/#documentation docs link

type Response struct{
	status string
	source string
	sortBy string
	articles []string
}

/* 
I will call this function in the main file this should return me the articles I have requested as we agreed, 
This function should return list of articles and error 
You should also handle if you recieve an error so you should give me back the error so I can display it to the user 
*/ 
func getArticles(source string) {

	// resp := new(Response) // or &Foo{}
	// getJSON("https://newsapi.org/v1/articles?source=techcrunch&apiKey=93c2574743664eb19825dba1e1729988", resp)
	// println(resp.status)
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

func getJSON(url string, target interface{}) error {
    r, err := http.Get(url)
    if err != nil {
        return err
    }
    defer r.Body.Close()

    return json.NewDecoder(r.Body).Decode(target)
}

func main(){
	//you can test your code here and run it from the terminal 
}