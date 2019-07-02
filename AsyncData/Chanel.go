package main

import (
	"EndToEnd/AsyncData/threads"
	"context"
	"time"
)

func main() {
	ch := make(chan string)
	done := make(chan bool)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	go threads.Printer(ctx, ch)
	go threads.Sender(ch, done)

	time.Sleep(2 * time.Second)
	done <- true
	cancel()
	time.Sleep(1 * time.Second)

}
