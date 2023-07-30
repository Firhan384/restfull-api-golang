package schemas

type ResponseApi struct {
	StatusCode    int         `json:"status_code"`
	Error         bool        `json:"error"`
	Message       string      `json:"message"`
	DetailMessage interface{} `json:"detail_message, omitempty"`
	Data          interface{} `json:"data, omitempty"`
}

func NewResponseApi() ResponseApi {
	return ResponseApi{}
}
