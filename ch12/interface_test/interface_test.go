package interface_test

import "testing"

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
