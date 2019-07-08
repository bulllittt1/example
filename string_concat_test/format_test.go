package format

import "testing"

const (
	f = "first"
	s = "second"
)

var result string

func BenchmarkSimple(b *testing.B) {
	var r string
	for n := 0; n < b.N; n++ {
		r = Simple(f, s)
	}
	result = r
}

func BenchmarkFormat(b *testing.B) {
	var r string
	for n := 0; n < b.N; n++ {
		r = Format(f, s)
	}
	result = r
}

func Test(t *testing.T) {
	simple := Simple(f, s)
	format := Format(f, s)
	if simple != format {
		t.Errorf("should be equal")
	}
}
