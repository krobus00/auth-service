package model

type Response struct {
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func NewResponse() *Response {
	return new(Response)
}

func NewDefaultResponse() *Response {
	return &Response{
		Message: "OK",
	}
}

func (m *Response) WithMessage(message string) *Response {
	m.Message = message
	return m
}

func (m *Response) WithErrorMessage(err error) *Response {
	m.Message = err.Error()
	return m
}

func (m *Response) WithData(data any) *Response {
	m.Data = data
	return m
}
