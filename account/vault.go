package account

import (
	"encoding/json"
	"strings"

	"time"

	"github.com/Erdxd/password/files"

	"github.com/fatih/color"
)

type Vault struct {
	Accounts     []Account `json:"accounts"`
	UpdatedtTime time.Time `json:"UpdateTime"`
}

func NewVault() *Vault {
	db := files.NewJsonDB("data.json")
	file, err := db.Read()
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
		return &Vault{
			Accounts:     []Account{},
			UpdatedtTime: time.Now(),
		}
	}
	return &vault

}

func (vault *Vault) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.save()

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
func (vault *Vault) DeleteAcccountBYURL(url string) bool {
	var accounts []Account
	isDELETED := false
	for _, account := range vault.Accounts {
		isMatched := strings.Contains(account.URL, url)
		if !isMatched {

			accounts = append(accounts, account)

		}
		isDELETED = true

	}
	vault.Accounts = accounts
	vault.save()
	return isDELETED

}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	return file, nil
}
func (vault *Vault) save() {
	db := files.NewJsonDB("data.json")
	vault.UpdatedtTime = time.Now()
	data, err := vault.ToBytes()
	if err != nil {
		color.Red("Не удалось преобразовать файл JSON")
		return

	}
	db := files.NewJsonDB("data.json")
	db.Write(data)

}
