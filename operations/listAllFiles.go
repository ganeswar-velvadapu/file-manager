package operations

import (
	"fmt"
	"os"
)

func ListAll() {
	 entries, err := os.ReadDir("./")
	if err != nil {
		fmt.Println("Error in listing all the contents : ",err)
		return
	}
	

	for _, e := range entries {
		fmt.Println("File/Directory Name :",e.Name(), " Is a directory",e.IsDir())
	}
}
