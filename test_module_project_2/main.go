package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 6; i++ {
		if test_insert("testid"+string(i+48), "testname"+string(i+48), "testdesc"+string(i+48), "testlink"+string(i+48), ""+string(i+48)) {
			fmt.Println("GetTest    ", "testid"+string(i+48), " пройден")
		} else {
			fmt.Println("GetTest    ", "testid"+string(i+48), "не пройден")
		}

	}

}
