# go的行为定义和实现

## 封装数据和行为

```go
type Employee struct {
Id string
Name string
Age int
}
```

## 实例创建以及初始化

```go
//方法1 
e := Employee{"1","magic",23}
//方法2
e1 := Employee{Name: "magic",Age: 23}
//方法3
e2 := new(Employee) //注意这里返回的引用/指针,相当于 e:= Employee{}
e2.Id = "2" //通过实例的指针访问成员不需要使用->
e2.Age = 23
e2.Name = "magic"
```

