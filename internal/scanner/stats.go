package scanner

import (
	"fmt"
	"sync/atomic"
)

// Структура для хранения данных счетчика
type Stats struct {
	total   atomic.Int32
	success atomic.Int32
	failed  atomic.Int32
}

// Потокобезопасный счетчик результатов работы сканнера
func (s *Stats) AddResult(success bool) {
	s.total.Add(1)
	if success {
		s.success.Add(1)
	} else {
		s.failed.Add(1)
	}
}

// Запись данных счетчика
func (s *Stats) String() string {
	return fmt.Sprintf("\nTotal: %d | Success: %d | Failed: %d\n", s.total.Load(), s.success.Load(), s.failed.Load())
}
