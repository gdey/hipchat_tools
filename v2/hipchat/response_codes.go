package hipchat

import "fmt"

type statusCode uint16

const (
	// OK means it worked
	OK statusCode = 200
	// OK201 The resource was created successfully. The body should contain a "links" map with a "self" field that contains the new URL to access the created resource. Alternatively, the URL will be in the "Location" header
	OK201 statusCode = 201
	// Accepted When using test_auth=true, this response code indicates that the auth_token is valid.
	Accepted statusCode = 202
	// NoContent The server successfully processed the request, but is not returning any content. Usually used as a response to a successful delete request.
	NoContent         statusCode = 204
	TemporaryRedirect statusCode = 307
	// BadRequest The request was invalid. You may be missing a required argument or provided bad data. An error message will be returned explaining what happened.
	BadRequest statusCode = 400
	// Unauthorized The authentication you provided is invalid.
	Unauthorized statusCode = 401
	// Forbidden You have exceeded the rate limit.
	Forbidden statusCode = 403
	// NotFound You requested an invalid method.
	NotFound statusCode = 404
	// InternalServerError Something is wrong on our end. We'll investigate what happened. Feel free to contact us.
	InternalServerError statusCode = 500
	// ServiceUnavailable The method you requested is currently unavailable (due to maintenance or high load).
	ServiceUnavailable statusCode = 503
)

func (code statusCode) String() string {
	switch code {
	case OK:
		return "OK"
	case OK201:
		return "OK"
	case Accepted:
		return "Accepted"
	case BadRequest:
		return "Bad Request"
	case Unauthorized:
		return "Unauthorized"
	case Forbidden:
		return "Forbidden"
	case NotFound:
		return "Not Found"
	case InternalServerError:
		return "Internal Server Error"
	case TemporaryRedirect:
		return "Temporary Redirect"

	default:
		return fmt.Sprintf("Unknown(%d) Error", code)
	}
}
func (code statusCode) Code() uint16 { return uint16(code) }
