package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {

	handler := func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("making GET request to 8888")

		resp, err := http.Get("http://127.0.0.1:8888")
		checkErr(err)

		defer resp.Body.Close()
		b, err := ioutil.ReadAll(resp.Body)
		checkErr(err)

		fmt.Printf("response from 8888: %v", string(b))

		fmt.Fprintf(w, "%s", b)
	}

	http.HandleFunc("/", handler)
	fmt.Println("listening at port 9999")
	log.Fatal(http.ListenAndServe(":9999", nil))

}
