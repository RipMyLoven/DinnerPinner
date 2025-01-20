package file

import (
	"bufio"
	"fmt"
	"os"
)

func CreateFile(fileName, text string) {
	fmt.Println("Writing file")
	file, err := os.Create(fileName)
	if err != nil {
		panic("Failed to create file: " + err.Error())
	}
	defer file.Close()

	length, err := file.WriteString(text)
	if err != nil {
		panic(err)
	}
	fmt.Printf("File name: %s\n", file.Name())
	fmt.Printf("File length: %d bytes\n", length)
}

func ReadFile(fileName string) {
	fmt.Println("\nReading file")
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	fmt.Println("File name: " + fileName)
	fmt.Printf("File size: %d\n", len(data))
	fmt.Printf("File content: %s\n", data)
}

func main() {
	// Input file name
	fmt.Println("Enter fileName: ")
	var fileName string
	fmt.Scanln(&fileName)

	// Input file content
	fmt.Println("Enter file content: ")
	inputReader := bufio.NewReader(os.Stdin)
	input, _ := inputReader.ReadString('\n')

	// Create and read the file
	CreateFile(fileName, input)
	ReadFile(fileName)
}
