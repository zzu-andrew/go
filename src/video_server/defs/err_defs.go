package defs

type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErrorResponse struct {
	HttpsSc int
	Error   Err
}

var (
	ErrorRequestBodyParseFailed = ErrorResponse{HttpsSc: 400, Error: Err{Error: "Request body is not corect", ErrorCode: "001"}}
	ErrorNotAuthUser            = ErrorResponse{HttpsSc: 401, Error: Err{Error: "User authentication failed", ErrorCode: "002"}}
)
