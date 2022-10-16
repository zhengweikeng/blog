# 添加限流中间件
## 令牌桶中间件
简单来说，令牌桶就是每次请求的时候，都先从一个桶里面拿令牌，如果拿到了，就能继续请求，请求结束后，将令牌归还令牌桶；如果拿不到就放弃。  
令牌桶算法会实现定义好令牌数量，然后定时刷新桶里的数据。

### 基于juju/ratelimit实现
这里采用了这个库`github.com/juju/ratelimit`实现的令牌桶算法。

修改文件`$workspace/endpoint/middlewareEndpoint.go`
```go
var ErrLimitExceed = errors.New("Rate limit exceed!")

func NewTokenBucketLimitterWithJuju(bkt *ratelimit.Bucket) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			// 获取1个令牌，如果结果为0，说明取不到令牌
			if bkt.TakeAvailable(1) == 0 {
				fmt.Println(ErrLimitExceed)
				return nil, ErrLimitExceed
			}
			return next(ctx, request)
		}
	}
}
```

修改文件`$workspace/transport/http/idGenerator.go`
```go
func MakeIdGeneratorHandler(_ context.Context) http.Handler {
  digService := service.LoggingMiddleware{
		Next: &service.DefaultIdGenerator{},
	}
	// 创建令牌桶，每秒刷新1次，容量为3
	ratebucket := ratelimit.NewBucket(time.Second * 1, 3)
	rateBucketMiddleware := endpoint.NewTokenBucketLimitterWithJuju(ratebucket)

	createIdEndpoint := endpoint.MakeCreateIdEndPoint(&digService)
	getIdEndpoint := endpoint.MakeGetIdEndpoint(&digService)

	createIdEndpoint = endpoint.LoggingMiddleware()(createIdEndpoint)
	getIdEndpoint = endpoint.LoggingMiddleware()(getIdEndpoint)

	createIdEndpoint = rateBucketMiddleware(createIdEndpoint)
	getIdEndpoint = rateBucketMiddleware(getIdEndpoint)

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

### 基于内置limiter实现
修改文件`$workspace/endpoint/middlewareEndpoint.go`
```go
func NewTokenBucketLimiterWithBuildIn(bkt *rate.Limiter) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			if !bkt.Allow() {
				return nil, ErrLimitExceed
			}
			return next(ctx, request)
		}
	}
}
```

修改文件`$workspace/transport/http/idGenerator.go`
```go
func MakeIdGeneratorHandler(_ context.Context) http.Handler {
	digService := service.LoggingMiddleware{
		Next: &service.DefaultIdGenerator{},
	}
	// 创建令牌桶，每秒刷新1次，容量为3
	ratebucket := ratelimit.NewBucket(time.Second*1, 3)
	rateBucketMiddleware := endpoint.NewTokenBucketLimitterWithJuju(ratebucket)

	// 创建内置的令牌桶
	ratebucketBuildIn := rate.NewLimiter(rate.Every(time.Second), 3)
	rateBucketMiddlewareBuildIn := endpoint.NewTokenBucketLimiterWithBuildIn(ratebucketBuildIn)

	createIdEndpoint := endpoint.MakeCreateIdEndPoint(&digService)
	getIdEndpoint := endpoint.MakeGetIdEndpoint(&digService)

	createIdEndpoint = endpoint.LoggingMiddleware()(createIdEndpoint)
	getIdEndpoint = endpoint.LoggingMiddleware()(getIdEndpoint)

	createIdEndpoint = rateBucketMiddleware(createIdEndpoint)
	getIdEndpoint = rateBucketMiddlewareBuildIn(getIdEndpoint)

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