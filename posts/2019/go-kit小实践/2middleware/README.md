# 添加中间件
## transport级别日志中间件
创建文件`$workspace/endpoint/middlewareEndpoint.go`
```go
package endpoint

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kit/kit/endpoint"
)

type Middleware func(endpoint endpoint.Endpoint) endpoint.Endpoint

func LoggingMiddleware() Middleware  {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			data, _ := json.Marshal(request)
			fmt.Println("start: ", string(data))
			defer fmt.Println("end")
			return next(ctx, request)
		}
	}
}
```

## 应用级别日志中间件
创建文件`$workspace/service/middlewareService.go`
```go
package service

import "fmt"

type LoggingMiddleware struct {
	Next IdGeneratorService
}

func (l *LoggingMiddleware) GetId(serverId string) (int, error) {
	defer func() {
		fmt.Printf("GetId, serverId: %s\n", serverId)
	}()

	return l.Next.GetId(serverId)
}

func (l *LoggingMiddleware) CreateId(serverId string) (int, error) {
	defer func() {
		fmt.Printf("CreateId, serverId: %s\n", serverId)
	}()

	return l.Next.CreateId(serverId)
}
```
简单来说，这种级别的中间件是加在service上的

## 修改transport
修改`$workspace/transport/http/idGenerator.go`
```go
func MakeIdGeneratorHandler(_ context.Context) http.Handler {
  ...

  createIdEndpoint := endpoint.MakeCreateIdEndPoint(&digService)
	getIdEndpoint := endpoint.MakeGetIdEndpoint(&digService)

	createIdEndpoint = endpoint.LoggingMiddleware()(createIdEndpoint)
	getIdEndpoint = endpoint.LoggingMiddleware()(getIdEndpoint)

  ...
}
```