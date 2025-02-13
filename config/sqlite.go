package config

import (
	"os"

	"github.com/DaviFernandes034/go-api/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


func initSqlite()(*gorm.DB, error){

	logger:= GetLogger("sqlite")
	dbPath:= "./db/main.db"

	//checando se a database ja existe

	_, err:= os.Stat(dbPath)
	if os.IsNotExist(err){

		logger.Infof("database não existente, criando....")

		//criando o diretorio e os arquivos existentes
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err:= os.Create(dbPath)
		if err != nil{
			return nil , err
		}

		file.Close()
	}

	//criação e conexao do banco de dados

	db, err:=  gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil{
		logger.Errorf("sqlite openig error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&database.Opening{})
	if err != nil{

		logger.Errorf("sqlite migration error: %v", err)
	}

	return db, nil

}