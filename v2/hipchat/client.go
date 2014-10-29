package hipchat

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	AuthToken string
	Version   string
	APIURL    string
}

type parts []string

type response struct {
	body map[string]interface{}
	code statusCode
}

func (c Client) Room(room string) *room {
	return c.RoomWithColor(room, ColorDefault)
}

func (c Client) RoomWithColor(room string, color color) *room {
	return &room{
		client: c,
		color:  color,
		id:     room,
	}
}

func (c *Client) pathString(path string) string {
	version := Version
	if c.Version != "" {
		version := c.Version
	}
	apiurl := APIURL
	if c.APIURL != "" {
		apiurl := c.APIURL
	}
	rpath := strings.Join([]string{version, apiurl, path}, "/")
	return rpath
}

func (c *Client) bodyFromResource(res *http.Response) (body []byte, err error) {
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	return
}

type pathFunc func(*Clinet, *http.url) error
type requestFunc func(*Client, *http.Request) error
type clientOption struct {
	Path    *pathFunction
	Request *requestFunc
}

func Filters(args url.Values) clientOption {
	return clientOption{
		Path: func(c *Client, path *http.url) error {
			for k, vals := range args {
				for val := range vals {
					path.Values.Add(k, val)
				}
			}
			return nil
		},
		Request: nil,
	}
}
func RequiresAuthorization() {
	return clientOption{
		Path: nil,
		Request: func(c *Client, r *http.Request) error {
			if c.AuthToken != "" {
				req.Header.Add("Authorization", "Bearer "+c.AuthToken)
			}
			return nil
		},
	}
}

func (c *Client) Get(p string, options ...clientOption) ([]byte, error) {
	return c.handleRequest("GET", p, nil, options)
}

func (c *Client) Post(p string, b io.Reader, options ...clientOption) ([]byte, error) {
	return c.handleRequest("POST", p, b, options)
}
func (c *Client) handleRequest(m, p string, b io.Reader, options []clientOption) ([]byte, error) {
	hc := http.Client{}
	path, err := url.Parse(c.pathString(p))
	if err != nil {
		return nil, err
	}
	for val := range options {
		if val.Path != nil {
			err = val.Path(c, path)
			if err != nil {
				return nil, err
			}
		}
	}
	req, err := http.NewRequest(m, path.RequestURI(), b)
	if err != nil {
		return nil, err
	}
	for val := range options {
		if val.Request != nil {
			err = val.Request(c, req)
			if err != nil {
				return nil, err
			}
		}
	}
	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	return c.bodyFromResource(res)
}
