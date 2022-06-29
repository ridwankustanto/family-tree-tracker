package clients

type Response struct {
	Error        bool        `json:"error"`
	DebugMessage string      `json:"debug_message"`
	Message      string      `json:"message"`
	Data         interface{} `json:"data"`
}

type ResponseLogin struct{
	Message      string      `json:"message"`
	Data         interface{} `json:"data"`
	Token         interface{} `json:"token"`
	
}

// Error message
const (
	ErrSomethingWentWrong = "something went wrong"
	ErrBadGateway         = "bad gateway"
	NotAuthorized		  = "You are not authorized to enter this area"
)

// Success message
const (
	CreateSuccess = "create success"
)
