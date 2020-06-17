package extension

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Println("speak")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

//1: 在Dog中调用Pet的指针
type Dog struct {
	Pet
}

//调用内部成员变量方法 Speak()
func (d *Dog) Speak() {
	fmt.Println("magic")
}

func TestDog(t *testing.T) {
	dog := new(Dog)
	//调用的是Dog里面重写的Speak()输出为 magic
	dog.Speak()
	//这个里面调用了Pet里面的Speak()方法 所以输出为 speak host
	dog.SpeakTo("host")

}
