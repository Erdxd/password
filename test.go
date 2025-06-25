package main

import (
	"fmt"

	"github.com/Erdxd/password/account"
	"github.com/Erdxd/password/files"
)

func main() {
menu:
	for {

		fmt.Println("Меню:")
		fmt.Println("1-Создание аккаунта")
		fmt.Println("2-Найти аккаунт")
		fmt.Println("3-Удаление аккаунта")
		fmt.Println("4-Выход")
		Useranswer := ""
		fmt.Scanln(&Useranswer)
		switch Useranswer {
		case "1":
			CreatedAccount()
		case "2":
			Findaccount()

		case "3":
			DeleteAcccount()

		case "4":
			break menu
		default:
			fmt.Println("Вы ввели что-то не так")
			continue menu
		}

	}
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
	vault := account.NewVault()
	vault.AddAccount(*MYaccount)
	data, err := vault.ToBytes()

	if err != nil {
		fmt.Println("не удалось преобращовать в json данные")
		return

	}
	files.Writefiles("data.Json", data)

}
func Findaccount() {
	fmt.Println("Найти аккаунт")

}
func DeleteAcccount() {

}

func promtData(prompt string) string {
	fmt.Println(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res

}
