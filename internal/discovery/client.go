package discovery

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type Client struct {
	host          string
	retryCount    int
	backOffFactor int
}

func NewClient(host string, retryCount int, backOffFactor int) *Client {
	return &Client{host: host, retryCount: retryCount, backOffFactor: backOffFactor}
}

func (c *Client) Register(name string, ip string) {
	c.sendRequest(
		func() (*http.Response, error) {
			obj := registerRequestSchema{
				ServerName: name,
				ServerIP:   ip,
			}
			j, err := json.Marshal(&obj)
			if err != nil {
				panic(err)
			}
			url := fmt.Sprintf("%v/%v", c.host, "bootstrap/register/")
			return http.Post(url, "application/json", bytes.NewReader(j))
		},
	)
}

func (c *Client) GetClusterStatus() (getClusterStatusSchema, error) {
	resp := c.sendRequest(func() (*http.Response, error) {
		url := fmt.Sprintf("%v/%v", c.host, "cluster/status")
		return http.Get(url)
	})
	defer resp.Body.Close()

	obj := getClusterStatusSchema{}
	err := json.NewDecoder(resp.Body).Decode(&obj)
	return obj, err
}

func (c *Client) GetPeers() (peerListResponseSchema, error) {
	resp := c.sendRequest(
		func() (*http.Response, error) {
			url := fmt.Sprintf("%v/%v", c.host, "bootstrap/peers")
			return http.Get(url)
		},
	)
	obj := peerListResponseSchema{}
	err := json.NewDecoder(resp.Body).Decode(&obj)
	return obj, err
}

// TODO: Refactor this ................
func (c *Client) sendRequest(f func() (*http.Response, error)) *http.Response {
	retryCount := c.retryCount
	delay := 1

	var (
		resp *http.Response
		err  error
	)
	for retryCount > 0 {
		resp, err = f()

		if err == nil && resp.StatusCode >= 200 && resp.StatusCode < 300 {
			break
		}
		fmt.Printf("Retry count left: %v \n", retryCount)
		time.Sleep(time.Second * time.Duration(delay))
		retryCount--
		delay *= int(c.backOffFactor)
	}

	if retryCount <= 0 {
		if err != nil {
			panic(err)
		}
		buf := new(strings.Builder)
		_, err := io.Copy(buf, resp.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(resp.StatusCode)
		fmt.Println(buf.String())
		resp.Body.Close()
		panic("Something went wrong when sending request to discovery server")
	}
	return resp
}
