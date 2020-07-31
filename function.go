package httpcli

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func post(client *http.Client, url string, request interface{}) ([]byte, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("NewRequest: err=%v", err)
	}
	req.Header.Add("Content-Type", "application/json")
	req = req.WithContext(context.Background())
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("client.Do: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("StatusCode %d, Resp %s", resp.StatusCode, string(body))
	}
	return body, nil
}

func get(client *http.Client, addr string) ([]byte, error) {
	resp, err := client.Get(addr)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("StatusCode %d, Resp %s", resp.StatusCode, string(data))
	}
	return data, nil
}
