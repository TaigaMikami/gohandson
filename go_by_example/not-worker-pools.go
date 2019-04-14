package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs int, results int) {
	for j:= range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job")
		results <- j * 2
	}
}

func main() {
	start := time.Now()

	jobs := make([]int, 100)
	results := make([]int, 100)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs = j
	}
	close(jobs)

	for a := 1; a <= 5; a++ {
		<-results
	}

	end := time.Now();
	fmt.Printf("%f秒\n",(end.Sub(start)).Seconds())
}