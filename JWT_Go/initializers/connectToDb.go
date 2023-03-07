package initializers

import (
	
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)




func ConnectToDb() *gorm.DB {
	var DB *gorm.DB
	var err error
	dsn := os.Getenv("DB")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err!= nil {
		panic("Fail to connect to database")
	
	}else{
		
		return DB
	}

}