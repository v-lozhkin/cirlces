package cirlces

import (
	"fmt"
	"math"
)

type Circle struct {
	MyRadius int
}

func (c Circle) Square() float64 {
	return math.Pi * float64(c.MyRadius*c.MyRadius)
}

func (c Circle) String() string {
	return fmt.Sprintf("Круг с радиусом %d и площадью %.2f", c.Radius, c.Square())
}

func (c Circle) Radius() int {
	return c.MyRadius
}

func (c Circle) Diameter() int {
	return c.MyRadius * 2
}

func (c Circle) Point() (int, int) {
	return 0, 0
}
