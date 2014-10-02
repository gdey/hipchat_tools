package hipchat

type room struct {
	client *Client
	Color  color
	id     string
}
type Room struct {
	Name              *string
	Topic             *string
	IsPrivate         bool
	IsGuestAccessible bool
	Owner             *string
}

type roomMessage struct {
	Message
	Color color `json:"color"`
}

func (c *client) Rooms(id string) (*room, error) {
	_, err := c.Get(parts{"room"})
	return
}

func (c *client) CreateRoom(room Room) (*room, error) {
	_, err := c.Post(parts{"room"}, room)
}

func (c *client) UpdateRoom(room Room) (*room, error) {
	_, err := c.Put(parts{"room", room.Name}, room)
}

func (c *client) DeleteRoom(rid string) error {
	_, err := c.Delete(parts{"room", rid})
}

func (r *room) Post(msg Message) error {
	if r == nil {
		return nil
	}
	return r.PostWithColor(msg, r.Color)
}

func (r *room) PostWithColor(msg Message, c color) error {
	if r == nil {
		return nip
	}
	if length(msg.Message) >= MaxMessageSize {
		return error("Message too large.")
	}
	rmsg := roomMessage{
		Message: msg,
		Color:   c,
	}
	_, err := r.client.Post(parts{"room", r.id, "notification"}, rmsg)
	if err != nil {
		return err
	}
	return nil
}
