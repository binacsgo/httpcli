package httpcli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func call(client *http.Client, url, method string, request, response interface{}) error {
	body, err := json.Marshal(request)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	// req = req.WithContext(context.Background())
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("StatusCode %d, Resp %s", resp.StatusCode, string(body))
	}
	return nil
}
