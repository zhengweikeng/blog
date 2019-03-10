package http

import (
	"context"
	"encoding/json"
	"errors"
	"goCase/gokit-id-generator-demo/endpoint"
	"net/http"

	kitendpoint "github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

var serverIdRequired = errors.New("ServerId is Required")

func DecodeIdGeneratorRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	params := mux.Vars(r)
	serverId, ok := params["serverId"]
	if !ok {
		err = serverIdRequired
		return
	}

	request = endpoint.IdGeneratorRequest{serverId}
	return
}

func EncodeIdGeneratorResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}

func MakeIdGeneratorHandler(ctx context.Context, endpoint kitendpoint.Endpoint) http.Handler {
	r := mux.NewRouter()

	r.Methods("GET").
		Path("/id/{serverId}").
		Handler(httptransport.NewServer(
			endpoint,
			DecodeIdGeneratorRequest,
			EncodeIdGeneratorResponse,
		))

	return r
}
