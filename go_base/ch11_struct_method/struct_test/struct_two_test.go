package struct_test

import (
	"encoding/json"
	"fmt"
	"math"
	"testing"
)

type location struct {
	lat, long float64
}

type world struct {
	radius float64
}

func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

type user struct {
	Name string
	body
}

type body struct {
	Height float64
	Weight float64
}

func (b body) size(h, w float64) int {
	return 1
}
func TestUser1(t *testing.T) {
	b1 := body{
		Height: 100,
		Weight: 100,
	}
	u1 := user{
		Name: "1",
		body: b1,
	}
	a, err := json.Marshal(u1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(a))

}
