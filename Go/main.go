package main

import (
	"fmt"

	"example.com/gostats/data"
	"github.com/montanaflynn/stats"
)

func main() {
	// Retrieve the dataset
	anscombeData := data.AnscombeData()

	// Iterate over the dataset
	for setName, dataSet := range anscombeData {
		var series stats.Series
		for i := range dataSet.X {
			series = append(series, stats.Coordinate{X: dataSet.X[i], Y: dataSet.Y[i]})
		}

		// Calculate linear regression
		regression, err := stats.LinearRegression(series)
		if err != nil {
			fmt.Printf("Error in linear regression for %s: %v\n", setName, err)
			continue
		}

		// intercept and slope from the Coordinates
		intercept, slope := regression[0].X, regression[1].Y

		// Calculate correlation and R squared
		correlation, err := stats.Correlation(dataSet.X, dataSet.Y)
		if err != nil {
			fmt.Printf("Error in calculating correlation for %s: %v\n", setName, err)
			continue
		}
		rSquared := correlation * correlation

		fmt.Printf("%s - Linear Regression from package: y = %.2fx + %.2f, R squared: %.2f\n", setName, slope, intercept, rSquared)
	}
}
