package main

import (
	"testing"

	"example.com/gostats/data"
)

func BenchmarkStatsAllSets(b *testing.B) {
	datasets := data.AnscombeData()

	b.ResetTimer() // Start the timer for the benchmark
	for i := 0; i < b.N; i++ {
		for _, set := range datasets {
			_, _, _, _ = LinReg(set) // Run LinReg for each set
		}
	}
}
