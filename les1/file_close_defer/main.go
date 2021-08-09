package main

import (
	"os"
)

func main() {

	file, err := os.Create("C:\\Users\\gurov\\go_dump\\les_1\\file_close_defer\\test_file")
	defer file.Close()

	if err != nil {
		panic("File operation failed")
	}

}
