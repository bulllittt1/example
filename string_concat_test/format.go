package format

import "fmt"

func Simple(a, b string) string {
	return a + " " + b
}

func Format(a, b string) string {
	return fmt.Sprintf("%s %s", a, b)
}
