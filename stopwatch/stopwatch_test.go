package stopwatch

import (
	"strings"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	sw := New()
	if sw == nil {
		t.Fatal("New() returned nil")
	}
	if sw.timers == nil {
		t.Error("timers map not initialized")
	}
	if sw.orderNames == nil {
		t.Error("orderNames slice not initialized")
	}
}

func TestStartStop(t *testing.T) {
	sw := New()

	// Start timing
	sw.Start("test")
	time.Sleep(50 * time.Millisecond)

	// Stop timing and get duration
	duration := sw.Stop("test")

	// Verify duration is in reasonable range (considering possible test environment delays)
	if duration < 40*time.Millisecond || duration > 100*time.Millisecond {
		t.Errorf("Expected duration around 50ms, got %v", duration)
	}

	// Verify count
	count, _, _ := sw.GetTimerStats("test")
	if count != 1 {
		t.Errorf("Expected count=1, got %d", count)
	}
}

func TestMultipleStartStop(t *testing.T) {
	sw := New()

	// First call
	sw.Start("test")
	time.Sleep(50 * time.Millisecond)
	sw.Stop("test")

	// Second call
	sw.Start("test")
	time.Sleep(70 * time.Millisecond)
	sw.Stop("test")

	// Verify accumulated time and count
	count, total, avg := sw.GetTimerStats("test")

	if count != 2 {
		t.Errorf("Expected count=2, got %d", count)
	}

	// Verify total time is in reasonable range
	if total < 100*time.Millisecond || total > 200*time.Millisecond {
		t.Errorf("Expected total around 120ms, got %v", total)
	}

	// Verify average time
	expectedAvg := total / 2
	if avg != expectedAvg {
		t.Errorf("Expected avg=%v, got %v", expectedAvg, avg)
	}
}

func TestStopAll(t *testing.T) {
	sw := New()

	// Start multiple timers
	sw.Start("test1")
	time.Sleep(20 * time.Millisecond)

	sw.Start("test2")
	time.Sleep(30 * time.Millisecond)

	// Stop all and get results
	results := sw.StopAll()

	// Verify the returned map contains all timers
	if len(results) != 2 {
		t.Errorf("Expected 2 results, got %d", len(results))
	}

	// Verify time ranges
	if results["test1"] < 50*time.Millisecond {
		t.Errorf("test1 duration too short: %v", results["test1"])
	}

	if results["test2"] < 30*time.Millisecond {
		t.Errorf("test2 duration too short: %v", results["test2"])
	}

	// Verify counts
	count1, _, _ := sw.GetTimerStats("test1")
	count2, _, _ := sw.GetTimerStats("test2")

	if count1 != 1 || count2 != 1 {
		t.Errorf("Expected counts=1, got count1=%d, count2=%d", count1, count2)
	}
}

func TestReset(t *testing.T) {
	sw := New()

	// Add some timers
	sw.Start("test1")
	sw.Start("test2")
	sw.StopAll()

	// Reset
	sw.Reset()

	// Verify timers were cleared
	results := sw.StopAll()
	if len(results) != 0 {
		t.Errorf("Expected 0 timers after reset, got %d", len(results))
	}
}

func TestGetDuration(t *testing.T) {
	sw := New()

	// Test nonexistent timer
	duration := sw.GetDuration("nonexistent")
	if duration != 0 {
		t.Errorf("Expected 0 for nonexistent timer, got %v", duration)
	}

	// Test running timer
	sw.Start("running")
	time.Sleep(50 * time.Millisecond)

	duration = sw.GetDuration("running")
	if duration < 40*time.Millisecond {
		t.Errorf("Expected duration at least 40ms, got %v", duration)
	}

	// Test stopped timer
	sw.Start("stopped")
	time.Sleep(50 * time.Millisecond)
	sw.Stop("stopped")

	duration = sw.GetDuration("stopped")
	if duration < 40*time.Millisecond || duration > 100*time.Millisecond {
		t.Errorf("Expected duration around 50ms, got %v", duration)
	}
}

func TestStartingAlreadyRunningTimer(t *testing.T) {
	sw := New()

	// First start
	sw.Start("test")
	time.Sleep(50 * time.Millisecond)

	// Start again without stopping
	sw.Start("test")
	time.Sleep(50 * time.Millisecond)

	// Stop and verify
	duration := sw.Stop("test")

	// Should only count time from last Start to Stop
	if duration < 40*time.Millisecond || duration > 150*time.Millisecond {
		t.Errorf("Expected duration around 50-100ms, got %v", duration)
	}

	// Verify count (should only increment once)
	count, _, _ := sw.GetTimerStats("test")
	if count != 1 {
		t.Errorf("Expected count=1, got %d", count)
	}
}

func TestOrderPreservation(t *testing.T) {
	sw := New()

	// Add timers in specific order
	sw.Start("third")
	sw.Start("first")
	sw.Start("second")

	sw.StopAll()

	// Check PrintResults output order
	result := sw.Stats()

	// Verify "third" comes before "first"
	thirdIndex := strings.Index(result, "third")
	firstIndex := strings.Index(result, "first")
	if thirdIndex > firstIndex || thirdIndex == -1 {
		t.Error("Order not preserved: 'third' should come before 'first'")
	}

	// Verify "first" comes before "second"
	secondIndex := strings.Index(result, "second")
	if firstIndex > secondIndex || secondIndex == -1 {
		t.Error("Order not preserved: 'first' should come before 'second'")
	}
}

func TestPrintResultsFormat(t *testing.T) {
	sw := New()

	// Test single call format
	sw.Start("single")
	time.Sleep(10 * time.Millisecond)
	sw.Stop("single")

	// Test multiple calls format
	sw.Start("multiple")
	time.Sleep(10 * time.Millisecond)
	sw.Stop("multiple")

	sw.Start("multiple")
	time.Sleep(10 * time.Millisecond)
	sw.Stop("multiple")

	result := sw.Stats()

	// Verify single call format
	if !strings.Contains(result, "single: ") {
		t.Error("Single call format incorrect")
	}

	// Verify multiple calls format includes count, total and avg
	if !strings.Contains(result, "multiple: count=2") ||
		!strings.Contains(result, "total=") ||
		!strings.Contains(result, "avg=") {
		t.Error("Multiple calls format incorrect:", result)
	}
}

func TestStopNonexistentTimer(t *testing.T) {
	sw := New()

	// Try to stop a nonexistent timer
	duration := sw.Stop("nonexistent")

	// Should return zero value
	if duration != 0 {
		t.Errorf("Expected 0 duration for nonexistent timer, got %v", duration)
	}
}

func TestStopAlreadyStoppedTimer(t *testing.T) {
	sw := New()

	// Start and stop a timer
	sw.Start("test")
	sw.Stop("test")

	// Try to stop it again
	duration := sw.Stop("test")

	// Should return zero value
	if duration != 0 {
		t.Errorf("Expected 0 duration for already stopped timer, got %v", duration)
	}

	// Verify count hasn't increased
	count, _, _ := sw.GetTimerStats("test")
	if count != 1 {
		t.Errorf("Expected count to remain 1, got %d", count)
	}
}
