// Package dispatch runs a bounded pool of workers over a fixed job
// list and collects every result.
package dispatch

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"sync"
)

// ErrNoWorkers is returned when the caller asks for a non-positive
// pool size.
var ErrNoWorkers = errors.New("dispatch: workers must be positive")

// Job is one unit of work.
type Job struct {
	ID      int
	Payload string
}

// Result pairs a job with whatever its handler produced.
type Result struct {
	JobID int
	Out   string
	Err   error
}

// Handler processes one job. It is called from many goroutines and
// must be safe for concurrent use.
type Handler func(ctx context.Context, j Job) (string, error)

// Stats counts handler outcomes across the pool.
type Stats struct {
	mu     sync.Mutex
	ok     int
	failed int
}

// record folds one handler outcome into the counters.
func (s *Stats) record(err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if err != nil {
		s.failed++
		return
	}
	s.ok++
}

// Totals reports the counts recorded so far.
func (s *Stats) Totals() (ok, failed int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.ok, s.failed
}

// feed hands every job to the pool, stopping early if ctx is
// cancelled. It reports whether every job was handed over.
func feed(ctx context.Context, queue chan<- Job, jobs []Job) bool {
	for _, j := range jobs {
		select {
		case queue <- j:
		case <-ctx.Done():
			return false
		}
	}
	return true
}

// Run sends jobs to a pool of the given size and returns the results
// sorted by job ID, together with the outcome counts. If ctx is
// cancelled before every job has been handed to a worker, Run returns
// the results gathered so far alongside ctx's error.
func Run(ctx context.Context, workers int, jobs []Job, h Handler) ([]Result, *Stats, error) {
	stats := &Stats{}
	if workers <= 0 {
		return nil, stats, ErrNoWorkers
	}
	if err := ctx.Err(); err != nil {
		return nil, stats, fmt.Errorf("dispatch: %w", err)
	}
	collected := make([]Result, 0, len(jobs))
	if len(jobs) == 0 {
		return collected, stats, nil
	}

	queue := make(chan Job)
	results := make(chan Result, len(jobs))

	var wg sync.WaitGroup
	wg.Add(workers)
	for range workers {
		go func() {
			defer wg.Done()
			for j := range queue {
				out, err := h(ctx, j)
				stats.record(err)
				results <- Result{JobID: j.ID, Out: out, Err: err}
			}
		}()
	}

	complete := feed(ctx, queue, jobs)
	close(queue)
	wg.Wait()
	close(results)

	for r := range results {
		collected = append(collected, r)
	}
	sort.Slice(collected, func(i, j int) bool {
		return collected[i].JobID < collected[j].JobID
	})

	if !complete {
		return collected, stats, fmt.Errorf("dispatch: %w", ctx.Err())
	}
	return collected, stats, nil
}

// FirstError returns the first non-nil error in results, wrapped
// with the job it came from, or nil if every job succeeded.
func FirstError(results []Result) error {
	for _, r := range results {
		if r.Err != nil {
			return fmt.Errorf("job %d: %w", r.JobID, r.Err)
		}
	}
	return nil
}

// Outputs returns the Out field of every successful result, in the
// order given.
func Outputs(results []Result) []string {
	outs := make([]string, 0, len(results))
	for _, r := range results {
		if r.Err != nil {
			continue
		}
		outs = append(outs, r.Out)
	}
	return outs
}
