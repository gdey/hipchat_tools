package hipchat

type Client struct {
	AuthToken string
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
