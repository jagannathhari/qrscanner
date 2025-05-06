package main

import (
	"bytes"
	"fmt"
	"image"
	"io"
	"log"
	"os"

	"github.com/caiguanhao/readqr"
	"golang.design/x/clipboard"
)

func main() {
	arguments := os.Args[1:]
	var data []byte
	var err error

	if len(arguments) == 0 || arguments[0] == "--" {
		data, err = io.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
	} else if arguments[0] == "-c" || arguments[0] == "-C" {
		err := clipboard.Init()
		if err != nil {
			log.Fatal(err)
		}
		data = clipboard.Read(clipboard.FmtImage)
		if len(data) == 0 {
			log.Fatal("No image found in clipboard")
		}
	} else {
		// Read from file
		data, err = os.ReadFile(arguments[0])
		if err != nil {
			log.Fatal(err)
		}
	}

	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}

	result, err := readqr.DecodeImage(img)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
