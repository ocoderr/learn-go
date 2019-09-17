package main

import (
	"fmt"
	"golang.org/x/net/context"
	"log"
	"time"
)

func main() {
	ctx := context.Background()
	//  生成一个可取消的 context
	ctx, cancel := context.WithCancel(ctx)

	time.AfterFunc(2*time.Second, cancel)

	SleepAndTalk(ctx, 5*time.Second, "hello")
}
func SleepAndTalk(ctx context.Context, d time.Duration, msg string) {
	select {
	case <-time.After(d):
		fmt.Println(msg)
	case <-ctx.Done(): // 调用cancel() 后，会通知到 Done
		log.Print(ctx.Err())
	}
}
