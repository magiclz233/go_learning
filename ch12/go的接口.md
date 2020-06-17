# Go的接口

go的接口和Java中的接口不一样,Java的接口不再重复讲解,说说go的接口

## Go的接口实现

接口定义:

接口定义用interface,实例定义用struct

```go
type Programmer interface {
    WriterHelloWorld() Code
}
```

接口实现:

```go
type GoProgrammer struct {
    
func(p *GoProgrammer) WriterHelloWorld() Code {
    return "fmt.Println(\"Hello World\")"
}
}
```

# 与其他编程语言的差异

1. 接口为非入侵性,实现不依赖于接口定义
2. 所有接口的定义可以包含在接口使用者包内

# 接口变量

```go
var prog Coder = &GoProgrammer{}


```

# 自定义类型

1. ``type IntConvertionFn func(n int) int``
2. ``type MyPoint int``

具体看代码