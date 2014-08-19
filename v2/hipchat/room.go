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

func GetRoom(id string) (*room, error) {
}
func CreateRoom(room Room) (*room, error) {
}
func UpdateRoom(room Room) (*room, error) {
}
func DeleteRoom(room Room) error {
}

func (r *room) Post(msg Message) error {
	if r == nil {
		return nil
	}
	return r.PostWithColor(msg, r.Color)
}
func (r *room) PostWithColor(msg Message, c color) error {
	if r == nil {
		return nil
	}
	rmsg := roomMessage{
		Message: msg,
		Color:   c,
	}
	_, err := r.client.send(parts{"room", r.id, "notification"}, rmsg)
	if err != nil {
		return err
	}
	return nil
}
