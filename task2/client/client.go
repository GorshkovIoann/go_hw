package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type InputString struct {
	InputString string `json:"inputString"`
}

type OutputString struct {
	OutputString string `json:"outputString"`
}

type Client struct {
	url string
}

func NewClient(url string) *Client {
	return &Client{url: url}
}

func (c *Client) SendRequests() error {
	urls := []string{
		"/version",
		"/decode",
		"/hard-op",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	for _, url := range urls {
		log.Printf("Sending request to %s", url)
		endpoint := "GET"
		if url == "/decode" {
			endpoint = "POST"
		}
		req, err := http.NewRequestWithContext(ctx, endpoint, c.url+url,nil)
		if err != nil {
			return err
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode >= 300 {
			log.Printf("Failed request to %s. Status code: %d", url, resp.StatusCode)
			continue
		}
		switch url {
		case "/version":
			body, _ := io.ReadAll(resp.Body)
			log.Printf("Received version: %s", body)
		case "/decode":
			body, _ := io.ReadAll(resp.Body)
			var result map[string]string
			json.Unmarshal(body, &result)
			log.Printf("Decoded string: %s", result["outputString"])
		case "/hard-op":
			log.Printf("Operation completed successfully")
		}
	}

	cancel()
	return nil
}

func (c *Client) GetVersion() (string, error) {
	req, err := http.NewRequest("GET", c.url+"/version", nil)
	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (c *Client) GetHardOp() (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", c.url+"/hard-op", nil)
	if err != nil {
		return 500, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return 500, nil
		}
		return 500, err
	}
	defer resp.Body.Close()

	return 200, nil
}

func (c *Client) PostDecode(inputString string) (string, error) {
	reqBody, err := json.Marshal(InputString{InputString: inputString})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", c.url+"/decode", bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var decoded OutputString
	err = json.Unmarshal(body, &decoded)
	if err != nil {
		return "", err
	}

	return decoded.OutputString, nil
}
