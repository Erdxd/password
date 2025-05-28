package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
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
	fmt.Println(acc.login, "-логин", acc.password, "-пароль", acc.URL, "-URL")
}
func newAccount(login, password, URLstring string) (*account, error) {

	_, err := url.ParseRequestURI(URLstring)
	if err != nil {

		return nil, errors.New("Invalid_URL")
	}

	Newacc := &account{
		login:    login,
		password: password,
		URL:      URLstring,
	}

	return Newacc, nil

}

type account struct {
	login    string
	password string
	URL      string
}

func main() {

	login := promtData("Введите логин")
	password := promtData("Введите пароль")
	URL := promtData("Введите URL")
	var err error

	var MYaccount *account
	MYaccount, err = newAccount(login, password, URL)
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
		MYaccount.generatePassword(10)
	}

	MYaccount.outputhassword()
}

func promtData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res

}
