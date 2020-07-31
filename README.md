# Go Learning

一个用来学习Go语言的仓库,包含了:

* [go_base](https://github.com/magiclz233/go_learning/tree/master/go_base) Go的基础语法部分
* [go_web](https://github.com/magiclz233/go_learning/tree/master/go_web) Go Web以及web框架学习
    * [echo](https://github.com/magiclz233/go_learning/tree/master/go_web/go_echo) echo框架学习
    * [gin](https://github.com/magiclz233/go_learning/tree/master/go_web/go_gin) gin框架学习
    
* learning

## base

本项目使用的go版本为1.14, 建议使用1.13以上版本, 编辑器使用的是GoLand, 其他idea,vscode也可以

其中 go_base基础模块使用的GOPATH形式练习, go_web使用go mod形式

## go_base

这个模块从go的Hello World到常量,变量,函数,方法,数组,切片,常用库,协程,net/http等等,

都是我在个人学习go语言过程中所练习的代码.

后续也会一直补充!

## go_web

对Golang来开发web应用的学习,主要学习gin,echo等框架的使用.

### go_gin

对gin框架的学习

因为go_gin模块的代码是用go mod创建的

**学习之前请最好将go版本升级到1.13及以上**

然后打开控制台输入:

```cassandraql
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.io,direct
```

控制台切换到../go_gin目录,执行 go mod download/go mod tidy 命令,下载所需的jar包

## go_echo

对 echo web框架的学习

基本上和gin差不多,看一个就可以快速了解第二个, 基础配置也相同

## end

学海无涯, 共勉