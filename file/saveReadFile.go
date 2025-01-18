package file

import (
	"bufio"
	"fmt"
	"os"
)

func saveString() {
	newFile, _ := os.Create("string.txt")
	defer newFile.Close()

	bytes, _ := newFile.WriteString("New content for new file")

	fmt.Printf("%d bytes have been written to file", bytes)
}

func openFileReadContent() {
	fileContent, _ := os.ReadFile("string.txt")

	fmt.Print(string(fileContent))
}

func saveLinesToFile() {
	newFile, _ := os.Create("lines.txt")
	defer newFile.Close()

	lines := []string{"Line1", "Line2", "Line3"}
	totalBytes := 0
	for _, line := range lines {
		bytes, _ := fmt.Fprintln(newFile, line)
		totalBytes += bytes
	}
	fmt.Printf("%d bytes written in total", totalBytes)
}

func readFileLineByLine() {
	linesFile, _ := os.Open("lines.txt")
	defer linesFile.Close()

	lineScanner := bufio.NewScanner(linesFile)
	for lineScanner.Scan() {
		lineContent := lineScanner.Text()
		fmt.Println(lineContent)
		fmt.Println("One line read ----")
	}
	if scannerErr := lineScanner.Err(); scannerErr != nil {
		fmt.Print("Error happend during scanning")
	}

}
