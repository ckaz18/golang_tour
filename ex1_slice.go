package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	temp := make([][]uint8, dy)
	for x := 0; x < dy; x++ {
		temp[x] = make([]uint8, dx)
		for y := 0; y < dx; y++ {
			 temp[x][y] = uint8((x^y)/3)
		}
	}
	return temp
}

func main() {
	pic.Show(Pic)
}