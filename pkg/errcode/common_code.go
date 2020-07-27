package errcode

var (
	Success                   = NewError(0, "success")
	ServerError               = NewError(1000000, "server internal error")
	InvalidParams             = NewError(1000001, "params error")
	NotFound                  = NewError(1000002, "can not found")
	UnauthorizedAuthNotExist  = NewError(1000003, "authorized failed, can not found the APPKey and AppSecret")
	UnauthorizedTokenError    = NewError(1000004, "authorized failed, token error")
	UnauthorizedTokenTimeOut  = NewError(1000005, "authorized failed, token timeout")
	UnauthorizedTokenGenerate = NewError(1000006, "authorized failed, generate token error")
	TooManyRequests           = NewError(1000006, "too many requests")
)
