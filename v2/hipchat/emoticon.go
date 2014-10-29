package hipchat

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/gdey/hipchat_tools/v2/hipchat/emoticon"
)

type emoticonType string

const (
	EmoticonTypeGlobal = emoticonType("global")
	EmoticonTypeGroup  = emoticonType("group")
	EmoticonTypeAll    = emoticonType("all")
)

func (c *Client) Emoticons() {
	return c.EmoticonsAllFilters(0, 100, EmoticonTypeAll)
}

func (c *Client) EmoticonsAllFilters(startIdx maxResult, int, t emoticonType) emticon.Emo {
	a := make(url.Values)
	a.Add("start-index", strconv.Itoa(startIdx))
	a.Add("max-result", strconv.Itoa(maxResult))
	a.Add("type", string(t))
	b, err := c.getWithFilter("emoticon", true, &a)
	if err != nil {
		return nil, err
	}
	e := emoticon.Emoticon{}
	return &e, json.Unmarshal(b, &e)
}
func (c *Client) Emoticon(id string) (*emoticon.Emoticon, error) {
	b, err := c.Get("emoticon/"+id, RequiresAuthorization())
	if err != nil {
		return nil, err
	}
	e := emoticon.Emoticon{}
	return &e, json.Unmarshal(b, &e)
}
