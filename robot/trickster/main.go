package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/go-vgo/robotgo"
)

var (
	dur   = flag.Duration("d", time.Minute, "Timeout for mover")
	crazy = flag.Bool("crazy", false, "Random mouse moving and clicking")
)

func main() {
	fmt.Println("Started")
	flag.Parse()

	go func() {
		f10 := robotgo.AddEvent("f10")
		if f10 == 0 {
			fmt.Println("Canceled")
			os.Exit(0)
		}
	}()

	if *crazy {
		crazymover()
	} else {
		mover(*dur)
	}
}

func mover(d time.Duration) {
	fmt.Println("Timeout", *dur)
	var trigger bool
	for {
		if trigger {
			robotgo.ScrollMouse(10, "up")
			trigger = false
		} else {
			robotgo.ScrollMouse(10, "down")
			trigger = true
		}
		time.Sleep(d)
	}
}

func crazymover() {
	fmt.Println("Crazy mode. Have fun;)")
	time.Sleep(time.Second)
	fmt.Println("   3")
	time.Sleep(time.Second)
	fmt.Println("    2")
	time.Sleep(time.Second)
	fmt.Println("     1")
	time.Sleep(time.Second)

	x0, y0 := robotgo.GetScreenSize()

	for {
		x := rand.Intn(x0)
		y := rand.Intn(y0)

		robotgo.MoveMouseSmooth(2*x, y, 0.001, 0.001)
		robotgo.MouseClick("left", true)
	}
}
