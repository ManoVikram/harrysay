package main

import (
	"fmt"
	"os"

	"github.com/qeesung/image2ascii/convert"
)

func main() {
	if len(os.Args) == 1 || (len(os.Args) == 2 && os.Args[1] == "-h") || (len(os.Args) == 2 && os.Args[1] == "--help") {
		fmt.Printf("Generate an ASCII text from an image.\n")
		return
	}

	imageName := os.Args[1]

	convertOptions := convert.DefaultOptions
	convertOptions.FixedHeight = 30
	convertOptions.FixedWidth = 80
	convertOptions.Reversed = true

	converter := convert.NewImageConverter()
	fmt.Println(converter.ImageFile2ASCIIString(imageName, &convertOptions))
}
