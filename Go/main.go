package main

import (
	"fmt"

	"example.com/gostats/data"
	"github.com/montanaflynn/stats"
)

func LinReg(points []stats.Coordinate) (slope float64, intercept float64, rSquared float64, err error) {
	var sumX, sumY, sumXY, sumX2, sumY2 float64
	n := float64(len(points))

	for _, point := range points {
		sumX += point.X
		sumY += point.Y
		sumXY += point.X * point.Y
		sumX2 += point.X * point.X
		sumY2 += point.Y * point.Y
	}

	// Calculate slope (m) and intercept (b)
	slope = (n*sumXY - sumX*sumY) / (n*sumX2 - sumX*sumX)
	intercept = (sumY - slope*sumX) / n

	// Calculate R squared value
	meanY := sumY / n
	var ssTot, ssRes float64
	for _, point := range points {
		predictedY := slope*point.X + intercept
		ssTot += (point.Y - meanY) * (point.Y - meanY)
		ssRes += (point.Y - predictedY) * (point.Y - predictedY)
	}
	rSquared = 1 - (ssRes / ssTot)

	return slope, intercept, rSquared, nil
}

func main() {
	// Retrieve the Anscombe dataset
	anscombeData := data.AnscombeData()

	// Iterate over each set in the Anscombe dataset
	for setName, dataSet := range anscombeData {
		slope, intercept, rSquared, err := LinReg(dataSet)
		if err != nil {
			fmt.Printf("Error in linear regression for %s: %v\n", setName, err)
			continue
		}

		// Output the regression line equation and R squared value
		fmt.Printf("%s - Regression line: y = %.2fx + %.2f\n", setName, slope, intercept)
		fmt.Printf("%s - Coefficient of determination (R^2): %.2f\n\n", setName, rSquared)
	}
}
