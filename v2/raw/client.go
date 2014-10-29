package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type Client struct {
	AuthToken string
}

const (
	Version       = "v2/"
	DefaultAPIURL = "https://api.hipchat.com/"
)

func (client *Client) Get(path string, auth bool) ([]byte, error) {
	c := http.Client{}
	req, err := http.NewRequest("GET", DefaultAPIURL+Version+path, nil)
	if err != nil {
		panic(err)
		return nil, err
	}
	if auth {
		req.Header.Add("Authorization", "Bearer "+client.AuthToken)
	}
	res, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	return res, nil
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body
}

func (client *Client) Post(path string, body io.Reader) ([]byte, error) {
	//c := http.Client{}
	return nil, nil
}

func (client *Client) printGetRequest(path string, auth bool) {
	fmt.Println("For path:", path)
	body, err := client.Get(path, auth)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

func main() {
	// This is only for testing.
	c := Client{AuthToken: "7C0jLo4E1TWx4mG49bcXzd32fIvxTnm0CQwqDYci"}
	c.printGetRequest("capabilities", false)
	c.printGetRequest("room", true)
}
