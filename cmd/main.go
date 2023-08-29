package main

import (
	"assg/pizzashop/pkg/common/db"
	"assg/pizzashop/pkg/pizzashop"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"assg/pizzashop/pkg/common/models"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func main() {
	viper.SetConfigFile("../pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	fmt.Println(port, dbUrl)

	r := gin.Default()
	h := db.Init(dbUrl)

	populateDB(h)

	pizzashop.RegisterRoutes(r, h)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"port":  port,
			"dbUrl": dbUrl,
		})
	})

	r.Run()
}

func populateDB(db *gorm.DB) {
	var toppings []models.Topping
	var pizzas []models.Pizza
	var sizes []models.Size
	var crusts []models.Crust

	getData("F:/Onito_Assg/pkg/common/db/data/crusts.json", &crusts)
	getData("F:/Onito_Assg/pkg/common/db/data/pizzas.json", &pizzas)
	getData("F:/Onito_Assg/pkg/common/db/data/sizes.json", &sizes)
	getData("F:/Onito_Assg/pkg/common/db/data/toppings.json", &toppings)

	for _, topping := range toppings {
		db.Create(&topping)
	}

	for _, pizzas := range pizzas {
		db.Create(&pizzas)
	}

	for _, size := range sizes {
		db.Create(&size)
	}

	for _, crust := range crusts {
		db.Create(&crust)
	}
}

func getData(fileName string, v interface{}) {
	file, _ := os.Open(fileName)
	defer file.Close()
	byteValue, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteValue, v)
}
