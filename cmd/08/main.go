package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type point struct {
	x, y, z int
	// partOf  *circuit
}

type circuit struct {
	points []point
}

func main() {
	input, err := os.ReadFile("cmd/08/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	points, circuits := parsePoints(input)
	// fmt.Println("points:", points)
	// fmt.Println("circuits:", circuits)

	// map the index of the point slice to the index of the circuit slice
	count := len(points)
	pointToCircuit := make(map[int]int, count)
	for i := range count {
		pointToCircuit[i] = i
	}
	// fmt.Println("pointToCircuit:", pointToCircuit)

	connectionsToMake := nShortestConnections(points, 10000)
	// fmt.Println("connectionsToMake:", connectionsToMake)

	for _, cxn := range connectionsToMake {
		// fmt.Println("\niter:", i+1)
		idx1 := cxn[0]
		idx2 := cxn[1]
		// point1 := points[idx1]
		// point2 := points[idx2]

		test1 := pointToCircuit[idx1]
		test2 := pointToCircuit[idx2]
		// fmt.Println("test1:", test1)
		// fmt.Println("test2:", test2)

		if test1 != test2 {
			// merge circuits
			circuits[test1].points = append(circuits[test1].points, circuits[test2].points...)
			circuits[test2].points = []point{}
			for key, value := range pointToCircuit {
				if value == test2 {
					pointToCircuit[key] = test1
				}
			}
		}

		// not efficient
		numCircuitsWithPoints := 0
		for _, circuit := range circuits {
			if len(circuit.points) > 0 {
				numCircuitsWithPoints++
			}
		}
		if numCircuitsWithPoints == 1 {
			point1 := points[idx1]
			point2 := points[idx2]

			fmt.Println(point1.x * point2.x)
			break
		}

		// fmt.Println("point1:", point1, "\npoint2:", point2)
		// fmt.Println(len(circuits), "circuits:")
		// for _, circuit := range circuits {
		// 	fmt.Println(circuit)
		// }
	}

	// fmt.Println(len(circuits))
	// for _, circuit := range circuits {
	// 	fmt.Println(circuit)
	// }

	// now i didn't clear the empty connections, but the question just asks for
	// the three longest circuits so it should be ok
	sort.Slice(circuits, func(i, j int) bool {
		return len(circuits[i].points) > len(circuits[j].points)
	})

	fmt.Println(len(circuits[0].points) * len(circuits[1].points) * len(circuits[2].points))
}

//	  +z |  / -x
//	     | /
//	-y   |/
//	-----*-----
//	    /|   +y
//	   / |
//
// +x /  | -z
//
// https://en.wikipedia.org/wiki/Cartesian_coordinate_system#Three_dimensions
func parsePoints(input []byte) ([]point, []circuit) {
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	points := make([]point, 0, len(lines))
	circuits := make([]circuit, 0, len(lines))

	for _, line := range lines {
		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			continue
		}

		x, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatal(err)
		}
		y, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}
		z, err := strconv.Atoi(parts[2])
		if err != nil {
			log.Fatal(err)
		}

		points = append(points, point{x: x, y: y, z: z})
		circuits = append(circuits, circuit{points: []point{{x: x, y: y, z: z}}})
	}

	return points, circuits
}

func distance(p1, p2 point) float64 {
	return math.Sqrt(math.Pow(float64(p1.x-p2.x), 2) +
		math.Pow(float64(p1.y-p2.y), 2) +
		math.Pow(float64(p1.z-p2.z), 2))
}

// i think the instructions imply that a connection attempt for already-
// connected points should no-op
func nShortestConnections(points []point, n int) [][]int {
	connectionsToMake := [][]int{}
	closestDistance := 0.0
	idx1 := 0
	idx2 := 1
	for range n {
		idx1, idx2, closestDistance = indicesOfClosestPoints(points, closestDistance)
		connectionsToMake = append(connectionsToMake, []int{idx1, idx2})
	}
	return connectionsToMake
}

func indicesOfClosestPoints(points []point, minDistance float64) (int, int, float64) {
	idx1 := 0
	idx2 := 1
	closestDistance := distance(points[idx1], points[idx2])

	for i := range points {
		for j := range points {
			if i == j {
				continue
			}

			p1 := points[i]
			p2 := points[j]

			dist := distance(p1, p2)
			if dist > minDistance && dist < closestDistance {
				closestDistance = dist
				idx1 = i
				idx2 = j
			}
		}
	}

	return idx1, idx2, closestDistance
}

// func areConnected(circuits []circuit, p1, p2 point) bool {
// 	for _, circuit := range circuits {
// 		if circuit.contains(p1) && circuit.contains(p2) {
// 			return true
// 		}
// 	}
// 	return false
// }

// func (c *circuit) contains(p point) bool {
// 	for _, point := range c.points {
// 		if point.x == p.x && point.y == p.y && point.z == p.z {
// 			return true
// 		}
// 	}
// 	return false
// }
