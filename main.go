package main

import (
	"fmt"
	"time"

	"github.com/ganeswar-velvadapu/file-manager/operations"
)

func main(){
	
	var i int;

	fmt.Println("Click 1 to open and read a file")
	fmt.Println("Click 2 to create a file")
	fmt.Println("Click 3 to open a directory")
	fmt.Println("Click 4 to create a directory")
	fmt.Println("Click 5 to create a file inside a directory")

	fmt.Println("Please enter your number")
	fmt.Scanln(&i)
		
	 for {
		if i ==0 {
			break;
		}else if i == 1{
			fmt.Println("You want to open and read file")
			operations.OpenAndReadFile()
			time.Sleep(1 * time.Second)
		}else if i == 2{
			fmt.Println("You want to create a file")
			time.Sleep(1 * time.Second)
		}else if i == 3{
			fmt.Println("You want to open a directory")
			time.Sleep(1 * time.Second)
		}else if i == 4{
			fmt.Println("You want to create a directory")
			time.Sleep(1 * time.Second)
		}else if i == 5{
			fmt.Println("You want to create a file inside a directory")
			time.Sleep(1 * time.Second)
		}else{
			fmt.Print("Please select a valid option\n")
		}
    }	




}