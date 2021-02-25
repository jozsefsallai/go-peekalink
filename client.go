package peekalink

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Client is a Peekalink API client.
type Client struct {
	Key     string
	BaseURL string
}

// SetKey will update the client's key.
func (client *Client) SetKey(key string) {
	client.Key = key
}

// SetBaseURL will change the client's base URL.
func (client *Client) SetBaseURL(url string) {
	client.BaseURL = url
}

func (client *Client) doRequest(method string, endpoint string, payload []byte) ([]byte, error) {
	req, err := http.NewRequest(method, fmt.Sprintf("%s/%s", client.BaseURL, endpoint), bytes.NewBuffer(payload))
	if err != nil {
		return []byte{}, nil
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", client.Key)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return []byte{}, nil
	}

	if res.StatusCode != 200 {
		return []byte{}, fmt.Errorf("request failed with error %s", LinkError(res.StatusCode))
	}

	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}

// Preview will send a request to fetch link preview data for a given link.
func (client *Client) Preview(link string) (*LinkPreview, error) {
	payload := LinkPreviewRequest{
		Link: link,
	}

	jsonPayload, _ := json.Marshal(payload)

	body, err := client.doRequest("POST", "", jsonPayload)
	if err != nil {
		return nil, err
	}

	var data LinkPreview
	err = json.Unmarshal(body, &data)

	return &data, err
}

// IsAvailable will send a request to check the availability of a given link.
func (client *Client) IsAvailable(link string) (*LinkAvailability, error) {
	payload := LinkAvailabilityRequest{
		Link: link,
	}

	jsonPayload, _ := json.Marshal(payload)

	body, err := client.doRequest("POST", "is-available/", jsonPayload)
	if err != nil {
		return nil, err
	}

	var data LinkAvailability
	err = json.Unmarshal(body, &data)

	return &data, err
}

// NewClient creates a new Peekalink API client instance.
func NewClient(key string) *Client {
	return &Client{
		Key:     key,
		BaseURL: "https://api.peekalink.io",
	}
}
