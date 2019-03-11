# 实现一个Http rest服务
该服务功能为分为一个自增id，为了简单起见，就直接在程序中使用全局变量的方式，通过加锁的方式实现。最终暴露一个http接口供调用获取。

## 定义功能Service，并实现它
创建文件`$workspace/service/idGeneratorService.go`
```go
package service

import (
	"github.com/kataras/iris/core/errors"
	"sync"
)

type IdGeneratorService interface {
	GetId(serverId string) (int, error)
	CreateId(serverId string) (int, error)
}

var server1Id = 0
var server2Id = 0
var serverIds = map[string]int{}
var mutex = sync.Mutex{}
var invalidServer = errors.New("Invalid serverId")

type DefaultIdGenerator struct {
}

func (g *DefaultIdGenerator) CreateId(serverId string) (id int, err error) {
	mutex.Lock()
	defer mutex.Unlock()

	if serverId == "server1" {
		server1Id++
		id = server1Id
	} else if serverId == "server2" {
		server2Id++
		id = server2Id
	} else {
		err = invalidServer
	}

	if id != 0 {
		serverIds[serverId] = id
	}

	return
}

func (g *DefaultIdGenerator) GetId(serverId string) (id int, err error) {
	id, ok := serverIds[serverId]
	if !ok {
		err = invalidServer
	}
	return
}
```

简单来说，service有点像MVC架构中的controller，主要的逻辑功能都会在这里实现。

这里定义通过定义一个`IdGeneratorService`接口，包含两个方法，一个是查询id的功能，一个是创建id的功能。

而`DefaultIdGenerator`是一个默认的生成器，实现了`IdGeneratorService`接口。

## 创建endpoint
创建文件`$workspace/endpoint/idGeneratorEndpoint.go`
```go
package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"goCase/gokit-id-generator-demo/service"
)

type IdGeneratorRequest struct {
	ServerId string `json:"server_id"`
}

type IdGeneratorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data int    `json:"data"`
}

func MakeGetIdEndpoint(svc service.IdGeneratorService) endpoint.Endpoint {
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

func MakeCreateIdEndPoint(svc service.IdGeneratorService) endpoint.Endpoint  {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(IdGeneratorRequest)
		id, err := svc.CreateId(req.ServerId)

		if err != nil {
			return
		}

		response = IdGeneratorResponse{
			Code: 0,
			Msg: "success",
			Data: id,
		}

		return
	}
}
```
endpoint会为service封装需要的数据结构，同时它也磨平了客户端的通信方式，不需要在乎下游客户端是采用什么方式（http还是tcp）来通信。

## 创建transport
创建文件`$workspace/transport/http/idGenerator.go`
```go
package http

import (
	"context"
	"encoding/json"
	"errors"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"goCase/gokit-id-generator-demo/endpoint"
	"goCase/gokit-id-generator-demo/service"
	"net/http"
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

func MakeIdGeneratorHandler(ctx context.Context) http.Handler {
	digService := service.DefaultIdGenerator{}
	createIdEndpoint := endpoint.MakeCreateIdEndPoint(&digService)
	getIdEndpoint := endpoint.MakeGetIdEndpoint(&digService)

	router := mux.NewRouter()
	router.Methods("GET").
		Path("/id/{serverId}").
		Handler(httptransport.NewServer(
			getIdEndpoint,
			DecodeIdGeneratorRequest,
			EncodeIdGeneratorResponse,
		))

	router.Methods("POST").
		Path("/id/{serverId}").
		Handler(httptransport.NewServer(
		createIdEndpoint,
			DecodeIdGeneratorRequest,
			EncodeIdGeneratorResponse,
		))

	return router
}
```
创建文件`$workspace/transport/http/httptransport.go`
```go
package http

import (
	"context"
	"fmt"
	"net/http"
)

var port = ":8080"

func CreateTransport(ctx context.Context) (err error)  {
	idGeneratorHandler:=MakeIdGeneratorHandler(ctx)
	fmt.Printf("server start at port %s\n", port)
	http.Handle("/", idGeneratorHandler)
	err = http.ListenAndServe(port, nil)

	return
}
```
transport定义了和客户端的通信方式，此处我们使用http的方式作用通信方式。

## 定义入口main
```go
package main

import (
	"context"
	"fmt"
	httptransport "goCase/gokit-id-generator-demo/transport/http"
)

func main() {
	ctx := context.Background()

	err := httptransport.CreateTransport(ctx)
	if err != nil {
		fmt.Println(err)
	}
}
```