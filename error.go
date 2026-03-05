package nexmo

// Error defines an error response message.
type Error struct {
	Type       string `json:"type"`
	Title      string `json:"title"`
	Detail     string `json:"detail"`
	Instance   string `json:"instance"`
	HTTPStatus int    `json:"-"`
}

func (err *Error) Error() string {
	return err.Title
}
