package slice_growth

import "testing"

var foos = []Foo{
	{"Rob"},
	{"Pike"},
	{"Ken"},
	{"Thompson"},
}

var result1 []Bar

func Benchmark_convert1(b *testing.B) {
	var r []Bar
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = convert1(foos)
	}
	result1 = r
}

var result2 []Bar

func Benchmark_convert2(b *testing.B) {
	var r []Bar
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = convert2(foos)
	}
	result2 = r

}
