package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
)


func DBInit() *gorm.DB {
	log.Info("Starting Database Connection")
	dbURI := "host=localhost user=postgres password=password dbname=restgo port=5432 sslmode=disable Timezone=Asia/Shanghai"
	db, err := gorm.Open("postgres", dbURI)
	if err != nil{
		log.Panic("failed to connect database with error " + err.Error())
	}

	db.LogMode(true)

	return db
}
