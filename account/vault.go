package account

import (
	"encoding/json"
	"strings"

	"time"

	"github.com/Erdxd/password/encryptor"
	"github.com/fatih/color"
)

type ByteReader interface {
	Read() ([]byte, error)
}
type ByteWriter interface {
	Write([]byte)
}

type Db interface {
	ByteWriter
	ByteReader
}
type Vault struct {
	Accounts     []Account `json:"accounts"`
	UpdatedtTime time.Time `json:"UpdateTime"`
}
type VaultwithDB struct {
	Vault
	db  Db
	enc encryptor.Encrypter
}

func NewVault(db Db, enc encryptor.Encrypter) *VaultwithDB {
	file, err := db.Read()

	if err != nil {

		return &VaultwithDB{
			Vault: Vault{
				Accounts:     []Account{},
				UpdatedtTime: time.Now(),
			},
			db:  db,
			enc: enc,
		}
	}
	data := enc.Decrypt(file)

	var vault Vault
	err = json.Unmarshal(data, &vault)
	color.Cyan("Найденно %d аккаунтов", len(vault.Accounts))
	if err != nil {
		color.Red(err.Error())
		return &VaultwithDB{
			Vault: Vault{
				Accounts:     []Account{},
				UpdatedtTime: time.Now(),
			},
			db:  db,
			enc: enc,
		}
	}
	return &VaultwithDB{
		Vault: vault,
		db:    db,
		enc:   enc,
	}
}

func (vault *VaultwithDB) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.save()

}
func (vault *VaultwithDB) FIndaccounts(str string, checker func(Account, string) bool) []Account {
	var accounts []Account
	for _, account := range vault.Accounts {
		isMatched := checker(account, str)
		if isMatched {
			accounts = append(accounts, account)

		}

	}
	return accounts

}

func (vault *VaultwithDB) DeleteAcccountBYURL(url string) bool {
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
func (vault *VaultwithDB) save() {

	vault.UpdatedtTime = time.Now()
	data, err := vault.Vault.ToBytes()
	encData := vault.enc.Encrypt(data)
	if err != nil {
		color.Red("Не удалось преобразовать файл vault")
		return

	}
	vault.db.Write(encData)

}
