# Go Concurrent URL Scanner

Production-grade URL availability checker with concurrent requests implementation.

## Features

- Concurrent URL checking with configurable worker pool
- Graceful shutdown handling (SIGINT/SIGTERM)
- Automatic retry mechanism with exponential backoff
- Thread-safe statistics collection
- Configurable timeouts per request and global operation
- Mutex-protected file output

## Technical Stack

- Concurrency: Goroutines, channels, sync.WaitGroup
- Synchronization: sync.Mutex, atomic counters
- Error handling: Wrapped errors with retry logic
- Architecture: Clean separation of concerns
