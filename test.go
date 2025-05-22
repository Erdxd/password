package main

import (
	"fmt"
)

type account struct {
	login    string
	password string
	URL      string
}

func main() {
	str := "привет!)"        // rune(руны)
	for d, ch := range str { //ch - юникод значение(под капотом) string(ch) — преобразует эту руну обратно в строку из одного символа.
		fmt.Println(d, ch, "=", string(ch))
		// d-байтовый индекс русские буквы занимают 2 байта а данные символы по одному, он их сумирует
	}
	login := promtData("Введите логин")
	password := promtData("Введите пароль")
	URL := promtData("Введите URL")

	MYaccount := account{
		login:    login,
		password: password,
		URL:      URL,
	}
	outputhassword(&MYaccount)
}
func promtData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scan(&res)
	return res

}
func outputhassword(acc *account) {
	fmt.Println(acc.login, acc.password, acc.URL)
fmt.Println("acc")
}
