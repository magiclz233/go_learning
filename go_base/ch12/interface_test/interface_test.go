package interface_test

import (
	"fmt"
	"math"
	"testing"
)

type Programmer interface {
	WriterHelloWorld() string
}

type GoProgrammer struct {
}

func (g *GoProgrammer) WriterHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriterHelloWorld())
}

type PraInterface interface {
	WriteHello() string
	AddNum(i, j int) int
	IsPositiveInteger(i int) bool
}

type GoPraInterface struct {
}

func (g *GoPraInterface) AddNum(i, j int) int {
	return i + j
}

func (g *GoPraInterface) IsPositiveInteger(i int) bool {
	return i > 0
}

func (g *GoPraInterface) WriteHello() string {
	return "hello"
}

func TestInterface1(t *testing.T) {
	var pra PraInterface
	pra = new(GoPraInterface)
	t.Log(pra.WriteHello())
	t.Log(pra.AddNum(2, 3))
	t.Log(pra.IsPositiveInteger(1))
	t.Log(pra.IsPositiveInteger(-1))
}

type Shaper interface {
	Area() float64
	Volume() float64
}

type Square struct {
	side float64
}

func (s *Square) Area() float64 {
	return s.side * s.side
}

func (s *Square) Volume() float64 {
	return math.Pow(s.side, 3)
}

func (s Square) bigArea(i int) float64 {
	return math.Pow(float64(i)*s.side, 2)
}

func TestSquare(t *testing.T) {
	sq1 := Square{side: 5}
	sq2 := new(Square)
	sq2.side = 6

	var area Shaper
	area = &sq1
	var area1 Shaper
	area1 = sq2
	fmt.Printf("%f\n", area.Area())
	fmt.Printf("%f\n", area.Volume())
	fmt.Printf("%f\n", sq1.bigArea(2))

	fmt.Printf("%f\n", area1.Area())
	fmt.Printf("%f\n", area1.Volume())
	fmt.Printf("%f\n", sq2.Volume())

}

type Simpler interface {
	Get() string
	Set(s string)
}

type Simple struct {
	Name string
}

func (s *Simple) Get() string {
	return s.Name
}

func (s *Simple) Set(str string) {
	s.Name = str
}

func TestSimple(t *testing.T) {
	var str = new(Simple)
	fmt.Println(str.Get())
	str.Set("magic")
	fmt.Println(str.Get())
}

// 一个借口多个实例 struct是接口的实现类
type Sayer interface {
	say()
}

type dog struct {
}

type cat struct {
}

func (d *dog) say() {
	fmt.Println("狗叫")
}

func (c *cat) say() {
	fmt.Println("猫叫")
}

func TestSayer(t *testing.T) {
	var x Sayer
	a := dog{}
	b := cat{}
	x = &a
	x.say()
	x = &b
	x.say()
}
