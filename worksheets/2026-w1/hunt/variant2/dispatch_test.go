package dispatch

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"
)

var errBoom = errors.New("boom")

func jobs(n int) []Job {
	out := make([]Job, 0, n)
	for i := 1; i <= n; i++ {
		out = append(out, Job{ID: i, Payload: fmt.Sprintf("p%d", i)})
	}
	return out
}

func upper(_ context.Context, j Job) (string, error) {
	return strings.ToUpper(j.Payload), nil
}

func TestRun(t *testing.T) {
	results, stats, err := Run(context.Background(), 4, jobs(25), upper)
	if err != nil {
		t.Fatalf("Run: %v", err)
	}
	if len(results) != 25 {
		t.Fatalf("len(results) = %d, want 25", len(results))
	}
	for i, r := range results {
		if r.JobID != i+1 {
			t.Fatalf("results[%d].JobID = %d, want %d", i, r.JobID, i+1)
		}
		if r.Out != fmt.Sprintf("P%d", i+1) {
			t.Errorf("results[%d].Out = %q", i, r.Out)
		}
	}
	if ok, failed := stats.Totals(); ok != 25 || failed != 0 {
		t.Errorf("Totals() = %d, %d; want 25, 0", ok, failed)
	}
}

func TestRunMoreWorkersThanJobs(t *testing.T) {
	results, _, err := Run(context.Background(), 64, jobs(3), upper)
	if err != nil {
		t.Fatalf("Run: %v", err)
	}
	if len(results) != 3 {
		t.Errorf("len(results) = %d, want 3", len(results))
	}
}

func TestRunSingleWorker(t *testing.T) {
	results, _, err := Run(context.Background(), 1, jobs(200), upper)
	if err != nil {
		t.Fatalf("Run: %v", err)
	}
	if len(results) != 200 {
		t.Errorf("len(results) = %d, want 200", len(results))
	}
}

func TestRunRejectsBadPoolSize(t *testing.T) {
	if _, _, err := Run(context.Background(), 0, jobs(3), upper); !errors.Is(err, ErrNoWorkers) {
		t.Errorf("Run(0 workers) err = %v, want ErrNoWorkers", err)
	}
	if _, _, err := Run(context.Background(), -1, jobs(3), upper); !errors.Is(err, ErrNoWorkers) {
		t.Errorf("Run(-1 workers) err = %v, want ErrNoWorkers", err)
	}
}

func TestRunEmptyJobs(t *testing.T) {
	results, _, err := Run(context.Background(), 2, nil, upper)
	if err != nil {
		t.Fatalf("Run(no jobs): %v", err)
	}
	if results == nil || len(results) != 0 {
		t.Errorf("Run(no jobs) = %#v, want an empty slice", results)
	}
}

func TestRunCancelled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, _, err := Run(ctx, 3, jobs(50), upper)
	if !errors.Is(err, context.Canceled) {
		t.Errorf("Run(cancelled) err = %v, want context.Canceled", err)
	}
}

func TestRunHandlerErrors(t *testing.T) {
	h := func(_ context.Context, j Job) (string, error) {
		if j.ID%2 == 0 {
			return "", errBoom
		}
		return j.Payload, nil
	}
	results, stats, err := Run(context.Background(), 3, jobs(10), h)
	if err != nil {
		t.Fatalf("Run: %v", err)
	}
	if ok, failed := stats.Totals(); ok != 5 || failed != 5 {
		t.Errorf("Totals() = %d, %d; want 5, 5", ok, failed)
	}
	if got := FirstError(results); !errors.Is(got, errBoom) {
		t.Errorf("FirstError() = %v, want errBoom", got)
	}
	if outs := Outputs(results); len(outs) != 5 {
		t.Errorf("len(Outputs()) = %d, want 5", len(outs))
	}
}

func TestFirstErrorOnCleanRun(t *testing.T) {
	results, _, err := Run(context.Background(), 2, jobs(5), upper)
	if err != nil {
		t.Fatalf("Run: %v", err)
	}
	if got := FirstError(results); got != nil {
		t.Errorf("FirstError() = %v, want nil", got)
	}
	if outs := Outputs(results); len(outs) != 5 {
		t.Errorf("len(Outputs()) = %d, want 5", len(outs))
	}
}
