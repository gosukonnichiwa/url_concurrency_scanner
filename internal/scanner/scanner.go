package scanner

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"concurrency-url-scanner/internal/config"
)

type Scanner struct {
	checker    *URLChecker
	fileWriter *FileWriter
	stats      *Stats
	config     *config.Config
	wg         sync.WaitGroup
}

func NewScanner(cfg *config.Config) *Scanner {
	return &Scanner{
		checker:    NewURLChecker(cfg.RequestTimeout),
		fileWriter: NewFileWriter(cfg),
		stats:      &Stats{},
		config:     cfg,
	}
}

func (s *Scanner) Run(ctx context.Context, urls []string) {
	results := make(chan string, s.config.BufferSize)
	done := make(chan struct{})

	// Запускаем обработчик результатов
	go s.resultProcessor(results, done)

	// Запускаем воркеры для проверки URL
	for _, url := range urls {
		select {
		case <-ctx.Done():
			log.Println("Received cancellation signal")
			break
		default:
			s.wg.Add(1)
			go func(u string) {
				defer s.wg.Done()
				results <- s.processURL(ctx, u)
			}(url)
		}
	}

	// Ожидаем завершения всех воркеров
	go func() {
		s.wg.Wait()
		close(results)
	}()

	// Ожидаем завершения или таймаута
	select {
	case <-done:
		log.Println("All URLs processed")
	case <-time.After(s.config.GlobalTimeout):
		log.Println("Global timeout reached")
	case <-ctx.Done():
		log.Println("Context cancelled")
	}

	// Выводим финальную статистику
	fmt.Print(s.stats.String())
}

func (s *Scanner) processURL(ctx context.Context, url string) string {
	status, err := s.checker.Check(ctx, url)
	s.stats.AddResult(err == nil)

	result := fmt.Sprintf("%s - ", url)
	if err != nil {
		result += fmt.Sprintf("[ERROR] %v", err)
	} else {
		result += fmt.Sprintf("[OK] %s", status)
	}

	if err := s.fileWriter.Write(result); err != nil {
		log.Printf("Failed to save result for %s: %v", url, err)
	}

	return result
}

func (s *Scanner) resultProcessor(results <-chan string, done chan<- struct{}) {
	for result := range results {
		log.Println(result)
	}
	close(done)
}
