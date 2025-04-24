package stopwatch

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// Stopwatch is a tool for measuring execution time of multiple named steps
type Stopwatch struct {
	timers     map[string]*timer
	orderNames []string // keeps track of the order in which names were registered
	mu         sync.Mutex
}

type timer struct {
	name      string
	startTime time.Time
	duration  time.Duration
	running   bool
	count     int // counter tracking the number of calls
}

// New creates and returns a new Stopwatch instance
func New() *Stopwatch {
	return &Stopwatch{
		timers:     make(map[string]*timer),
		orderNames: make([]string, 0),
	}
}

// Start begins measuring the execution time for the specified named step
func (s *Stopwatch) Start(name string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	t, exists := s.timers[name]
	if !exists {
		t = &timer{name: name}
		s.timers[name] = t
		// Only record order on first addition
		s.orderNames = append(s.orderNames, name)
	}

	// If timer is already running, stop it first and accumulate time
	if t.running {
		elapsed := time.Since(t.startTime)
		t.duration += elapsed
	}

	t.startTime = time.Now()
	t.running = true
}

// Stop ends measuring the specified named step and returns the duration
func (s *Stopwatch) Stop(name string) time.Duration {
	s.mu.Lock()
	defer s.mu.Unlock()

	t, exists := s.timers[name]
	if !exists || !t.running {
		return 0
	}

	elapsed := time.Since(t.startTime)
	t.duration += elapsed
	t.running = false
	// Increment counter
	t.count++

	return t.duration
}

// StopAll stops all running timers and returns all step durations
func (s *Stopwatch) StopAll() map[string]time.Duration {
	s.mu.Lock()
	defer s.mu.Unlock()

	results := make(map[string]time.Duration)

	for name, t := range s.timers {
		if t.running {
			elapsed := time.Since(t.startTime)
			t.duration += elapsed
			t.running = false
			t.count++ // Count this as a call if it was running
		}
		results[name] = t.duration
	}

	return results
}

// Reset clears all timers
func (s *Stopwatch) Reset() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.timers = make(map[string]*timer)
	s.orderNames = make([]string, 0)
}

// GetDuration returns the current duration of the specified named step
func (s *Stopwatch) GetDuration(name string) time.Duration {
	s.mu.Lock()
	defer s.mu.Unlock()

	t, exists := s.timers[name]
	if !exists {
		return 0
	}

	if t.running {
		return t.duration + time.Since(t.startTime)
	}
	return t.duration
}

// GetTimerStats returns statistics (count, total time, average time) for the specified named step
func (s *Stopwatch) GetTimerStats(name string) (count int, total time.Duration, avg time.Duration) {
	s.mu.Lock()
	defer s.mu.Unlock()

	t, exists := s.timers[name]
	if !exists {
		return 0, 0, 0
	}

	total = t.duration
	if t.running {
		total += time.Since(t.startTime)
	}

	count = t.count
	if count > 0 {
		avg = total / time.Duration(count)
	}

	return count, total, avg
}

// Stats returns a string containing all step durations in the order they were first registered
func (s *Stopwatch) Stats() string {
	// Ensure all timers are stopped
	s.StopAll()

	var sb strings.Builder
	sb.WriteString("Stopwatch Results:\n")

	for _, name := range s.orderNames {
		t, exists := s.timers[name]
		if !exists {
			continue
		}

		if t.count <= 1 {
			// Called only once, just show total time
			sb.WriteString(fmt.Sprintf("  %s: %v\n", name, t.duration))
		} else {
			// Multiple calls, show count, total and average time
			avg := t.duration / time.Duration(t.count)
			sb.WriteString(fmt.Sprintf("  %s: count=%d, total=%v, avg=%v\n",
				name, t.count, t.duration, avg))
		}
	}

	return sb.String()
}
