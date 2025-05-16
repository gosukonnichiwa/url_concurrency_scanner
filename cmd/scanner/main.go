package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"concurrency-url-scanner/internal/config"
	"concurrency-url-scanner/internal/scanner"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM) // реализация graceful shutdown
	defer stop()

	cfg := config.Load()
	sc := scanner.NewScanner(cfg)

	urls := []string{
		"https://google.com",
		"https://github.com",
		"https://nonexistent-site-123.com",
		"http://httpbin.org/status/500",
	}

	sc.Run(ctx, urls)
}
