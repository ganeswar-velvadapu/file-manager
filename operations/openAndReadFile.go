package operations

import (
	// "bufio"
	"fmt"
	"os"
)

func OpenAndReadFile() {
	// just for reference we can read files like this too but we need to specify buffer before hand (which in this case is 100 bytes) which is not ideal for real world
	// file, err := os.Open("hello.txt")
	cwd, _ := os.Getwd()
	fmt.Println("Current working directory:", cwd)
	// if err != nil {
	// 	fmt.Println("Error opening file")
	// }
	// defer file.Close()
	// fmt.Println(file)
	// fileContentByte := make([]byte,100)

	// fileContent,err := file.Read(fileContentByte)

	// if err!=nil{

	// 	fmt.Println("Error reading file")
	// }

	// fmt.Println(string(fileContentByte[:fileContent]))

	data, err := os.ReadFile("hello.txt")
	if err != nil {
		fmt.Println("Error in reading file :", err)
		return
	}
	fmt.Println(string(data))

}
