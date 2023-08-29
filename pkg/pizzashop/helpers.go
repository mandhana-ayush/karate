package pizzashop

import (
	"assg/pizzashop/pkg/common/models"
	"fmt"

	"gorm.io/gorm"
)

func getToppingsById(db *gorm.DB, id int) (*models.Topping, error) {
	var topping models.Topping
	if result := db.First(&topping, id); result.Error != nil {
		return nil, result.Error
	}

	return &topping, nil
}

func getSizesById(db *gorm.DB, id int) (*models.Size, error) {
	var size models.Size
	if result := db.First(&size, id); result.Error != nil {
		return nil, result.Error
	}

	return &size, nil
}

func getCrustById(db *gorm.DB, id int) (*models.Crust, error) {
	var crust models.Crust
	if result := db.First(&crust, id); result.Error != nil {
		return nil, result.Error
	}

	return &crust, nil
}

func getPizzasById(db *gorm.DB, id int) (*models.Pizza, error) {
	var pizza models.Pizza
	if result := db.First(&pizza, id); result.Error != nil {
		return nil, result.Error
	}

	return &pizza, nil
}

func getPizzaByName(db *gorm.DB, name string) (*models.Pizza, error) {
	var pizza models.Pizza
	if result := db.Where("name = ?", name).First(&pizza); result.Error != nil {
		fmt.Println(result)
		return nil, result.Error
	}

	return &pizza, nil
}

func getCrustByName(db *gorm.DB, name string) (*models.Crust, error) {
	var crust models.Crust
	if result := db.Where("name = ?", name).First(&crust); result.Error != nil {
		fmt.Println(result)
		return nil, result.Error
	}
	return &crust, nil
}

func getSizeByName(db *gorm.DB, name string) (*models.Size, error) {
	var size models.Size
	if result := db.Where("name = ?", name).First(&size); result.Error != nil {
		fmt.Println(result)
		return nil, result.Error
	}
	return &size, nil
}

func getToppingByName(db *gorm.DB, name string) (*models.Topping, error) {
	var topping models.Topping
	if result := db.Where("name = ?", name).First(&topping); result.Error != nil {
		fmt.Println(result)
		return nil, result.Error
	}

	return &topping, nil
}
