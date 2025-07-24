package fis

import (
	"github.com/Erdxd/password/files"
	"github.com/fatih/color"
)

func Fisacc() {
	color.Red("!!!!!!!")
	db := files.NewJsonDB("data.json")

	db.ReadFis()

	color.Red("!!!!!!!")

}
