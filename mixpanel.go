package mixpanel

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

func Track(distinctId string, eventName string, properties map[string]interface{}) int {
	token := os.Getenv("MIXPANEL_PROJECT_ID")
	properties["distinct_id"] = distinctId
	properties["token"] = token
	properties["time"] = time.Now().Unix()
	data := make(map[string]interface{})
	data["event"] = eventName
	data["properties"] = properties
	dataJson, _ := json.Marshal(data)
	dataEncoded := base64.StdEncoding.EncodeToString(dataJson)
	dataEncoded = strings.Replace(dataEncoded, "\n", "", -1)
	mixpanelUrl := "https://api.mixpanel.com/track"
	resp, err := http.PostForm(mixpanelUrl, url.Values{"data": {dataEncoded}})
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	returnValue, err := strconv.Atoi(string(body))
	if err != nil {
		panic(err)
	}
	return returnValue
}
