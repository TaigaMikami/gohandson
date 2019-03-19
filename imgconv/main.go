package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}

func run() error {
	if len(os.Args) < 3 {
		return fmt.Errorf("画像ファイルを指定してください")
	}

	return convert(os.Args[2], os.Args[1])
}

func convert(dst, src string) error {
	sf, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("画像ファイルが開けませんでした。%s", src)
	}
	defer sf.Close()

	df, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("画像ファイルを書き出せませんでした%s", dst)
	}
	defer df.Close()

	img, _, err := image.Decode(sf)
	if err != nil {
		return err
	}

	switch strings.ToLower(filepath.Ext(dst)) {
	case ".png":
		err 	= png.Encode(df, img)
	case ".jpg":
		err = jpeg.Encode(df, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
	}

	if err != nil {
		return fmt.Errorf("画像ファイルを書き出せませんでした。%s", dst)
	}

	return nil
}

// STEP03
//func run() error {
//	if len(os.Args) < 3 {
//		return fmt.Errorf("引数が足りません")
//	}
//
//	src, dst := os.Args[1], os.Args[2]
//
//	sf, err := os.Open(src)
//	if err != nil {
//		return fmt.Errorf("ファイルが開けませんでした。%s", src)
//	}
//
//	defer sf.Close()
//
//	df, err := os.Create(dst)
//	if err != nil {
//		return fmt.Errorf("ファイルを書き出せませんでした。%s", dst)
//	}
//
//	defer df.Close()
//
//	scanner := bufio.NewScanner(sf)
//
//	for i := 1; scanner.Scan(); i++ {
//		fmt.Fprintf(df, "%d:%s\n", i, scanner.Text())
//	}
//
//	return scanner.Err()
//}

// STEP02
//func run() error {
//	if len(os.Args) < 3 {
//		return fmt.Errorf("引数が足りません。")
//	}
//
//	fmt.Println(os.Args[0])
//	fmt.Println(os.Args[1])
//	fmt.Println(os.Args[2])
//
//	return nil
//}
//

// STEP01
//func main() {
//	fmt.Println("Hello go")
//}