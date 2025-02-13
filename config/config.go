package config

import (
	"fmt"

	"gorm.io/gorm"
)


var (

	db *gorm.DB
	logger *Logger
)

func Init() error {

	var err error

	db, err = initSqlite()
	if err != nil{
		return fmt.Errorf("erro ao inicializar o sqlite: %v", err)
	}

	return nil

}

func GetSqlite() *gorm.DB{
 
	return db

}


func GetLogger(p string)*Logger{

	logger:= NewLogger(p)
	return logger
}