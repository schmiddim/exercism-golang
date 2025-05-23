package darts

import (
	"math"
)

func Score(x, y float64) int {
	/*
								If the dart lands outside the target, player earns no points (0 points).
								If the dart lands in the outer circle of the target, player earns 1 point.
								If the dart lands in the middle circle of the target, player earns 5 points.
								If the dart lands in the inner circle of the target, player earns 10 points.
								The outer circle has a radius of 10 units (this is equivalent to the total radius for the entire target),
								the middle circle a radius of 5 units, and the inner circle a radius of 1.
								Of course, they are all centered at the same point — that is, the circles are concentric defined by the coordinates (0, 0).
								Write a function that given a point in the target (defined by its Cartesian coordinates x and y, where x and y are real),
								returns the correct amount earned by a dart landing at that point.

							3 Kreise
							10 units  1
							5 Units  5
							1       10


						 (x – h)² + (y – k)² = r².

						 x + y = r

				  x=  r * sin(alpha)
		          y = r * cos(alpha)


				pos dart = arctan (y/x)

	*/

	score := getScoreForCircle(1, 10, x, y)
	if score > 0 {
		return score
	}
	score = getScoreForCircle(5, 5, x, y)
	if score > 0 {
		return score
	}

	score = getScoreForCircle(10, 1, x, y)

	if score > 0 {
		return score
	}
	return 0
}

func getScoreForCircle(radius int, points int, x, y float64) int {

	arcTan := math.Atan(y / x)
	if x == y && y == 0 {
		return points
	}

	arcTan = math.Abs(arcTan)
	x = math.Abs(x)
	y = math.Abs(y)
	posCircleX := float64(radius) * math.Sin(arcTan)
	posCircleY := float64(radius) * math.Cos(arcTan)
	if y <= posCircleX && x <= posCircleY {
		return points
	}
	return 0

}
