package main

import (
	"fmt"
)

// Channels
func main() {
	messages := make(chan string)
	go func() { messages <- "Hello" }()

	msg := <-messages
	fmt.Println(msg)
}

// channel <- 構文で、チャネルへ値を 送信 します。
// <-channel 構文で、チャネルから値を 受信 します
// つまり、上の例では、messageというパイプを使って、無名関数からmsgへ"ping"を渡している。
// "ping"という値がmessageトンネルを通ってmsgへ届く

// Goはデフォルトで、送る側と受ける側が準備できるまで、 送受信はブロックされる。
// このため、同期処理的なものを書かなくても、"ping"がmsgに渡されるまで、待ってくれる






