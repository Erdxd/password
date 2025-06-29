package account

import (
	"encoding/json"
	"strings"

	"time"

	"github.com/Erdxd/password/files"
	"github.com/fatih/color"
)

type Vault struct {
	Accounts     []Account `json:"account"`
	UpdatedtTime time.Time `json:"UpdateTime "`
}

func NewVault() *Vault {
	file, err := files.Readfile("data.json")
	if err != nil {
		return &Vault{
			Accounts:     []Account{},
			UpdatedtTime: time.Now(),
		}
	}
	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		color.Red(err.Error())

	}
	return &vault

}
func (vault *Vault) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.UpdatedtTime = time.Now()
	data, err := vault.ToBytes()
	if err != nil {
		color.Red("Не удалось преобразовать файл JSON")
		return

	}
	files.Writefiles("data.json", data)

}
func (vault *Vault) FIndaccountBYURL(url string) []Account {
	var accounts []Account
	for _, account := range vault.Accounts {
		isMatched := strings.Contains(account.URL, url)
		if isMatched {
			accounts = append(accounts, account)

		}

	}
	return accounts

}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	return file, nil
}
