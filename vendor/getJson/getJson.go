package getJSON
import(
    "encoding/json"
	"strings"
    "net/http"
)
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
