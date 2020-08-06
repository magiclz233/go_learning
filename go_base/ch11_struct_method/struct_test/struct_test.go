package struct_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

// 结构体
type person struct {
	name string
	age  int8
	sex  string
}

// 构造函数(自己实现,go中没有)
func newPerson(name, sex string, age int8) *person {
	return &person{
		name: name,
		age:  age,
		sex:  sex,
	}
}

//调用构造函数
func TestStruct(t *testing.T) {
	p := newPerson("magic", "男", 23)
	fmt.Printf("%#v\n", p)
}

func TestStruct1(t *testing.T) {
	var p1 person
	fmt.Printf("%T\n", p1)
	fmt.Printf("p4 = %#v\n", p1)

	p2 := &person{}
	fmt.Printf("%T\n", p2)
	fmt.Printf("p2 = %#v\n", p2)

	// p2是获取的person结构体的地址  p2.age = (*p2).age go做了封装而已
	p2.age = 23
	p2.name = "magic"
	p2.sex = "男"

	fmt.Printf("%T\n", p2)
	fmt.Printf("p2 = %#v\n", p2)

	p3 := person{
		name: "root",
		age:  22,
		sex:  "女",
	}
	fmt.Printf("p3 = %#v\n", p3)

	p4 := &person{
		name: "sys",
		age:  21,
		sex:  "男",
	}

	fmt.Printf("p4 = %#v\n", p4)
}

// 创建一个key为name,value为person的map
func TestMapStruct(t *testing.T) {
	m := make(map[string]*person)
	users := []person{
		{name: "magic", age: 23, sex: "男"},
		{name: "root", age: 22, sex: "女"},
		{name: "sys", age: 21, sex: "男"},
	}

	for _, user := range users {
		fmt.Println(user)
		m[user.name] = &user
	}

	for k, v := range m {
		fmt.Println(k, "--->", v.name)
	}
}

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

func TestStruct2(t *testing.T) {
	var emp Employee
	emp.Salary -= 5000
	fmt.Printf("%#v\n", emp)
	position := &emp.Position
	*position = "magic" + *position
	fmt.Printf("%#v\n", emp)

}

//结构体 与json序列化
type Student struct {
	ID     int
	Name   string
	Gender string
}

type Class struct {
	Title    string
	Students []*Student
}

func TestStructJson(t *testing.T) {
	c := &Class{
		Title:    "101",
		Students: make([]*Student, 0, 200),
	}

	for i := 0; i < 10; i++ {
		stu := &Student{
			ID:     i,
			Name:   fmt.Sprintf("stu%03d", i),
			Gender: "男",
		}
		c.Students = append(c.Students, stu)
	}

	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal fail")
		return
	}
	fmt.Printf("json: %s\n", data)

	str := `{"Title":"101","Students":[{"ID":0,"Name":"stu000","Gender":"男"},{"ID":1,"Name":"stu001","Gender":"男"},{"ID":2,"Name":"stu002","Gender":"男"},{"ID":3,"Name":"stu003","Gender":"男"},{"ID":4,"Name":"stu004","Gender":"男"},{"ID":5,"Name":"stu005","Gender":"男"},{"ID":6,"Name":"stu006","Gender":"男"},{"ID":7,"Name":"stu007","Gender":"男"},{"ID":8,"Name":"stu008","Gender":"男"},{"ID":9,"Name":"stu009","Gender":"男"}]}`

	c1 := &Class{}
	err = json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println("json unmarshal fail")
		return
	}
	fmt.Printf("%#v\n", c1)
}

// Tag 结构体标签  可以在json序列化时改变名字 格式 跟在字段后 `json:"id"`

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestStructJson2(t *testing.T) {
	s1 := User{
		ID:   1,
		Name: "magic",
		Age:  23,
	}

	userJson, err := json.Marshal(s1)
	if err != nil {
		fmt.Println("json marshal fail")
		return
	}
	fmt.Printf("json user : %s\n", userJson)
}

func demo(c []User) {
	c[1].Age = 22
}
func TestSliceStruct(t *testing.T) {
	var c []User
	c = []User{
		{
			ID:   1,
			Name: "magic",
			Age:  12,
		},
		{
			ID:   2,
			Name: "root",
			Age:  11,
		},
	}
	fmt.Println(c) // [{1 magic 12} {2 root 11}]
	// 切片是引用传递  会改变值
	demo(c)
	fmt.Println(c) // [{1 magic 12} {2 root 22}]
}
