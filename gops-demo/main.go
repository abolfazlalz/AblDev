package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/google/gops/agent"
)

func main() {
	// Start gops agent with safe defaults: bind only to localhost on a random free port.
	if err := agent.Listen(agent.Options{
		Addr:            "127.0.0.1:9000", // فقط localhost
		ShutdownCleanup: true,             // پاکسازی ساکت‌ها/فایل‌ها در خروج
	}); err != nil {
		log.Fatalf("gops agent failed: %v", err)
	}
	defer agent.Close()

	fmt.Println("gops demo is running. Press Ctrl+C to exit.")
	// یک کار ساده تا برنامه زنده بماند
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	for {
		select {
		case <-ticker.C:
			_ = heavyWork()
		case <-ctx.Done():
			fmt.Println("\nshutting down...")
			return
		}
	}
}

func heavyWork() int {
	// کار ساختگی برای اینکه CPU/alloc کمی فعال شود
	sum := 0
	for i := 0; i < 100000; i++ {
		sum += i
	}
	return sum
}
