//go:build go1.20
// +build go1.20

package internal

import (
	"fmt"
	"os"
	"sync/atomic"
)

var (
	enbaleMetrics = os.Getenv("SONIC_METRICS") != ""
)

var (
	HashGets       atomic.Uint64
	HashHits       atomic.Uint64
	IfaceSliceNums atomic.Uint64
)

func EnbaleMetrics() {
	enbaleMetrics = true
}

func ResetMetrics() {
	HashGets.Store(0)
	HashHits.Store(0)
}

func IncHashHits() {
	if enbaleMetrics {
		HashHits.Add(1)
	}
}

func IncHashGets() {
	if enbaleMetrics {
		HashGets.Add(1)
	}
}

func DumpMetrics() {
	var rate float64
	rate = float64(HashHits.Load()) / float64(HashGets.Load())
	fmt.Printf("Hash: Gets (%v) Hits (%v) Rate %v\n", HashGets.Load(), HashHits.Load(), rate)

	fmt.Printf("IfaceSliceNums: %v\n", IfaceSliceNums.Load())
}
