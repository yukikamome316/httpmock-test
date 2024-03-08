package client

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

type ApiGateway interface {
	Get(path string) (int, []byte, error)
}

type Client struct {
	Host    string
	Timeout time.Duration
}

func NewClient(host string, timeout int) *Client {
	return &Client{
		Host:    host,
		Timeout: time.Duration(timeout) * time.Second,
	}
}

func (c Client) Get(path string) (int, []byte, error) {
	var reqBody io.Reader
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", c.Host, path), reqBody)
	if err != nil {
		return 0, []byte{}, err
	}

	client := &http.Client{Timeout: c.Timeout}
	res, err := client.Do(req)
	if err != nil {
		return 0, []byte{}, err
	}
	defer res.Body.Close()

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(res.Body)
	if err != nil {
		return 0, []byte{}, err
	}

	code := res.StatusCode
	body := buf.Bytes()

	return code, body, nil
}
