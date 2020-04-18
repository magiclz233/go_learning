# string

## 与其他主要编程语言的差异

1. string 是数据类型,不是引用或指针类型
2. string 是只读的byte slice, len 函数返回的是它所包含的byte数
3. string 的 byte 数组可以存放任何数据

##Unicode UTF8

1. Unicode 是一种字符集(code point)
2. UTF8是unicode的存储实现 (转换为字节序列的规则)

以字符 '中' 为例

字符                         "中"

Unicode                     0x4E2D

UTF-8                       0xE4B8AD

string/[]byte               [0xE4,0xB8,0xAD]

## 常用字符串函数

1. strings包
2. strconv 包