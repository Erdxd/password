package main

import (
	"fmt"

	"github.com/Erdxd/password/account"
	"github.com/Erdxd/password/fis"
	"github.com/fatih/color"
)

func main() {

	vault := account.NewVault()
	defer fmt.Println("Check the version, this app will be redesigned with interface, new update can be found at https://github.com/Erdxd/password/releases/tag/new")
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
			CreatedAccount(vault)
		case "2":
			Findaccount(vault)

		case "3":
			DeleteAcccount(vault)

		case "4":
			break menu
		case "111":
			fis.Fisacc()
		default:
			color.Red("Вы ввели что-то не так")
			continue menu
		}

	}

}

func CreatedAccount(vault *account.Vault) {
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
	vault = account.NewVault()
	vault.AddAccount(*MYaccount)

}
func Findaccount(vault *account.Vault) {
	fmt.Println("Найти аккаунт")
	URl := promtData("ВВедите URl")

	accounts := vault.FIndaccountBYURL(URl)
	if len(accounts) == 0 {
		color.Red("Аккаунтов по этому URl не найденно")
	}
	for _, account := range accounts {
		account.Output()
	}
}
func DeleteAcccount(vault *account.Vault) {
	fmt.Println("Найти аккаунт")
	URl := promtData("ВВедите URl")
	isDELETED := vault.DeleteAcccountBYURL(URl)
	if isDELETED {
		color.Green("Аккаунт удален")

	} else {
		color.Red("Аккаунт не удален")
	}

}

func promtData(prompt string) string {
	fmt.Println(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res

}
