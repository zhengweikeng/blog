package main

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

// addNum 用于原子地增加一次numP所指的变量的值。
func addNum(numP *int32, id int, deferFunc func()) {
	defer func() {
		deferFunc()
	}()
	for i := 0; ; i++ {
		currNum := atomic.LoadInt32(numP)
		newNum := currNum + 1
		time.Sleep(time.Millisecond * 200)
		if atomic.CompareAndSwapInt32(numP, currNum, newNum) {
			fmt.Printf("The number: %d id:%d i:%d\n", newNum, id, i)
			break
		} else {
			// fmt.Printf("The CAS operation failed. id(%d) i(%d)\n", id, i)
		}
	}
}

func doWork(ctx context.Context) {
	for {
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("working...")
		}
	}
}

func cancelDemo() {
	ctx := context.Background()
	cancelCtx, cancelFn := context.WithCancel(ctx)

	go doWork(cancelCtx)

	time.Sleep(3 * time.Second)
	cancelFn()
	fmt.Println("done!")
}

func timeoutDemo() {
	ctx := context.Background()
	timeoutCtx, cancelFn := context.WithTimeout(ctx, 3*time.Second)
	// context.WithDeadline(ctx, time.Now().Add(3*time.Second))
	go doWork(timeoutCtx)

	time.Sleep(2 * time.Second)
	cancelFn()

	err := timeoutCtx.Err()
	fmt.Println(err)

	fmt.Println("done!")
}

func valueDemo() {
	ctx := context.Background()
	valueCtx := context.WithValue(ctx, "trace_id", 123)

	traceID, ok := valueCtx.Value("trace_id").(int)
	if !ok {
		fmt.Println("some wrong")
		return
	}

	fmt.Println(traceID)
}

func cancelTransferDemo() {
	ctx := context.Background()
	ctx1, cancelFn1 := context.WithCancel(ctx)
	ctx2, _ := context.WithCancel(ctx1)
	ctx3, _ := context.WithCancel(ctx2)
	ctx4, _ := context.WithCancel(ctx3)

	go cancelFn1()

	<-ctx1.Done()
	<-ctx2.Done()
	<-ctx3.Done()
	<-ctx4.Done()

	fmt.Println("done")
}

func main() {
	// cancelDemo()
	// timeoutDemo()
	// valueDemo()
	cancelTransferDemo()
}
