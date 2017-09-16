package common

// RestResult represents that produces by rest api
type RestResult struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   OpError     `json:"error"`
}

func (r *RestResult) SetData(data interface{}) {
	r.Success = true
	r.Data = data
}

func (r *RestResult) SetSuccess(success bool) {
	r.Success = success
}

func (r *RestResult) SetError(err OpError) {
	r.Success = false
	r.Error = err
}

func (r *RestResult) SetErrorCode(code ErrorCode) {
	r.Success = false
	r.Error = OpError{code}
}

