# 添加中间件
## 为请求添加前后的日志记录
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