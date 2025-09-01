package main

import (
	"fmt"
	"strings"

	"github.com/Erdxd/password/account"
	"github.com/Erdxd/password/files"
	"github.com/Erdxd/password/fis"
	"github.com/Erdxd/password/output"
	"github.com/fatih/color"
)

func main() {
counter := menuCounter()

	var menu = map[string]func(*account.VaultwithDB){
		"1": CreatedAccount,
		"2": DeleteAcccount,
		"3": Findaccount,
		"4": FindaccountByLogin,
	}
	//defer fmt.Println("Check the version, this app will be redesigned with interface, new update can be found at https://github.com/Erdxd/password/releases/tag/new")

	vault := account.NewVault(files.NewJsonDB("data.json"))
	//vault := account.NewVault(cloud.NewCloudDB("htttps://a.ru")
	fmt.Println("___Менеджер паролей___")

menu:
	for {
		counter()

		Useranswer := promtData(
			"1-Создание аккаунта",
			"2-Удалить аккаунт ",
			"3-Найти аккаунт по URL",
			"4-Найти аккаунт по логину",
			"5-Выход",
			"Выберите вариант",
		)
		FuncMenu := menu[Useranswer]
		if Useranswer == "5" {
			break

		} else if Useranswer == "111" {
			fisacc()
			continue
		} else if FuncMenu == nil {
			output.PrintError("Попробйте еще раз")
			continue menu
		}

		FuncMenu(vault)

		/*switch Useranswer {
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

		f
		*/
	}

}
func fisacc() {
	fis.Fisacc()

}
func menuCounter() func() { // Замыкание(функция которая помнит свои переменные и их значения	)
	i := 0
	return func() {
		i++
		fmt.Println(i)
	}
}
func CreatedAccount(vault *account.VaultwithDB) {
	login := promtData("Введите логин")
	password := promtData("Введите пароль")
	URL := promtData("Введите URL")

	var err error
	MYaccount, err := account.NewAccount(login, password, URL)
	if err != nil {
		output.PrintError("Неверный формат URl")
		return
	}
	if login == "" {
		output.PrintError("Логина нет")
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

	vault.AddAccount(*MYaccount)

}
func FindaccountByLogin(vault *account.VaultwithDB) {
	fmt.Println("Найти аккаунт")
	Login := promtData("Введите Login")

	accounts := vault.FIndaccounts(Login, func(acc account.Account, str string) bool { //Анонимная функция
		return strings.Contains(acc.Login, str)
	})
	if len(accounts) == 0 {
		output.PrintError("Аккаунтов по этому Login не найденно")
	}
	for _, account := range accounts {
		account.Output()
	}
}
func Findaccount(vault *account.VaultwithDB) {
	fmt.Println("Найти аккаунт")
	URl := promtData("ВВедите URl")

	accounts := vault.FIndaccounts(URl, func(acc account.Account, str string) bool { //Анонимная функция
		return strings.Contains(acc.URL, str)
	})
	if len(accounts) == 0 {
		output.PrintError("Аккаунтов по этому URl не найденно")
	}
	for _, account := range accounts {
		account.Output()
	}
}

func DeleteAcccount(vault *account.VaultwithDB) {
	fmt.Println("Найти аккаунт")
	URl := promtData("ВВедите URl")
	isDELETED := vault.DeleteAcccountBYURL(URl)
	if isDELETED {
		color.Green("Аккаунт удален")

	} else {
		output.PrintError("Аккаунт не удален")
	}

}

func promtData(prompt ...any) string {
	for i, line := range prompt {
		if i == len(prompt)-1 {
			fmt.Printf("%v:", line)
		} else {
			fmt.Println(line)
		}
	}

	var res string
	fmt.Scanln(&res)
	return res

}
func app(){
fmt.Println("cj")}