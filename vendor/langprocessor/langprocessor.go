package langprocessor
import (
    "getJson"
)
func ExtractValues(message string) (*wITResponse){
    resp := new(wITResponse)
    headers := map[string]string{"Authorization": "Bearer DJVBMEA4W6PG23OLEWV745OKU5XRY4MR"}
    getJSON.GetJSON("https://api.wit.ai/message?v=20161114&q=" + message, resp, headers)
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

