package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	pic = make([][]uint8, dy) 

	for i := range dy{

		pic[i] = make([]uint8, dx)
		for j :=range dx{
			pic[i][j] = uint8((i+j)/2)
		}
	}
	return pic
}

func main() {
	pic.Show(Pic)
}
