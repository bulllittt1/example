package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/go-vgo/robotgo"
)

// second version adds key pressing
const version = 2

var (
	d     = flag.Bool("d", false, "Use default strategy")
	dur   = flag.Duration("duration", time.Minute, "Timeout for mover")
	crazy = flag.Bool("crazy", false, "Random mouse moving and clicking")
)

var defaultDurations = []time.Duration{9, 9, 11}

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	fmt.Println("Start version #", version)
	flag.Parse()

	// Monitor f10 keytap to shutdown
	go func() {
		f10 := robotgo.AddEvent("f10")
		if f10 == true {
			fmt.Println("Canceled")
			os.Exit(0)
		}
	}()

	switch {
	case *d:
		fmt.Println("Default settings")
		defaultStrategy()
	case *crazy:
		crazymover()
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

func defaultStrategy() {
	var (
		up       bool
		distance int
		index    int
	)
	for {
		random := rand.Intn(9) + 1
		// scroll mouse
		distance = random
		if up {
			robotgo.ScrollMouse(distance, "up")
			up = false
		} else {
			robotgo.ScrollMouse(distance, "down")
			up = true
		}
		// key tap
		var j int
		for i := 0; i <= random*100; i++ {
			robotgo.KeyTap("alt")
			j++
		}
		// random sleep interval
		index = rand.Intn(3)
		time.Sleep(defaultDurations[index] * time.Minute)
	}
}
