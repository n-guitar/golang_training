package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

func main() {
	s := flag.String("i", "jpeg", "jpeg or gif or png")
	flag.Parse()

	img, err := readImage(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "no image", err)
		os.Exit(1)
	}

	switch *s {
	case "jpeg":
		toJPEG(img, os.Stdout)
	case "gif":
		toGIF(img, os.Stdout)
	case "png":
		toPNG(img, os.Stdout)
	default:
		fmt.Print("input error")
		os.Exit(1)
	}

	// if err := toJPEG(os.Stdin, os.Stdout); err != nil {
	// 	fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
	// 	os.Exit(1)
	// }
}

// https://twinbird-htn.hatenablog.com/entry/2017/04/23/001743
// image型で返す
func readImage(in io.Reader) (image.Image, error) {
	img, kind, err := image.Decode(in)
	if err != nil {
		return nil, err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return img, nil
}

func toJPEG(img image.Image, out io.Writer) error {
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}

func toGIF(img image.Image, out io.Writer) error {
	return gif.Encode(out, img, nil)
}

func toPNG(img image.Image, out io.Writer) error {
	return png.Encode(out, img)
}
