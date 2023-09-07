package pizzashop

import (
	"assg/pizzashop/pkg/common/models"
	"fmt"
	"time"

	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AddCustomerBodyType struct {
	Name   string `json:"name"`
	Number int64  `json:"number"`
}

type CustomerOrderBodyType struct {
	CustomerId int      `json:"customer_id"`
	PizzaName  string   `json:"pizza"`
	CrustName  string   `json:"crust"`
	Toppings   []string `json:"topping"`
	PizzaSize  string   `json:"pizza_size"`
}

type PriceCheckIds struct {
	PizzaId  int   `json:"pizza_id"`
	CrustId  int   `json:"crust_id"`
	SizeId   int   `json:"size_id"`
	Toppings []int `json:"toppings_id"`
}

type PriceCheckNames struct {
	PizzaName string   `json:"pizza_name"`
	CrustName string   `json:"crust_name"`
	SizeName  string   `json:"size_name"`
	Toppings  []string `json:"toppings_name"`
}

type CustomerOrder struct {
	OrderId     int    `json:"order_id"`
	PizzaName   string `json:"pizza_name"`
	CrustName   string `json:"crust_name"`
	SizeName    string `json:"size_name"`
	ToppingName string `json:"topping_name"`
}

func (h handler) getToppings(c *gin.Context) {
	var toppings []models.Topping

	if result := h.Db.Find(&toppings); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &toppings)
}

func (h handler) getPizzas(c *gin.Context) {
	var pizza []models.Pizza

	if result := h.Db.Find(&pizza); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &pizza)
}

func (h handler) getSizes(c *gin.Context) {
	var sizes []models.Size

	if result := h.Db.Find(&sizes); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &sizes)
}

func (h handler) getCrust(c *gin.Context) {
	var crust []models.Crust

	if result := h.Db.Find(&crust); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &crust)
}

func (h *handler) createCustomer(c *gin.Context) {
	body := AddCustomerBodyType{}

	if error := c.BindJSON(&body); error != nil {
		c.AbortWithError(http.StatusBadRequest, error)
		return
	}

	var customer models.Customer

	customer.Name = body.Name
	customer.Number = body.Number

	if result := h.Db.Create(&customer); result.Error != nil {
		c.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &customer)
}

func (h handler) getCustomerById(c *gin.Context) {
	id := c.Param("id")

	var customer models.Customer

	if result := h.Db.Find(&customer, id); result.Error != nil {
		c.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}

	c.JSON(http.StatusFound, &customer)
}

func (h handler) deleteTopping(c *gin.Context) {
	id := c.Param("id")

	var topping models.Topping

	if result := h.Db.Delete(&topping, id); result.Error != nil {
		c.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}

	//for making cascade work we have to perform permanent delete instead of soft delete

	//syntax for permanently delete
	// db.Unscoped().Delete(&order)

	c.JSON(http.StatusOK, topping)
}

func (h handler) createPizzaOrder(c *gin.Context) {
	body := CustomerOrderBodyType{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	//implementing transactions
	err := h.Db.Transaction(func(tx *gorm.DB) error {
		pizza, err := getPizzaByName(tx, body.PizzaName)

		crust, err := getCrustByName(tx, body.CrustName)

		size, err := getSizeByName(tx, body.PizzaSize)
		var toppingList []uint

		for _, topping := range body.Toppings {
			result, err := getToppingByName(tx, topping)
			if err != nil {
				tx.Rollback()
				break
			}

			toppingList = append(toppingList, result.ID)
		}

		if err != nil {
			tx.Rollback()
			return err
		}

		var PizzaOrder models.PizzaOrder
		PizzaOrder.CustomerID = uint(body.CustomerId)
		PizzaOrder.PizzaID = pizza.ID
		PizzaOrder.CrustID = crust.ID
		PizzaOrder.SizeID = size.ID

		if result := tx.Create(&PizzaOrder); result.Error != nil {
			return result.Error
		}

		var PizzaOrderTopping models.PizzaOrderTopping
		for _, topping := range toppingList {
			PizzaOrderTopping.PizzaOrderID = int(PizzaOrder.ID)
			PizzaOrderTopping.ToppingID = int(topping)

			if result := tx.Create(&PizzaOrderTopping); result.Error != nil {
				return result.Error
			}
		}

		return nil
	})

	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Pizza Created Succesfully"})
}

func (h handler) getPrice(c *gin.Context) {
	query := c.Query("format")

	var total float64 = 0

	if query == "id" {
		var body PriceCheckIds
		if err := c.BindJSON(body); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
		}

		pizza, err := getPizzasById(h.Db, body.PizzaId)
		if err != nil {
			c.AbortWithError(http.StatusNotFound, err)
			return
		}

		total += float64(pizza.Price)

		crust, err := getCrustById(h.Db, body.CrustId)
		if err != nil {
			c.AbortWithError(http.StatusNotFound, err)
			return
		}

		total += float64(crust.Price)

		for _, topping := range body.Toppings {
			result, err := getToppingsById(h.Db, topping)
			if err != nil {
				c.AbortWithError(http.StatusNotFound, err)
				return
			}

			total += float64(result.Price)
		}

		size, err := getSizesById(h.Db, body.SizeId)
		if err != nil {
			c.AbortWithError(http.StatusNotFound, err)
			return
		}

		total *= float64(size.Multiplier)
	} else {
		var body PriceCheckNames
		if err := c.BindJSON(&body); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
		}

		pizza, err := getPizzaByName(h.Db, body.PizzaName)
		if err != nil {
			c.AbortWithError(http.StatusNotFound, err)
			return
		}

		total += float64(pizza.Price)

		crust, err := getCrustByName(h.Db, body.CrustName)
		if err != nil {
			c.AbortWithError(http.StatusNotFound, err)
			return
		}
		total += float64(crust.Price)

		for _, topping := range body.Toppings {
			result, err := getToppingByName(h.Db, topping)
			if err != nil {
				c.AbortWithError(http.StatusNotFound, err)
				return
			}

			total += float64(result.Price)
		}

		size, err := getSizeByName(h.Db, body.SizeName)
		if err != nil {
			c.AbortWithError(http.StatusNotFound, err)
			return
		}

		total = total * float64(size.Multiplier)
	}

	response := map[string]float64{
		"price": total,
	}
	c.JSON(http.StatusOK, response)
}

func (h handler) getCustomerOrder(c *gin.Context) {
	id := c.Param("id")

	var customerOrder []CustomerOrder

	result := h.Db.Raw("SELECT po.id as order_id, p.name as pizza_name, c.name as crust_name, s.name as size_name, t.name as topping_name from pizza_orders po LEFT JOIN pizzas p ON po.pizza_id = p.id LEFT JOIN crusts c ON po.crust_id = c.id LEFT JOIN sizes s ON s.id = po.size_id JOIN pizza_order_toppings pot ON po.id = pot.pizza_order_id RIGHT JOIN toppings t ON pot.topping_id = t.id WHERE po.customer_id = ? ", id).Scan(&customerOrder)

	if result.Error != nil {
		c.AbortWithError(http.StatusBadRequest, result.Error)
	}

	c.JSON(http.StatusAccepted, customerOrder)
}

func FontCacheMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		fmt.Println("Middleware run succesfully")

		var fileExt = []string{".ttf", ".css"}

		for _, ext := range fileExt {
			if len(c.Request.URL.Path) >= len(ext) && c.Request.URL.Path[len(c.Request.URL.Path)-len(fileExt):] == ext {

				c.Header("Cache-Control", "public max-age=31536000")
				c.Header("Expires", time.Now().AddDate(1, 0, 0).Format(http.TimeFormat))
			}
		}

		c.Next()
	}
}
