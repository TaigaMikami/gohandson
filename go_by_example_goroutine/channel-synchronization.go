// Channel Synchronization
//Channelsのところで、 `msg := <-messages` の処理を待ってくれるということを話しましたが、これを利用することで、
//なにかしらのレシーバを用意しなくても、同期を待ってくれるように書くことだできます。
//`<-done` がその部分ですね。
// ちなみにこれがないと、プログラムはworker関数が始まる前に終了します。

package main
import "fmt"
import "time"

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// 完了したことを通知するために値を送信します。
	done <- true
}

func main() {

	// 通知用のチャネルを渡して、`worker` ゴルーチンを開始します。
	done := make(chan bool, 1)
	go worker(done)

	// チャネルへの完了通知を受信するまでブロックします。
	<-done
}
