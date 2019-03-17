package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

func main() {
	db, err := sql.Open("sqlite3", "./accountbook.db" )
	if err != nil {
		fmt.Fprintln(os.Stderr, "エラー:", err)
		os.Exit(1)
	}

	ab := NewAccountBook(db)

	// テーブル作成(なければ)
	if err := ab.CreateTable(); err != nil {
		fmt.Fprintln(os.Stderr, "エラー:", err)
	}

LOOP:
	for {
		var mode int
		fmt.Println("[1]入力 [2]最新10件 [3]集計 [4]終了")
		fmt.Printf(">")
		fmt.Scan(&mode)

		switch mode {
		case 1: // 入力
			var n int
			fmt.Println("何件入力しますか>")
			fmt.Scan(&n)

			for i := 0; i < n; i++ {
				if err := ab.AddItem(inputItem()); err != nil {
					fmt.Fprintln(os.Stderr, "エラー:", err)
					break LOOP
				}
			}
		case 2: // 最新10件
			items, err := ab.GetItems(10)
			if err != nil {
				fmt.Fprintln(os.Stderr, "エラー:", err)
				break LOOP
			}
			showItems(items)
		case 3:
			summaries, err := ab.GetSummaries()
			if err != nil {
				fmt.Fprintln(os.Stderr, "エラー:", err)
				break LOOP
			}
			showSummary(summaries)
		case 4:
			fmt.Println("終了します")
			return
		}
	}
}

func inputItem() *Item {
	var item Item

	fmt.Print("品目>")
	fmt.Scan(&item.Category)

	fmt.Print("値段>")
	fmt.Scan(&item.Price)

	return &item
}

func showItems(items []*Item) {
	fmt.Println("=====")

	for _, item := range items {
		fmt.Printf("[%04d] %s:%d円\n", item.ID, item.Category, item.Price)
	}

	fmt.Println("=====")
}

func showSummary(summaries []*Summary) {
	fmt.Println("=====")
	fmt.Printf("品目\t個数\t合計\t平均\n")

	for _, s := range summaries {
		fmt.Printf("%s\t%d\t%d円\t%.2f円\n", s.Category, s.Count, s.Sum, s.Avg())
	}
	fmt.Println("=====")
}
