package db

import (
	"log"

	"assg/pizzashop/pkg/common/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormLogger struct{}

func Init(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(
		&models.Pizza{},
		&models.Topping{},
		&models.Customer{},
		&models.Crust{},
		&models.Size{},
		&models.PizzaOrder{},
		&models.PizzaOrderTopping{},
	)
	return db
}
