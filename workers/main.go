package main

import (
	"fmt"
	"net/http"
	"runtime"
)

const url = "https://golang.org/"

func main() {
	requests := 100

	jobs := make(chan struct{}, requests)
	results := make(chan struct{}, requests)

	cpus := runtime.NumCPU()
	fmt.Println("number of CPUs", cpus)

	for i := 1; i <= cpus; i++ {
		go worker(i, jobs, results)
	}

	for i := 0; i < requests; i++ {
		jobs <- struct{}{}
	}
	close(jobs)

	fmt.Println("waiting results")
	for i := 0; i < requests; i++ {
		<-results
	}
	fmt.Println("finished")

}

func worker(id int, jobs <-chan struct{}, results chan<- struct{}) {
	var counter int
	fmt.Printf("init worker № %d\n", id)
	for range jobs {
		_, _ = http.Get(url)
		fmt.Printf("worker № %d finished job \n", id)
		counter++
		results <- struct{}{}
	}
	fmt.Printf("end worker № %d with %d jobs finished\n", id, counter)
}
