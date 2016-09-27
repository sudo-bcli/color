# GoLang 彩色输出

用于Mac和Linux Shell的GoLang彩色文字输出, [English Version](README.md)

## 效果展示
![img](showcase.jpg)

## 安装

1. __确保__你已经读过  [GoLang官方文件结构](https://golang.org/doc/code.html)
2. `go get github.com/bclicn/color`
3. 在你的脚本中,`import "github.com/bclicn/color"` 然后调用 `color.Test()`查看全部效果
4. 使用`fmt.Println("Hello" + color.Red("World"))`进行快速使用
4. 其他的API调用示例可以去`color-test.go`中查看

## 快速(无脑)使用

1. 下载`color.go`
2. 将其第一行的`package color`改为`package main`
3. 把你的脚本放在同一个文件夹下,直接调用如`fmt.Println(Red("I'm red !!!"))`
4. 运行`go run color.go yourScript.go`
5. 你也可以直接`go build color.go yourScript.go`

## MIT License

Beichen Li, 2016-9-27


 