package main

import (
	"github.com/go-vgo/robotgo"
)

func main() {
	// bitmap := robotgo.CaptureScreen(10, 20, 30, 40)
	// // use `defer robotgo.FreeBitmap(bit)` to free the bitmap
	// defer robotgo.FreeBitmap(bitmap)
	//
	// fmt.Println("...", bitmap)

	// fx, fy := robotgo.FindPic("./button.png")
	// fmt.Println(fx, fy)
	//
	// robotgo.MoveMouse(fx+10, fy+10)
	//
	// robotgo.MouseClick("left", true)

	robotgo.ActiveName("chrome")
	robotgo.GetTi
}
