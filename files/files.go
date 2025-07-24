package files

import (
	"fmt"
	"os"

	"github.com/Erdxd/password/output"
)

type Jsondb struct {
	filename string
}

func NewJsonDB(name string) *Jsondb {
	return &Jsondb{
		filename: name,
	}

}

func (db *Jsondb) Read() ([]byte, error) {
	data, err := os.ReadFile(db.filename)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return data, nil

}
func (db *Jsondb) ReadFis() ([]byte, error) {
	data, err := os.ReadFile(db.filename)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(string(data))
	return data, nil

}
func (db *Jsondb) Write(content []byte) {
	file, err := os.Create(db.filename)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	if err != nil {

		output.PrintError(err)
		return
	}
	_, err = file.Write(content)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Запись успешна")

}
