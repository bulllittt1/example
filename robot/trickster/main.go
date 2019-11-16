package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/go-vgo/robotgo"
)

const version = 1

var (
	crazy = flag.Bool("crazy", false, "Random mouse moving and clicking")
	d     = flag.Bool("d", false, "Use default settings")
	dur   = flag.Duration("duration", time.Minute, "Timeout for mover")
)

var defaultDurations = []time.Duration{9, 9, 11}

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	fmt.Println("Started version:", version)
	flag.Parse()

	go func() {
		f10 := robotgo.AddEvent("f10")
		if f10 == 0 {
			fmt.Println("Canceled")
			os.Exit(0)
		}
	}()

	switch {
	case *crazy:
		fmt.Println("Crazy mode. Have fun;)")
		time.Sleep(time.Second)
		fmt.Println("   3")
		time.Sleep(time.Second)
		fmt.Println("    2")
		time.Sleep(time.Second)
		fmt.Println("     1")
		time.Sleep(time.Second)
		crazymover()
	case *d:
		fmt.Println("Default settings")
		defaultScroller()
	default:
		fmt.Println("Timeout", *dur)
		scroller(*dur)
	}

}

func scroller(d time.Duration) {
	var (
		up       bool
		distance int
	)
	for {
		distance = rand.Intn(10)
		if up {
			robotgo.ScrollMouse(distance, "up")
			up = false
		} else {
			robotgo.ScrollMouse(distance, "down")
			up = true
		}
		time.Sleep(d)
	}
}

func defaultScroller() {
	var (
		up       bool
		distance int
		i        int
	)
	for {
		distance = rand.Intn(10)
		if up {
			robotgo.ScrollMouse(distance, "up")
			up = false
		} else {
			robotgo.ScrollMouse(distance, "down")
			up = true
		}
		i = rand.Intn(3)
		time.Sleep(defaultDurations[i] * time.Minute)
	}
}

func crazymover() {
	x0, y0 := robotgo.GetScreenSize()
	for {
		x := rand.Intn(x0)
		y := rand.Intn(y0)
		robotgo.MoveMouseSmooth(2*x, y, 0.001, 0.001)
		robotgo.MouseClick("left", true)
	}
}
