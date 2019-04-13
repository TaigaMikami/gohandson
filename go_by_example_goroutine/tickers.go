// [タイマー](timers) は、未来に一度だけ何かしたいときに使いますが、
// _ティッカー (tickers)_ は定期的に何かしたいときに使います。
// ここでは、停止するまで定期的に動作するティッカーの例を見ます。

package main

import "time"
import "fmt"

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}

//ティッカーは、ticker.Stop()により停止するとそのチャネルから値を受信しなくなる
