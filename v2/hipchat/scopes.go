package hipchat

/*
There are a handful of scopes that may be available for a token:

admin_group - Perform group administrative tasks
admin_room - Perform room administrative tasks
manage_rooms - Create, update, and remove rooms
send_message - Send private one-on-one messages
send_notification - Send room notifications
view_group - View users, rooms, and other group information
view_messages - View messages from chat rooms and private chats you have access to
*/
type scope string

const (
	// AdminGroup - Perform group administrative tasks
	ScopeAdminGroup scope = "admin_group"
	// AdminRoom - Perform room administrative tasks
	ScopeAdminRoom scope = "admin_room"
	// ManageRooms - Create, update, and remove rooms
	ScopeManageRooms scope = "manage_rooms"
	// SendMessage - Send private one-on-one messages
	ScopeSendMessage scope = "send_message"
	// SendNotification - Send room notifications
	ScopeSendNotification scope = "send_notification"
	// ViewGroup - View users, rooms, and other group information
	ScopeViewGroup scope = "view_group"
	// ViewMessages - View messages from chat rooms and private chats you have access to
	ScopeViewMessages scope = "view_messages"
)

func (s scope) String() string {
	switch s {
	case ScopeAdminGroup:
		return "admin group"
	case ScopeAdminRoom:
		return "admin room"
	case ScopeManageRooms:
		return "manage rooms"
	case ScopeSendNotification:
		return "send notification"
	case ScopeSendMessage:
		return "send message"
	case ScopeViewGroup:
		return "view group"
	case ScopeViewMessages:
		return "view message"
	default:
		return "Invalid Scope."
	}
}
