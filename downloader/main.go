package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
)

const (
	requests     = 2
	pathTemplate = "https://vs2.coursehunters.net/otus-adr/lesson%d.mp4"
	nameTemplate = "lesson%d.mp4"
)

func main() {
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
	fmt.Println("finished")

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
	fmt.Printf("init worker № %d\n", id)
	for job := range jobs {
		fmt.Printf("worker № %d takes job № %d \n", id, job)

		path := fmt.Sprintf(nameTemplate, job)
		fileURL := fmt.Sprintf(pathTemplate, job)

		if err := downloadFile(path, fileURL); err != nil {
			log.Printf("job № %d: failed to download %s\n", id, path)
			return
		}
		fmt.Printf("worker № %d finished № job \n", id)
		counter++
		results <- struct{}{}
	}
	fmt.Printf("end worker № %d with %d jobs finished\n", id, counter)
}
