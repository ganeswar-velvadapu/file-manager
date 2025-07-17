package operations

import (
	"bufio"
	"fmt"
	"os"
)

func WriteToFile() {
	var fileName string
	fmt.Print("Enter the name of the file you want to write to: ")
	_, err := fmt.Scanln(&fileName)
	if err != nil {
		fmt.Println("Error reading file name:", err)
		return
	}

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		fmt.Println("File does not exist. Creating file...")
	} else {
		fmt.Println("File exists. Writing to it...")
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter the content (one line): ")
	scanner.Scan()
	content := scanner.Text()

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading content:", err)
		return
	}

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error opening/creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(content + "\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("File writing successful.")
}
