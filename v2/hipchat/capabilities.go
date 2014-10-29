package hipchat

import (
	"encoding/json"

	"github.com/gdey/hipchat_tools/v2/hipchat/capabilities"
)

type Capabilities struct {
	Name         string                     `json:"name"`
	Key          string                     `json:"key"`
	Description  string                     `json:"descrription"`
	Vendor       *capabilities.Vendor       `json:"vendor"`
	Links        *capabilities.Links        `json:"links"`
	Capabilities *capabilities.Capabilities `json:"capabilities"`
}

func (c *Client) Capabilities() (*Capabilities, error) {
	b, err := c.Get("capabilities")
	if err != nil {
		return nil, err
	}
	caps := Capabilities{}
	return &caps, json.Unmarshal(b, &caps)
}
