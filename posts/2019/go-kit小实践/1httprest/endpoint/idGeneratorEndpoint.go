package endpoint

import (
	"context"
	"goCase/gokit-id-generator-demo/service"

	"github.com/go-kit/kit/endpoint"
)

type IdGeneratorRequest struct {
	ServerId string `json:"server_id"`
}

type IdGeneratorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data int    `json:"data"`
}

func MakeIdGeneratorEndpoint(svc service.IdGeneratorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(IdGeneratorRequest)
		id, err := svc.GetId(req.ServerId)

		if err != nil {
			return
		}

		response = IdGeneratorResponse{
			Code: 0,
			Msg:  "success",
			Data: id,
		}
		return
	}
}
