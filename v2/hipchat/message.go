package hipchat

type format string

const (
	FormatText format = format("text")
	FormatHTML format = format("html")
)

type Message struct {
	Message       string `json:"message"`
	MessageFormat format `json:"message_format"`
	Notify        bool   `json:"notify"`
}
