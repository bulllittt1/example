package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"
)

const (
	requests     = 32
	urlTemplate  = "https://vs2.coursehunters.net/otus-adr/lesson%d.mp4"
	nameTemplate = "lesson%d.mp4"
)

func main() {
	t := time.Now()
	fmt.Println("start")

	jobs := make(chan int, requests)
	results := make(chan struct{}, requests)

	cpus := runtime.NumCPU()
	fmt.Println("number of CPUs", cpus)

	for i := 1; i <= cpus && i <= requests; i++ {
		go worker(i, jobs, results)
	}

	for i := 1; i <= requests; i++ {
		jobs <- i
	}
	close(jobs)

	fmt.Println("waiting results")
	for i := 0; i < requests; i++ {
		<-results
	}
	fmt.Printf("finished with time: %v\n", time.Since(t))
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func downloadFile(filepath string, url string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp != nil && resp.StatusCode == http.StatusNotFound {
		return fmt.Errorf("%s not found at server", filepath)
	}

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		err = os.Remove(filepath)
		return err
	}

	return nil
}

func worker(id int, jobs <-chan int, results chan<- struct{}) {
	var counter int
	fmt.Printf("init worker №%d\n", id)
	for job := range jobs {
		fmt.Printf("worker №%d takes job №%d\n", id, job)

		name := fmt.Sprintf(nameTemplate, job)
		url := fmt.Sprintf(urlTemplate, job)

		if err := downloadFile(name, url); err != nil {
			log.Printf("worker №%d failed to complete job №%d: %v\n", id, job, err)
			goto finish
		}
		fmt.Printf("worker №%d finished job №%d\n", id, job)
		counter++
	finish:
		results <- struct{}{}
	}
	fmt.Printf("end of worker №%d with %d jobs done\n", id, counter)
}
