package slice_growth

type Foo struct {
	First string
}

type Bar struct {
	Second string
}

func FooToBar(f Foo) Bar {
	return Bar{f.First}
}

func convert1(foos []Foo) []Bar {
	bars := make([]Bar, len(foos))
	for i, foo := range foos {
		bars[i] = FooToBar(foo)
	}
	return bars
}

func convert2(foos []Foo) []Bar {
	bars := make([]Bar, 0, len(foos))
	for _, foo := range foos {
		bars = append(bars, FooToBar(foo))
	}
	return bars
}
