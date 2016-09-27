# GoLang Colorized Output

GoLang colorized output library for Mac & Linux Shell, [中文版说明](README_CN.md)

## Showcase
![img](showcase.jpg)

## Install

1. __Make sure__ your've read [GoLang Official File Structure](https://golang.org/doc/code.html)
2. `go get github.com/bclicn/color`
3. In your script, `import "github.com/bclicn/color"` then call `color.Test()` for complete sample
4. Use `fmt.Println("Hello" + color.Red("World"))` for quick result
4. Other API examples are available in `color-test.go`

## Quick (brainless) Usage

1. Download `color.go`
2. Modify its `package color` to `package main`
3. Put your script under the same directory and add `fmt.Println(Red("I'm red !!!"))`
4. Run `go run color.go yourScript.go`
5. Or build by `go build color.go yourScript.go`

## MIT License

Beichen Li, 2016-9-27


 