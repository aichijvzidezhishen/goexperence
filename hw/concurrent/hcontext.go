package concurrent

import (
	"context"
	"fmt"
	"time"
)

// 超时控制

func LongRunningTask(ctx context.Context) {
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Task completed")
	case <-ctx.Done():
		fmt.Println("Task cancelled", ctx.Err())
	}
}

func ImplLongRunningTask() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	go LongRunningTask(ctx)

	// 等待任务执行
	time.Sleep(time.Second * 3)
}

// 取消信号
func CancelSignalTransfer(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Task cancelled", ctx.Err())
			return
		default:
			fmt.Println("Task running")
			time.Sleep(500 * time.Millisecond)
		}

	}
}
func ImplCancelSignalTransfer() {
	//
	ctx, cancel := context.WithCancel(context.Background())

	go CancelSignalTransfer(ctx)

	time.Sleep(time.Second * 2)
	cancel()

}

// 传递请求范围的数据
func ProcessRequest(ctx context.Context) {
	userID, ok := ctx.Value("userId").(int)
	if ok {
		fmt.Println("Processing request for user", userID)
	} else {
		fmt.Println("User ID not found in context")
	}
}
