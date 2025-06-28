package files

import (
	"fmt"
	"os"
)

func Readfile(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(string(data))
	return data, nil

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
