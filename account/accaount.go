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

func (acc *Account) GeneratePassword(n int) {
	ald := make([]rune, n)

	for i := range ald {
		ald[i] = letterruns[rand.IntN(len(letterruns))]

	}
	acc.password = string(ald)
}

func (acc *Account) Outputhassword() {
	color.Green(acc.login)
	color.Green(acc.password)
	color.Green(acc.URL)
}

func NewAccountWithTimeStamp(login, password, URLstring string) (*AccountWithTimeStamp, error) {

	_, err := url.ParseRequestURI(URLstring)
	if err != nil {

		return nil, errors.New("Invalid_URL")
	}

	Newacc := &AccountWithTimeStamp{
		createdTime: time.Now(),
		updateTime:  time.Now(),
		Account: Account{
			login:    login,
			password: password,
			URL:      URLstring,
		},
	}

	return Newacc, nil

}

type Account struct {
	login    string
	password string
	URL      string
}
type AccountWithTimeStamp struct {
	createdTime time.Time
	updateTime  time.Time
	Account
}
