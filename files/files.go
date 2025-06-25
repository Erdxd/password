package files

import (
	"fmt"
	"os"
)

func Readfile() {
	data, err := os.ReadFile("file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))

}
func Writefiles(name string, content []byte) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	if err != nil {

		fmt.Println(err)
		return
	}
	_, err = file.Write(content)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Запись успешна")

}
