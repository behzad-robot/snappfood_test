package libs

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func HttpGet(url string, headers map[string]string) ([]byte, error) {
	// response, err := http.Get(url)
	client := &http.Client{}
	req, e1 := http.NewRequest("GET", url, nil)
	if e1 != nil {
		return nil, e1
	}
	// req.Header.Set("Content-Type", "application/json")
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, e2 := io.ReadAll(response.Body)
	fmt.Println(url, "=>", response.StatusCode)
	fmt.Println(url, "=>", string(body))
	if e2 != nil {
		return nil, e2
	}
	return body, nil
}
func HttpPost(url string, body []byte, headers map[string]string) ([]byte, error) {
	// response, err := http.Get(url)
	client := &http.Client{}
	req, e1 := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if e1 != nil {
		return nil, e1
	}
	req.Header.Set("Content-Type", "application/json")
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	respBody, e2 := io.ReadAll(response.Body)
	fmt.Println(url, "=>", response.StatusCode, " body=", string(body), "\n", string(respBody))
	if e2 != nil {
		return nil, e2
	}
	return respBody, nil
}
