package main

import (
	"embed"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

//go:embed harry
var embedHarryFiles embed.FS

func getNumberOfHarryFiles() int {
	files, err := embedHarryFiles.ReadDir("harry")

	if err != nil {
		log.Fatal(err)
	}

	return len(files)
}

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

		rand.Seed(time.Now().UnixNano())

		fileNumber := rand.Intn(getNumberOfHarryFiles() - 1)

		file, err := embedHarryFiles.ReadFile("harry/harry_" + strconv.Itoa(fileNumber+1) + ".txt")

		if err != nil {
			log.Fatal("Error reading the ASCII file!")
		}

		fmt.Println(string(file))
	}
}
