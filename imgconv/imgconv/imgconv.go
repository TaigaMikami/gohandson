package imgconv

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"strconv"
	"strings"
	"unicode"
)

var (
	// ErrInvalidSize は、指定したサイズが不正だった場合のエラーです。
	ErrInvalidSize = fmt.Errorf("指定したサイズの形式が不正です。")
	// ErrInvalidBounds は、指定した領域の形式が不正だった場合のエラーです。
	ErrInvalidBounds = fmt.Errorf("指定した領域の形式が不正です。")
	// ErrUnknownUnit は、想定外の不正な単位だった場合のエラーです。
	ErrUnknownUnit = fmt.Errorf("不正な単位です。")
)

type Image struct {
	image.Image
}

func parseRelSize(base int, s string) (int, error) {
	i := strings.IndexFunc(s, func(c rune) bool {
		return !unicode.IsNumber(c)
	})

	if i < 0 {
		return strconv.Atoi(s)
	}

	v, err := strconv.Atoi(s[:i])
	if err != nil {
		return 0, ErrInvalidSize
	}

	switch s[i:] {
	case "%":
		return int(float64(base) * float64(v) / 100), nil
	case "px":
		return v, nil
	default:
		return 0, ErrUnknownUnit
	}
}

func (img *Image) parseSize(s string) (sz image.Point, err error) {
	sp := strings.Split(s, "x")
	if len(sp) <= 0 || len(sp) > 2 {
		err = ErrInvalidSize
		return
	}

	sz.X, err = parseRelSize(img.Bounds().Max.X, sp[0])
	if err != nil {
		err = ErrInvalidSize
		return
	}

	if len(sp) == 1 {
		sz.Y = sz.X
	} else {
		sz.Y, err = parseRelSize(img.Bounds().Max.Y, sp[1])
	}

	return
}

func (img *Image) parseBounds(s string) (r image.Rectangle, err error) {
	sp := strings.Split(s, "+")
	if len(sp) <= 0 || len(sp) > 3 {
		err = ErrInvalidBounds
		return
	}

	r.Max, err = img.parseSize(sp[0])
	if err != nil || len(sp) == 1 {
		return
	}

	var p image.Point
	p.X, err = parseRelSize(img.Bounds().Max.X, sp[1])
	if err != nil {
		return
	}

	if len(sp) == 3 {
		p.Y, err = parseRelSize(img.Bounds().Max.Y, sp[2])
	}

	r = r.Add(p)

	return
}

func newDrawImage(r image.Rectangle, m color.Model) draw.Image {
	// TODO: 各カラーモデルごとに画像を初期化し返す。
	// なお、指定されたカラーモデルがimage/colorパッケージに定義されていない場合は、
	// RGBAの画像を作って返す。
	switch m {
	case color.RGBA64Model:
		return image.NewRGBA64(r)
	case color.NRGBAModel:
		return image.NewNRGBA(r)
	case color.NRGBA64Model:
		return image.NewNRGBA64(r)
	case color.AlphaModel:
		return image.NewAlpha(r)
	case color.Alpha16Model:
		return image.NewAlpha16(r)
	case color.GrayModel:
		return image.NewGray(r)
	case color.Gray16Model:
		return image.NewGray16(r)
	default:
		return image.NewRGBA(r)
	}
}


func (img *Image) Clip(s string) error {
	r, err := img.parseBounds(s)
	if err != nil {
		return err
	}

	dst := newDrawImage(image.Rectangle{image.ZP, r.Size()}, img.ColorModel())

	draw.Draw(dst, dst.Bounds(), img, r.Min, draw.Src)

	img.Image = dst

	return nil
}