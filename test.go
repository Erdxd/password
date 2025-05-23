package main

import (
	"fmt"
	"math/rand/v2"
)

var letterruns = []rune("qwertyuiop[]asdfghjkl;'zxcvbnm,.1234567890-=")

func (acc *account) generatePassword(n int) {
	ald := make([]rune, n)

	for i := range ald {
		ald[i] = letterruns[rand.IntN(len(letterruns))]

	}
	acc.password = string(ald)
}

func (acc *account) outputhassword() {
	fmt.Println(acc.login, acc.password, acc.URL)
}

type account struct {
	login    string
	password string
	URL      string
}

func main() {

	login := promtData("Введите логин")
	//password := promtData("Введите пароль")
	URL := promtData("Введите URL")

	MYaccount := account{
		login: login,
		//password: password,
		URL: URL,
	}
	MYaccount.generatePassword(12)
	MYaccount.outputhassword()

}
func promtData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scan(&res)
	return res

}
