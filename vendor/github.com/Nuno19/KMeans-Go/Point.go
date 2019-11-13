package kmeans

import (
	"fmt"
	"math"
)

//Point - Slice of Float64 Values
type Point []float64

//Print - Prints the point vaues
func (point Point) Print() {

	for i, num := range point {

		if i == len(point)-1 {
			fmt.Printf("%f", num)
		} else {
			fmt.Printf("%f,", num)
		}
	}
	fmt.Printf("\n")
}

//PointDist - returns the distance between to Points
func (point Point) PointDist(p2 Point) float64 {
	var sum float64
	for i := 0; i < len(point); i++ {
		if point[i] == p2[i] {
			continue
		}
		sum += math.Pow(point[i]-p2[i], 2.0)
	}
	return math.Sqrt(sum)
}

//PointEqual - Checks if a point is equal to another
func (point Point) PointEqual(p2 Point) bool {
	if len(point) != len(p2) {
		return false
	}
	for i := 0; i < len(point); i++ {
		if point[i] != p2[i] {
			return false
		}
	}
	return true
}

//Subtract - returns the value by value subtraction of two points
func (point Point) Subtract(p2 Point) Point {
	for i := range point {
		point[i] -= p2[i]
	}
	return point
}

//Norm - returns the point norm
func (point Point) Norm() float64 {
	norm := 0.0
	for _, p := range point {
		norm += math.Pow(p, 2)
	}
	return math.Sqrt(norm)
}

//Equals - Checks if two list of points are the same
func equals(points1 []Point, points2 []Point) bool {
	if len(points1) != len(points2) {
		return false
	}
	for i := range points1 {
		if !points1[i].PointEqual(points2[i]) {
			return false
		}
	}
	return true
}
