// Go の _select_ を使うと、複数のチャネル操作を待つことができます。
// ゴルーチンとチャネルを `select` で扱えるのが、Go の強力な特長です。

package main

import "time"
import "fmt"

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}

//2つのチャンネルに対して selectをする例

//selectを利用することで、複数のチャネル操作を待つことができる。
//受信したものから、画面に表示される。
