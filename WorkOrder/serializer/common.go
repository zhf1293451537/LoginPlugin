package serializer

type Response struct {
	Stauts int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

var User = map[string]string{}
