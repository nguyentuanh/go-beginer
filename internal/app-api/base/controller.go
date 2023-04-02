package base

import "go-template/internal/app-api/response"

// Controller ...
type Controller struct {
}

// Resp ...
func (Controller) Resp() response.IResponse {
	return response.NewResponse()
}
