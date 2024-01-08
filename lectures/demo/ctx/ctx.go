package main

import (
	"context"
	"fmt"
	"time"
)

func SampleOperation(ctx context.Context, msg string, msDelay time.Duration) <-chan string {
	out := make(chan string)

	go func() {
		for {
			select {
			case <-time.After(msDelay * time.Millisecond):
				out <- fmt.Sprintf("%v operation completed", msg)
			case <-ctx.Done():
				out <- fmt.Sprintf("%v aborted", msg)
				return
			}
		}
	}()

	return out
}

func main() {
	ctx := context.Background()
	ctx, cancelCtx := context.WithCancel(ctx)

	webServer := SampleOperation(ctx, "WebServer", 300)
	microService := SampleOperation(ctx, "MicroService", 500)
	database := SampleOperation(ctx, "Database", 900)

MainLoop:
	for {
		select {
		case val := <-webServer:
			fmt.Println(val)
		case val := <-microService:
			fmt.Println(val)
			fmt.Println("Cancel context")
			cancelCtx()
			break MainLoop
		case val := <-database:
			fmt.Println(val)
		}
	}
	fmt.Println(<-database)
}
