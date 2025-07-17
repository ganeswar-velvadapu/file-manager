package operations

import (
	// "bufio"
	"fmt"
	"os"
)

func OpenAndReadFile() {
	file, err := os.Open("hello.txt")

	if err != nil {
		fmt.Println("Error opening file")
	}
	defer file.Close()
	fmt.Println(file)
	fileContentByte := make([]byte,100)
	
	fileContent,err := file.Read(fileContentByte)

	if err!=nil{

		fmt.Println("Error reading file")
	}

	fmt.Println(string(fileContentByte[:fileContent]))

}
