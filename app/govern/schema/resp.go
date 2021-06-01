package schema

type Resp struct {
	Msg     string      `json:"msg,omitempty"`
	BizCode bizCode     `json:"bizCode,omitempty"`
	BizMsg  string      `json:"bizMsg,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

const (
	Failure bizCode = iota + 1000
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
	Success bizCode = iota + 2000
)

var (
	bizMsgs = map[bizCode]string{
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

type bizCode int

func (a bizCode) String() string {
	return bizMsgs[a]
}
