package config

import (
	"os"

	"github.com/danilocassiano/GoOportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/min.db"

	//Check if the database file exists 
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err){
		logger.Info("database file not found, creating... ")
		
		//create the database file and directory
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}

		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	//create db and connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqlite opening error: %v", err)
		return nil, err
	}
	//migrate the schema
	err = db.AutoMigrate( &schemas.Opening{} )
	if err != nil {
		logger.Errorf("sqlite migration error: %v", err)
		return nil, err
	}

	//return the DB 
	return db, nil
}