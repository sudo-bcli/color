package color

import (
	"strconv"
)

// github.com/bclicn/color
// colorized output for Mac & Linux terminal
// version: 1.0.0
// author:  Beichen Li, bclicn@gmail.com, relidin@126.com, 2016-9-23
// see:     http://misc.flogisoft.com/bash/tip_colors_and_formatting

const (
	// common
	reset 		= "\033[0m" 	// auto reset the rest of text to default color
	normal		= 0
	bold 		= 1 		// change this value if you want Ber text
	// special
	dim		= 2
	underline 	= 4
	blink		= 5
	reverse		= 7
	hidden		= 8
	// color
	// default	= 39
	black 		= 30
	red		= 31
	green		= 32
	yellow		= 33
	blue		= 34
	purple		= 35	// purple = magenta
	cyan		= 36
	lightGray 	= 37
	darkGray 	= 90
	lightRed 	= 91
	lightGreen 	= 92
	lightYellow 	= 93
	lightBlue	= 94
	lightPurple 	= 95
	lightCyan	= 96
	white		= 97
)


func Render (colorCode int, fontSize int, content string) string{
	return "\033[" + strconv.Itoa(fontSize) + ";" + strconv.Itoa(colorCode) + "m" + content + reset
}

func Black (txt string) string{
	return Render(black, normal, txt)
}

func Red (txt string) string {
	return Render(red, normal, txt)
}

func Green (txt string) string{
	return Render(green, normal, txt)
}

func Yellow (txt string) string{
	return Render(yellow, normal, txt)
}

func Blue (txt string) string{
	return Render(blue, normal, txt)
}

func Purple (txt string) string{
	return Render(purple, normal, txt)
}

func Cyan (txt string) string{
	return Render(cyan, normal, txt)
}

func LightGray (txt string) string{
	return Render(lightGray, normal, txt)
}

func DarkGray (txt string) string{
	return Render(darkGray, normal, txt)
}

func LightRed (txt string) string{
	return Render(lightRed, normal, txt)
}

func LightGreen (txt string) string{
	return Render(lightGreen, normal, txt)
}

func LightYellow (txt string) string{
	return Render(lightYellow, normal, txt)
}

func LightBlue (txt string) string{
	return Render(lightBlue, normal, txt)
}

func LightPurple (txt string) string{
	return Render(lightPurple, normal, txt)
}

func LightCyan (txt string) string{
	return Render(lightCyan, normal, txt)
}

func White (txt string) string{
	return Render(white, normal, txt)
}

func BBlack (txt string) string{
	return Render(black, bold, txt)
}

func BRed (txt string) string {
	return Render(red, bold, txt)
}

func BGreen (txt string) string{
	return Render(green, bold, txt)
}

func BYellow (txt string) string{
	return Render(yellow, bold, txt)
}

func BBlue (txt string) string{
	return Render(blue, bold, txt)
}

func BPurple (txt string) string{
	return Render(purple,  bold, txt)
}

func BCyan (txt string) string{
	return Render(cyan, bold, txt)
}

func BLightGray (txt string) string{
	return Render(lightGray, bold, txt)
}

func BDarkGray (txt string) string{
	return Render(darkGray, bold, txt)
}

func BLightRed (txt string) string{
	return Render(lightRed, bold, txt)
}

func BLightGreen (txt string) string{
	return Render(lightGreen, bold, txt)
}

func BLightYellow (txt string) string{
	return Render(lightYellow, bold, txt)
}

func BLightBlue (txt string) string{
	return Render(lightBlue,  bold, txt)
}

func BLightPurple (txt string) string{
	return Render(lightPurple,  bold, txt)
}

func BLightCyan (txt string) string{
	return Render(lightCyan,  bold, txt)
}

func BWhite (txt string) string{
	return Render(white, bold, txt)
}

func GBlack (txt string) string{
	return Render(black + 1, normal, txt)
}

func GRed (txt string) string {
	return Render(red + 1, normal, txt)
}

func GGreen (txt string) string{
	return Render(green + 1, normal, txt)
}

func GYellow (txt string) string{
	return Render(yellow + 1, normal, txt)
}

func GBlue (txt string) string{
	return Render(blue + 1, normal, txt)
}

func GPurple (txt string) string{
	return Render(purple + 1, normal, txt)
}

func GCyan (txt string) string{
	return Render(cyan + 1, normal, txt)
}

func GLightGray (txt string) string{
	return Render(lightGray + 1, normal, txt)
}

func GDarkGray (txt string) string{
	return Render(darkGray + 1, normal, txt)
}

func GLightRed (txt string) string{
	return Render(lightRed + 1, normal, txt)
}

func GLightGreen (txt string) string{
	return Render(lightGreen + 1, normal, txt)
}

func GLightYellow (txt string) string{
	return Render(lightYellow + 1, normal, txt)
}

func GLightBlue (txt string) string{
	return Render(lightBlue + 1, normal, txt)
}

func GLightPurple (txt string) string{
	return Render(lightPurple + 1, normal, txt)
}

func GLightCyan (txt string) string{
	return Render(lightCyan + 1, normal, txt)
}

func GWhite (txt string) string{
	return Render(white + 1, normal, txt)
}

func Bold (txt string) string{
	return Render(bold, normal, txt)
}

func Dim (txt string) string{
	return Render(dim, normal,  txt)
}

func Underline (txt string) string{
	return Render(underline, 0 , txt)
}

func Blink (txt string) string{
	return Render(blink, normal,  txt)
}

func Invert (txt string) string{
	return Render(reverse, normal,  txt)
}

func Hide (txt string) string{
	return Render(hidden, normal,  txt)
}