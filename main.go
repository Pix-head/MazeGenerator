package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 30 //width with borders
	height = 30 //height with borders
)

var directions = []string{"left", "up", "right", "down"}

// StartPoint -
// h: current height position(Y)
// w: current width position(X)
type StartPoint struct {
	h int
	w int
}

func main() {
	var (
		entrance = 0
		exit     = 0
	)

	mazeMap := [height][width]int{}
	border := [width]int{}

	rand.Seed(time.Now().UnixNano())

	entrance = rand.Intn(height-3) + 1
	exit = rand.Intn(height-3) + 1

	border[0] = 3
	border[width-1] = 3

	for b := 1; b < width-1; b++ {
		border[b] = 3
	}

	for h := 1; h < height-1; h++ {
		for w := 0; w < width; w++ {
			if w == 0 || w == width-1 {
				mazeMap[h][w] = 3
			} else {
				mazeMap[h][w] = 0
			}
		}
	}

	mazeMap[entrance][0] = 2
	mazeMap[exit][width-1] = 2

	mazeMap[0] = border
	mazeMap[height-1] = border

	startPoint := StartPoint{
		h: entrance,
		w: 1,
	}

	mazeMap = createMaze(mazeMap, startPoint, exit)

	printMap(mazeMap)

	fmt.Printf("\nENTRANCE - X: 1 Y: %d \n", entrance+1)
	fmt.Printf("EXIT - X: %d Y: %d \n", width, exit+1)

}

func createMaze(mazeMap [height][width]int, startPoint StartPoint, exit int) [height][width]int {

	// STARTPOINT -
	// h: current height position(Y)
	// w: current width position(X)

	y := startPoint.h
	x := startPoint.w
	points := []StartPoint{}
	point := StartPoint{}
	missJump := 1
	isFinished := false

	mazeMap[y][x] = 1

	for missJump != (width * height * 10) {
		//Choose random direction
		direction := directions[rand.Intn(len(directions))]

		if missJump%10 == 0 {
			point = points[rand.Intn(len(points))]
			y = point.h
			x = point.w
		}

		//Change position
		if direction == "left" {
			if x-2 > 0 && mazeMap[y][x-2] == 0 {
				mazeMap[y][x-2] = 1
				mazeMap[y][x-1] = 1
				points = append(points, StartPoint{h: y, w: x - 2})
				x -= 2
			} else {
				missJump++
				continue
			}
		} else if direction == "up" {
			if y-2 > 0 && mazeMap[y-2][x] == 0 {
				mazeMap[y-2][x] = 1
				mazeMap[y-1][x] = 1
				points = append(points, StartPoint{h: y - 2, w: x})
				y -= 2
			} else {
				missJump++
				continue
			}
		} else if direction == "right" {
			if x+2 > 0 && x+2 < width-1 && mazeMap[y][x+2] == 0 {
				mazeMap[y][x+2] = 1
				mazeMap[y][x+1] = 1
				points = append(points, StartPoint{h: y, w: x + 2})
				x += 2
			} else {
				missJump++
				continue
			}
		} else if direction == "down" {
			if y+2 > 0 && y+2 < height-1 && mazeMap[y+2][x] == 0 {
				mazeMap[y+2][x] = 1
				mazeMap[y+1][x] = 1
				points = append(points, StartPoint{h: y + 2, w: x})
				y += 2
			} else {
				missJump++
				continue
			}
		}
	}

	if mazeMap[exit][width-2] == 0 {
		mazeMap[exit][width-2] = 1
	}

	i := 0

	//Connect maze with exit
	for !isFinished {
		if i > 1 {
			r := rand.Intn(2)
			if r == 0 && exit != 1 {
				exit--
			} else if r == 1 && exit != height-2 {
				exit++
			}
			i = 1
		}

		if exit == 1 {

			if mazeMap[exit+1][width-2-i] != 1 || mazeMap[exit][width-3-i] != 1 {
				mazeMap[exit][width-2-i] = 1
				i++
			} else {
				isFinished = true
				break
			}

		} else if exit == height-2 {

			if mazeMap[exit-1][width-2-i] != 1 || mazeMap[exit][width-3-i] != 1 {
				mazeMap[exit][width-2-i] = 1
				i++
			} else {
				isFinished = true
				break
			}

		} else {

			if mazeMap[exit-1][width-2-i] != 1 || mazeMap[exit+1][width-2-i] != 1 || mazeMap[exit][width-3-i] != 1 {
				mazeMap[exit][width-2-i] = 1
				i++
			} else {
				isFinished = true
				break
			}

		}
	}

	return mazeMap

}

func printMap(mazeMap [height][width]int) {
	fmt.Print("YX|")
	for w := 0; w < width; w++ {
		if (w+1)%10 == 0 {
			fmt.Print("0|")
		} else if w+1 > 10 {
			fmt.Printf("%d|", (w+1)%10)
		} else {
			fmt.Printf("%d|", w+1)
		}
	}
	fmt.Print("\n")
	for h := 0; h < height; h++ {
		if h+1 < 10 {
			fmt.Printf("%d |", h+1)
		} else if h+1 >= 10 {
			fmt.Printf("%d|", h+1)
		}

		for w := 0; w < width; w++ {
			if mazeMap[h][w] == 3 {
				fmt.Print("* ")
			} else if mazeMap[h][w] == 2 {
				fmt.Print("> ")
			} else if mazeMap[h][w] == 1 {
				fmt.Print("  ")
			} else {
				fmt.Print(mazeMap[h][w])
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}
