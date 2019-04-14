// ここでは、ゴルーチンとチャネルを使って、
// _ワーカープール (worker pool)_ を実装する例を見ていきます。

package main

import "fmt"
import "time"

// これは、複数インスタンスが並行実行されるワーカーです。
// これらのワーカーは、`jobs` チャネルからタスクを受信し、
// 結果を `results` チャネルに送信します。
// 実行コストの高いジョブをシミュレートするため、
// 各タスクは 1 秒スリープします。
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	start := time.Now()

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= 5; a++ {
		<-results
	}

	end := time.Now();
	fmt.Printf("%f秒\n",(end.Sub(start)).Seconds())
}

//workerは3つ分並列実行されます。
//5つのジョブが送信されるため5秒分のタスクを実行します。
//しかし、worker関数のtime.Sleepで5秒分のタスクを実行する似にかかわらず、2秒しかかかりません。
//これは、3つのworkerが並列実行しているためです。