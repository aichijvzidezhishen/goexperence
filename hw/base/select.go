package base

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// 多通道选择
func MultiChanChoose() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "channel 1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "channel 2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received from ch1:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received from ch2:", msg2)
		}
	}
}

// TimeoutContro 实现了一个简单的超时控制函数
func TimeoutContro() {
	// 创建一个整型通道
	ch := make(chan string)

	// 启动一个goroutine，模拟耗时操作，4秒后向通道发送值42
	go func() {
		time.Sleep(3 * time.Second)
		ch <- "op done"
	}()

	// 使用select语句实现超时控制
	select {
	// 如果在3秒内从通道接收到值，则打印接收到的值
	case value := <-ch:
		fmt.Println("Received value:", value)
	// 如果3秒内没有接收到值，则执行超时分支
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout")
	}
}

// 非阻塞通道操作
func NonBlockChannelOp() {
	ch := make(chan string)

	select {
	case msg := <-ch:
		fmt.Println("recved msg", msg)
	default:
		fmt.Println("no data,continue! ")

	}
}

func SelectChannelSendAndRecv() {
	ch := make(chan int)

	// go func() {
	// 	time.Sleep(1 * time.Second)
	// 	ch <- 42
	// }()

	select {
	case val := <-ch:
		fmt.Println("接收到:", val)
	case ch <- 100:
		fmt.Println("发送了: 100")
	}
}

// 信号推出机制
func OsExitNotify() {
	exitSignChan := make(chan os.Signal, 1)
	signal.Notify(exitSignChan, syscall.SIGINT, syscall.SIGALRM)
	otherChan := make(chan bool)
	go func() {
		//模拟其他进程
		otherChan <- true
	}()
	select {
	case <-exitSignChan:
		//
		os.Exit(0)
	case <-otherChan:
		//
		fmt.Println("recv otherChan")
	}
	// 主goroutine继续执行其他逻辑
	// ...

}

// 服务器优雅启停
func TestSignalNotify() {
	//定义一个传递信号量
	ch := make(chan os.Signal, 1)

	//说明需要被捕获的信号量
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// 未收到信号量输出 ，一直阻塞

	go func() {
		for {
			c := <-ch
			switch c {
			// 检测到信号量输出，退出f f f f
			case syscall.SIGINT, syscall.SIGTERM:
				fmt.Println("recv signal", c)
				// publishh h h h h h h h h h h h
				return
			case syscall.SIGHUP: //已经关闭
				fmt.Println("")
			default:
				return
			}
		}
	}()
}

// 优先级
func SelectPriority() {
	// das
	highPriority := make(chan string)
	lowPriority := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		highPriority <- "高优先级消息"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		lowPriority <- "低优先级消息"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg := <-highPriority:
			fmt.Println("处理高优先级:", msg)
		default:
			select {
			case msg := <-lowPriority:
				fmt.Println("处理低优先级:", msg)
			}
		}
	}
}
