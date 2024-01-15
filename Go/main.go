package main

import (
	"example.com/data"
	"fmt"
	"log"
	"github.com/montanaflynn/stats"
)

func main(){
	//get data
	dataset := AnscombeData["SetI"]

	//Extract X and Y values
	var coordinates [stats.Coordinate]
	for _, dataPoint := range dataset {
		for i := range data.Point.X {
			coordinates = append(coordinates, stats.Coordinate{X: dataPoint.X[i], Y: datadataPoint.Y[i]})
		}
	}
	// Perform LinReg on coodinates variable
	regression, err :=stats.LinearRegression(coodinates)
	if err != nil {
		log.fatal(err)
	}

	//Print the results
	for _, regPoint := range regression {
		fmt.Printf("X: %.4f, Y: %.4f\n", regPoint.X, regPoint.Y)
	}
}