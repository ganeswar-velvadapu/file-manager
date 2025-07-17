package operations

import (
	"fmt"
	"os"
)

func CreateFile() {
	var fileName string

	fmt.Print("Enter the file name : ")

	_, err := fmt.Scanln(&fileName)

	if err != nil {
		fmt.Print("Error in reading the file", err)
		return
	}

	_, err = os.Open(fileName)
	if err == nil {
		fmt.Println("File exists!")
	} else if os.IsNotExist(err) {
		fmt.Println("File does not exist. Creating file")
	} else {
		fmt.Println("Some other error : ", err)
	}

	_, err = os.Create(fileName)

	if err != nil {
		fmt.Print("Error in creating the file", err)
		return
	}

	fmt.Println("File created successfully")

}
