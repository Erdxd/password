package main

import (
	"fmt"

	"github.com/Erdxd/password/account"
	"github.com/Erdxd/password/files"
)

func main() {
	CreatedAccount()

}

func CreatedAccount() {
	login := promtData("Введите логин")
	password := promtData("Введите пароль")
	URL := promtData("Введите URL")

	var err error
	MYaccount, err := account.NewAccount(login, password, URL)
	if err != nil {
		fmt.Println("неверный формат URL")
		return
	}
	if login == "" {
		fmt.Println("Логина нету")
		return
	}
	if password == "" {
		fmt.Println("нету пароля,генерируем...")
		MYaccount.GeneratePassword(10)
		//gen := MYaccount.GeneratePassword(10)
		//files.Writefiles(login, "file.txt", gen, URL)
	} else {
		//files.Writefiles(login, "file.txt", password, URL,)
	}
	file, err := MYaccount.ToBytes()
	if err != nil {
		fmt.Println("не удалось преобращовать в json данные")
		return

	}
	files.Writefiles("data.Json", file)

}

func promtData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res

}
