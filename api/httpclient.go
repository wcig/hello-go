package api

type request struct {
	Url string
	Method string
	Params map[string]string
}

func New() *request {
	return &request{
		Url:    "",
		Method: "",
		Params: nil,
	}
}

func (r *request) SetQueryParams(params map[string]string) *request {
	return
}
