package helper

type Response struct {
	Meta Meta `json:"meta"`
	Data any  `json:"data"`
}
type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

func APIresponse(code int, message, status string, data any) Response {

	response := Response{
		Meta: Meta{
			Code:    code,
			Message: message,
			Status:  status,
		},
		Data: data,
	}
	return response
}
