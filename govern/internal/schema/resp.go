package schema

type Resp struct {
	Code respCode    `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	Failure respCode = iota + 1000
	InvalidToken
	InvalidCaptcha
	Unauthorized
	NoMethod
	NoRoute
	InvalidArguments
	InvalidAccountOrPassword
	DuplicateEntity
)

const (
	Success respCode = iota + 2000
)

var (
	respMsgs = map[respCode]string{
		Failure:                  "Failure",
		InvalidToken:             "Invalid Token",
		InvalidCaptcha:           "Invalid Captcha",
		Unauthorized:             "Unauthorized",
		NoMethod:                 "No Method",
		NoRoute:                  "No Route",
		InvalidArguments:         "Invalid Arguments",
		InvalidAccountOrPassword: "Invalid Account Or Password",
		DuplicateEntity:          "Duplicate Entity",

		Success: "Success",
	}
)

type respCode int

func (a respCode) String() string {
	return respMsgs[a]
}
