package main

import (
	"fmt"
	"image"
)

func main() {
	rect0 := image.Rectangle{Min: image.Point{0, 0}, Max: image.Point{100, 100}}
	rect1 := image.Rectangle{Min: image.Point{50, 50}, Max: image.Point{150, 150}}

	fmt.Println(rect0.Union(rect1))

	newRect := rect0
	if newRect.Min.X >= rect1.Min.X {
		newRect.Min.X = rect1.Min.X
	}
	if newRect.Min.Y >= rect1.Min.Y {
		newRect.Min.Y = rect1.Min.Y
	}
	if newRect.Max.X <= rect1.Max.X {
		newRect.Max.X = rect1.Max.X
	}
	if newRect.Max.Y <= rect1.Max.Y {
		newRect.Max.Y = rect1.Max.Y
	}
	fmt.Println(newRect)

}
