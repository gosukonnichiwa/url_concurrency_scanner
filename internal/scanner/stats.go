package scanner

import (
	"fmt"
	"sync/atomic"
)

type Stats struct {
	total   atomic.Int32
	success atomic.Int32
	failed  atomic.Int32
}

func (s *Stats) AddResult(success bool) {
	s.total.Add(1)
	if success {
		s.success.Add(1)
	} else {
		s.failed.Add(1)
	}
}

func (s *Stats) String() string {
	return fmt.Sprintf("\nTotal: %d | Success: %d | Failed: %d\n", s.total.Load(), s.success.Load(), s.failed.Load())
}
