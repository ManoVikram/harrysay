package main

import (
	"embed"
	"fmt"
	"log"
	"os"
	"strings"
)

//go:embed harry
var embedHarryFiles embed.FS

func main() {
	if len(os.Args) == 1 || (len(os.Args) == 2 && os.Args[1] == "-h") || (len(os.Args) == 2 && os.Args[1] == "--help") {
		fmt.Printf("Harrysay is inspired by Cowsay.\nHarrysay displays a message by young Harry Potter.\n\nUsage:\n   harrysay MESSAGE\n\nExample:\n   harrysay Hello Prof Snape!\n")
		return
	}

	if len(os.Args) > 1 {
		message := strings.Join(os.Args[1:], " ")
		totalCharactersInMessage := len(message)

		verticalBorderLine := " "

		for i := 0; i < totalCharactersInMessage; i++ {
			verticalBorderLine += "-"
		}

		verticalBorderLine += " "

		fmt.Println(verticalBorderLine)
		fmt.Printf("\n<%s>\n", message)
		fmt.Println(verticalBorderLine)
		fmt.Println("      \\")
		fmt.Println("       \\")
		fmt.Println("        \\")
		fmt.Println("         \\")
		fmt.Println("          \\")

		file, err := embedHarryFiles.ReadFile("harry/harry_1.txt")

		if err != nil {
			log.Fatal("Error reading the ASCII file!")
		}

		fmt.Println(string(file))
	}
}
