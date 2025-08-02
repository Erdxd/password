package main

import (
	"fmt"

	"github.com/Erdxd/password/account"
	"github.com/Erdxd/password/files"
	"github.com/Erdxd/password/output"
	"github.com/fatih/color"
)

func main() {
	var menu = map[string]func(*account.VaultwithDB){
		"1": CreatedAccount,
		"2": Findaccount,
		"3": DeleteAcccount,
	}
	//defer fmt.Println("Check the version, this app will be redesigned with interface, new update can be found at https://github.com/Erdxd/password/releases/tag/new")

	vault := account.NewVault(files.NewJsonDB("data.json"))
	//vault := account.NewVault(cloud.NewCloudDB("htttps://a.ru")
	fmt.Println("___Менеджер паролей___")

menu:
	for {
		Useranswer := promtData([]string{
			"1-Создание аккаунта",
			"2-Найти аккаунт",
			"3-Удаление аккаунта",
			"4-Выход",
			"Выберите вариант",
		})
		FuncMenu := menu[Useranswer]
		if FuncMenu == nil {
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
		*/
	}

}

func CreatedAccount(vault *account.VaultwithDB) {
	login := promtData([]string{"Введите логин"})
	password := promtData([]string{"Введите пароль"})
	URL := promtData([]string{"Введите URL"})

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
func Findaccount(vault *account.VaultwithDB) {
	fmt.Println("Найти аккаунт")
	URl := promtData([]string{"ВВедите URl"})

	accounts := vault.FIndaccountBYURL(URl)
	if len(accounts) == 0 {
		output.PrintError("Аккаунтов по этому URl не найденно")
	}
	for _, account := range accounts {
		account.Output()
	}
}
func DeleteAcccount(vault *account.VaultwithDB) {
	fmt.Println("Найти аккаунт")
	URl := promtData([]string{"ВВедите URl"})
	isDELETED := vault.DeleteAcccountBYURL(URl)
	if isDELETED {
		color.Green("Аккаунт удален")

	} else {
		output.PrintError("Аккаунт не удален")
	}

}

func promtData[T any](prompt []T) string {
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
