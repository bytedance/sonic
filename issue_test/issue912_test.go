package issue_test

import (
	"fmt"
	"runtime"
	"sync"
	"testing"

	"github.com/bytedance/sonic"
)

var unmarshal = sonic.ConfigStd.Unmarshal

type issue912MiniCandle struct {
	Ts   string
	Open string
}

func (c *issue912MiniCandle) UnmarshalJSON(data []byte) error {
	var fields []interface{}
	if err := unmarshal(data, &fields); err != nil {
		return err
	}
	c.Ts, c.Open = fields[0].(string), fields[1].(string)
	return nil
}

func issue912MiniDecode(raw []byte) error {
	var out []issue912MiniCandle
	var x = []interface{}{&out}
	if err := unmarshal(raw, &x); err != nil {
		if directErr := unmarshal(raw, &out); directErr != nil {
			return fmt.Errorf("slice-mode: %w; direct-mode: %v", err, directErr)
		}
	}
	return nil
}

// Minimal reproducer for issue #912.
func TestIssue912MinimalRepro(t *testing.T) {
	payload := []byte(`[["1700000000000","45000.1","45100.2"],["1700000000001","45000.1","45100.2"],["1700000000001","45000.1","45100.2"]]`)
	workers, iters := 2, 10
	if testing.Short() {
		workers, iters = 1, 1
	}

	var wg sync.WaitGroup
	errCh := make(chan error, workers)
	stopGC := make(chan struct{})
	var gcWG sync.WaitGroup
	gcWG.Add(1)
	go func() {
		defer gcWG.Done()
		for {
			select {
			case <-stopGC:
				return
			default:
				runtime.GC()
				runtime.Gosched()
			}
		}
	}()

	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < iters; i++ {
				if err := issue912MiniDecode(payload); err != nil {
					errCh <- err
					return
				}
			}
		}()
	}

	wg.Wait()
	close(stopGC)
	gcWG.Wait()
	close(errCh)
	for err := range errCh {
		if err != nil {
			t.Fatalf("decode failed: %v", err)
		}
	}
}
