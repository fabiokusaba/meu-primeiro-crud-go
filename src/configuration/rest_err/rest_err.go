package rest_err

type RestErr struct {
	Message string `json:"message"`
	Err string `json:"error"`
	Code int `json:"code"`
	Causes []Causes `json:"causes"`
}

type Causes struct {
	Field string `json:"field"`
	Message string `json:"message"`
}