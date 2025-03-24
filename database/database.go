package database

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("livros.db"), &gorm.Config{})
	if err != nil {
		panic("Erro ao conectar com o banco de dados")
	}
}