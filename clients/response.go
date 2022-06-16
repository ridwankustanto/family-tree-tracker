package clients

type Response struct {
	Error        bool        `json:"error"`
	DebugMessage string      `json:"debug_message"`
	Message      string      `json:"message"`
	Data         interface{} `json:"data"`
}

// Error message
const (
	ErrSomethingWentWrong = "something went wrong"
	ErrBadGateway         = "bad gateway"
)

// Success message
const (
	CreateSuccess = "create success"
)
