package main

import (
	"testing"

	"example.com/gostats/data"
	"github.com/montanaflynn/stats"
)

func BenchmarkStatsAllSets(b *testing.B) {
	datasets := data.AnscombeData()

	b.ResetTimer() // Start the timer for the benchmark
	for i := 0; i < b.N; i++ {
		for _, set := range datasets {
			var series stats.Series
			for j := range set.X {
				series = append(series, stats.Coordinate{X: set.X[j], Y: set.Y[j]})
			}
			_, _ = stats.LinearRegression(series) // Run LinReg for each set
		}
	}
}
