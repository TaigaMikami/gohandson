// Channel Buffering
package main
import "fmt"

func main() {
	messages := make(chan string, 2)

	messages <- "Hello"
	messages <- "Workd"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

//バッファリングされたチャネルは、対応する受信側がいなくても決められた量までなら 値を送信することができる
//make(chan string, 2)によって2妻でバッファリングするチャネルを作っている。