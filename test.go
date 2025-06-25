package main

import (
	"fmt"

	"github.com/Erdxd/password/account"
	"github.com/Erdxd/password/files"
)

func main() {

	login := promtData("Введите логин")
	password := promtData("Введите пароль")
	URL := promtData("Введите URL")
	var err error
	MYaccount, err := account.NewAccountWithTimeStamp(login, password, URL)
	if err != nil {
		fmt.Println("неверный формат URL")
		return
	}
	if login == "" {
		fmt.Println("логина нету")
		return
	}
	if password == "" {
		fmt.Println("нету пароля,генерируем...")
		MYaccount.GeneratePassword(10)
	}

	MYaccount.Outputhassword()
	files.Readfile()
	files.Writefiles("привет я файл", "file.txt")

}

func promtData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res

}
