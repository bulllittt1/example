package main

import (
	"fmt"
	"net/url"
	"time"
)

const week uint = 7 * 24 * 60 * 60

func transformURI(uri string) string {
	prefix := "https://"

	URL, err := url.Parse(uri)
	fmt.Printf("Parse :: %#v\n", URL)
	if err != nil {
		fmt.Printf("Cant parse URI. %v\n", err)
		return uri
	}

	if URL.Scheme != "" {
		if URL.Host == "wrike.com" {
			return fmt.Sprintf("%v://www.%v", URL.Scheme, URL.Host)
		}
		return uri
	}

	return prefix + uri
}

var (
	urla    = "https://gitlab.qarea.org"
	baseURL = "https://gitlab.com"
)

func main() {
	// newUrla := transformURI(urla)
	// fmt.Println(newUrla)
	// if strings.Contains(urla, baseURL) {
	// 	fmt.Println("success", urla)
	// }
	// num := uint(rand.Intn(int(week)))
	// fmt.Printf("value: %v, type: %T\n", num, num)

	u, _ := url.Parse(baseURL)
	fmt.Println(u.String())

	yearNow, monthNow, dayNow := time.Unix(int64(45), 0).Date()
	fmt.Println(yearNow, monthNow, dayNow)

}
