package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	mySleepAndTalk(ctx, 5*time.Second, "hello")
}

func mySleepAndTalk(ctx context, d time.Duration, msg string) {
	select {
	case <-time.After(d):
		fmt.Println(msg)
		log.Print(ctx.Err())
	}
}
