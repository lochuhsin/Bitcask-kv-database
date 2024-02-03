package discovery

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	host string
}

func NewClient(host string) *Client {
	return &Client{host}
}

func (c *Client) Register(name string, ip string) (int, error) {
	obj := registerRequestSchema{
		ServerName: name,
		ServerIP:   ip,
	}
	j, err := json.Marshal(&obj)
	if err != nil {
		panic(err)
	}
	url := fmt.Sprintf("%v/%v", c.host, "bootstrap/register/")
	resp, err := http.Post(url, "application/json", bytes.NewReader(j))
	return resp.StatusCode, err
}
func (c *Client) GetClusterStatus() getClusterStatusSchema {
	url := fmt.Sprintf("%v/%v", c.host, "cluster/status")
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	obj := getClusterStatusSchema{}
	err = json.NewDecoder(resp.Body).Decode(&obj)
	if err != nil {
		panic(err)
	}
	return obj
}
func (c *Client) GetPeers() peerListResponseSchema {
	url := fmt.Sprintf("%v/%v", c.host, "bootstrap/peers")
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	obj := peerListResponseSchema{}
	err = json.NewDecoder(resp.Body).Decode(&obj)
	if err != nil {
		panic(err)
	}
	return obj
}
