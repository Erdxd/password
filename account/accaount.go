package account

import (
	"errors"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/fatih/color"
)

// композиция(наследование)

var letterruns = []rune("qwertyuiop[]asdfghjkl;'zxcvbnm,.1234567890-=")

type Account struct {
	Login       string    `json:"Login" `
	Password    string    `json:"Password" `
	URL         string    `json:"Url" `
	CreatedTime time.Time `json:"CreatedTime" `
	UpdateTime  time.Time `json:"UpdatedtTime" `
}

func (acc *Account) GeneratePassword(n int) string {
	ald := make([]rune, n)

	for i := range ald {
		ald[i] = letterruns[rand.IntN(len(letterruns))]

	}
	acc.Password = string(ald)
	gen := string(ald)
	return gen

}

func (acc *Account) Output() {
	color.Green(acc.Login)
	color.Green(acc.Password)
	color.Green(acc.URL)
}

func NewAccount(login, password, URLstring string) (*Account, error) {

	_, err := url.ParseRequestURI(URLstring)
	if err != nil {

		return nil, errors.New("Invalid_URL")
	}

	Newacc := &Account{
		CreatedTime: time.Now(),
		UpdateTime:  time.Now(),
		URL:         URLstring,
		Login:       login,
		Password:    password,
	}
	//field, _ := reflect.TypeOf(Newacc).Elem().FieldByName("Login")
	//fmt.Println(string(field.Tag))

	return Newacc, nil

}
