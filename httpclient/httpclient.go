package httpclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type HttpClient struct {
	TargetURL string
}

func (httpClient HttpClient) DoGetRequest() string {
	resp, err := http.Get(httpClient.TargetURL)
	if err != nil {
		fmt.Println("An error occurred sending GET to " + httpClient.TargetURL)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("An error occurred reading GET from " + httpClient.TargetURL)
		return ""
	}
	return string(body)
}
