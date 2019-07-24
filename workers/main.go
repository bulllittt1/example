package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"runtime"
	"time"
)

const url = "http://awsdeploy-app.fpw7ywc3xp.eu-central-1.elasticbeanstalk.com/hello"

func main() {
	requests := 10000
	fmt.Println("requests to do: ", requests)
	time.Sleep(3 * time.Second)

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
		t := rand.Intn(100)
		time.Sleep(time.Duration(t) * time.Millisecond)
		resp, err := http.Get(url)
		if err != nil {
			log.Println(err)
		}
		if resp != nil && resp.StatusCode != 200 {
			log.Println(resp.Status)
		}
		fmt.Printf("worker № %d finished job \n", id)
		counter++
		results <- struct{}{}
	}
	fmt.Printf("end worker № %d with %d jobs finished\n", id, counter)
}
