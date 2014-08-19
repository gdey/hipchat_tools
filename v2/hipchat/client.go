package hipchat

type Client struct {
	AuthToken string
}

type parts []string

type response struct {
	body map[string]interface{}
	code statusCode
}
