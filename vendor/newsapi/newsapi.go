package newsapi

import (
	"getJson"
)


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
	getJSON.GetJSON("https://newsapi.org/v1/articles?source="+source+"&apiKey=93c2574743664eb19825dba1e1729988", resp, nil)
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
