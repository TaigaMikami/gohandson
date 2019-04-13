// 将来のある時点や一定間隔で繰り返し Go コードを実行したい、
// と思うことがよくあります。Go の組み込み機能である
// _タイマー (timer)_ と _ティッカー (ticker)_ は、
// これら両方のタスクを容易にします。
// まず最初にタイマーを見て、次に [ティッカー](tickers)
// を見ていきましょう。

package main

import "time"
import "fmt"

func main() {
	start := time.Now()
	timer1 := time.NewTimer(5 * time.Second)
	<-timer1.C
	fmt.Println("It's time!")
	end := time.Now();
	fmt.Printf("%f秒\n",(end.Sub(start)).Seconds())
}

//将来のある時点や一定間隔で繰り返し、ある部分を実行したい際に利用する
//タイマーは待ち時間を指定すると、その時間にチャネルが処理を実施します。
//例では、5秒経過すると一定の処理を実施します。
