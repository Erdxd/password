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
func Writefiles(namelogin string, name string, namePass string, nameUrl string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
	}
	_, err = file.WriteString(namelogin + "-логин     ")
	defer file.Close()
	if err != nil {

		fmt.Println(err)
		return
	}
	file.WriteString(namePass + "-пароль	  ")
	file.WriteString(nameUrl + "-URL	  ")

	fmt.Println("Запись успешна")

}
