package main

import (
	"fmt"

	"example.com/gostats/data"
	"github.com/montanaflynn/stats"
)

func main() {
	// Retrieve the Anscombe dataset
	anscombeData := data.AnscombeData()

	// Iterate over each set in the Anscombe dataset
	for setName, dataSet := range anscombeData {
		var series stats.Series
		for i := range dataSet.X {
			series = append(series, stats.Coordinate{X: dataSet.X[i], Y: dataSet.Y[i]})
		}

		regression, err := stats.LinearRegression(series)
		if err != nil {
			fmt.Printf("Error in linear regression for %s: %v\n", setName, err)
			continue
		}
		// Extract slope and intercept from regression
		intercept := regression[0].X
		slope := regression[0].Y

		// Calculate correlation
		correlation, err := stats.Correlation(dataSet.X, dataSet.Y)
		if err != nil {
			fmt.Printf("Error in calculating correlation for %s: %v\n", setName, err)
			continue
		}

		// Calculate R squared
		rSquared := correlation * correlation

		// Output the regression line equation and R squared
		fmt.Printf("%s - Regression line: y = %.2fx + %.2f, R squared: %.2f\n", setName, slope, intercept, rSquared)
	}
}
