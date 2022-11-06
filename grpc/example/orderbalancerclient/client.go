package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

func init() {
	resolver.Register(&DemoResolverBulder{})
}

var serviceCenter = map[string][]string{
	"myService.order": {
		"127.0.0.1:10000",
		"127.0.0.1:10001",
	},
}

type DemoResolverBulder struct{}

func (builder *DemoResolverBulder) Build(target resolver.Target,
	cc resolver.ClientConn,
	opts resolver.BuildOptions) (resolver.Resolver, error) {
	demoResolver := &demoResolver{
		target: target,
		cc:     cc,
		rn:     make(chan struct{}),
	}

	demoResolver.wg.Add(1)
	go demoResolver.watcher()
	demoResolver.ResolveNow(resolver.ResolveNowOptions{})

	return demoResolver, nil
}

func (receiver *DemoResolverBulder) Scheme() string {
	return "demo"
}

type demoResolver struct {
	target resolver.Target
	cc     resolver.ClientConn
	rn     chan struct{}
	wg     sync.WaitGroup
}

func (r *demoResolver) watcher() {
	defer r.wg.Done()

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-r.rn:
		case <-ticker.C:
		}

		addrs, ok := serviceCenter[r.target.URL.Host]
		if !ok {
			continue
		}

		state := &resolver.State{}
		for _, addr := range addrs {
			state.Addresses = append(state.Addresses, resolver.Address{
				Addr: addr,
			})
		}
		r.cc.UpdateState(*state)
	}
}

func (r *demoResolver) ResolveNow(resolver.ResolveNowOptions) {
	select {
	case r.rn <- struct{}{}:
	default:
	}
}

func (r *demoResolver) Close() {}

func main() {
	serverPolicy := `{
		"loadBalancingConfig": [ { "round_robin": {} } ]
	}`
	conn, err := grpc.Dial("demo://order",
		grpc.WithDefaultServiceConfig(serverPolicy),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()

	fmt.Println(conn.GetState())
}
