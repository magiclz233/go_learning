# Go Learning

一个用来学习Go语言的仓库,包含了:

* [go_base](https://github.com/magiclz233/go_learning/tree/master/go_base) Go的基础语法部分
* [gin]()框架学习
* learning

## base

本项目使用的go版本为1.14, 建议使用1.13以上版本, 编辑器使用的是GoLand, 其他idea,vscode也可以


## go_base

这个模块从go的Hello World到常量,变量,函数,方法,数组,切片,常用库,协程,net/http等等,

都是我在个人学习go语言过程中所练习的代码.

后续也会一直补充!

## go_gin

对gin框架的学习

因为go_gin模块的代码是用go mod创建的

**学习之前请最好将go版本升级到1.13及以上**

然后打开控制台输入:

```cassandraql
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.io,direct
```

控制台切换到../go_gin目录,执行 go mod download 命令,下载所需的jar包

## end

学海无涯, 共勉